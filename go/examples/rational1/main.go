// rational1.go

//
// Rational numbers.
//

package main

import (
	"errors"
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"time"
)

//////////////////////////////////////////////////////

// An interface is a list of function headers.
type Rationalizer interface {
	// Makes a rational a/b, where the denominator is guaranteed to be
	// positive. If b is not 0, then the a nil error value is returned. If b
	// is non-zero, then a nil error value is returned.
	// func MakeRational(a, b int) (Rational, error)

	// Returns a random rational number, with the denominator guaranteed to be
	// non-zero. The numerator and denominator range from at least 0 to 1000.
	RandRational() Rational


	// Converts an int to an equivalent Rational.
	MakeRationalInt(n int) Rational

	// Converts a rational number to the equivalent float64.
	toFloat64(r Rational) float64

	// Returns the numerator of a Rational.
	Numerator(r Rational) int

	// Returns the denominator of a Rational.
	Denominator(r Rational) int

	// Returns the numerator, denominator.
	Split(r Rational) (int, int)

	// Returns true if s and r equal Rationals, and false otherwise.
	Equal(r Rational, s Rational) bool

	// Returns true if r is strictly less than s. *Don't* do this by, say,
	// using toFloat64, since that could have round-off errors.
	LessThan(r Rational, s Rational) bool

	// Returns true if r is less than, or equal to, s. *Don't* do this by,
	// say, using toFloat64, since that could have round-off errors.
	LessThanEqual(r Rational, s Rational) bool

	// Returns the reciprocal of r. If r's numerator is zero, then a non-nil
	// error is returned. If the numerator is not 0, then the error value is
	// nil.
	Invert(r Rational) (Rational, error)

	// Returns a new Rational that is the sum of r and s.
	Add(r Rational, s Rational) Rational

	// Returns a new Rational that is the product of r and s.
	Multiply(r Rational, s Rational) Rational

	// Returns a new Rational that is the quotient of r and s. If the
	// rationals can't be divided due to division by 0, then return a non-nil
	// error value. If the division is successful, the error should be nil.
	Divide(r Rational, s Rational) (Rational, error)

	// Returns a new Rational that is in lowest terms and is equal to r.
	Reduce(r Rational) Rational
} // Rationalizer interface

//////////////////////////////////////////////////////

type Rational struct {
	a, b int
}

// Makes a rational a/b, where the denominator is guaranteed to be positive.
func MakeRational(a, b int) (Rational, error) {
	if b == 0 {
		return Rational{0, 0}, errors.New("MakeRational: 0 denominator")
		// } else if a < 0 && b < 0 {
		//     return Rational{-a, -b}, nil
	} else if b < 0 {
		return Rational{-a, -b}, nil
	} else {
		return Rational{a, b}, nil
	}
}

func RandRational() Rational {
	return Rational{rand.Intn(1000), 1 + rand.Intn(1000)}
}

func RandRationalList(n int) []Rational {
	result := []Rational{}
	for i := 0; i < n; i++ {
		result = append(result, RandRational())
	}
	return result
}

func Normalize(r Rational) Rational {
	if r.b < 0 {
		return Rational{-r.a, -r.b}
	} else {
		return Rational{r.a, r.b}
	}
}

func toFloat64(r Rational) float64 {
	return float64(r.a) / float64(r.b)
}

// helper function to convert an int to a Rational
func MakeRationalInt(n int) Rational {
	result, _ := MakeRational(n, 1)
	return result
}

func Numerator(r Rational) int {
	return r.a
}

func Denominator(r Rational) int {
	return r.b
}

func Good(r Rational) bool {
	return Denominator(r) != 0
}

// returns the numerator and the denominator
func Split(r Rational) (int, int) {
	return r.a, r.b
}

// r.a/r.b = s.a/s.b --> r.a * s.b = s.a * r.b
func Equal(r Rational, s Rational) bool {
	return r.a*s.b == r.b*s.a
}

func LessThan(r Rational, s Rational) bool {
	return r.a*s.b < r.b*s.a
}

func LessThanEqual(r Rational, s Rational) bool {
	return Equal(r, s) || LessThan(r, s)
}

func GreaterThan(r Rational, s Rational) bool {
	return !LessThanEqual(r, s)
}

func Invert(r Rational) (Rational, error) {
	if Denominator(r) == 0 {
		return Rational{0, 0}, errors.New("invert: zero denominator")
	} else {
		return Rational{Denominator(r), Numerator(r)}, nil
	}
}

