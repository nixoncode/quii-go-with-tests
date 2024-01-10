// Package _2_integers
// Created by GoLand.
// User: nixon
// Date: 10/1/2024
// Time: 07:01
package _2_integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	got := Add(2, 2)
	expect := 4

	if got != expect {
		t.Errorf("expected '%d' but got '%d'", expect, got)
	}
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
