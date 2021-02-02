// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// FlatfeestackABI is the input ABI used to generate the binding from.
const FlatfeestackABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"PaymentReleased\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"address_\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"addresses_\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"balances_\",\"type\":\"uint256[]\"}],\"name\":\"fill\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"release\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// FlatfeestackFuncSigs maps the 4-byte function signature to its string representation.
var FlatfeestackFuncSigs = map[string]string{
	"70a08231": "balanceOf(address)",
	"3bc4baa4": "fill(address[],uint256[])",
	"86d1a69f": "release()",
}

// FlatfeestackBin is the compiled bytecode used for deploying new contracts.
var FlatfeestackBin = "0x608060405234801561001057600080fd5b50600180546001600160a01b03191633179055610585806100326000396000f3fe6080604052600436106100345760003560e01c80633bc4baa41461003957806370a082311461016257806386d1a69f146101a7575b600080fd5b6101606004803603604081101561004f57600080fd5b81019060208101813564010000000081111561006a57600080fd5b82018360208201111561007c57600080fd5b8035906020019184602083028401116401000000008311171561009e57600080fd5b91908080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525092959493602081019350359150506401000000008111156100ee57600080fd5b82018360208201111561010057600080fd5b8035906020019184602083028401116401000000008311171561012257600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295506101bc945050505050565b005b34801561016e57600080fd5b506101956004803603602081101561018557600080fd5b50356001600160a01b031661035f565b60408051918252519081900360200190f35b3480156101b357600080fd5b5061016061037a565b6001546001600160a01b031633146102055760405162461bcd60e51b815260040180806020018281038252602281526020018061052e6022913960400191505060405180910390fd5b80518251146102455760405162461bcd60e51b81526004018080602001828103825260368152602001806104ce6036913960400191505060405180910390fd5b6000805b835181101561031a576102af83828151811061026157fe5b602002602001015160008087858151811061027857fe5b60200260200101516001600160a01b03166001600160a01b031681526020019081526020016000205461044590919063ffffffff16565b6000808684815181106102be57fe5b60200260200101516001600160a01b03166001600160a01b03168152602001908152602001600020819055506103108382815181106102f957fe5b60200260200101518361044590919063ffffffff16565b9150600101610249565b503481111561035a5760405162461bcd60e51b815260040180806020018281038252602a815260200180610504602a913960400191505060405180910390fd5b505050565b6001600160a01b031660009081526020819052604090205490565b336000908152602081905260409020546103c55760405162461bcd60e51b81526004018080602001828103825260278152602001806104a76027913960400191505060405180910390fd5b33600081815260208190526040808220805490839055905190929183156108fc02918491818181858888f19350505050158015610406573d6000803e3d6000fd5b50604080513381526020810183905281517fdf20fd1e76bc69d672e4814fafb2c449bba3a5369d8359adf9e05e6fde87b056929181900390910190a150565b60008282018381101561049f576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b939250505056fe5061796d656e7453706c69747465723a206163636f756e7420686173206e6f2062616c616e636541646472657373657320616e642062616c616e636573206172726179206d7573742068617665207468652073616d65206c656e67746853756d206f662062616c616e63657320697320686967686572207468616e207061696420616d6f756e744f6e6c7920746865206f776e65722063616e20616464206e6577207061796f757473a26469706673582212202b8cf79a36e7f57cce5c7a341bb9923cfb47ef355ee9dc62e6b7f5754c61938a64736f6c63430007060033"