// \frac{a}{b} + \frac{c}{d} = \frac{ad + bc}{bd}
func Add(r Rational, s Rational) Rational {
	a, b := Split(r)
	c, d := Split(s)
	return Rational{a*d + b*c, b * d}
}

// \frac{a}{b} \cdot \frac{c}{d} = \frac{ac}{bd}
func Multiply(r Rational, s Rational) Rational {
	a, b := Split(r)
	c, d := Split(s)
	return Rational{a * c, b * d}
}

// \frac{\frac{a}{b}}{\frac{c}{d}} = \frac{ad}{bc}
// b, c, d non-zero
func Divide(r Rational, s Rational) (Rational, error) {
	a, b := Split(r)
	c, d := Split(s)
	if b == 0 || c == 0 || d == 0 {
		return Rational{0, 0}, errors.New("divide: zero denominator")
	} else {
		return Rational{a * d, b * c}, nil
	}
}

// greatest common divisor
func Gcd(a, b int) int {
	for b != 0 {
		temp := b
		b = a % b
		a = temp
	}
	return a
}

func Reduce(r Rational) Rational {
	a, b := Split(r)
	d := Gcd(a, b)
	if d == 1 || d == -1 {
		return r
	} else {
		return Rational{a / d, b / d}
	}
}

// implements the Stringer interface
func (r Rational) String() string {
	return strconv.Itoa(r.a) + "/" + strconv.Itoa(r.b)
}

func insertionSort(lst []Rational) {
	n := len(lst)
	if n < 2 {
		return
	}
	for i := 1; i < n; i++ {
		for j := i; j > 0 && GreaterThan(lst[j-1], lst[j]); j-- {
			lst[j], lst[j-1] = lst[j-1], lst[j]
		}
	}
}

func isSorted(lst []int) bool {
	for i := 1; i < len(lst); i++ {
		if lst[i-1] > lst[i] {
			return false
		}
	}
	return true
}

//////////////////////////////////////////////////////

// returns 1/1 + 1/2 + ... + 1/n
func harmonic(n int) float64 {
	total := 0.0
	for i := 1; i <= n; i++ {
		total += 1 / float64(i)
	}
	return total
}

//////////////////////////////////////////////////////
//
// Sorting
//
//////////////////////////////////////////////////////

type BySize []Rational

func (a BySize) Len() int           { return len(a) }
func (a BySize) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a BySize) Less(i, j int) bool { return LessThan(a[i], a[j]) }

//////////////////////////////////////////////////////

func main() {
	r, _ := MakeRational(2, 3)
	fmt.Printf("%v\n", r)
	fmt.Printf("%v\n", Add(r, r))
	fmt.Println()
	for i := -5; i <= 5; i++ {
		r := MakeRationalInt(i)
		r, err := Invert(r)
		if err != nil {
			fmt.Printf("can't invert %v: %v\n", r, err)
		} else {
			fmt.Printf("%v\n", r)
		}
	}

	fmt.Println()
	N := 25
	total, _ := MakeRational(0, 1)
	for i := 1; i <= N; i++ {
		r := MakeRationalInt(i)
		r, err := Invert(r)
		if err != nil {
			fmt.Printf("can't invert %v: %v\n", r, err)
		} else {
			// fmt.Printf("%v\n", r)
			total = Reduce(Add(total, r))

		}
	}
	fmt.Printf("total: %v\n", total)
	fmt.Printf("total: %v\n", Reduce(total))
	fmt.Printf("harmonic(%v): %v\n", N, harmonic(N))

	//
	// test sorting
	//
	fmt.Println()
	nums := []Rational{
		Rational{2, 3}, Rational{-2, 3}, Rational{2, 3},
		Rational{20, 3}, Rational{0, 3}, MakeRationalInt(5),
	}
	nums2 := make([]Rational, len(nums))
	copy(nums2, nums)
	fmt.Println(nums)
	sort.Sort(BySize(nums))
	fmt.Println(nums, "sort.Sort")

	fmt.Println()
	fmt.Println(nums2)
	insertionSort(nums2)
	fmt.Println(nums2, "insertionSort")

	//
	// test random rationals
	//
	rand.Seed(time.Now().UnixNano())
	lst := RandRationalList(10)
	fmt.Println(lst)
	insertionSort(lst)
	fmt.Println(lst)

	for _, r := range lst {
		fmt.Printf("%v ", toFloat64(r))
	}
	fmt.Println()
} // main
