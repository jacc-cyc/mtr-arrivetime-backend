package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var lines = []string{
	"AEL: Airport Express",
	"TCL: Tung Chung Line",
	"WRL: West Rail Line",
	"TKL: Tseung Kwan O Line",
}

func init() {
	//cobra.OnInitialize(initConfig)
	rootCmd.AddCommand(mtrlineCmd)
}

var mtrlineCmd = &cobra.Command{
	Use:   "line",
	Short: "Check out the available MTR lines",
	Long:  `Check out the available MTR lines. Due to the MTR API limitation, only 4 lines is available.`,

	Run: func(cmd *cobra.Command, args []string) {
		for i := 0; i < 4; i++ {
			fmt.Println(lines[i])
		}
	},
}
