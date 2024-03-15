package persistance

import (
	"errors"
	"time"
)

type Db struct {
	pollution map[time.Time]Pollution
	weather   map[time.Time]Weather
}

type Pollution struct {
	Maincn string
	Mainus string
	Aqius  int
	Aqicn  int
}

type Weather struct {
	Temperature   int
	Pressure      int
	Humidity      int
	WindSpeed     float64
	WindDirection int
}

func NewDb() *Db {
	return &Db{
		pollution: make(map[time.Time]Pollution),
		weather:   make(map[time.Time]Weather),
	}
}

func (d *Db) GetClosestPolution(closest time.Time) (best Pollution, err error) {
	if len(d.pollution) == 0 {
		return Pollution{}, errors.New("no pollution data")
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

func (d *Db) GetClosestWeather(closest time.Time) (best Weather, err error) {
	if len(d.weather) == 0 {
		return Weather{}, errors.New("no weather data")
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

func (d *Db) AddPollution(p Pollution, t time.Time) {
	d.pollution[t] = p
}

func (d *Db) AddWeather(w Weather, t time.Time) {
	d.weather[t] = w
}
