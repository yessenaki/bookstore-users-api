package datetime

import "time"

const layout = "2006-01-02 15:04:05"

func GetCurrentDatetime() string {
	now := time.Now().UTC()

	return now.Format(layout)
}
