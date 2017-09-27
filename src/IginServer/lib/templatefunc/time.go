package templatefunc

import (
	"IginServer/lib/other"
)

func todate(t interface{}, n ...string) string {
	defer other.Thow()
	return other.ToDate(t, n...)
}

func totime(t string, n ...string) int {
	defer other.Thow()
	return other.ToTime(t, n...)
}
