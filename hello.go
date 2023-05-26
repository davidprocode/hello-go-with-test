package main

import "fmt"

const prefixHelloEnglish = "Hello, "
const prefixHelloSpanish = "Hola, "
const prefixHelloFrench = "Bonjour, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	prefix:= prefixHelloEnglish

	switch language {
	case "Spanish":
		prefix = prefixHelloSpanish
	case "French":
	prefix = prefixHelloFrench
	}

	return prefix + name
}

func main() {
	fmt.Println(Hello("World", ""))
}
