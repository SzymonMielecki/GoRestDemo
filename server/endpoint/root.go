package endpoint

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/SzymonMielecki/air_qual/server/logic"
	"github.com/SzymonMielecki/air_qual/server/persistance"
	"github.com/labstack/echo/v4"
)

func GetClosestPolution(a logic.AppState) echo.HandlerFunc {
	return func(c echo.Context) error {
		time, err := time.Parse(time.RFC3339, c.Param("time"))
		if err != nil {
			return c.String(http.StatusBadRequest, "Error parsing time")
		}
		p, err := a.GetClosestPolution(time)
		if err != nil {
			return c.String(http.StatusInternalServerError, "Error getting closest pollution")
		}
		json, err := json.Marshal(p)
		if err != nil {
			return c.String(http.StatusInternalServerError, "Error marshalling pollution")
		}
		return c.JSON(http.StatusOK, json)
	}
}

func GetClosestWeather(a logic.AppState) echo.HandlerFunc {
	return func(c echo.Context) error {
		time, err := time.Parse(time.RFC3339, c.Param("time"))
		if err != nil {
			return c.String(http.StatusBadRequest, "Error parsing time")
		}
		w, err := a.GetClosestWeather(time)
		if err != nil {
			return c.String(http.StatusInternalServerError, "Error getting closest weather")
		}
		json, err := json.Marshal(w)
		if err != nil {
			return c.String(http.StatusInternalServerError, "Error marshalling weather")
		}
		return c.JSON(http.StatusOK, json)
	}
}

func AddPollution(a logic.AppState) echo.HandlerFunc {
	return func(c echo.Context) error {
		pollution := persistance.Pollution{}
		body, err := io.ReadAll(c.Request().Body)
		if err != nil {
			return c.String(http.StatusInternalServerError, "Error reading the request body")
		}
		if err := json.Unmarshal(body, &pollution); err != nil {
			return c.String(http.StatusInternalServerError, "Error unmarshalling pollution")
		}
		time := time.Now()
		a.AddPollution(time, pollution)
		return c.String(http.StatusOK, "Pollution added")
	}
}

func AddWeather(a logic.AppState) echo.HandlerFunc {
	return func(c echo.Context) error {
		weather := persistance.Weather{}
		body, err := io.ReadAll(c.Request().Body)
		if err != nil {
			return c.String(http.StatusInternalServerError, "Error reading the request body")
		}
		if err := json.Unmarshal(body, &weather); err != nil {
			return c.String(http.StatusInternalServerError, "Error unmarshalling weather")
		}
		time := time.Now()
		a.AddWeather(time, weather)
		return c.String(http.StatusOK, "Weather Added")
	}
}
