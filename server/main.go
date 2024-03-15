package main

import (
	"os"

	"github.com/SzymonMielecki/air_qual/server/endpoint"
	"github.com/SzymonMielecki/air_qual/server/logic"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	state := logic.NewAppState()

	e.POST("/pollution", endpoint.AddPollution(*state))
	e.POST("/weather", endpoint.AddWeather(*state))
	e.GET("/pollution/:time", endpoint.GetClosestPolution(*state))
	e.GET("/weather/:time", endpoint.GetClosestWeather(*state))

	port := os.Getenv("PORT")

	if port == "" {
		port = ":1323"
	}

	e.Logger.Fatal(e.Start(port))
}
