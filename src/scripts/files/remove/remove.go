package remove

import (
  "GoBun/srb2kart/addons"
  "GoBun/scripts/common/subcommand"
  "GoBun/functional/strings"
  "fmt"
)

var Remove = subcommand.Subcommand{"remove", []string{"remove", "rm"}, remove}

func remove(args []string) (fmt.Stringer, error) {
  if len(args) == 0 {
    return strings.Stringer{"{filename}"}, nil
  }

  for n, filename := range args {
    err := addons.Remove(addons.Addon{filename})
    if err != nil {
      return nil, fmt.Errorf("Could not delete file nr %v: %s", n, err)
    }
  }

  return nil, nil
}
