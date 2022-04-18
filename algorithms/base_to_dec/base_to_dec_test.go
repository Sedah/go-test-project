package base_to_dec

import (
	"fmt"
	"testing"
)

func TestBaseToDec(t *testing.T) {
	tests := []struct {
		want, base int
		have       string
	}{
		{1, 2, "1"},
		{2, 2, "10"},
		{7, 3, "21"},
		{14, 2, "1110"},
		{14, 16, "E"},
		{17, 16, "11"},
		{3735928559, 2, "11011110101011011011111011101111"},
		{3735928559, 3, "100122100210211112102"},
		{3735928559, 4, "3132223123323233"},
		{3735928559, 5, "30122344203214"},
		{3735928559, 6, "1414413525315"},
		{3735928559, 7, "161402603666"},
		{3735928559, 8, "33653337357"},
		{3735928559, 9, "10570724472"},
		{3735928559, 10, "3735928559"},
		{3735928559, 11, "164791A470"},
		{3735928559, 12, "8831A383B"},
		{3735928559, 13, "476CC321C"},
		{3735928559, 14, "276253DDD"},
		{3735928559, 15, "16CEB1BDE"},
		{3735928559, 16, "DEADBEEF"},
	}
	for _, tc := range tests {
		t.Run(fmt.Sprintf("%v in base %v", tc.have, tc.base), func(t *testing.T) {
			got := BaseToDec(tc.have, tc.base)
			if got != tc.want {
				t.Fatalf("BaseToDec() = %v; want %v", got, tc.want)
			}
		})
	}
}
