package backend

import (
	"fmt"
	"time"
)

func TimeSince(t *time.Time) string {
	timeSince := time.Since(*t)
	timeStr := "now"

	if timeSince > 0 {
		if timeSince < time.Hour {
			minutes := int(timeSince.Minutes())
			if minutes == 1 {
				timeStr = fmt.Sprintf("%d minute ago", minutes)
			} else {
				timeStr = fmt.Sprintf("%d minutes ago", minutes)
			}
		} else if timeSince < 24*time.Hour {
			hours := int(timeSince.Hours())
			if hours == 1 {
				timeStr = fmt.Sprintf("%d hour ago", hours)
			} else {
				timeStr = fmt.Sprintf("%d hours ago", hours)
			}
		} else if timeSince < 365*24*time.Hour {
			days := int(timeSince.Hours() / 24)
			if days == 1 {
				timeStr = fmt.Sprintf("%d day ago", days)
			} else {
				timeStr = fmt.Sprintf("%d days ago", days)
			}
		} else {
			years := int(timeSince.Hours() / (365 * 24))
			if years == 1 {
				timeStr = fmt.Sprintf("%d year ago", years)
			} else {
				timeStr = fmt.Sprintf("%d years ago", years)
			}
		}
	}

	return timeStr
}
