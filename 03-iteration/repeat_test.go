// Package _3_iteration
// Created by GoLand.
// User: nixon
// Date: 11/1/2024
// Time: 07:44
package _3_iteration

import "testing"

func TestRepeat(t *testing.T) {
	expect := "aaaaa"
	got := Repeat("a", 5)

	if expect != got {
		t.Errorf("expected %q but got %q", expect, got)
	}
}

func BenchmarkRepeat(b *testing.B) {
	{
		for i := 0; i < b.N; i++ {
			Repeat("a", 5)
		}
	}
}
