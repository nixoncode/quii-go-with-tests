// Package _3_iteration
// Created by GoLand.
// User: nixon
// Date: 11/1/2024
// Time: 07:44
package _3_iteration

import "testing"

func TestRepeat(t *testing.T) {
	expect := "aaaa"
	got := Repeat("a")

	if expect != got {
		t.Errorf("expected %q but got %q", expect, got)
	}
}
