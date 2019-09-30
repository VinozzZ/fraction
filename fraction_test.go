package multiplyfraction

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type expected struct {
	result   string
	hasError bool
}

func TestMultiplyFractions(t *testing.T) {
	tests := []struct {
		data     [2]Fraction
		expected expected
	}{
		{
			[2]Fraction{
				Fraction{1, 2},
				Fraction{1, 3},
			},
			expected{
				result: "1/6",
			},
		},
		{
			[2]Fraction{
				Fraction{6, 10},
				Fraction{4, 10},
			},
			expected{
				result: "6/25",
			},
		},
		{
			[2]Fraction{
				Fraction{6, 5},
				Fraction{4, 10},
			},
			expected{
				result: "12/25",
			},
		},
		{
			[2]Fraction{
				Fraction{100, 7},
				Fraction{20, 3},
			},
			expected{
				result: "2000/21",
			},
		},
		{
			[2]Fraction{
				Fraction{10, 20},
				Fraction{4, 10},
			},
			expected{
				result: "1/5",
			},
		},
		{
			[2]Fraction{
				Fraction{50, 0},
				Fraction{1, 10},
			},
			expected{
				result:   "",
				hasError: true,
			},
		},
		{
			[2]Fraction{
				Fraction{0, 5},
				Fraction{1, 10},
			},
			expected{
				result: "0",
			},
		},
		{
			[2]Fraction{
				Fraction{100, 5},
				Fraction{1, 10},
			},
			expected{
				result: "2",
			},
		},
		{
			[2]Fraction{
				Fraction{-1, 5},
				Fraction{-2, 10},
			},
			expected{
				result: "1/25",
			},
		},
		{
			[2]Fraction{
				Fraction{-3, 7},
				Fraction{7, 3},
			},
			expected{
				result: "-1",
			},
		},
	}

	for _, test := range tests {
		result, err := test.data[0].MultiplyFractions(&test.data[1])
		if test.expected.hasError {
			require.Error(t, err)
		} else {
			require.NoError(t, err)
		}

		require.Equal(t, test.expected.result, result.String())
	}
}

func BenchmarkMultiplyFractions(b *testing.B) {
	for n := 0; n < b.N; n++ {
		f := &Fraction{2, 6}
		m := &Fraction{10, 4}
		f.MultiplyFractions(m)
	}
}
