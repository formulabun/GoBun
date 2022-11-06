package addons

func (a AddonGroup) String() string {
	return a.GroupName
}

func (a Addon) String() string {
	return a.File
}
