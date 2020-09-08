package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Line string

var AEL = []string{
	"HOK: Hong Kong",
	"KOW: Kowloon",
	"TSY: Tsing Yi",
	"AIR: Airport",
	"AWE: AsiaWorld Expo",
}

var TCL = []string{
	"HOK: Hong Kong",
	"KOW: Kowloon",
	"OLY: Olympic",
	"NAC: Nam Cheong",
	"LAK: Lai King",
	"TSY: Tsing Yi",
	"SUN: Sunny Bay",
	"TUC: Tung Chung",
}

var WRL = []string{
	"HUH: Hung Hom",
	"ETS: East Tsim Sha Tsui",
	"AUS: Austin",
	"NAC: Nam Cheong",
	"MEF: Mei Foo",
	"TWW: Tsuen Wan West",
	"KSR: Kam Sheung Road",
	"YUL: Yuen Long",
	"LOP: Long Ping",
	"TIS: Tin Shui Wai",
	"SIH: Siu Hong",
	"TUM: Tuen Mun",
}

var TKL = []string{
	"NOP: North Point",
	"QUB: Quarry Bay",
	"YAT: Yau Tong",
	"TIK: Tiu Keng Leng",
	"TKO: Tseung Kwan O",
	"LHP: LOHAS Park",
	"HAH: Hang Hau",
	"POA: Po Lam",
}

func init() {
	//cobra.OnInitialize(initConfig)
	mtrstationCmd.Flags().StringVarP(&Line, "line", "l", "", "your MTR Line (required)")
	mtrstationCmd.MarkFlagRequired("line")
	rootCmd.AddCommand(mtrstationCmd)
}

var mtrstationCmd = &cobra.Command{
	Use:   "station",
	Short: "Check out the available MTR station for a specific line",
	Long:  `Check out the available MTR station for a specific line. Due to the MTR API limitation, only 4 lines is available.`,
	Run: func(cmd *cobra.Command, args []string) {
		switch Line {
		case "AEL":
			for i := 0; i < len(AEL); i++ {
				fmt.Println(AEL[i])
			}
		case "TCL":
			for i := 0; i < len(TCL); i++ {
				fmt.Println(TCL[i])
			}
		case "WRL":
			for i := 0; i < len(WRL); i++ {
				fmt.Println(WRL[i])
			}
		case "TKL":
			for i := 0; i < len(TKL); i++ {
				fmt.Println(TKL[i])
			}
		default:
			fmt.Println("Incorrect line name. Please check your spelling (in provided short form).")
		}
	},
}
