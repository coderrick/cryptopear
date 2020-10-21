package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/ethclient"
	ens "github.com/wealdtech/go-ens/v3"
)

func main() {
	// Replace SECRET with your own access token for this example to work.
	client, err := ethclient.Dial("http://127.0.0.1:8545") //https://mainnet.infura.io/v3/SECRET
	if err != nil {
		panic(err)
	}

	// Resolve a name to an address
	domain := "cryptopear.eth"
	address, err := ens.Resolve(client, domain)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Address of %s is %s\n", domain, address.Hex())

	// Reverse resolve an address to a name
	reverse, err := ens.ReverseResolve(client, address)
	if err != nil {
		panic(err)
	}
	if reverse == "" {
		fmt.Printf("%s has no reverse lookup\n", address.Hex())
	} else {
		fmt.Printf("Name of %s is %s\n", address.Hex(), reverse)
	}
}
