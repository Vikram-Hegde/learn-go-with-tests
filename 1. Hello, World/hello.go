package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const kannada = "Kannada"

const englishPrefix = "Hello, "
const spanishPrefix = "Hola, "
const frenchPrefix = "Bonjour, "
const kannadaPrefix = "Namaskara, "

func Hello(name, language string) string {
	if name == "" {
		name = "world"
	}

	prefix := englishPrefix

	switch language {
	case spanish:
		prefix = spanishPrefix
	case french:
		prefix = frenchPrefix
	case kannada:
		prefix = kannadaPrefix
	}

	return prefix + name
}

func main() {
	fmt.Println(Hello("Vikram", ""))
}