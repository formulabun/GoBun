package get

import (
	"GoBun/constants"
	"GoBun/functional/strings"
	"GoBun/http"
	"GoBun/scripts/common/subcommand"
	"fmt"
)

var Get = subcommand.Subcommand{"get", []string{"get"}, get}

func get(args []string) (fmt.Stringer, error) {
	if len(args) != 1 {
		return strings.Stringer{"<url>"}, nil
	}
	err := http.Download(args[0], constants.FilePath)
	return nil, err
}
