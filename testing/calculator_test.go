package main

import (
	"testing"
	"log"
)

func TestAdd(t *testing.T) {
	if res := Add(10, 20); res != 30 {
		log.Fatalf("Add(10, 20) fail. Expected %d, get %d\n", 30, res)
	}
}