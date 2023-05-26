package main

import "fmt"

const prefixHelloEnglish = "Hello, "
const prefixHelloSpanish = "Hola, "
const prefixHelloFrench = "Bonjour, "

func PrefixSelector(language string)(prefix string)  {
	switch language {
		case "Spanish":
			prefix = prefixHelloSpanish
		case "French":
			prefix = prefixHelloFrench
		default:
			prefix = prefixHelloEnglish
	}
	return
}

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	return PrefixSelector(language) + name
}

func main() {
	fmt.Println(Hello("World", ""))
}
