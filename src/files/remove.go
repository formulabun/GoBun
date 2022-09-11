package files

import (
	"GoBun/constants"
	"os"
	"path/filepath"
)

func RemoveAddonFile(fileName string) error {
	return os.Remove(filepath.Join(constants.FilePath, fileName))
}
