package persistance

import (
	"errors"
	"time"

	"github.com/SzymonMielecki/GoRestDemo/server/utils"
)

type Db struct {
	pollution map[time.Time]utils.Pollution
	weather   map[time.Time]utils.Weather
}

func NewDb() *Db {
	return &Db{
		pollution: make(map[time.Time]utils.Pollution),
		weather:   make(map[time.Time]utils.Weather),
	}
}

func (d *Db) GetClosestPolution(closest time.Time) (best utils.Pollution, err error) {
	if len(d.pollution) == 0 {
		return utils.Pollution{}, errors.New("no pollution data")
	}
	timeDiff := time.Hour * 24 * 365
	for k, v := range d.pollution {
		if time.Time.Sub(k, closest) < timeDiff {
			timeDiff = time.Time.Sub(k, closest)
			best = v
		}
	}
	return best, nil
}

func (d *Db) GetClosestWeather(closest time.Time) (best utils.Weather, err error) {
	if len(d.weather) == 0 {
		return utils.Weather{}, errors.New("no weather data")
	}
	timeDiff := time.Hour * 24 * 365
	for k, v := range d.weather {
		if time.Time.Sub(k, closest) < timeDiff {
			timeDiff = time.Time.Sub(k, closest)
			best = v
		}
	}
	return best, nil
}

func (d *Db) AddPollution(p utils.Pollution, t time.Time) {
	d.pollution[t] = p
}

func (d *Db) AddWeather(w utils.Weather, t time.Time) {
	d.weather[t] = w
}
