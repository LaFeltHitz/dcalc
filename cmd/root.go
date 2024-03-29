package cmd

import (
	// "fmt"
	"github.com/spf13/cobra"
	// "os"
)

const (
	_1m float64 = 60000
)

var tempo float64

var rootCmd = &cobra.Command{
	Use:   "dcalc",
	Short: "Delay Calculator calculates note delay based on a given tempo",
	Run: func(cmd *cobra.Command, args []string) {
		
	}
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	var flagName = "t"
	rootCmd.Flags().Float64Var(&tempo, flagName, 120, "beat tempo")
	rootCmd.MarkFlagRequired(flagName)
}

func calculateDelay(tempo *float64) float64 {
	return _1m / *tempo
}
