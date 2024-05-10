package backend

import (
	"context"
	"sort"
	"time"

	"github.com/mmcdole/gofeed"
)

func GetRssContent() []gofeed.Item {
	rssList := GetRssList()

	var rssContent []gofeed.Item

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	feedParser := gofeed.NewParser()

	for _, rss := range rssList {
		link := rss.Link
		feed, _ := feedParser.ParseURLWithContext(link, ctx)

		for _, item := range feed.Items {
			rssContent = append(rssContent, *item)
		}
	}

	sort.Slice(rssContent, func(i, j int) bool {
		return rssContent[i].PublishedParsed.After(*rssContent[j].PublishedParsed)
	})

	return rssContent
}
