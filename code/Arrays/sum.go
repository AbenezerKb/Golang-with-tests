package code

func Sum(numbers [10]int) int {
	total := 0

	for _, nums := range numbers {
		total += nums
	}

	return total

}
