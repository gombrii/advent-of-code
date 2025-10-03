package registry

var Registry = make(map[string]map[string]map[string]func(string) any)

func Register(year string, day string, part string, partFunc func(string) any) {
	ensureExists(year, day)
	Registry[year][day][part] = partFunc
}

func ensureExists(year string, day string) {
	if _, exists := Registry[year]; !exists {
		Registry[year] = make(map[string]map[string]func(string) any)
	}
	if _, exists := Registry[year][day]; !exists {
		Registry[year][day] = make(map[string]func(string) any)
	}
}
