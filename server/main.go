package main

import (
	"os"

	"github.com/SzymonMielecki/GoRestDemo/server/endpoint"
	"github.com/SzymonMielecki/GoRestDemo/server/logic"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	state := logic.NewAppState()

	e.POST("/pushData", endpoint.AddUniversal(*state))
	e.GET("/pollution/:time", endpoint.GetClosestPolution(*state))
	e.GET("/weather/:time", endpoint.GetClosestWeather(*state))

	port := os.Getenv("PORT")

	if port == "" {
		port = ":1323"
	}

	e.Logger.Fatal(e.Start(port))
}
