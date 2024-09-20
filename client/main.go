package main

import (
	"fmt"
	"os"

	"github.com/SzymonMielecki/air_qual/client/cmd"
)

func Execute() {
	var apiKey string
	var city string
	var state string
	var country string

	RootCmd := cmd.RootCommand(city, state, country, apiKey)
	RootCmd.Flags().StringVarP(&city, "city", "c", "Warsaw", "City")
	RootCmd.Flags().StringVarP(&state, "state", "s", "Mazovia", "State")
	RootCmd.Flags().StringVarP(&country, "country", "t", "Poland", "Country")
	RootCmd.Flags().StringVarP(&apiKey, "apiKey", "k", "", "API key")
	RootCmd.MarkFlagRequired("apiKey")

	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	Execute()
}
