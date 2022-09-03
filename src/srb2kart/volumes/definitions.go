package volumes

type VolumeSet struct {
	LocalFile string `bson::"localFiles"`
	Luafiles  string `bson::"luaFiles"`
}
