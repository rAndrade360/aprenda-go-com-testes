package main

import "testing"

func TestHello(t *testing.T) {

	checkResultMessage := func(result string, expected string, t *testing.T) {
		if result != expected {
			t.Errorf("resultado '%s', esperado '%s'", result, expected)
		}
	}

	t.Run("Say hello to the people", func(t *testing.T) {
		result := Hello("Renan", "")
		expected := "Hello, Renan"

		checkResultMessage(result, expected, t)
	})

	t.Run("Say hello world when a empty string is passed", func(t *testing.T) {
		result := Hello("", "")
		expected := "Hello, world"

		checkResultMessage(result, expected, t)
	})

	t.Run("Say hello in spanish", func(t *testing.T) {
		result := Hello("Renan", "spanish")
		expected := "Hola, Renan"

		checkResultMessage(result, expected, t)
	})

	t.Run("Say hello in french", func(t *testing.T) {
		result := Hello("Renan", "french")
		expected := "Bonjour, Renan"

		checkResultMessage(result, expected, t)
	})

}
