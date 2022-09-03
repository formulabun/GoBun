package remove

import (
	"GoBun/database"
	"GoBun/functional/strings"
	"GoBun/scripts/common/subcommand"
	"fmt"
)

var Remove = subcommand.Subcommand{"remove", []string{"remove", "rm"}, remove}

func remove(args []string) (fmt.Stringer, error) {
	client, disconnect, _ := database.CreateClient()
	defer disconnect()

	if len(args) != 1 {
		return strings.Stringer{"<server name>"}, nil
	}

	client.RemoveServer(args[0])
	return nil, nil
}
