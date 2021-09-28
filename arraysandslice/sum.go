package arraysandslice

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	length := len(numbersToSum)

	result := make([]int, length)

	for i, numbers := range numbersToSum {
		result[i] = Sum(numbers)
	}

	return result
}

func SumTails(numbersToSum ...[]int) []int {
	length := len(numbersToSum)
	result := make([]int, length)

	for i, numbers := range numbersToSum {
		if len(numbers) == 0 {
			result[i] = 0
			continue
		}
		result[i] = Sum(numbers[1:])
	}

	return result
}
