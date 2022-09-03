package addons

import (
	"GoBun/constants"
	"GoBun/functional/array"
	"fmt"
	"os"
	"path/filepath"
)

const readErr = "Could not list addons: %s\nDid you mount /data correctly?"

func fullPathTo(addon Addon) string {
	_, err := os.Stat(constants.BasePath)
	if err != nil {
		fmt.Printf("Could not find the directory \"%s\". Either the docker mount is incorrect or you are working locally. Writing to current directory.\n", constants.BasePath)
		return addon.File
	} else {
		path := filepath.Join(constants.BasePath, addon.File)
		return path
	}
}

func All() ([]Addon, error) {
	file, err := os.Open(constants.BasePath)
	defer file.Close()
	if err != nil {
		return nil, fmt.Errorf(readErr, err)
	}

	entries, err := file.ReadDir(0)
	if err != nil {
		return nil, fmt.Errorf(readErr, err)
	}
	return array.Map(entries, func(entry os.DirEntry) Addon {
		return Addon{entry.Name()}
	}), nil
}

func Exists(addon Addon) bool {
	path := fullPathTo(addon)
	stat, err := os.Stat(path)
	return err == nil && stat.Mode().IsRegular()
}

func Remove(addon Addon) error {
	err := os.Remove(fullPathTo(addon))
	if err != nil {
		return fmt.Errorf("Could not delete file %s: %s", addon.File, err)
	}
	return nil
}
