package backend

type RssInfo struct {
	Title string
	Link  string
}

// GetRssList returns a list of RSS feeds
func GetRssList() []RssInfo {
	rssList := []RssInfo{
		{
			Title: "Kawabangga",
			Link:  "https://www.kawabangga.com/feed",
		},
		{
			Title: "Julia Evans",
			Link:  "https://jvns.ca/atom.xml",
		},
		{
			Title: "Ruanyifeng",
			Link:  "https://www.ruanyifeng.com/blog/atom.xml",
		},
	}

	return rssList
}
