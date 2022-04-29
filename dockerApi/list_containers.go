package main

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
)

func ListContainers() {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	filterList := filters.NewArgs()
	// filterList.Add("name", "python")

	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{
		All:     true,
		Filters: filterList,
	})

	if err != nil {
		panic(err)
	}

	for _, container := range containers {
		fmt.Print(container.ID, "\t")
		fmt.Println(container.Image)
	}

}

func main() {

	ListContainers()
}
