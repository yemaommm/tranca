package conf

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type XMLlog struct {
	XMLName   xml.Name  `xml:"mapper"`
	Namespace string    `xml:"namespace,attr"`
	Svs       []context `xml:"string"`
}

type context struct {
	Id          string `xml:"id,attr"`
	Description string `xml:",innerxml"`
}

var GET map[string]map[string]string
var MyStr map[string]string

func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func Walk(dir string, f func(filename string, fi os.FileInfo, err error)) {
	path := dir
	if !PathExists(path) {
		return
	}
	filepath.Walk(path, func(filename string, fi os.FileInfo, err error) error { //遍历目录
		if err != nil { //忽略错误
			fmt.Println(err)
			return err
		}

		if !fi.IsDir() { // 忽略目录
			f(filename, fi, err)
		}
		return err
	})
}

func Update() {
	Walk("./config", func(filename string, fi os.FileInfo, err error) {
		attr := strings.Split(filename, ".")
		if len(attr) != 2 || attr[1] != "ini" {
			return
		}
		fmt.Println("logfile: " + filename)
		conf := SetConfig(filename)
		list := conf.ReadList()
		// c := list[0]["config"]

		for _, j := range list {
			for x, y := range j {
				GET[x] = y
			}
		}
	})
}

func UpdateXML() {
	Walk("./XML", func(filename string, fi os.FileInfo, err error) {
		attr := strings.Split(filename, ".")
		if len(attr) != 2 || strings.ToLower(attr[1]) != "xml" {
			return
		}
		fmt.Println("xmlfile: " + filename)
		content, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Println(err)
			return
		}
		v := XMLlog{}
		if xml.Unmarshal(content, &v) == nil {
			for _, j := range v.Svs {
				j.Description = strings.Replace(j.Description, "&lt;", "<", -1)
				j.Description = strings.Replace(j.Description, "&gt;", ">", -1)
				j.Description = strings.Replace(j.Description, "&amp;", "&", -1)
				MyStr[v.Namespace+"."+j.Id] = j.Description
			}
		} else {
			fmt.Println("error:", filename)
		}

	})
}

func GetString(key, v string) string {
	if GET[key] == nil {
		return ""
	}
	return GET[key][v]
}

func GetInt(key, v string) int {
	if GET[key] == nil {
		return 0
	}
	i, _ := strconv.Atoi(GET[key][v])
	return i
}

func GetBool(key, v string) bool {
	if GET[key] == nil {
		return false
	}
	i, _ := strconv.ParseBool(GET[key][v])
	return i
}

func XmlGet(key string) string {
	return MyStr[key]
}

func init() {
	GET = make(map[string]map[string]string)
	MyStr = make(map[string]string)
	Update()
	UpdateXML()
}
