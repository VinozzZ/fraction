package fraction_test

import (
	"testing"

	"github.com/vinozzZ/multiplyFraction/multiplyfraction"

	"github.com/stretchr/testify/require"
)

func TestMultiplyFractions(t *testing.T) {
	tests := []struct {
		data     [2]multiplyfraction.Fraction
		expected string
	}{
		{
			[2]multiplyfraction.Fraction{
				multiplyfraction.Fraction{1, 2},
				multiplyfraction.Fraction{1, 3},
			},
			"1/6",
		},
		{
			[2]multiplyfraction.Fraction{
				multiplyfraction.Fraction{6, 10},
				multiplyfraction.Fraction{4, 10},
			},
			"6/25",
		},
		{
			[2]multiplyfraction.Fraction{
				multiplyfraction.Fraction{6, 5},
				multiplyfraction.Fraction{4, 10},
			},
			"12/25",
		},
	}

	for _, test := range tests {
		result := test.data[0].MultiplyFractions(test.data[1])

		require.Equal(result.String(), test.expected)
	}
}
