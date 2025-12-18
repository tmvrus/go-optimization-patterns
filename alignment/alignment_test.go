// go test -bench=. -benchmem
package alignment

import "testing"

func BenchmarkGood(b *testing.B) {
	for b.Loop() {
		result := make([]CircleAligned, 1_000_000)
		for _, item := range result {
			item.X = 1
		}

	}
}

func BenchmarkBad(b *testing.B) {
	for b.Loop() {
		result := make([]Circle, 1_000_000)
		for _, item := range result {
			item.X = 1
		}
	}
}
