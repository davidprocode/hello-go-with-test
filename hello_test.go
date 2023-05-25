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
}
