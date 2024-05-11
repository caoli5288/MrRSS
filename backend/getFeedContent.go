package backend

import (
	"context"
	"net/url"
	"sort"
	"time"

	"github.com/mmcdole/gofeed"
)

type FeedContentInfo struct {
	Title string
	Image string
	Item  gofeed.Item
}

func GetFeedContent() []FeedContentInfo {
	feedList := GetFeedList()

	var feedContent []FeedContentInfo

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	feedParser := gofeed.NewParser()

	for _, feed := range feedList {
		link := feed.Link
		feed, err := feedParser.ParseURLWithContext(link, ctx)
		if err != nil {
			continue
		}

		for _, item := range feed.Items {
			u, err := url.Parse(feed.Link)
			if err != nil {
				panic(err)
			}

			imageUrl := "https://www.google.com/s2/favicons?sz=16&domain=" + u.Host
			if feed.Image != nil {
				imageUrl = feed.Image.URL
			}

			feedContent = append(feedContent, FeedContentInfo{
				Title: feed.Title,
				Image: imageUrl,
				Item:  *item,
			})
		}
	}

	sort.Slice(feedContent, func(i, j int) bool {
		return feedContent[i].Item.PublishedParsed.After(*feedContent[j].Item.PublishedParsed)
	})

	return feedContent
}
