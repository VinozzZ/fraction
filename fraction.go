package multiplyfraction

import (
	"errors"
	"fmt"
)

// Fraction represents ratio.
type Fraction struct {
	Numerator, Denominator int
}

// MultiplyFractions multiplies fraction by another fraction.
func (f *Fraction) MultiplyFractions(m *Fraction) (*Fraction, error) {
	if f.Denominator == 0 {
		return nil, errors.New("invalid argument: denominator should not be zero")
	}

	// get fraction's absolute value
	f, isFNegative := abs(f)
	m, isMNegative := abs(m)

	f = f.reduce()
	m = m.reduce()
	f.Numerator *= m.Numerator
	f.Denominator *= m.Denominator

	if f.Numerator == 0 {
		return &Fraction{0, 0}, nil
	}

	// reduce the result
	f = f.reduce()

	if isFNegative != isMNegative {
		f.Numerator = -f.Numerator
	}

	return f, nil
}

func (f *Fraction) String() string {
	if f == nil {
		return ""
	}
	if f.Numerator == 0 {
		return "0"
	}
	if f.Denominator == 1 {
		return fmt.Sprintf("%v", f.Numerator)
	}

	return fmt.Sprintf("%v/%v", f.Numerator, f.Denominator)
}

// Reduce makes given fraction's denominator reduced.
func (f *Fraction) reduce() *Fraction {
	// get the greatest common factor
	nums := []int{f.Numerator, f.Denominator}
	gcf := gcf(nums)
	if gcf == 0 {
		return f
	}
	// reduce the fraction by the greatest common factor(gcf)
	f.Denominator /= gcf
	f.Numerator /= gcf

	return f
}

// gcf returns the greatest common factor of an array of numbers.
func gcf(nums []int) int {
	// find all factors of both numbers
	var commons []int
	uniques := make(map[int]bool)

	for _, num := range nums {
		for i := 1; i <= num; i++ {
			if num%i == 0 {
				if uniques[i] {
					// find the ones that are common to both
					commons = append(commons, i)
				}
				uniques[i] = true
			}
		}
	}

	// choose the greatest
	var greatest int
	for _, common := range commons {
		if common > greatest {
			greatest = common
		}
	}

	return greatest
}

// abs returns the absolute value of x.
func abs(x *Fraction) (*Fraction, bool) {
	var isNegative bool
	if x.Numerator < 0 {
		x.Numerator = -x.Numerator
		isNegative = true
	}
	if x.Denominator < 0 {
		x.Denominator = -x.Denominator
		if isNegative {
			isNegative = false
		}
	}
	return x, isNegative
}
