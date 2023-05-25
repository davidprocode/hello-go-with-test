package main

import "fmt"

const prefixHelloEnglish = "Hello, "
const prefixHelloSpanish = "Hola, "
const prefixHelloFrench = "Bonjour, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	if language == "Spanish" {
		return prefixHelloSpanish + name
	}
	if language == "French" {
		return prefixHelloFrench + name
	}
	return prefixHelloEnglish + name
}

func main() {
	fmt.Println(Hello("World", ""))
}
