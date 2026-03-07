package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	got := Repeat("a", 5)
	want := "aaaaa"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for range b.N {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	output := Repeat("a", 3)
	fmt.Println(output)
	// Output: aaa
}

func TestReverse(t *testing.T) {
	got := Reverse("Cat")
	want := "taC"

	if got != want {
		t.Errorf("Got %q, want %q", got, want)
	}
}

func ExampleReverse() {
	output := Reverse("Example")
	fmt.Println(output)
	// Output: elpmaxE
}