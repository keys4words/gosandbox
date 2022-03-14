package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

func Piping() {
	c1 := exec.Command("cat", "runWithArgs.go")
	c2 := exec.Command("grep", "firefox")
	r, w := io.Pipe()
	c1.Stdout = w
	c2.Stdin = r
	var b2 bytes.Buffer
	c2.Stdout = &b2

	c1.Start()
	c2.Start()
	c1.Wait()
	w.Close()

	c2.Wait()
	io.Copy(os.Stdout, &b2)

	fmt.Println(b2)
}
func InputPiping() {
	cmd := exec.Command("cat")
	stdin, err := cmd.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		defer stdin.Close()
		io.WriteString(stdin, "write from goroutine")
	}()
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", out)
}
func OutputPiping() {
	cmd := exec.Command("echo", "piping in output pipe")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	data, err := ioutil.ReadAll(stdout)

	if err != nil {
		log.Fatal(err)
	}

	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("%s\n", string(data))
}
func main() {
	// Piping()
	// InputPiping()
	OutputPiping()
}
