package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func read(filepath string) {
	f, err := os.Open(filepath)
	check(err)

	b1 := make([]byte, 9)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1))

	b2 := make([]byte, 9)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes: %s\n", n2, string(b2[:n2]))

	o3, err := f.Seek(5, 1)
	check(err)
	b3 := make([]byte, 9)
	n3, err := io.ReadAtLeast(f, b3, 2)
	check(err)
	fmt.Printf("%d bytes @ %d %s\n", n3, o3, string(b3))

}

func write(filepath string) {
	f, err := os.Create(filepath)
	check(err)

	defer f.Close()

	o2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(o2)
	check(err)
	fmt.Printf("wrote %d bytes\n", n2)

	o3, err := f.WriteString("new line\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", o3)

	_ = f.Sync()

	w := bufio.NewWriter(f)
	o4, err := w.WriteString("buffered string\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", o4)
	w.Flush()
}

func main() {
	// read("test.txt")
	write("output.txt")
}
