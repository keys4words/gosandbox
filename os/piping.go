package main

import (
	"bytes"
	"fmt"
	"io"
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

func main() {
	Piping()
}
