package register

var Registry = make(map[string]map[string]map[string]func(string))

func Part1(year string, day string, partFunc func(string)) {
	ensureExists(year, day)
	Registry[year][day]["part1"] = partFunc
}
func Part2(year string, day string, partFunc func(string)) {
	ensureExists(year, day)
	Registry[year][day]["part2"] = partFunc
}

func ensureExists(year string, day string) {
	if _, exists := Registry[year]; !exists {
		Registry[year] = make(map[string]map[string]func(string))
	}
	if _, exists := Registry[year][day]; !exists {
		Registry[year][day] = make(map[string]func(string))
	}
}
