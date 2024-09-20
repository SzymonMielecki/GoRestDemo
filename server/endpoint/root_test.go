package endpoint

import (
	"reflect"
	"testing"

	"github.com/SzymonMielecki/GoRestDemo/server/logic"
	"github.com/labstack/echo/v4"
)

func TestGetClosestPolution(t *testing.T) {
	type args struct {
		a logic.AppState
	}
	tests := []struct {
		want echo.HandlerFunc
		args args
		name string
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
		want echo.HandlerFunc
		args args
		name string
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

func TestAddUniversal(t *testing.T) {
	type args struct {
		a logic.AppState
	}
	tests := []struct {
		want echo.HandlerFunc
		args args
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddUniversal(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddUniversal() = %v, want %v", got, tt.want)
			}
		})
	}
}
