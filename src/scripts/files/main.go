package main

import (
	"GoBun/scripts/common/subcommand"
  "GoBun/scripts/files/add"
  "GoBun/scripts/files/list"
  "GoBun/scripts/files/remove"
  "os"
  "fmt"
  "path/filepath"
)

func printUsage(options fmt.Stringer) {
  fmt.Printf("usage: %s %s\n", filepath.Base(os.Args[0]), options)
}

func main() {
	subcommands := subcommand.Registry{}

  subcommands.Register(add.Add).Register(list.List).Register(remove.Remove)

  if len(os.Args) < 2 {
    printUsage(subcommands)
    return
  }


  help, err := subcommands.Run(os.Args[1], os.Args[2:]);
  if err != nil {
    fmt.Printf("Something went wrong with that command: %s\n", err)
  }

  if help != nil {
    printUsage(help)
  }
}
