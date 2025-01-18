package arraysandslices

func Sum(numbers [5]int) (sum int) {
	for _, number := range numbers {
		sum += number
	}

	return
}

func SumSlice(numbers []int) (sum int) {
	for _, number := range numbers {
		sum += number
	}
	return
}

func SumAll(slices ...[]int) []int {
	var sums []int
	for _, numbers := range slices {
		sums = append(sums, SumSlice(numbers))
	}

	return sums
}
