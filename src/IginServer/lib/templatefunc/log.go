package templatefunc

import (
	"fmt"
)

func ilog(i ...interface{}) string {
	if len(i) <= 0 {
		return ""
	} else {
		return fmt.Sprintf(i[0].(string), i[1:]...)
	}
}
