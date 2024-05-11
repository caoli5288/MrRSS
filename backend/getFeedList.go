package backend

type FeedInfo struct {
	Title string
	Link  string
}

func GetFeedList() []FeedInfo {
	feedList := []FeedInfo{
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
		{
			Title: "Appinn",
			Link:  "https://www.appinn.com/feed/",
		},
	}

	return feedList
}
