package cookie

import (
	// "code.google.com/p/mahonia"
	"fmt"
	"net/http"
	"net/url"
	// "IginServer/conf"
	"IginServer/lib/other"
)

func Get(req *http.Request, key string) string {
	str, err := req.Cookie(url.QueryEscape(key))
	if err != nil {
		return ""
	} else {
		s, _ := url.QueryUnescape(str.Value)
		return s
	}
}

func Set(res http.ResponseWriter, key string, value string) {
	/*
		type Cookie struct {
		        Name       string
		        Value      string
		        Path       string
		        Domain     string
		        Expires    time.Time
		        RawExpires string

		        // MaxAge=0 means no 'Max-Age' attribute specified.
		        // MaxAge<0 means delete cookie now, equivalently 'Max-Age: 0'
		        // MaxAge>0 means Max-Age attribute present and given in seconds
		        MaxAge   int
		        Secure   bool
		        HttpOnly bool
		        Raw      string
		        Unparsed []string // Raw text of unparsed attribute-value pairs
		}
	*/
	c := http.Cookie{Name: key, Value: value, Path: "/"}
	http.SetCookie(res, &c)
}

func Secure_set(res http.ResponseWriter, key string, value string, base string) {
	value = string(other.Base64Encode([]byte(value), base))
	c := http.Cookie{Name: key, Value: value, Path: "/"}
	http.SetCookie(res, &c)
}

func Secure_get(req *http.Request, key string, base string) string {
	c, err := req.Cookie(key)
	if err == nil {
		value := c.Value
		v, _ := other.Base64Decode([]byte(value), base)
		strv := fmt.Sprintf("%s", v)
		return strv
	} else {
		// fmt.Println(err)
		return ""
	}
}
