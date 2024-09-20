package cmd

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

func RootCommand(city, state, country, apiKey string) *cobra.Command {
	return &cobra.Command{
		Use:   "air_qual",
		Short: "Send air quality data to the server",
		Long:  `Send air quality data to the server, which then stores it in a basic database.`,
		Run: func(cmd *cobra.Command, args []string) {
			formatedApiUrl := fmt.Sprintf("http://api.airvisual.com/v2/city?city=%s&state=%s&country=%s&key=%s", city, state, country, apiKey)
			resp, err := http.Get(formatedApiUrl)
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

			if _, err = http.Post(url_weather, contentType, bytes.NewBuffer(body)); err != nil {
				fmt.Println("Error: ", err)
				return
			}
			fmt.Println("Data sent to server")
		},
	}
}

func init() {
}
