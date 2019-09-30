package multiplyfraction

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMultiplyFractions(t *testing.T) {
	tests := []struct {
		data     [2]Fraction
		expected string
	}{
		{
			[2]Fraction{
				Fraction{1, 2},
				Fraction{1, 3},
			},
			"1/6",
		},
		{
			[2]Fraction{
				Fraction{6, 10},
				Fraction{4, 10},
			},
			"6/25",
		},
		{
			[2]Fraction{
				Fraction{6, 5},
				Fraction{4, 10},
			},
			"12/25",
		},
	}

	for _, test := range tests {
		result := test.data[0].MultiplyFractions(test.data[1])

		require.Equal(t, result.String(), test.expected)
	}
}
