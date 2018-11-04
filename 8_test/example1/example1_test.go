// go test -v -bench . -benchmem
package example1

import (
	"testing"
)

func TestSum(t *testing.T) {
	if Sum(1, 2) != 3 {
		t.Fail()
	}
}

func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if Sum(1, 2) != 3 {
			b.Fail()
		}
	}
}
