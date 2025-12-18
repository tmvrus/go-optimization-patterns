package zerocopy

//go test -bench=. -benchmem
import (
	"testing"
)

func BenchmarkGood(b *testing.B) {
	for b.Loop() {
		r := newGoodReader()
		dst := make([]byte, 100)
		for {
			err := r.Read(dst)
			if err != nil {
				break
			}
		}

	}
}

func BenchmarkBad(b *testing.B) {
	for b.Loop() {
		r := newBadReader()
		for {
			data := r.Read()
			if data == nil {
				break
			}
			_ = data
		}
	}
}
