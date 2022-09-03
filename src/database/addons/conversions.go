package addons

import (
	"GoBun/srb2kart/addons"
)

func contentToCollection(in content) addons.AddonCollection {
	if in.kind == fileType {
		return addons.Addon{File: in.value}
	} else if in.kind == groupType {
		return addons.AddonGroup{GroupName: in.value}
	}

	return addons.Addon{}
}
