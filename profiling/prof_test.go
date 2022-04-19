package main

import "testing"

func TestFibZero(t *testing.T) {
	res := fib(0)
	if res != 0 {
		t.Fatal("must be 0")
	}
}

func TestFibOne(t *testing.T) {
	res := fib(1)
	if res != 1 {
		t.Fatal("must be 1")
	}
}

func TestFibThree(t *testing.T) {
	res := fib(3)
	if res != 2 {
		t.Fatal("must be 2")
	}
}
