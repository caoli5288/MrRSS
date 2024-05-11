package backend

type FeedContentFilterInfo struct {
	FeedTitle string
	FeedImage string
	Title     string
	Link      string
	TimeSince string
	Time      string
	Image     string
	Content   string
}

func FilterFeedContent() []FeedContentFilterInfo {
	feedContent := GetFeedContent()

	var feedContentInfo []FeedContentFilterInfo

	for _, item := range feedContent {
		imageURL := ""
		filterImageUrl := FilterImage(item.Item.Content)
		if item.Item.Image != nil {
			imageURL = item.Item.Image.URL
		} else if filterImageUrl != nil {
			imageURL = *filterImageUrl
		}

		timeSinceStr := TimeSince(item.Item.PublishedParsed)

		feedContentInfo = append(feedContentInfo, FeedContentFilterInfo{
			FeedTitle: item.Title,
			FeedImage: item.Image,
			Title:     item.Item.Title,
			Link:      item.Item.Link,
			TimeSince: timeSinceStr,
			Time:      item.Item.PublishedParsed.Format("2006-01-02 15:04"),
			Image:     imageURL,
			Content:   item.Item.Content,
		})
	}

	return feedContentInfo
}
