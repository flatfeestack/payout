# Flatfeestack Payout

## Compile contract
Install abigen:

``cd $GOPATH/src/github.com/ethereum/go-ethereum && make && make devtools``

Then use it to compile the contract

``abigen --sol=Flatfeestack.sol --pkg=main  --out=contract.go``