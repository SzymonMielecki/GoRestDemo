package endpoint

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/SzymonMielecki/air_qual/server/logic"
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

func AddUniversal(a logic.AppState) echo.HandlerFunc {
	return func(c echo.Context) error {
		body, err := io.ReadAll(c.Request().Body)
		if err != nil {
			return c.String(http.StatusInternalServerError, "Error reading the request body")
		}
		a.AddUniversal(body)
		return c.String(http.StatusOK, "Data added")
	}
}
