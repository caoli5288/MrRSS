package backend

import (
	"regexp"
)

func FilterImage(content string) *string {
	imgRegex := regexp.MustCompile(`img[^>]*src="([^"]*)`)

	var firstImageURL *string
	imgMatches := imgRegex.FindStringSubmatch(content)
	if len(imgMatches) > 1 {
		firstImageURL = &imgMatches[1]
	}

	return firstImageURL
}
