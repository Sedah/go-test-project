package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAreaofSegment(t *testing.T) {
	tests := []struct {
		radius, angle float64
		want          string
	}{
		{10.0, 90.0, "28.5398, 285.6192"},
		{10.0, 30.0, "1.1799, 312.9793"},
		// {20.0, 180.0, []float64{628.3174692820413, 628.3174692820413}},
	}

	for _, tc := range tests {

		t.Run(fmt.Sprintf("radius %v angle %v", tc.radius, tc.angle), func(t *testing.T) {
			got := CalArea(tc.radius, tc.angle)
			fmt.Println(reflect.TypeOf(got), reflect.TypeOf(tc.want))
			if got != tc.want {
				t.Fatalf("CalArea() = %v; want %v", got, tc.want)
			}

		})
	}

}
