package main

import "testing"

func TestHello(t *testing.T) {
	checkString := func(t *testing.T, expected string, result string) {
		t.Helper()
		if result != expected {
			t.Errorf("result:'%s', expected:'%s'", result, expected)
		}
	}

	t.Run("Should say 'Hello' ", func(t *testing.T) {
		result := Hello("David", "")
		expected := "Hello, David"

		checkString(t, result, expected)
	})

	t.Run("Should say 'Hello World' if empty string ", func(t *testing.T) {
		result := Hello("", "")
		expected := "Hello, World"

		checkString(t, expected, result)
	})

	t.Run("Should support spanish language", func(t *testing.T) {
		result := Hello("David", "Spanish")
		expected := "Hola, David"

		checkString(t, expected, result)
	})

	t.Run("Should support french language", func(t *testing.T) {
		result := Hello("David", "French")
		expected := "Bonjour, David"

		checkString(t, expected, result)
	})
}
