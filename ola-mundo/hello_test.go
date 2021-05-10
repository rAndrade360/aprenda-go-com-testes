package main

import "testing"

func TestHello(t *testing.T) {

	t.Run("Say hello to the people", func(t *testing.T) {
		result := Hello("Renan")
		expected := "Hello, Renan"

		if result != expected {
			t.Errorf("resultado '%s', esperado '%s'", result, expected)
		}
	})

	t.Run("Say hello world when a empty string is passed", func(t *testing.T) {
		result := Hello("")
		expected := "Hello, world"

		if result != expected {
			t.Errorf("resultado '%s', esperado '%s'", result, expected)
		}
	})

}
