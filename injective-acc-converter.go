package main

import (
	"fmt"
	"flag"
	cosmtypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

func main() {
	// Define a new command-line flag "address"
	addressPtr := flag.String("address", "inj1hkhdaj2a2clmq5jq6mspsggqs32vynpk228q3r", "the account address to encode")

	// Parse the command-line arguments
	flag.Parse()

	cosmtypes.GetConfig().SetBech32PrefixForAccount("inj", "injpub")
	// Convert the account address from Bech32 string to AccAddress type
	accAddress, err := cosmtypes.AccAddressFromBech32(*addressPtr)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s%024d", hexutil.Encode(accAddress.Bytes()[:20]), 0)
}
