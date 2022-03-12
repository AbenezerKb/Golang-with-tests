package code

func Sum(numbers []int) int {
	total := 0

	for _, nums := range numbers {
		total += nums
	}

	return total

}
