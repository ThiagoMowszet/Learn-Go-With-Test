package iteration

import "testing"

func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("\nRepeated = %s\nExpected = %s", repeated, expected)
	}
}

// NOTE: By default Benchmarks are run sequentially.
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}

func TestExampleRepeat(t *testing.T) {
	repeated := ExampleRepeat("a", 5)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("\nRepeated = %s\nExpected = %s", repeated, expected)
	}
}

func BenchmarkExampleRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ExampleRepeat("a", 5)
	}
}
