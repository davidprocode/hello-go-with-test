package main

import "fmt"

const prefixHelloEnglish = "Hello, "
const prefixHelloSpanish = "Hola, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	if language == "Spanish" {
		return prefixHelloSpanish + name
	}
	return prefixHelloEnglish + name
}

func main() {
	fmt.Println(Hello("World", ""))
}
