package templatefunc

import (
	"IginServer/conf"
)

func config() map[string]map[string]string {
	return conf.GET
}
