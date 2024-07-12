package arrayslice

func Sum(numbers []int) int {
	var sum int
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	// My method: 1
	// sumAll := make([]int, 0)
	// for _, numbersSlice := range numbersToSum {
	// 	var total int
	// 	for _, number := range numbersSlice {
	// 		total += number
	// 	}
	// 	sumAll = append(sumAll, total)
	// }

	// Method 2
	// lengthsOfNumber := len(numbersToSum)
	// sumAll := make([]int, lengthsOfNumber)

	// for i, numbers := range numbersToSum {
	// 	sumAll[i] = Sum(numbers)
	// }

	var sumAll []int
	for _, numbers := range numbersToSum {
		sumAll = append(sumAll, Sum(numbers))
	}
	return sumAll
}

func SumAllTails(numbersToSumTail ...[]int) []int {
	var sumsTail []int

	for _, numbers := range numbersToSumTail {
		if len(numbers) == 0 {
			sumsTail = append(sumsTail, 0)
		} else {
			tail := numbers[1:]
			sumsTail = append(sumsTail, Sum(tail))

		}
	}

	return sumsTail
}