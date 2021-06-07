package iterations

import (
	"fmt"
	"strings"
	"testing"
)

func TestRepeat(t *testing.T) {
	repetitionsNumber := 6
	testChar := "a"
	repetitions := Repeat(testChar, repetitionsNumber)
	expected := strings.Repeat(testChar, repetitionsNumber)

	if repetitions != expected {
		t.Errorf("Expected: %s, but received: %s", expected, repetitions)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	repetitionsNumber := Repeat("A", 7)
	fmt.Println(repetitionsNumber)
	// Output: AAAAAAA
}
