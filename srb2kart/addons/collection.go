package addons

func (a Addon) Addons() []string {
	return []string{a.File}
}

func (a Addon) Format() []string {
	return []string{a.File}
}

func (a AddonGroup) Addons() []string {
	result := []string{}
	for _, addon := range a.Items {
		result = append(result, addon.Addons()...)
	}
	return result
}

func (a AddonGroup) Format() (result []string) {
	result = []string{a.GroupName}
	for _, item := range a.Items {
		for i, line := range item.Format() {
			var padding string
			if i == 0 {
				padding = "|- "
			} else {
				padding = "|  "
			}
			result = append(result, padding+line)
		}
	}
	return result
}
