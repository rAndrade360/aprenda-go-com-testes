package main

import "fmt"

const englishPrefix = "Hello, "
const spanishPrefix = "Hola, "
const frenchPrefix = "Bonjour, "
const spanish = "spanish"
const french = "french"

func Hello(name string, language string) string {
	var prefix string

	if name == "" {
		name = "world"
	}

	switch language {
	case spanish:
		prefix = spanishPrefix
	case french:
		prefix = frenchPrefix
	default:
		prefix = englishPrefix
	}

	return prefix + name
}

func main() {
	fmt.Println(Hello("Renan", ""))
}
