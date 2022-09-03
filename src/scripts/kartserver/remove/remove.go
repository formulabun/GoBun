package remove

import (
  "GoBun/scripts/common/subcommand"
  "fmt"
  "GoBun/functional/strings"
  "GoBun/database"
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
