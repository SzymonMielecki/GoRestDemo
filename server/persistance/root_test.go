package persistance

import (
	"reflect"
	"testing"
	"time"

	"github.com/SzymonMielecki/GoRestDemo/server/utils"
)

func TestNewDb(t *testing.T) {
	tests := []struct {
		want *Db
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDb(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDb() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDb_GetClosestPolution(t *testing.T) {
	type fields struct {
		pollution map[time.Time]utils.Pollution
		weather   map[time.Time]utils.Weather
	}
	type args struct {
		closest time.Time
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantBest utils.Pollution
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Db{
				pollution: tt.fields.pollution,
				weather:   tt.fields.weather,
			}
			gotBest, err := d.GetClosestPolution(tt.args.closest)
			if (err != nil) != tt.wantErr {
				t.Errorf("Db.GetClosestPolution() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotBest, tt.wantBest) {
				t.Errorf("Db.GetClosestPolution() = %v, want %v", gotBest, tt.wantBest)
			}
		})
	}
}

func TestDb_GetClosestWeather(t *testing.T) {
	type fields struct {
		pollution map[time.Time]utils.Pollution
		weather   map[time.Time]utils.Weather
	}
	type args struct {
		closest time.Time
	}
	tests := []struct {
		args     args
		fields   fields
		name     string
		wantBest utils.Weather
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Db{
				pollution: tt.fields.pollution,
				weather:   tt.fields.weather,
			}
			gotBest, err := d.GetClosestWeather(tt.args.closest)
			if (err != nil) != tt.wantErr {
				t.Errorf("Db.GetClosestWeather() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotBest, tt.wantBest) {
				t.Errorf("Db.GetClosestWeather() = %v, want %v", gotBest, tt.wantBest)
			}
		})
	}
}

func TestDb_AddPollution(t *testing.T) {
	type fields struct {
		pollution map[time.Time]utils.Pollution
		weather   map[time.Time]utils.Weather
	}
	type args struct {
		t time.Time
		p utils.Pollution
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Db{
				pollution: tt.fields.pollution,
				weather:   tt.fields.weather,
			}
			d.AddPollution(tt.args.p, tt.args.t)
		})
	}
}

func TestDb_AddWeather(t *testing.T) {
	type fields struct {
		pollution map[time.Time]utils.Pollution
		weather   map[time.Time]utils.Weather
	}
	type args struct {
		t time.Time
		w utils.Weather
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Db{
				pollution: tt.fields.pollution,
				weather:   tt.fields.weather,
			}
			d.AddWeather(tt.args.w, tt.args.t)
		})
	}
}
