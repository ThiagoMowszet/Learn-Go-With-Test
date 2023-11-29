package main

import (
	"fmt"
)

const (
	spanish = "Spanish"
	french  = "French"

	SpanishHelloPrefix = "Hola, "
	EnglishHelloPrefix = "Hello, "
	FrenchHelloPrefix  = "Bonjour, "
)

func Hello() string {
	return "Hello, world!"
}

func HelloWithParameter(name string) string {
	return "Hello, " + name
}

func HelloEnglishPrefix(name string) string {
	return EnglishHelloPrefix + name
}

func Hello2(name, language string) string {
	if name == "" {
		name = "World"
	}

	if language == spanish {
		return SpanishHelloPrefix + name
	}

	if language == "French" {
		return "Bonjour, " + name
	}
	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case "French":
		prefix = FrenchHelloPrefix
	case "Spanish":
		prefix = SpanishHelloPrefix
	default:
		prefix = EnglishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello())
}
