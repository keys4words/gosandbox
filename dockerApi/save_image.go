package main

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/docker/docker/client"
)

func SaveImage(imageid string, filepath string) {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	images := make([]string, 1)
	images[0] = imageid
	reader, err := cli.ImageSave(ctx, images)
	if err != nil {
		panic(err)
	}
	defer reader.Close()

	file, err := os.Create(filepath)
	defer file.Close()

	writtenBytes, err := io.Copy(file, reader)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Bytes written %vn", writtenBytes)

}

func main() {

	SaveImage("0ac33e5f5afa", "./image.tar")
}
