package main

import "testing"

func TestHello(T *testing.T) {
	T.Run("Should say 'Hello' ", func(t *testing.T) {
		result := Hello("David")
		expected := "Hello, David"

		if result != expected {
			T.Errorf("result:'%s', expected:'%s'", result, expected)
		}
	})

	T.Run("Should say 'Hello World' if empty string ", func(t *testing.T) {
		result := Hello("")
		expected := "Hello, World"

		if result != expected {
			T.Errorf("result:'%s', expected:'%s'", result, expected)
		}
	})
}
