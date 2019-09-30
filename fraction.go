package multiplyfraction

import "fmt"

// Fraction represents ratio.
type Fraction struct {
	Numerator, Denominator int
}

// MultiplyFractions multiplies fraction by another fraction.
func (f Fraction) MultiplyFractions(m Fraction) Fraction {
	f = f.Reduce()
	m = m.Reduce()
	f.Numerator *= m.Numerator
	f.Denominator *= m.Denominator

	return f
}

// Reduce makes given fraction's denominator reduced.
func (f Fraction) Reduce() Fraction {
	// get the greatest common factor
	nums := []int{f.Numerator, f.Denominator}
	gcf := gcf(nums)
	// reduce the fraction by the greatest common factor(gcf)
	f.Denominator /= gcf
	f.Numerator /= gcf

	return f
}

func (f Fraction) String() string {
	return fmt.Sprintf("%v/%v", f.Numerator, f.Denominator)
}

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
