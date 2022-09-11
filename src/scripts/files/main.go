package main

import (
	"GoBun/scripts/common/subcommand"
	"GoBun/scripts/files/get"
	"GoBun/scripts/files/list"
	"GoBun/scripts/files/remove"
	"fmt"
	"os"
)

func main() {
	var r = subcommand.Registry{}
	r.Register(list.List).Register(remove.Remove).Register(get.Get)

	if len(os.Args) < 2 {
		fmt.Printf("usage: %s %s\n", os.Args[0], r)
		return
	}

	help, err := r.Run(os.Args[1], os.Args[2:])
	if err != nil {
		fmt.Println(err)
		return
	}
	if help != nil {
		fmt.Printf("usage: %s %s\n", os.Args[0], help)
		return
	}

}
