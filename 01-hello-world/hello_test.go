// Package main
// Created by GoLand.
// User: nixon
// Date: 22/12/2023
// Time: 12:00
package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello()
	want := "Hello world!"

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
