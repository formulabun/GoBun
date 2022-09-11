package list

import (
	"GoBun/scripts/common/subcommand"
  "GoBun/files"
	"fmt"
)

var List = subcommand.Subcommand{"list", []string{"ls", "list"}, list}

func list(args []string) (fmt.Stringer, error) {
	fmt.Println(files.ListAddonFiles())
	return nil, nil
}
