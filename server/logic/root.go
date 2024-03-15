package logic

import (
	"time"

	"github.com/SzymonMielecki/air_qual/server/persistance"
)

type AppState struct {
	db *persistance.Db
}

func NewAppState() *AppState {
	return &AppState{
		db: persistance.NewDb(),
	}
}

func (a *AppState) AddPollution(t time.Time, p persistance.Pollution) {
	a.db.AddPollution(p, t)
}

func (a *AppState) AddWeather(t time.Time, w persistance.Weather) {
	if w.Temperature < -50 || w.Temperature > 50 {
		return
	}
	if w.Pressure < 800 || w.Pressure > 1200 {
		return
	}
	if w.Humidity < 0 || w.Humidity > 100 {
		return
	}
	if w.WindSpeed < 0 || w.WindSpeed > 300 {
		return
	}

	a.db.AddWeather(w, t)
}

func (a *AppState) GetClosestPolution(t time.Time) (persistance.Pollution, error) {
	return a.db.GetClosestPolution(t)
}

func (a *AppState) GetClosestWeather(t time.Time) (persistance.Weather, error) {
	return a.db.GetClosestWeather(t)
}
