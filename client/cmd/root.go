package cmd

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

var (
	apiKey  string
	city    string
	state   string
	country string
)

var rootCmd = &cobra.Command{
	Use:   "air_qual",
	Short: "Check air quality in your area",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		formatedUrl := fmt.Sprintf("http://api.airvisual.com/v2/city?city=%s&state=%s&country=%s&key=%s", city, state, country, apiKey)
		resp, err := http.Get(formatedUrl)
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
		port := os.Getenv("PORT")
		if port == "" {
			port = ":1323"
		}

		url_weather := fmt.Sprintf("http://localhost%s/pushData", port)

		contentType := "application/json"
		_, err = http.Post(url_weather, contentType, bytes.NewBuffer(body))
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
		fmt.Println("Data sent to server")
	},
}

func Connect() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&city, "city", "c", "Warsaw", "City")
	rootCmd.PersistentFlags().StringVarP(&state, "state", "s", "Mazovia", "State")
	rootCmd.PersistentFlags().StringVarP(&country, "country", "t", "Poland", "Country")
	rootCmd.PersistentFlags().StringVarP(&apiKey, "apiKey", "k", "", "API key")

	rootCmd.MarkFlagRequired("apiKey")
}
