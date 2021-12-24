package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSearchAnagrams(t *testing.T) {
	tests := map[string]struct {
		words []string
		want  map[string][]string
	}{
		"general": {
			words: []string{"Пятак", "тяпка", "пятка", "листок", "тяпка", "Слиток", "столик", "приказ", "каприз", "одно"},
			want:  map[string][]string{"листок": {"слиток", "столик"}, "приказ": {"каприз"}, "пятак": {"тяпка", "пятка"}},
		},
	}

	for name, testCase := range tests {
		t.Run(name, func(t *testing.T) {
			req := require.New(t)
			res := SearchAnagrams(testCase.words)
			req.Equal(testCase.want, res)
		})
	}
}
