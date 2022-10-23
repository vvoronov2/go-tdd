package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("5 times a", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"

		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	})
	t.Run("7 times b", func(t *testing.T) {
		repeated := Repeat("b", 7)
		expected := "bbbbbbb"

		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	})
}

func TestStringRepeat(t *testing.T) {
	t.Run("Repeat `test` 10 times", func(t *testing.T) {
		repeated := StringRepeat("test", 10)
		expected := "testtesttesttesttesttesttesttesttesttest"

		if StringEqual(repeated, expected) != true {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 25)
	}
}

func ExampleRepeat() {
	repeated := Repeat("x", 8)
	fmt.Println(repeated)
	// Output: xxxxxxxx
}
