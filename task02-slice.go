package homework

func reverse(input []int64) (result []int64) {
	for _, num := range input {
		result = append([]int64{num}, result...)
	}

	return
}
