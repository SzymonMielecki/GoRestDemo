package logic

import (
	"reflect"
	"testing"
	"time"

	"github.com/SzymonMielecki/air_qual/server/persistance"
)

func TestNewAppState(t *testing.T) {
	tests := []struct {
		want *AppState
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAppState(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAppState() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAppState_AddPollution(t *testing.T) {
	type fields struct {
		db *persistance.Db
	}
	type args struct {
		t time.Time
		p persistance.Pollution
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
			a := &AppState{
				db: tt.fields.db,
			}
			a.AddPollution(tt.args.t, tt.args.p)
		})
	}
}

func TestAppState_AddWeather(t *testing.T) {
	type fields struct {
		db *persistance.Db
	}
	type args struct {
		t time.Time
		w persistance.Weather
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
			a := &AppState{
				db: tt.fields.db,
			}
			a.AddWeather(tt.args.t, tt.args.w)
		})
	}
}

func TestAppState_GetClosestPolution(t *testing.T) {
	type fields struct {
		db *persistance.Db
	}
	type args struct {
		t time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    persistance.Pollution
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &AppState{
				db: tt.fields.db,
			}
			got, err := a.GetClosestPolution(tt.args.t)
			if (err != nil) != tt.wantErr {
				t.Errorf("AppState.GetClosestPolution() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AppState.GetClosestPolution() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAppState_GetClosestWeather(t *testing.T) {
	type fields struct {
		db *persistance.Db
	}
	type args struct {
		t time.Time
	}
	tests := []struct {
		fields  fields
		args    args
		name    string
		want    persistance.Weather
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &AppState{
				db: tt.fields.db,
			}
			got, err := a.GetClosestWeather(tt.args.t)
			if (err != nil) != tt.wantErr {
				t.Errorf("AppState.GetClosestWeather() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AppState.GetClosestWeather() = %v, want %v", got, tt.want)
			}
		})
	}
}
