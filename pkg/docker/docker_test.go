package docker

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
	"testing"
)

func TestDockerCli(t *testing.T) {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := cli.Close(); err != nil {
			return
		}
	}()

	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{
		Filters: filters.NewArgs(filters.Arg("name", "redis"), filters.Arg("name", "test")),
	})
	if err != nil {
		panic(err)
	}

	for _, container := range containers {
		fmt.Println(container.Names[0], container.State, "?", container.Status, container)
	}
}

func TestDockerControl(t *testing.T) {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := cli.Close(); err != nil {
			return
		}
	}()

	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{
		//All:     true,
		Filters: filters.NewArgs(filters.Arg("name", "test")),
	})
	if err != nil {
		panic(err)
	}

	//duration := time.Second * 10
	//if err := cli.ContainerStop(ctx, containers[0].ID, &duration); err != nil {
	//	panic(err)
	//}

	if err := cli.ContainerStart(ctx, containers[0].ID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}
}
