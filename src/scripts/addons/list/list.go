package list

import (
	"GoBun/database"
	"fmt"
	"os"
)

func printListUsage() {
	fmt.Printf("usage: %s list", os.Args[0])
}

func listGroups() {
	client, disconnect, err := database.CreateClient()
	if err != nil {
		fmt.Printf("Could not list groups: %s", err)
		return
	}

	result, err := client.ListAddonGroups()

	if err != nil {
		fmt.Printf("Error while reading database: %s", err)
	}

	for _, item := range result.Items {
		fmt.Println(item)
	}

	if err := disconnect(); err != nil {
		fmt.Printf("Error while ending list: %s", err)
	}
}

func listGroup(groupName string) {
	client, disconnect, err := database.CreateClient()
	if err != nil {
		fmt.Printf("Could not list groups: %s", err)
		return
	}
	group, _ := client.AddonGroup(groupName)
	for _, line := range group.Format() {
		fmt.Println(line)
	}

	if err := disconnect(); err != nil {
		fmt.Printf("Error while ending reading a group: %s", err)
	}
}

func List() {
	switch len(os.Args) {
	case 2:
		listGroups()
	case 3:
		listGroup(os.Args[2])
	default:
		printListUsage()
	}
}
