package utils

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func RefreshTeku() {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	for _, container := range containers {
		if container.Image == "consensys/teku:latest" {
			fmt.Printf("SIGNALLING HUP TO VALIDATOR")
			err := cli.ContainerKill(context.Background(), container.ID, "HUP")
			if err != nil {
				panic(err)
			}
		}
	}
}