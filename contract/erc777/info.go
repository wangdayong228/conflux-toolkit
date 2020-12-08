package erc777

import (
	"fmt"
	"math/big"

	"github.com/Conflux-Chain/conflux-toolkit/contract/common"
	"github.com/Conflux-Chain/conflux-toolkit/util"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:   "info",
		Short: "Query token info",
		Run:   queryInfo,
	})
}

func queryInfo(cmd *cobra.Command, args []string) {
	contract := common.MustCreateContract(abiJSON)
	defer contract.Client.Close()

	var result *big.Int

	common.MustCall(contract, &result, "gas_total_limit")
	fmt.Println("Total gas:", util.DisplayValueWithUnit(result))
}
