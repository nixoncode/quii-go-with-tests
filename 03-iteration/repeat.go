// Package _3_iteration
// Created by GoLand.
// User: nixon
// Date: 11/1/2024
// Time: 07:48
package _3_iteration

func Repeat(c string, count int) string {
	var repeated string

	for i := 0; i < count; i++ {
		repeated = repeated + c
	}
	return repeated
}
