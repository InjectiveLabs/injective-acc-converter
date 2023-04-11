package main

import (
	"fmt"
	cosmtypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

func main() {
	cosmtypes.GetConfig().SetBech32PrefixForAccount("inj", "injpub")
	accAddressString := "inj1hkhdaj2a2clmq5jq6mspsggqs32vynpk228q3r"
	accAddress, err := cosmtypes.AccAddressFromBech32(accAddressString)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s%024d", hexutil.Encode(accAddress.Bytes()[:20]), 0)
}
