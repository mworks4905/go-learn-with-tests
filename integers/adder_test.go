package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	total := 4

	if sum != total {
		t.Errorf("expected %d, but got %d", total, sum)
	}
}

func ExampleAdd() {
	sum := Add(2, 3)
	fmt.Println(sum)
	//Output: 5
}