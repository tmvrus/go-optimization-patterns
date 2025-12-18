package preallocation

import (
	"testing"
)

func BenchmarkGood(b *testing.B) {
	for b.Loop() {
		Good()
	}
}

func BenchmarkBad(b *testing.B) {
	for b.Loop() {
		Bad()
	}
}
