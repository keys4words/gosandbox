package main

import (
	"log"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
	"unsafe"
)

var letterRunes = []rune("qwertyuiopasdfghjklzxcvbnmQAZWSXEDCRFVTGBYHNUJMIKOLP")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func Clone(s string) string {
	b := make([]byte, len(s))
	copy(b, s)
	return *(*string)(unsafe.Pointer(&b))
}

func main() {
	rand.Seed(time.Now().UnixNano())
	dateHolder := make([]string, 0)
	filter := RandStringRunes(104567)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGHUP)
	run := true

	go func() {
		oscall := <-c
		log.Printf("system call:%+v", oscall)
		run = false
		os.Exit(1)
	}()

	for {
		if !run {
			break
		}
		dateHolder = append(dateHolder, Clone(filter))
		time.Sleep(1 * time.Second)
	}
}
