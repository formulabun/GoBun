package remove

import (
	"GoBun/files"
	"GoBun/functional/strings"
	"GoBun/scripts/common/subcommand"
	"fmt"
)

var Remove = subcommand.Subcommand{"remove", []string{"rm", "remove"}, remove}

func remove(args []string) (fmt.Stringer, error) {
	if len(args) != 1 {
		return strings.Stringer{"<filename>"}, nil
	}
	err := files.RemoveAddonFile(args[0])
	return nil, err
}
