package endpoint

import (
	"reflect"
	"testing"

	"github.com/SzymonMielecki/air_qual/server/logic"
	"github.com/labstack/echo/v4"
)

func TestGetClosestPolution(t *testing.T) {
	type args struct {
		a logic.AppState
	}
	tests := []struct {
		name string
		args args
		want echo.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetClosestPolution(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetClosestPolution() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetClosestWeather(t *testing.T) {
	type args struct {
		a logic.AppState
	}
	tests := []struct {
		name string
		args args
		want echo.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetClosestWeather(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetClosestWeather() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddPollution(t *testing.T) {
	type args struct {
		a logic.AppState
	}
	tests := []struct {
		name string
		args args
		want echo.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddPollution(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddPollution() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddWeather(t *testing.T) {
	type args struct {
		a logic.AppState
	}
	tests := []struct {
		name string
		args args
		want echo.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddWeather(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddWeather() = %v, want %v", got, tt.want)
			}
		})
	}
}
