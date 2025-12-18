package preallocation

func Bad() {
	var result []int
	for i := 0; i < 1_000_000; i++ {
		result = append(result, i*i)
	}
}

func Good() {
	result := make([]int, 0, 1_000_000)
	for i := 0; i < 1_000_000; i++ {
		result = append(result, i*i)
	}
}
