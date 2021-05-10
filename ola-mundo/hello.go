package main

import "fmt"

const englishPrefix = "Hello, "
const spanishPrefix = "Hola, "
const frenchPrefix = "Bonjour, "
const spanish = "spanish"
const french = "french"

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}

	prefix := switchPrefix(language)

	return prefix + name
}

func switchPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishPrefix
	case french:
		prefix = frenchPrefix
	default:
		prefix = englishPrefix
	}

	return
}

func main() {
	fmt.Println(Hello("Renan", ""))
}
