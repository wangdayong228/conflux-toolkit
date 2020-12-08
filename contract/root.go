package contract

import (
	"github.com/Conflux-Chain/conflux-toolkit/contract/common"
	"github.com/Conflux-Chain/conflux-toolkit/contract/erc777"
	"github.com/Conflux-Chain/conflux-toolkit/contract/sponsorfaucet"
	"github.com/Conflux-Chain/conflux-toolkit/rpc"
	"github.com/Conflux-Chain/conflux-toolkit/util"
	"github.com/spf13/cobra"
)

var rootCmd = util.CreateUsageCommand("contract", "Contract subcommand")

func init() {
	rpc.AddURLVar(rootCmd)
	common.AddContractVar(rootCmd)

	sponsorfaucet.SetParent(rootCmd)
	erc777.SetParent(rootCmd)
}

// SetParent sets parent command
func SetParent(parent *cobra.Command) {
	parent.AddCommand(rootCmd)
}
