package positiveSum

func PositiveSum(numbers []int) int {
	sum := 0

	// for i := 0; i < len(numbers); i++ {
	// 	if numbers[i] > 0 {
	// 		sum += numbers[i]
	// 	}
	// }

	for _, value := range numbers {
		if value > 0 {
			sum += value
		}
	}
	return sum
}
