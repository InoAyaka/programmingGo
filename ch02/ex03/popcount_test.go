package popcount

import (
	"fmt"
	"testing"
)

func BenchmarkPopcounts(t *testing.B) {
	fmt.Println(Popcounts(12 + 16))
}

func BenchmarkPopcountsLoop(t *testing.B) {
	fmt.Println(PopcountsLoop(12 + 16))
}

func BenchmarkPopcountsShift(t *testing.B) {
	fmt.Println(PopcountsShift(12 + 16))
}
