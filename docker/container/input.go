package container

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

var (
	options = types.ContainerAttachOptions{
		Stream: true,
		Stdin:  true,
		Stdout: false,
		Stderr: true,
	}
)

func SendInput(client *client.Client, container types.Container, input string) error {
	connection, err := client.ContainerAttach(context.Background(), container.ID, options)
	if err != nil {
		return fmt.Errorf("could not attach to the container: %s", err)
	}

	_, err = connection.Conn.Write([]byte(input))
	if err != nil {
		return fmt.Errorf("could not write to the container: %s", err)
	}
	connection.Close()

	return nil
}
