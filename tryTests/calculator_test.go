package main

import (
	"testing"
	_ "log"
)

func TestAdd(t *testing.T) {
	if res := Add(10, 20); res != 30 {
		t.Errorf("Add(10, 20) fail. Expected %d, get %d\n", 30, res)
	}
	if res := Add(30, 30); res != 60 {
		t.Errorf("Add(30, 30) fail. Expected %d, get %d\n", 60, res)
	}
}

func TestSub(t *testing.T) {
	if res := Sub(30, 20); res != 10 {
		t.Errorf("Sub(30,20) fail. Expected %d, get %d\n", 10, res)
	}
}