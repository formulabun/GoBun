package add

import (
	"GoBun/database"
	"fmt"
	"os"
	"path/filepath"
)

const receiveGroupsFlag = "-G"
const supplyFileFlag = "-a"
const supplyGroupsFlag = "-g"

type arguments struct {
	groups    []string
	addGroups []string
	addFiles  []string
}

func printUsage() {
	var base = filepath.Base(os.Args[0])
	fmt.Printf("usage: %s add -G {groups} -a {files} -g {groups}\n", base)
	fmt.Printf("short: %s add <group> {files} -a {files} -g {groups}\n", base)
	fmt.Printf("show this help: %s add\n", base)
	fmt.Printf(
		`When the normal usage is used, add items to one or more groups.
When the short form is used, only add items to one group.

Optional flags: 
    -G {groups} : groups to add the other arguments to
    -a {files}  : files to add
    -g {groups} : groups to add
`)
}

func parseArguments(args []string) (result arguments) {
	if len(args) == 0 {
		fmt.Println()
		printUsage()
		fmt.Println()
		os.Exit(1)
	}
	var receive, supplyItems, supplyGroups bool
	for _, arg := range args {
		receive = receive || arg == receiveGroupsFlag
		supplyItems = supplyItems || arg == supplyFileFlag
		supplyGroups = supplyGroups || arg == supplyGroupsFlag
	}
	if !receive {
		var groupName = args[0]
		args = args[1:]
		if groupName == supplyFileFlag || groupName == supplyGroupsFlag {
			fmt.Println("err: The group name cannot be the same as a flag")
			fmt.Println()
			printUsage()
			fmt.Println()
			os.Exit(1)
		}
		result.groups = []string{groupName}
	}

	var state = "supplyFile"
	for _, arg := range args {
		if arg == receiveGroupsFlag {
			state = "receive"
			continue
		}
		if arg == supplyFileFlag {
			state = "supplyFile"
			continue
		}
		if arg == supplyGroupsFlag {
			state = "supplyGroup"
			continue
		}

		fmt.Println(arg)

		switch state {
		case "receive":
			result.groups = append(result.groups, arg)
		case "supplyFile":
			result.addFiles = append(result.addFiles, arg)
		case "supplyGroup":
			result.addGroups = append(result.addGroups, arg)
		default:
		case "files":
		}
	}

	return result
}

func Add() {
	arguments := parseArguments(os.Args[2:])

	client, disconnect, err := database.CreateClient()
	defer disconnect()

	if err != nil {
		fmt.Printf("Could not add addons/groups: %s", err)
		return
	}

	for _, group := range arguments.groups {
		err = client.InsertAddonFile(group, arguments.addFiles)
		if err != nil {
			fmt.Printf("%s\n", err)
		}
		err = client.InsertAddonGroup(group, arguments.addGroups)
		if err != nil {
			fmt.Printf("%s\n", err)
		}
	}

	if err := disconnect(); err != nil {
		fmt.Printf("Error while ending list: %s", err)
	}
}
