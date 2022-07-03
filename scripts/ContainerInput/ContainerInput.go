package main

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"os"
  "strings"
)

var (
	options = types.ContainerAttachOptions{
		Stream: true,
		Stdin:  true,
		Stdout: false,
		Stderr: true,
	}
)

func exitOnError(err error, reason string) {
  if err != nil {
    fmt.Println(reason)
    os.Exit(1)
  }
}

func findContainer(cli *client.Client) (types.Container) {
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

	connection, err := cli.ContainerAttach(context.Background(), container.ID, options)
  exitOnError(err, "Couldn't attach to the container.")

  command := strings.Join(os.Args[1:], " ") + "\n"
	_, err = connection.Conn.Write([]byte(command))
  exitOnError(err, "Couldn't write to the container.")

	connection.Close()
}
