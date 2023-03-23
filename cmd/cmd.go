package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var cmd = &cobra.Command{
	Use:   "rayx",
	Short: "xray poc executor",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var pocs string
var url string
var headers = newMapValue()
var cookies = newMapValue()

func init() {
	flags := cmd.Flags()
	flags.StringVarP(&pocs, "poc", "p", "", "POC file/folder (eg. path/to/poc,rce.yaml)")
	flags.StringVarP(&url, "url", "u", "", "target url (eg. https://httpbin.org)")

	flags.Var(&headers, "header", "(eg header1=value1,header1=value2)")
	flags.Var(&cookies, "cookie", "(eg cookie1=value1,cookie2=value2)")

}

func Execute() {
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
