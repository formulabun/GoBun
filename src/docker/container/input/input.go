package input

import (
	"context"
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

func SendInput(client *client.Client, container types.Container, input string) (error, string) {
	connection, err := client.ContainerAttach(context.Background(), container.ID, options)
	if err != nil {
		return err, "Couldn't attach to the container."
	}

	_, err = connection.Conn.Write([]byte(input))
	if err != nil {
		return err, "Couldn't write to the container."
	}
	connection.Close()

	return nil, ""
}
