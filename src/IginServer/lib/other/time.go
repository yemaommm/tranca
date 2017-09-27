package other

import (
	"fmt"
	"strconv"
	"time"
)

func ToDate(t interface{}, n ...string) string {
	var err error
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("M", err)
		}
	}()
	if t == nil {
		return ""
	}
	var tt int64
	switch t.(type) {
	case string:
		tt, err = strconv.ParseInt(t.(string), 10, 0)
		if err != nil {
			return "0"
		}
	case int:
		tt = int64(t.(int))
	case int16:
		tt = int64(t.(int16))
	case int8:
		tt = int64(t.(int8))
	case int64:
		tt = t.(int64)
	}
	tm := time.Unix(tt, 0)
	if len(n) == 0 {
		return tm.Format("2006-01-02 15:04:05")
	} else {
		return tm.Format(n[0])
	}
}

func ToTime(t string, n ...string) int {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("M", err)
		}
	}()
	format := "2006-01-02 15:04:05"
	if len(n) != 0 {
		format = n[0]
	} else {
		format = format[0:len(t)]
	}
	// fmt.Println(format)
	// fmt.Println(t)
	// fmt.Println(time.Now())
	loc, _ := time.LoadLocation("Local")
	if the_time, err := time.ParseInLocation(format, t, loc); err != nil {
		fmt.Println("M", err)
	} else {
		// fmt.Println(the_time)
		// fmt.Println(time.Unix(the_time.Unix(), 0).Format(format))
		// fmt.Println(the_time.Unix())
		return int(the_time.Unix())
	}
	return 0
}
