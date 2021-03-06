package account

import (
	"encoding/hex"
	"fmt"

	"github.com/Conflux-Chain/go-conflux-sdk/types"
	"github.com/Conflux-Chain/go-conflux-sdk/utils"
	"github.com/spf13/cobra"
)

var (
	nonce        uint32
	to           string
	gasLimit     uint32
	storageLimit uint64
	epoch        uint64
	chain        uint
	data         string
)

func init() {
	signCmd := &cobra.Command{
		Use:   "sign",
		Short: "Sign transaction to send",
		Run:   sign,
	}

	AddFromVar(signCmd)

	signCmd.PersistentFlags().Uint32Var(&nonce, "nonce", 0, "Transaction nonce")
	signCmd.MarkPersistentFlagRequired("nonce")

	signCmd.PersistentFlags().StringVar(&to, "to", "", "To address in HEX format")
	signCmd.MarkPersistentFlagRequired("to")

	AddGasPriceVar(signCmd)

	signCmd.PersistentFlags().Uint32Var(&gasLimit, "gas", 21000, "Gas limit")

	signCmd.PersistentFlags().StringVar(&ValueCfx, "value", "", "Value to transfer in CFX")
	signCmd.MarkPersistentFlagRequired("value")

	signCmd.PersistentFlags().Uint64Var(&storageLimit, "storage", 0, "Storage limit")

	signCmd.PersistentFlags().Uint64Var(&epoch, "epoch", 0, "Transaction epoch height")
	signCmd.MarkPersistentFlagRequired("epoch")

	signCmd.PersistentFlags().UintVar(&chain, "chain", 1029, "Conflux chain ID")

	signCmd.PersistentFlags().StringVar(&data, "data", "", "Transaction data or encoded ABI data in HEX format")

	rootCmd.AddCommand(signCmd)
}

func sign(cmd *cobra.Command, args []string) {
	tx := types.UnsignedTransaction{
		UnsignedTransactionBase: types.UnsignedTransactionBase{
			From:         types.NewAddress(MustParseAccount()),
			Nonce:        types.NewBigInt(int64(nonce)),
			GasPrice:     types.NewBigIntByRaw(MustParsePrice()),
			Gas:          types.NewBigInt(int64(gasLimit)),
			Value:        types.NewBigIntByRaw(MustParseValue()),
			StorageLimit: types.NewUint64(storageLimit),
			EpochHeight:  types.NewUint64(epoch),
			ChainID:      types.NewUint(chain),
		},
		To: types.NewAddress(to),
	}

	if len(data) > 0 {
		txData, err := utils.HexStringToBytes(data)
		if err != nil {
			fmt.Println("Invalid tx data:", err.Error())
			return
		}
		tx.Data = txData
	}

	password := MustInputPassword("Enter password: ")

	encoded, err := am.SignAndEcodeTransactionWithPassphrase(tx, password)
	if err != nil {
		fmt.Println("Failed to sign transaction:", err.Error())
		return
	}

	fmt.Println("=======================================")
	fmt.Println("0x" + hex.EncodeToString(encoded))
}
