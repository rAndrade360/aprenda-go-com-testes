package iterators

import "testing"

func TestIterator(t *testing.T) {
	result := Repeat("a")
	expected := "aaaa"

	if result != expected {
		t.Errorf("expected '%s', but got '%s'", expected, result)
	}
}

func BenckmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("b")
	}
}
