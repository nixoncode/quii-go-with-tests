// Package _1_hello_world
// Created by GoLand.
// User: nixon
// Date: 22/12/2023
// Time: 11:56
package main

import (
	"fmt"
)

func main() {
	fmt.Println(Hello("world", ""))
}

const (
	spanish       = "Spanish"
	french        = "French"
	swahili       = "Swahili"
	englishPrefix = "Hello, "
	spanishPrefix = "Hola, "
	frenchPrefix  = "Bonjour, "
	swahiliPrefix = "Jambo, "
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {

	switch language {
	case spanish:
		prefix = spanishPrefix
	case french:
		prefix = frenchPrefix
	case swahili:
		prefix = swahiliPrefix
	default:
		prefix = englishPrefix

	}
	return prefix
}
