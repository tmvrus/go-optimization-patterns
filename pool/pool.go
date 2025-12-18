package pool

import "sync"

var pool = sync.Pool{
	New: func() any {
		b := make([]int, 0, 1_000_000)
		return b
	},
}

func Good() {
	result := pool.Get().([]int)
	for i := 0; i < 1_000_000; i++ {
		result = append(result, i*i)
	}
	// Важно, мы возвращем обратно в пул "грязную" память
	pool.Put(result)
}

func Bad() {
	result := make([]int, 0, 1_000_000)
	for i := 0; i < 1_000_000; i++ {
		result = append(result, i*i)
	}
}
