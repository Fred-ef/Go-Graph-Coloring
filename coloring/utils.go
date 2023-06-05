package coloring

import (
	"errors"
	"math"
)

// converts a given number in base k in array-format,
// returning an array having in each position the respective
// digit of the result number
func ToBaseArray(num int, k int) ([]int, error) {
	if k == 0 {
		return nil, errors.New("base cannot be 0")
	}

	number := int(math.Abs(float64(num)))

	converted := []int{}
	if number == 0 {
		for i := 0; i < k; i++ {
			converted = append(converted, 0)
		}
	}
	for number > 0 {
		remainder := number % k
		converted = append([]int{remainder}, converted...)
		number = number / k
	}

	return converted, nil
}

// converts a given number in base k in an array-format
// on a fixed number of digits
func ToBaseArrayFixed(num int, k int, digitsNumber int) ([]int, error) {
	if k == 0 {
		return nil, errors.New("base cannot be 0")
	}

	number := int(math.Abs(float64(num)))

	if number == 0 {
		return createZeroArray(digitsNumber), nil
	}

	converted := []int{}
	for number > 0 {
		remainder := number % k
		converted = append([]int{remainder}, converted...)
		number = number / k
	}

	for len(converted) < digitsNumber {
		converted = append([]int{0}, converted...)
	}

	return converted, nil
}

// creates an array of a given length filled with 0s
func createZeroArray(length int) []int {
	result := make([]int, length)
	for i := 0; i < length; i++ {
		result[i] = 0
	}
	return result
}
