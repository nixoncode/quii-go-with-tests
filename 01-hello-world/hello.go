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
	fmt.Println(Hello("world"))
}

const englishPrefix = "Hello, "

func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return englishPrefix + name
}
