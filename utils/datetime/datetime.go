package datetime

import "time"

const layout = "2006-01-02T15:04:05Z"

func GetCurrentDatetime() string {
	now := time.Now().UTC()

	return now.Format(layout)
}
