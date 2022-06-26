package main

import "fmt"

// declare Number interface type to use as a type constraint
type Number interface {
	int64 | float64
}

func main() {
	// Init a map for int values
	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}

	// Init a map for float values
	floats := map[string]float64{
		"first":  35.98,
		"second": 26.99,
	}

	fmt.Printf("Non-Generic Sums: %v and %v\n", SumInts(ints), SumFloats(floats))
	fmt.Printf("Generic Sums: %v and %v\n", SumIntsOrFloats(ints), SumIntsOrFloats(floats))
	fmt.Printf("Generic Sums with Constraint: %v and %v\n", SumNumbers(ints), SumNumbers(floats))
}

// adds together the value of m.
func SumInts(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}

// adds together the value of m.
func SumFloats(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}
	return s
}

// sum the values of map . supporting both int64 and float64 as types of the values
func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

func SumNumbers[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}

	return s
}
