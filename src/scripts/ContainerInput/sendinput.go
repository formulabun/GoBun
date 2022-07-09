package main

import (
	"GoBun/docker/container/input"
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"os"
	"strings"
)

func exitOnError(err error, reason string) {
	if err != nil {
		fmt.Println(reason)
		os.Exit(1)
	}
}

func findContainer(cli *client.Client) types.Container {
	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	exitOnError(err, "Couldn't list containers.")
	for _, c := range containers {
		if c.Image == "srb2kart-fbun" {
			return c
		}
	}
	exitOnError(err, "Couldn't find container.")
	return types.Container{}
}

func main() {
	cli, err := client.NewClientWithOpts()
	exitOnError(err, "Coudn't start the client.")

	container := findContainer(cli)
	input.SendInput(cli, container, strings.Join(os.Args[1:], " "))
}
