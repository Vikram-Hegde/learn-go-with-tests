package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "

func printHello(language, name string) string {
	greetings := map[string]struct {
		prefix, defaultName string
	}{
		"Spanish": {spanishHelloPrefix, "Amigo"},
		"English": {englishHelloPrefix, "World"},
	}

	if name == "" {
		if val, ok := greetings[language]; ok {
			name = val.defaultName
		} else {
			name = greetings["English"].defaultName
		}
	}

	if val, ok := greetings[language]; ok {
		return val.prefix + name + "!"
	}

	return englishHelloPrefix + name + "!"
}

func main() {
	message := printHello("", "")
	fmt.Println(message)
}