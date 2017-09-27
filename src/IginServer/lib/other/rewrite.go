package other

import (
	// "fmt"
	// "github.com/go-martini/martini"
	// "io/ioutil"
	"net/http/httputil"
	// "net"
	"net/http"
	"net/url"
	// "time"
)

func Rewrite(res http.ResponseWriter, req *http.Request, iurl string) {
	remote, err := url.Parse(iurl)
	if err != nil {
		panic(err)
	}
	proxy := httputil.NewSingleHostReverseProxy(remote)
	proxy.ServeHTTP(res, req)
}
