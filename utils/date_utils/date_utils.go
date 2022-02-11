package date_utils

import "time"

const (
	apiDateLayout = "02-01-2006T15:04:05Z"
)

func GetNow() time.Time {
	return time.Now()
}

func GetNowString() string {
	return GetNow().Format(apiDateLayout)
}
