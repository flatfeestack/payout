package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/nspcc-dev/neo-go/pkg/core/native/nativenames"
	"github.com/nspcc-dev/neo-go/pkg/core/state"
	"github.com/nspcc-dev/neo-go/pkg/core/transaction"
	"github.com/nspcc-dev/neo-go/pkg/crypto/keys"
	"github.com/nspcc-dev/neo-go/pkg/encoding/address"
	"github.com/nspcc-dev/neo-go/pkg/io"
	neo "github.com/nspcc-dev/neo-go/pkg/rpcclient"
	"github.com/nspcc-dev/neo-go/pkg/smartcontract"
	"github.com/nspcc-dev/neo-go/pkg/smartcontract/callflag"
	"github.com/nspcc-dev/neo-go/pkg/smartcontract/manifest"
	"github.com/nspcc-dev/neo-go/pkg/smartcontract/nef"
	"github.com/nspcc-dev/neo-go/pkg/util"
	"github.com/nspcc-dev/neo-go/pkg/vm/emit"
	"github.com/nspcc-dev/neo-go/pkg/wallet"
	"io/ioutil"
	"log"
	"math/big"
)

func getNeoClient(endpoint string) (*neo.Client, error) {
	neoClient, err := neo.New(context.Background(), endpoint, neo.Options{})

	if err != nil {
		return nil, err
	}

	err = neoClient.Init()
	if err != nil {
		return nil, err
	}

	contractOwnerPrivateKey, err := keys.NewPrivateKeyFromWIF(opts.NEO.PrivateKey)
	if err != nil {
		return nil, err
	}
	// signatureBytes := signature_provider.NewSignatureNeo(dev, tea, contractOwnerPrivateKey)

	// Following the steps on the developer's side after receiving the signature bytes:
	// Create and initialize client
	// Developer received the signature bytes and can now create the transaction to withdraw funds
	owner := wallet.NewAccountFromPrivateKey(contractOwnerPrivateKey)

	if opts.NEO.Deploy {
		h, err := deploy(neoClient, owner)
		if err != nil {
			return nil, err
		} else {
			opts.NEO.Contract = h.StringLE()
		}
	}
	return neoClient, nil
}

func payoutNEO(addressValues []string, teas []*big.Int) (string, error) {
	var payoutNeoHash, err = util.Uint160DecodeStringLE(opts.NEO.Contract)
	if err != nil {
		log.Fatalf(err.Error())
		return "", err
	}
	contractOwnerPrivateKey, err := keys.NewPrivateKeyFromWIF(opts.NEO.PrivateKey)
	if err != nil {
		log.Fatalf(err.Error())
		return "", err
	}
	owner := wallet.NewAccountFromPrivateKey(contractOwnerPrivateKey)

	h := CreateBatchPayoutTx(neoClient, payoutNeoHash, owner, addressValues, teas)
	return h, nil
}

func CreateBatchPayoutTx(c *neo.Client, payoutNeoHash util.Uint160, acc *wallet.Account, addressValues []string, teas []*big.Int) string {
	var devP []interface{}
	for _, v := range addressValues {
		add, _ := address.StringToUint160(v)
		devP = append(devP, add)
	}
	var teaP []interface{}
	for _, v := range teas {
		teaP = append(teaP, v)
	}

	w := io.NewBufBinWriter()
	emit.AppCall(w.BinWriter, payoutNeoHash, "batchPayout", callflag.All, devP, teaP)
	script := w.Bytes()
	log.Printf("About to execute job [batchPayout]")
	sender := acc.PrivateKey().GetScriptHash()
	signer := transaction.Signer{
		Account: sender,
		Scopes:  transaction.CalledByEntry,
	}
	tx, err := c.CreateTxFromScript(script, acc, -1, 0, []neo.SignerAccount{{
		Signer: signer,
	}})
	if err != nil {
		log.Fatalf(err.Error())
	}
	net, err := c.GetNetwork()
	if err != nil {
		log.Fatalf(err.Error())
	}
	err = acc.SignTx(net, tx)
	if err != nil {
		log.Fatalf(err.Error())
	}
	hash, err := c.SendRawTransaction(tx)
	if err != nil {
		log.Fatalf("send raw transaction err: %v", err)
	}
	return hash.StringLE()
}

func readNEFFile(filename string) (*nef.File, []byte, error) {
	if len(filename) == 0 {
		return nil, nil, errors.New("no nef file was provided")
	}

	f, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf(err.Error())
		return nil, nil, err
	}

	nefFile, err := nef.FileFromBytes(f)
	if err != nil {
		return nil, nil, fmt.Errorf("can't parse NEF file: %w", err)
	}

	return &nefFile, f, nil
}

func readManifest(filename string) (*manifest.Manifest, []byte, error) {
	if len(filename) == 0 {
		return nil, nil, errors.New("no manifest file was provided")
	}

	manifestBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf(err.Error())
		return nil, nil, err
	}

	m := new(manifest.Manifest)
	err = json.Unmarshal(manifestBytes, m)
	if err != nil {
		return nil, nil, err
	}
	return m, manifestBytes, nil
}

func deploy(c *neo.Client, acc *wallet.Account) (util.Uint160, error) {
	nativeManagementContractHash, err := c.GetNativeContractHash(nativenames.Management)
	if err != nil {
		return util.Uint160{}, err
	}
	ne, nefB, err := readNEFFile("./PayoutNeo.nef")
	_, mfB, err := readManifest("./PayoutNeo.manifest.json")
	sender := acc.PrivateKey().GetScriptHash()
	pk := acc.PrivateKey().PublicKey().Bytes()
	appCallParams := []smartcontract.Parameter{
		{
			Type:  smartcontract.ByteArrayType,
			Value: nefB,
		},
		{
			Type:  smartcontract.ByteArrayType,
			Value: mfB,
		},
		{
			Type:  smartcontract.PublicKeyType,
			Value: pk,
		},
	}

	contractHash := state.CreateContractHash(sender, ne.Checksum, "PayoutNeo")
	signer := transaction.Signer{
		Account: sender,
		Scopes:  transaction.Global,
		// CustomContracts do not work with neo-go, if that scope is used for the sender when using the method CreateTxFromScript.
		// Same holds for CustomGroups...
		//Scopes:           transaction.CustomContracts,
		//AllowedContracts: []util.Uint160{contractHash},
	}
	resp, _ := c.InvokeFunction(nativeManagementContractHash, "deploy", appCallParams, []transaction.Signer{signer})
	tx, err := c.CreateTxFromScript(resp.Script, acc, -1, 0, []neo.SignerAccount{{Signer: signer}})
	if err != nil {
		return util.Uint160{}, err
	}
	txHash, err := c.SignAndPushTx(tx, acc, nil)
	if err != nil {
		return util.Uint160{}, err
	}
	fmt.Println("---------------------------------")
	fmt.Println("NEO Transaction: " + txHash.StringLE())
	fmt.Println("NEO smart contract deployed: " + contractHash.StringLE())
	fmt.Println("---------------------------------")
	return contractHash, nil
}
