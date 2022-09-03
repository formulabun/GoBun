package srb2kart

import (
	"GoBun/docker/volume"
)

type Srb2kart struct {
	Name       string
	Port       int
	Volumes    volume.VolumeSet
	AddonGroup string
}

func DefaultSrb2kart() Srb2kart {
	return Srb2kart{
		"srb2kart",
		5029,
		volume.VolumeSet{},
		"addons",
	}
}
