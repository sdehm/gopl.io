package main

import (
	"os/exec"
	"testing"
)

func BenchmarkEcho1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		exec.Command("go", "run", "../echo1.go", args(i)).Run()
	}
}

func BenchmarkEcho2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		exec.Command("go", "run", "../echo2.go", args(i)).Run()
	}
}

func BenchmarkEcho3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		exec.Command("go", "run", "../echo2.go", args(i)).Run()
	}
}

// Takes an integer num and returns a string of "test" repeated 1000^num times.
func args(num int) string {
	// build string of "test" repeated 1000^num times
	s := ""
	for i := 0; i < (1000 ^ num); i++ {
		s += "test "
	}
	return s
}
