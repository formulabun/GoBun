package main

import (
	"GoBun/scripts/common/subcommand"
	"GoBun/scripts/kartserver/list"
	"GoBun/scripts/kartserver/remove"
	"GoBun/scripts/kartserver/set"
	"fmt"
	"os"
	"path/filepath"
)

var listOptions = []string{"list"}

func printUsage(options fmt.Stringer) {
	base := filepath.Base(os.Args[0])
	fmt.Printf("%s %s\n", base, options)
}

func main() {
	commands := subcommand.Registry{}
	commands.Register(list.List).Register(set.Set).Register(remove.Remove)

	if len(os.Args) < 2 {
		printUsage(commands)
		os.Exit(0)
	}

	help, err := commands.Run(os.Args[1], os.Args[2:])
	if err != nil {
		fmt.Printf("command failed: %s\n", err)
		os.Exit(1)
	}
	if help != nil {
		printUsage(help)
		os.Exit(0)
	}
}
