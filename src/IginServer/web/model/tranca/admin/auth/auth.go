package auth

import (
// "encoding/json"
// "fmt"
// "IginServer/lib/session"
)

func Auth(sess map[string]interface{}) int {
	if sess["admin_info"] != nil {
		admin_info := sess["admin_info"].(map[string]string)
		if sess["admin_user"] == admin_info["username"] && sess["admin_passwd"] == admin_info["passwd"] {
			return 1
		}
	}
	return 0
}
