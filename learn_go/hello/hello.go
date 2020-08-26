package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const spanish = "Spanish"
const french = "French"

func Hello(name string, language string) string { 
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

// private functions starts with lower case, public functions are upper case
func greetingPrefix(language string) (prefix string) { // named return value, assigned a zero value
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return
}

func main() {
	fmt.Println(Hello("world", ""))
}