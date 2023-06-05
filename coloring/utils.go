package coloring

import "strconv"

// converts a given number to base k
func ToBase(num int, k int) int {
	if num == 0 {
		return 0
	}

	converted := ""
	for num > 0 {
		remainder := num % k
		converted = strconv.Itoa(remainder) + converted
		num = num / k
	}

	result, _ := strconv.Atoi(converted)
	return result
}

// converts a given number in base k, returning an array having in each position the respective digit of the result number
func ToBaseArray(num int, k int) []int {

	converted := []int{}
	if num == 0 {
		for i := 0; i < k; i++ {
			converted = append(converted, 0)
		}
	}
	for num > 0 {
		remainder := num % k
		converted = append([]int{remainder}, converted...)
		num = num / k
	}

	return converted
}

// converts a given number in base k in an array-format on 'digitsNumber' digits
func ToBaseArrayFixed(num int, k int, digitsNumber int) []int {
	if num == 0 {
		return createZeroArray(digitsNumber)
	}

	converted := []int{}
	for num > 0 {
		remainder := num % k
		converted = append([]int{remainder}, converted...)
		num = num / k
	}

	for len(converted) < digitsNumber {
		converted = append([]int{0}, converted...)
	}

	return converted
}

// creates an array of a given length filled with
func createZeroArray(length int) []int {
	result := make([]int, length)
	for i := 0; i < length; i++ {
		result[i] = 0
	}
	return result
}

// incr takes an array of integer n and an int k,
// that represent the digit of a number in base k,
// and returns the successor of n in base k
func Incr(n []int, k int) {
	for i := len(n) - 1; i >= 0; i-- {
		if n[i] < k-1 {
			n[i]++
			return
		}
		n[i] = 0 // n[i] = k
	}
}
