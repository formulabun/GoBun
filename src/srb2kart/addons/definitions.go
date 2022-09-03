package addons

type AddonCollection interface {
	Addons() []string
	Format() []string
}

type Addon struct {
	File string
}

type AddonGroup struct {
	GroupName string
	Items     []AddonCollection
}
