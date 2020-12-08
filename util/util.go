package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/big"
	"os"

	"github.com/shopspring/decimal"
	"github.com/spf13/cobra"
)

// CreateUsageCommand creates a command to display help.
func CreateUsageCommand(use, short string) *cobra.Command {
	return &cobra.Command{
		Use:   use,
		Short: short,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
}

// MustParseBigInt parses the specified value to number and returns value * 10 ^ exp.
func MustParseBigInt(value string, exp int32) *big.Int {
	num, err := decimal.NewFromString(value)
	if err != nil {
		fmt.Println("invalid decimal format for value:", err.Error())
		os.Exit(1)
	}

	return decimal.NewFromBigInt(big.NewInt(1), exp).Mul(num).BigInt()
}

// DisplayValueWithUnit returns the display format for given drip value.
func DisplayValueWithUnit(drip *big.Int) string {
	if big.NewInt(1_000).Cmp(drip) > 0 {
		return fmt.Sprintf("%v Drip", drip)
	}

	if big.NewInt(1_000_000).Cmp(drip) > 0 {
		return fmt.Sprintf("%v Kdrip", decimal.NewFromBigInt(drip, -3))
	}

	if big.NewInt(1_000_000_000).Cmp(drip) > 0 {
		return fmt.Sprintf("%v Mdrip", decimal.NewFromBigInt(drip, -6))
	}

	if big.NewInt(1_000_000_000_000).Cmp(drip) > 0 {
		return fmt.Sprintf("%v Gdrip", decimal.NewFromBigInt(drip, -9))
	}

	return fmt.Sprintf("%v CFX", decimal.NewFromBigInt(drip, -18))
}

// OsExitIfErr prints error msg and exit
func OsExitIfErr(err error, format string, a ...interface{}) {
	if err != nil {
		fmt.Printf(format, a...)
		fmt.Printf("--- error: %v", err)
		fmt.Println()
		os.Exit(1)
	}
}

// OsExit prints msg and exit
func OsExit(format string, a ...interface{}) {
	fmt.Printf(format, a...)
	fmt.Println()
	os.Exit(1)
}

// JSONFmt format json string with indent
func JSONFmt(jsonStr []byte) string {
	var str bytes.Buffer
	_ = json.Indent(&str, jsonStr, "", "    ")
	return str.String()
}

// JSONMarshalAndFmt marshal v to json and format it
func JSONMarshalAndFmt(v interface{}) string {
	j, e := json.Marshal(v)
	OsExitIfErr(e, "marshal %v to json error", v)
	return JSONFmt(j)
}