// DeployFlatfeestack deploys a new Ethereum contract, binding an instance of Flatfeestack to it.
func DeployFlatfeestack(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Flatfeestack, error) {
	parsed, err := abi.JSON(strings.NewReader(FlatfeestackABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(FlatfeestackBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Flatfeestack{FlatfeestackCaller: FlatfeestackCaller{contract: contract}, FlatfeestackTransactor: FlatfeestackTransactor{contract: contract}, FlatfeestackFilterer: FlatfeestackFilterer{contract: contract}}, nil
}

// Flatfeestack is an auto generated Go binding around an Ethereum contract.
type Flatfeestack struct {
	FlatfeestackCaller     // Read-only binding to the contract
	FlatfeestackTransactor // Write-only binding to the contract
	FlatfeestackFilterer   // Log filterer for contract events
}

// FlatfeestackCaller is an auto generated read-only Go binding around an Ethereum contract.
type FlatfeestackCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FlatfeestackTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FlatfeestackTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FlatfeestackFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FlatfeestackFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FlatfeestackSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FlatfeestackSession struct {
	Contract     *Flatfeestack     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FlatfeestackCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FlatfeestackCallerSession struct {
	Contract *FlatfeestackCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// FlatfeestackTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FlatfeestackTransactorSession struct {
	Contract     *FlatfeestackTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// FlatfeestackRaw is an auto generated low-level Go binding around an Ethereum contract.
type FlatfeestackRaw struct {
	Contract *Flatfeestack // Generic contract binding to access the raw methods on
}

// FlatfeestackCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FlatfeestackCallerRaw struct {
	Contract *FlatfeestackCaller // Generic read-only contract binding to access the raw methods on
}

// FlatfeestackTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FlatfeestackTransactorRaw struct {
	Contract *FlatfeestackTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFlatfeestack creates a new instance of Flatfeestack, bound to a specific deployed contract.
func NewFlatfeestack(address common.Address, backend bind.ContractBackend) (*Flatfeestack, error) {
	contract, err := bindFlatfeestack(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Flatfeestack{FlatfeestackCaller: FlatfeestackCaller{contract: contract}, FlatfeestackTransactor: FlatfeestackTransactor{contract: contract}, FlatfeestackFilterer: FlatfeestackFilterer{contract: contract}}, nil
}

// NewFlatfeestackCaller creates a new read-only instance of Flatfeestack, bound to a specific deployed contract.
func NewFlatfeestackCaller(address common.Address, caller bind.ContractCaller) (*FlatfeestackCaller, error) {
	contract, err := bindFlatfeestack(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FlatfeestackCaller{contract: contract}, nil
}

// NewFlatfeestackTransactor creates a new write-only instance of Flatfeestack, bound to a specific deployed contract.
func NewFlatfeestackTransactor(address common.Address, transactor bind.ContractTransactor) (*FlatfeestackTransactor, error) {
	contract, err := bindFlatfeestack(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FlatfeestackTransactor{contract: contract}, nil
}

// NewFlatfeestackFilterer creates a new log filterer instance of Flatfeestack, bound to a specific deployed contract.
func NewFlatfeestackFilterer(address common.Address, filterer bind.ContractFilterer) (*FlatfeestackFilterer, error) {
	contract, err := bindFlatfeestack(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FlatfeestackFilterer{contract: contract}, nil
}

// bindFlatfeestack binds a generic wrapper to an already deployed contract.
func bindFlatfeestack(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FlatfeestackABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Flatfeestack *FlatfeestackRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Flatfeestack.Contract.FlatfeestackCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Flatfeestack *FlatfeestackRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Flatfeestack.Contract.FlatfeestackTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Flatfeestack *FlatfeestackRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Flatfeestack.Contract.FlatfeestackTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Flatfeestack *FlatfeestackCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Flatfeestack.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Flatfeestack *FlatfeestackTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Flatfeestack.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Flatfeestack *FlatfeestackTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Flatfeestack.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address address_) view returns(uint256)
func (_Flatfeestack *FlatfeestackCaller) BalanceOf(opts *bind.CallOpts, address_ common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Flatfeestack.contract.Call(opts, &out, "balanceOf", address_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address address_) view returns(uint256)
func (_Flatfeestack *FlatfeestackSession) BalanceOf(address_ common.Address) (*big.Int, error) {
	return _Flatfeestack.Contract.BalanceOf(&_Flatfeestack.CallOpts, address_)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address address_) view returns(uint256)
func (_Flatfeestack *FlatfeestackCallerSession) BalanceOf(address_ common.Address) (*big.Int, error) {
	return _Flatfeestack.Contract.BalanceOf(&_Flatfeestack.CallOpts, address_)
}

// Fill is a paid mutator transaction binding the contract method 0x3bc4baa4.
//
// Solidity: function fill(address[] addresses_, uint256[] balances_) payable returns()
func (_Flatfeestack *FlatfeestackTransactor) Fill(opts *bind.TransactOpts, addresses_ []common.Address, balances_ []*big.Int) (*types.Transaction, error) {
	return _Flatfeestack.contract.Transact(opts, "fill", addresses_, balances_)
}

// Fill is a paid mutator transaction binding the contract method 0x3bc4baa4.
//
// Solidity: function fill(address[] addresses_, uint256[] balances_) payable returns()
func (_Flatfeestack *FlatfeestackSession) Fill(addresses_ []common.Address, balances_ []*big.Int) (*types.Transaction, error) {
	return _Flatfeestack.Contract.Fill(&_Flatfeestack.TransactOpts, addresses_, balances_)
}

// Fill is a paid mutator transaction binding the contract method 0x3bc4baa4.
//
// Solidity: function fill(address[] addresses_, uint256[] balances_) payable returns()
func (_Flatfeestack *FlatfeestackTransactorSession) Fill(addresses_ []common.Address, balances_ []*big.Int) (*types.Transaction, error) {
	return _Flatfeestack.Contract.Fill(&_Flatfeestack.TransactOpts, addresses_, balances_)
}

// Release is a paid mutator transaction binding the contract method 0x86d1a69f.
//
// Solidity: function release() returns()
func (_Flatfeestack *FlatfeestackTransactor) Release(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Flatfeestack.contract.Transact(opts, "release")
}

// Release is a paid mutator transaction binding the contract method 0x86d1a69f.
//
// Solidity: function release() returns()
func (_Flatfeestack *FlatfeestackSession) Release() (*types.Transaction, error) {
	return _Flatfeestack.Contract.Release(&_Flatfeestack.TransactOpts)
}

// Release is a paid mutator transaction binding the contract method 0x86d1a69f.
//
// Solidity: function release() returns()
func (_Flatfeestack *FlatfeestackTransactorSession) Release() (*types.Transaction, error) {
	return _Flatfeestack.Contract.Release(&_Flatfeestack.TransactOpts)
}

// FlatfeestackPaymentReleasedIterator is returned from FilterPaymentReleased and is used to iterate over the raw logs and unpacked data for PaymentReleased events raised by the Flatfeestack contract.
type FlatfeestackPaymentReleasedIterator struct {
	Event *FlatfeestackPaymentReleased // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FlatfeestackPaymentReleasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FlatfeestackPaymentReleased)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FlatfeestackPaymentReleased)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FlatfeestackPaymentReleasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FlatfeestackPaymentReleasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FlatfeestackPaymentReleased represents a PaymentReleased event raised by the Flatfeestack contract.
type FlatfeestackPaymentReleased struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPaymentReleased is a free log retrieval operation binding the contract event 0xdf20fd1e76bc69d672e4814fafb2c449bba3a5369d8359adf9e05e6fde87b056.
//
// Solidity: event PaymentReleased(address to, uint256 amount)
func (_Flatfeestack *FlatfeestackFilterer) FilterPaymentReleased(opts *bind.FilterOpts) (*FlatfeestackPaymentReleasedIterator, error) {

	logs, sub, err := _Flatfeestack.contract.FilterLogs(opts, "PaymentReleased")
	if err != nil {
		return nil, err
	}
	return &FlatfeestackPaymentReleasedIterator{contract: _Flatfeestack.contract, event: "PaymentReleased", logs: logs, sub: sub}, nil
}

// WatchPaymentReleased is a free log subscription operation binding the contract event 0xdf20fd1e76bc69d672e4814fafb2c449bba3a5369d8359adf9e05e6fde87b056.
//
// Solidity: event PaymentReleased(address to, uint256 amount)
func (_Flatfeestack *FlatfeestackFilterer) WatchPaymentReleased(opts *bind.WatchOpts, sink chan<- *FlatfeestackPaymentReleased) (event.Subscription, error) {

	logs, sub, err := _Flatfeestack.contract.WatchLogs(opts, "PaymentReleased")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FlatfeestackPaymentReleased)
				if err := _Flatfeestack.contract.UnpackLog(event, "PaymentReleased", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaymentReleased is a log parse operation binding the contract event 0xdf20fd1e76bc69d672e4814fafb2c449bba3a5369d8359adf9e05e6fde87b056.
//
// Solidity: event PaymentReleased(address to, uint256 amount)
func (_Flatfeestack *FlatfeestackFilterer) ParsePaymentReleased(log types.Log) (*FlatfeestackPaymentReleased, error) {
	event := new(FlatfeestackPaymentReleased)
	if err := _Flatfeestack.contract.UnpackLog(event, "PaymentReleased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
var SafeMathBin = "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212200be6ac6ac0047acd9b4634ad8e13bbe4642c1b4858b5fca9ee49ccbb9151a15b64736f6c63430007060033"

// DeploySafeMath deploys a new Ethereum contract, binding an instance of SafeMath to it.
func DeploySafeMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMath, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SafeMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// SafeMath is an auto generated Go binding around an Ethereum contract.
type SafeMath struct {
	SafeMathCaller     // Read-only binding to the contract
	SafeMathTransactor // Write-only binding to the contract
	SafeMathFilterer   // Log filterer for contract events
}

// SafeMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeMathSession struct {
	Contract     *SafeMath         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeMathCallerSession struct {
	Contract *SafeMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SafeMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeMathTransactorSession struct {
	Contract     *SafeMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SafeMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeMathRaw struct {
	Contract *SafeMath // Generic contract binding to access the raw methods on
}

// SafeMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeMathCallerRaw struct {
	Contract *SafeMathCaller // Generic read-only contract binding to access the raw methods on
}

// SafeMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeMathTransactorRaw struct {
	Contract *SafeMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeMath creates a new instance of SafeMath, bound to a specific deployed contract.
func NewSafeMath(address common.Address, backend bind.ContractBackend) (*SafeMath, error) {
	contract, err := bindSafeMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// NewSafeMathCaller creates a new read-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathCaller(address common.Address, caller bind.ContractCaller) (*SafeMathCaller, error) {
	contract, err := bindSafeMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathCaller{contract: contract}, nil
}

// NewSafeMathTransactor creates a new write-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeMathTransactor, error) {
	contract, err := bindSafeMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathTransactor{contract: contract}, nil
}

// NewSafeMathFilterer creates a new log filterer instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeMathFilterer, error) {
	contract, err := bindSafeMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeMathFilterer{contract: contract}, nil
}

// bindSafeMath binds a generic wrapper to an already deployed contract.
func bindSafeMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.SafeMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transact(opts, method, params...)
}
