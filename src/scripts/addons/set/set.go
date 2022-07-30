package set

import (
  "path/filepath"
  "os"
  "fmt"
  "GoBun/database"
)

func printUsage() {
  var base = filepath.Base(os.Args[0])
  fmt.Printf("%s set <groupname> [items]\n", base)
}

func Set() {
  if len(os.Args) < 4 {
    printUsage()
    os.Exit(0)
  }
  var groupName = os.Args[2]
  var items = os.Args[3:]
  var client, disconnect, err = database.CreateClient()
  defer disconnect()
  if err != nil {
    fmt.Printf("Could not create/set the contents of a group: %s\n", err)
    os.Exit(1)
  }

  if err = client.CreateAddonGroup(groupName, items); err != nil {
    fmt.Printf("Could not create/set contets of a group: %s\n", err)
    os.Exit(1)
  }
}
