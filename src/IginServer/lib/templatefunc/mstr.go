package templatefunc

import (
	"strconv"
)

func mstr(num, str string) string {
	ret := ""
	count, _ := strconv.Atoi(num)
	for i := 0; i < count; i++ {
		ret = ret + str
	}
	return ret
}
