package main

import "testing"

func TestHello(T *testing.T) {
	result := Hello()
	expected := "Hello, World"

	if result != expected {
		T.Errorf("result:'%s', expected:'%s'",result, expected)
	}
}
