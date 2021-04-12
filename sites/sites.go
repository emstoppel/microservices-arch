package sites

var sites = []string{"MLA", "MLB"}

func Exists(site string) bool {
	if site == "" {
		return false
	}
	for _, s := range sites {
		if s == site {
			return true
		}
	}
	return false
}
