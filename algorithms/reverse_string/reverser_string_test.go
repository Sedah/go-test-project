package main

import (
	"fmt"
	"testing"
)

func TestReverse(t *testing.T) {
	tests := []struct {
		word string
		want string
	}{
		{"cat", "tac"},
		{"alphabet", "tebahpla"},
		{"日本語", "語本日"},
		{"Hello world!", "!dlrow olleH"},
		{"The brown 狐 jumped over the 犬", "犬 eht revo depmuj 狐 nworb ehT"},
		{"愚公山を移す", "す移を山公愚"},
		{"Ο χρόνος είναι ακριβός", "ςόβιρκα ιανίε ςονόρχ Ο"},
	}
	for _, tc := range tests {
		t.Run(fmt.Sprintf("%v", tc.word), func(t *testing.T) {
			got := Reverse(tc.word)
			if got != tc.want {
				t.Fatalf("Reverse() = %v; want %v", got, tc.want)
			}
		})
	}
}
