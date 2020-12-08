package erc777

import (
	"fmt"
	"math/big"

	"github.com/Conflux-Chain/conflux-toolkit/account"
	"github.com/Conflux-Chain/conflux-toolkit/contract/common"
	"github.com/Conflux-Chain/conflux-toolkit/util"
	"github.com/Conflux-Chain/go-conflux-sdk/types"
	"github.com/spf13/cobra"
)

func init() {
	balanceofCommand := cobra.Command{
		Use:   "balance",
		Short: "Query balance of user",
		Run:   balanceOf,
	}
	account.AddAccountVar(&balanceofCommand)

	rootCmd.AddCommand(&balanceofCommand)
}

func balanceOf(cmd *cobra.Command, args []string) {
	contract := common.MustCreateContract(abiJSON)
	defer contract.Client.Close()

	var result *big.Int

	common.MustCall(contract, &result, "balanceOf", types.Address(account.MustParseAccount()).ToCommonAddress())
	fmt.Println("balance:", util.DisplayValueWithUnit(result))
}
