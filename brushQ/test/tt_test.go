package test

import (
	"fmt"
	"testing"
)

func Test_division(t *testing.T) {
	type args struct {
		a float64
		b float64
	}

	t.Run("a", func(t *testing.T) {
		got, err := division(6, 3)
		fmt.Println(got)
		if err != nil {
			t.Errorf("division() error = %v, wantErr %v", err, nil)
			return
		}
		if got != 2 {
			t.Errorf("division() = %v, want %v", got, 2)
		}
	})
}

func Benchmark_division(b *testing.B) {
	for i := 0; i < b.N; i++ {
		division(4, 6)
	}
}
