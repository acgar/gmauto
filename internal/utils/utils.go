package utils

import "time"

func DaysAgo(days int) string {
	return time.Now().AddDate(0, 0, -days).Format("2006/01/02")
}
