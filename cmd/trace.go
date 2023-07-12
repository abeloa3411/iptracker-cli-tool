/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/ipinfo/go/v2/ipinfo"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

type IP struct{
	IP string `json:"ip"`
	Country string `json:"country"`
	City string `json:"city"`
}

// traceCmd represents the trace command
var traceCmd = &cobra.Command{
	Use:   "trace",
	Short: "Trace the IP",
	Long: `Trace the IP.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			for _, ip := range args {
				getIP(ip)
			}
		}else{
			fmt.Println("Please provide an IP")
		}
	},
}

func init() {
	rootCmd.AddCommand(traceCmd)
}


func getIP(ip_address string){

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal(err)
	}
	
	// params: httpClient, cache, token. `http.DefaultClient` and no cache will be used in case of `nil`.
	client := ipinfo.NewClient(nil, nil, os.Getenv("TOKEN"))

	info, err := client.GetIPInfo(net.ParseIP(ip_address))

	
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(info)
	// Output: {8.8.8.8 dns.google false true Mountain View California US United States...
}