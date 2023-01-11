package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"go-http-cli/exception"
	"go-http-cli/formatter"
	"net/http"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "-",
	Long:  `-`,
	Run: func(cmd *cobra.Command, args []string) {
		isFlagNull := cmd.Flags().NFlag()
		isBody, _ := cmd.Flags().GetBool("body")
		isHeader, _ := cmd.Flags().GetBool("header")
		isStatus, _ := cmd.Flags().GetBool("status")

		var url string
		switch args[0] {
		case "http":
			url = "http://" + args[1]
		case "https":
			url = "https://" + args[1]
		default:
			url = args[0]

		}

		resp, err := http.Get(url)
		formatter := formatter.NewFormatter(resp)
		if err != nil {
			exception.ConnectionFailed(err)
		} else {
			if isStatus || isFlagNull == 0 {
				statusResponse := formatter.GetStatusResponse()
				fmt.Println(statusResponse)
			}
			if isHeader || isFlagNull == 0 {
				headerResponse := formatter.GetHeaderResponse()
				fmt.Println(headerResponse)
				fmt.Println("")
			}

			if isBody || isFlagNull == 0 {
				bodyResponse := formatter.GetBodyResponse()
				fmt.Println(bodyResponse)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
	getCmd.PersistentFlags().Bool("body", false, "show body response")
	getCmd.PersistentFlags().Bool("header", false, "show header response")
	getCmd.PersistentFlags().Bool("status", false, "show header response")
}
