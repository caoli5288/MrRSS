package backend

type RssContentInfo struct {
	Title   string
	Link    string
	Time    string
	Summary string
	Content string
}

func FilterRssContent() []RssContentInfo {
	rssContent := GetRssContent()

	var rssContentInfo []RssContentInfo

	for _, item := range rssContent {
		rssContentInfo = append(rssContentInfo, RssContentInfo{
			Title:   item.Title,
			Link:    item.Link,
			Time:    item.PublishedParsed.Format("2006-01-02 15:04:05"),
			Summary: item.Description,
			Content: item.Content,
		})
	}

	return rssContentInfo
}
