package date

import "time"

const (
	defaultLayout = "2006-01-02T15:04:05Z"
)

func GetNow() time.Time {
	return time.Now().UTC()
}

func GetNowString() string {
	return GetNow().Format(defaultLayout)
}

func GetNowDatetime() time.Time {
	return GetNow()
}
