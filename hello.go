package main

import "fmt"

const prefixHello = "Hello, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	if language == "Spanish" {
		return "Hola, " + name
	}
	return prefixHello + name
}

func main() {
	fmt.Println(Hello("World", ""))
}
