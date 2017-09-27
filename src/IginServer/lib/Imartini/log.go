package Imartini

import (
	"fmt"
	// "log"
	"IginServer/conf"
	"IginServer/lib/ierror"
	"IginServer/lib/redis"
	"encoding/json"
	ir "github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	panicHtml = `<html>
<head><title>PANIC: %s</title>
<style type="text/css">
html, body {
	font-family: "Roboto", sans-serif;
	color: #333333;
	background-color: #ea5343;
	margin: 0px;
}
h1 {
	color: #d04526;
	background-color: #ffffff;
	padding: 20px;
	border-bottom: 1px dashed #2b3848;
}
pre {
	margin: 20px;
	padding: 20px;
	border: 2px solid #2b3848;
	background-color: #ffffff;
}
</style>
</head><body>
<h1>PANIC</h1>
<pre style="font-weight: bold;">%s</pre>
<pre>%s</pre>
</body>
</html>`
)

var MyLog = Mlog{}

var file *os.File
var ofn string
var Replace = "!(B)"
var log_id = conf.GetString("config", "LOGO_ID")

type G map[string]interface{}

type Mlog struct{}

func (m *Mlog) Printf(str string, v ...interface{}) {
	m.Other("", str, v...)
}
func (m *Mlog) Err(str string, v ...interface{}) {
	m.Other("err.", str, v...)
}
func (m *Mlog) Other(f, str string, v ...interface{}) {
	l := strings.Replace(fmt.Sprintf(str, v...), Replace, "%", -1)
	l = "[" + time.Now().Format("2006-01-02 15:04:05") + "] " + l
	go m.tosave(f, l)
}

func (m *Mlog) tosave(str ...string) {
	data, _ := json.Marshal(map[string]interface{}{"name": str[0], "msg": str[1]})
	ris := redis.Get()
	defer ris.Close()
	ris.LPUSH(""+log_id, string(data))
	ris.PUB("i"+log_id, "1")
	data = nil
}

func save() {
	p := redis.PUB(redis.Get())
	p.Subscribe("i" + log_id)
R:
	for {
		data := p.Receive()
		switch data.(type) {
		case ir.Message:
			switch data {
			default:
				predis := redis.Get()
			N:
				for {
					stmp, _ := predis.RPOP("" + log_id)
					if stmp == "" {
						break N
					}
					var v map[string]string
					if err := json.Unmarshal([]byte(stmp), &v); err == nil {
						savefile(fmt.Sprintf("%s", v["name"]), fmt.Sprintf("%v", v["msg"]))
					}
					v = nil
				}
				predis.Close()
				predis = nil
			}
		case error:
			fmt.Println("log:", data)
			break R
		}
		data = nil
	}
	p.Close()
	p = nil
	go save()
}

func savefile(t string, str string) {
	f := time.Now().Format("2006-01-02")
	m := time.Now().Format("2006-01")
	f = LOGO_PATH + t + f + ".log"
	if f != ofn {
		file.Close()
		fs := strings.Split(f, "/")
		fs = append(fs[0:len(fs)-1], "/", m, "/", fs[len(fs)-1])
		f = strings.Join(fs, "/")
		ofn = f
		os.MkdirAll(strings.Join(fs[0:len(fs)-1], "/"), 0666)
		file, _ = os.OpenFile(f, os.O_APPEND|os.O_CREATE|os.O_RDWR|os.O_SYNC, 0666)
	}
	info, _ := file.Stat()
	if info.Size() > int64(LOGO_SIZE) {
		file.Close()
		os.Rename(f, f+"."+strconv.Itoa(time.Now().Hour())+"-"+strconv.Itoa(time.Now().Minute())+"-"+strconv.Itoa(time.Now().Second()))
		file, _ = os.OpenFile(f, os.O_APPEND|os.O_CREATE|os.O_RDWR|os.O_SYNC, 0666)
	}
	// logger := log.New(file, "", 0)
	// logger.Println(str)
	file.WriteString(str + "\n")

}

func AddLog(str string, param ...interface{}) {
	// g["MSG"] = append(g["MSG"].([]string), str)
	MyLog.Printf(str, param...)
}

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		g := make(G)
		g["starttime"] = time.Now()

		c.Request.ParseForm()                  //赋值所有post, put, get参数到c.RequestForm
		c.Request.ParseMultipartForm(32 << 20) //上传的文件存储在maxMemory大小的内存里面，如果文件大小超过了maxMemory，那么剩下的部分将存储在系统的临时文件中。

		addr := c.Request.Header.Get("X-Real-IP")
		if addr == "" {
			addr = c.Request.Header.Get("X-Forwarded-For")
			if addr == "" {
				addr = c.Request.RemoteAddr
			}
		}
		c.Request.RemoteAddr = addr
		for key, value := range c.Request.Form {
			for i, v := range value {
				value[i], _ = url.QueryUnescape(v)
			}
			c.Request.Form[key] = value
		}
		defer func() {
			if r := recover(); r != nil {
				stack := ierror.Stack(3)

				g["route"] = fmt.Sprintf("Started %s %s for %s %v %s in %v {%v} \n%s", c.Request.Method,
					strings.Replace(c.Request.RequestURI, "%", Replace, -1), c.Request.RemoteAddr,
					500, http.StatusText(500), time.Since(g["starttime"].(time.Time)), c.Request.Form, "%s")
				MyLog.Err(g["route"].(string), fmt.Sprintf("%s\n%s", r, stack))

				if DEBUG {
					body := []byte(fmt.Sprintf(panicHtml, r, r, stack))
					if nil != body {
						c.Data(500, "text/html", []byte(body))
					}
				}
			}
			g["route"] = fmt.Sprintf("Started %s %s for %s %v %s in %v {%v}", c.Request.Method,
				strings.Replace(c.Request.RequestURI, "%", Replace, -1), c.Request.RemoteAddr,
				c.Writer.Status(), http.StatusText(c.Writer.Status()), time.Since(g["starttime"].(time.Time)), c.Request.Form)
			MyLog.Printf(g["route"].(string))
			g = nil
		}()
		c.Next()
	}
}

func init() {
	if LOGO {
		go save()
	}
}
