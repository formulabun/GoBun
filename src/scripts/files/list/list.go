package list

import (
  "GoBun/scripts/common/subcommand"
  "GoBun/srb2kart/addons"
  "GoBun/functional/strings"
  "fmt"
)

var List = subcommand.Subcommand{"list", []string{"list", "ls"}, list}

func list(args []string) (fmt.Stringer, error) {
  if len(args) > 0 {
    return strings.Stringer{""}, nil
  }

  addons, err := addons.All();
  if err != nil {
    return nil, fmt.Errorf("Could not list addons: %s", err)
  }

  for _, addon := range addons {
    fmt.Println(addon)
  }

  return nil, nil
}
