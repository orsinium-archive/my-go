package main

// Sum return sum of passed numbers
func Sum(numbers []int) (sum int) {
	for _, number := range numbers {
		sum += number
	}
	return
}

// SumAll return sums of all passed arrays
func SumAll(numbersArrays ...[]int) []int {
	// prepare array with results
	sums := make([]int, len(numbersArrays))
	// count sum
	for i, numbers := range numbersArrays {
		sums[i] = Sum(numbers)
	}
	return sums
}

// SumAllTails return sum of all elements without first element.
func SumAllTails(numbersArrays ...[]int) (sums []int) {
	for _, numbers := range numbersArrays {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return
}
