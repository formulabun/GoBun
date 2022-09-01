package add

import (
  "os"
  "fmt"
  "GoBun/scripts/common/subcommand"
  "GoBun/constants"
  "GoBun/functional/strings"
  "net/http"
  "io"
  "path/filepath"
)

var Add = subcommand.Subcommand{"add", []string{"add", "fetch", "a"}, add}

func printUsage() fmt.Stringer {
  return strings.Stringer{"url <outfile>"}
}

func getPath(filename string) string {
  filepath.Base(filename)
  _, err := os.Stat(constants.BasePath)
  if err != nil {
    fmt.Printf("Could not find the directory \"%s\". Writing to working directory.\n", constants.BasePath)
    return filename
  } else {
    return filepath.Join(constants.BasePath, filename)
  }
}

func add(args []string) (help fmt.Stringer, err error) {
  if len (args) < 1 || len(args) > 2 {
    return printUsage(), nil
  }

  url := args[0]

  resp, err := http.Get(url)
  if err != nil {
    return nil, fmt.Errorf("Could not add file: %s\nIs it a valid url?", err)
  }

  defer resp.Body.Close()
  if resp.StatusCode != 200 {
    return nil, fmt.Errorf("Error fetching the file: %s", resp.Status)
  }

  var filename string
  if len(args) == 2 {
    filename = filepath.Base(args[1])
  } else {
    filename = filepath.Base(url)
  }

  file, err := os.Create(getPath(filename))

  if err != nil {
    return nil, fmt.Errorf("Could not save the file: %s\nDid you mount correctly?", err)
  }

  defer file.Close()
  _, err = io.Copy(file, resp.Body)
  if err != nil {
    return nil, fmt.Errorf("Could not write file: $s", err)
  }

  return nil, nil
}
