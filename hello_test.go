package main

import "testing"

func TestHello(T *testing.T) {
	result := Hello("David")
	expected := "Hello, David"

	if result != expected {
		T.Errorf("result:'%s', expected:'%s'", result, expected)
	}
}
