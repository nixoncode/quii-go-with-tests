// Package _3_iteration
// Created by GoLand.
// User: nixon
// Date: 11/1/2024
// Time: 07:48
package _3_iteration

const repeatCount = 5

func Repeat(c string) string {
	var repeated string

	for i := 0; i < repeatCount; i++ {
		repeated = repeated + c
	}
	return repeated
}
