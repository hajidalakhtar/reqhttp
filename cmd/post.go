package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"go-http-cli/exception"
	"go-http-cli/formatter"
	"net/http"
	"net/url"
	"strings"
)

var postCmd = &cobra.Command{
	Use:   "post",
	Short: "-",
	Long:  `-`,
	Run: func(cmd *cobra.Command, args []string) {
		isFlagNull := cmd.Flags().NFlag()
		isBody, _ := cmd.Flags().GetBool("body")
		isHeader, _ := cmd.Flags().GetBool("header")
		isStatus, _ := cmd.Flags().GetBool("status")

		var urlInput string
		var bodyInputRaw string
		switch args[0] {
		case "http":
			urlInput = "http://" + args[1]
			if len(args) == 3 {
				bodyInputRaw = args[2]
			}
		case "https":
			urlInput = "https://" + args[1]
			if len(args) == 3 {
				bodyInputRaw = args[2]
			}
		default:
			urlInput = args[0]
			if len(args) == 2 {
				bodyInputRaw = args[1]
			}
		}

		bodyInput := strings.Replace(bodyInputRaw, ",", "&", -1)
		v, err := url.ParseQuery(bodyInput)
		if err != nil {
			panic(err)
		}
		resp, err := http.PostForm(urlInput, v)
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
	rootCmd.AddCommand(postCmd)
	postCmd.PersistentFlags().Bool("body", false, "show body response")
	postCmd.PersistentFlags().Bool("header", false, "show header response")
	postCmd.PersistentFlags().Bool("status", false, "show header response")
}
