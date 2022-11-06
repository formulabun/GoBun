package files

import (
	"GoBun/constants"
	"GoBun/functional/array"
	"io/fs"
	"io/ioutil"
)

func ListAddonFiles() []string {
	info, _ := ioutil.ReadDir(constants.FilePath)
	return array.Map(info, func(file fs.FileInfo) string {
		return file.Name()
	})
}
