package factors_test

import (
	"factors/factors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNumberFactors(t *testing.T) {
	testCases := []struct {
		name string
		in   int
		out  []int
	}{
		{
			name: "case with 0",
			in:   0,
			out:  []int{0},
		},
		{
			name: "case with 90",
			in:   90,
			out:  []int{2, 3, 3, 5},
		},
		{
			name: "case with 1024",
			in:   1024,
			out:  []int{2, 2, 2, 2, 2, 2, 2, 2, 2, 2},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := factors.NumberFactors(tc.in)
			if !assert.Equal(t, tc.out, result) {
				t.Errorf("expected: %v, got: %v\n", tc.out, result)
			}
		})
	}
}
