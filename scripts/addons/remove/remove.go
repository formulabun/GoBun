package remove

import (
	"GoBun/database"
	"fmt"
	"os"
	"path/filepath"
)

func printUsage() {
	var base = filepath.Base(os.Args[0])
	fmt.Printf("delete a group: %s remove <group name>\n", base)
	fmt.Printf("delete an element: %s remove <group name> <item>\n", base)
}

func removeGroup(client *database.DBClient, groupName string) error {
	return client.RemoveGroup(groupName)
}

func removeItem(client *database.DBClient, groupName string, element string) error {
	return client.RemoveItemFromGroup(groupName, element)
}

func Remove() {
	if len(os.Args) < 3 {
		printUsage()
		os.Exit(0)
	}

	client, disconnect, err := database.CreateClient()
	defer disconnect()
	if err != nil {
		fmt.Printf("Could not delete: %s", err)
		os.Exit(1)
	}

	var group = os.Args[2]

	if len(os.Args) == 3 {
		err = removeGroup(client, group)
	} else {
		var item = os.Args[3]
		err = removeItem(client, group, item)
	}

	if err != nil {
		fmt.Printf("Could not remove item: %s\n", err)
		os.Exit(1)
	}

}
