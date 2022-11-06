package main

import (
	"GoBun/scripts/addons/add"
	"GoBun/scripts/addons/list"
	"GoBun/scripts/addons/remove"
	"GoBun/scripts/addons/set"
	"fmt"
	"os"
	"path/filepath"
)

var addMethod = []string{"add", "a"}
var removeMethod = []string{"remove", "rem", "r"}
var listMethod = []string{"list", "ls", "l"}
var setMethod = []string{"set", "s"}

func printUsage() {
	var base = filepath.Base(os.Args[0])
	fmt.Printf("usage: %s {add,remove,list,set}\n", base)
}

func main() {
	if len(os.Args) < 2 {
		printUsage()
		return
	}

	var method = os.Args[1]
	switch {
	case isAdd(method):
		add.Add()
	case isRemove(method):
		remove.Remove()
	case isList(method):
		list.List()
	case isSet(method):
		set.Set()
	default:
		printUsage()
	}
}

func isAdd(method string) bool {
	return contains(method, addMethod)
}

func isRemove(method string) bool {
	return contains(method, removeMethod)
}

func isList(method string) bool {
	return contains(method, listMethod)
}

func isSet(method string) bool {
	return contains(method, setMethod)
}

func contains(value string, array []string) bool {
	for _, a := range array {
		if a == value {
			return true
		}
	}
	return false
}
