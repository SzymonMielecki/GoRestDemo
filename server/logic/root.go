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

func (a *AppState) PutPollution(t time.Time, p persistance.Pollution) {
	a.db.AddPollution(p, t)
}

func (a *AppState) PutWeather(t time.Time, w persistance.Weather) {
	a.db.AddWeather(w, t)
}

func (a *AppState) GetClosestPolution(t time.Time) (persistance.Pollution, error) {
	return a.db.GetClosestPolution(t)
}

func (a *AppState) GetClosestWeather(t time.Time) (persistance.Weather, error) {
	return a.db.GetClosestWeather(t)
}
