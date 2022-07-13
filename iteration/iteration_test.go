package iteration

import (
	"fmt"
	"testing"
)

func TestIterator(t *testing.T){
	iterated := Iterate("a",5)
	expected := "aaaaa"

	if iterated != expected {
		t.Errorf("expected %q but got %q", expected, iterated)
	}
}

func BenchmarkIterator(b *testing.B)  {
	for i := 0; i < b.N; i++ {
		Iterate("a",1)
	}
}

func ExampleIterate() {
	str := Iterate("a",5)
	fmt.Println(str)
	//Output: aaaaa
}