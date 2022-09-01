package list

import (
  "GoBun/scripts/common/subcommand"
  "GoBun/functional/strings"
  "GoBun/database"
  "fmt"
)

var List = subcommand.Subcommand{"list", []string{"list", "ls"}, list}

func list(args []string) (fmt.Stringer, error) {
  client, disconnect, _ := database.CreateClient()
  defer disconnect()

  var err error = nil
  var r interface{}

  switch len(args) {
  case 0:
    r, err = client.ListServers()
  case 1:
    r, err = client.Server(args[0])
  default:
    return strings.Stringer{"<server name>"}, nil 
  }
  if err != nil {
    return nil, err
  }

  fmt.Println(r)
  return nil, nil
}
