package ueditor

import (
	// "fmt"
	// "regexp"
	"encoding/base64"
	"encoding/json"
	// "github.com/martini-contrib/render"
	"IginServer/conf"
	"IginServer/lib/Imartini"
	// "IginServer/lib/session"
	"IginServer/lib/upload"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	// "net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

/*
   {"", "GET", "/ueditor.config.js", ueditor.Ueditor_config},
   {"", "GET", "/controller.php", ueditor.Controller},
   {"", "POST", "/controller.php", ueditor.Controller},
*/

var config []byte
var defaultw, _ = strconv.Atoi(conf.GET["app"]["ueditorimgW"])
var defaulth, _ = strconv.Atoi(conf.GET["app"]["ueditorimgH"])
var uploadpath []string

func Ueditor_config(c *gin.Context) {
	url := c.Request.RequestURI
	iurl := strings.Split(url, "/")
	iurl = iurl[0 : len(iurl)-1]
	url = strings.Join(iurl, "/")

	c.Writer.Header().Set("Content-Type", "application/javascript; charset=utf-8")
	Imartini.HTML(c.Writer, c.Request, "framework/ueditor/ueditor_config.js", map[string]interface{}{
		"url": url + "/controller.php?path=" + c.Request.FormValue("path"),
	})
}

func Controller(c *gin.Context) {
	path := "-1"
	for _, i := range uploadpath {
		if c.Request.FormValue("path") == i {
			path = c.Request.FormValue("path")
		}
	}
	if path == "-1" {
		c.Writer.Write([]byte("{\"state\": \"参数有误\"}"))
		return
	}
	// info := session.GetSession(c.Writer, c.Request)

	action := c.Request.FormValue("action")
	var result []byte

	switch action {
	case "config":
		result = config //json_encode($CONFIG);
	/* 上传图片 */
	case "uploadimage":
		path, err := upload.ImgSave(c.Request, "upfile", defaultw, defaulth, path+"/img")
		result, err = resu(path, err)
		if err != nil {
			result = []byte("{\"state\": \"" + err.Error() + "\"}")
		}
	/* 上传涂鸦 */
	case "uploadscrawl":
		crawl := c.Request.FormValue("upfile")
		if crawl != "" {
			b, err := base64.StdEncoding.DecodeString(crawl)
			if err != nil {
				c.Writer.Write([]byte("{\"state\": \"" + err.Error() + "\"}"))
				return
			}
			path, err := upload.SaveFile(b, path+"/img")
			if err != nil {
				c.Writer.Write([]byte("{\"state\": \"" + err.Error() + "\"}"))
				return
			}
			m := make(map[string]interface{})
			m["url"] = "/" + path
			m["state"] = "SUCCESS"
			result, _ = json.Marshal(m)
		} else {
			path, err := upload.ImgBase64Save(c.Request, defaultw, defaulth, path+"/img")
			if err != nil {
				c.Writer.Write([]byte("{\"state\": \"" + err.Error() + "\"}"))
				return
			}
			m := make(map[string]interface{})
			m["url"] = "/" + path
			m["name"] = strings.Split(path, "/")[len(strings.Split(path, "/"))-1]
			m["type"] = strings.Split(path, ".")[len(strings.Split(path, "."))-1]
			m["state"] = "SUCCESS"
			result, _ = json.Marshal(m)
		}
	/* 上传视频 */
	case "uploadvideo":
		path, err := upload.Save(c.Request, "upfile", path+"/video")
		result, err = resu(path, err)
		if err != nil {
			result = []byte("{\"state\": \"" + err.Error() + "\"}")
		}
	/* 上传文件 */
	case "uploadfile":
		path, err := upload.Save(c.Request, "upfile", path+"/file")
		result, err = resu(path, err)
		if err != nil {
			// return r.JSON(200, map[string]interface{}{"state": err.Error()})
			result = []byte("{\"state\": \"" + err.Error() + "\"}")
		}
	/* 列出图片 */
	case "listimage":
		size := c.Request.FormValue("start")
		if size != "0" {
			c.Writer.Write([]byte("{\"state\": \"null\"}"))
			return
		}
		uploadpath := conf.GET["config"]["UPLOAD_PATH"] + "/" + path + "/img"
		stmp, err := WalkDir(uploadpath, "")
		if err != nil {
			result = []byte("{\"state\": \"" + err.Error() + "\"}")
		} else {
			result, _ = json.Marshal(map[string]interface{}{"state": "SUCCESS", "list": stmp})
		}
	/* 列出文件 */
	case "listfile":
		size := c.Request.FormValue("start")
		if size != "0" {
			c.Writer.Write([]byte("{\"state\": \"null\"}"))
			return
		}
		uploadpath := conf.GET["config"]["UPLOAD_PATH"] + "/" + path + "/file"
		stmp, err := WalkDir(uploadpath, "")
		if err != nil {
			result = []byte("{\"state\": \"" + err.Error() + "\"}")
		} else {
			result, _ = json.Marshal(map[string]interface{}{"state": "SUCCESS", "list": stmp})
		}
	/* 抓取远程文件 */
	case "catchimage":
		// $result = include("action_crawler.php");
	case "uploaddel":
		pupload := strings.Split(conf.GET["config"]["UPLOAD_PATH"], "/"+path+"/")
		url := strings.Split(c.Request.FormValue("url"), "/")
		isremove := false
		path := make([]string, 0)
		for i, j := range url {
			if isremove {
				path = append(path, j)
			} else {
				for x, y := range pupload {
					if url[i+x] == y {
						isremove = true
					} else {
						isremove = false
					}
				}
				if isremove {
					path = append(path, j)
				}
			}
		}
		err := os.Remove(strings.Join(path, "/"))
		Imartini.MyLog.Printf("%v:%v", path, err)
	default:
		result = []byte("{\"state\": \"请求地址出错\"}")
	}
	/* 输出结果 */
	// if (isset($_GET["callback"])) {
	//     if (preg_match("/^[\w_]+$/", $_GET["callback"])) {
	//         echo htmlspecialchars($_GET["callback"]) . '(' . $result . ')';
	//     } else {
	//         echo json_encode(array(
	//             'state'=> 'callback参数不合法'
	//         ));
	//     }
	// } else {
	//     echo $result;
	// }
	c.Writer.Write(result)
}

//获取指定目录及所有子目录下的所有文件，可以匹配后缀过滤。
func WalkDir(dirPth, suffix string) (files []map[string]interface{}, err error) {
	files = make([]map[string]interface{}, 0)
	// files = make([]string, 0, 30)
	// suffix = strings.ToUpper(suffix) //忽略后缀匹配的大小写

	err = filepath.Walk(dirPth, func(filename string, fi os.FileInfo, err error) error { //遍历目录
		if err != nil { //忽略错误
			return err
		}
		if fi.IsDir() { // 忽略目录
			return nil
		}
		// 匹配后缀名
		// if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) {
		// 	files = append(files, filename)
		// }
		// files = append(files, filename)
		files = append(files, map[string]interface{}{"url": conf.GetString("config", "IMAGE_HOST") + filename})
		return nil
	})
	// return files, err
	return files, err
}

func resu(path string, err error) ([]byte, error) {
	if err != nil {
		return nil, err
	}
	m := make(map[string]interface{})
	m["url"] = conf.GetString("config", "IMAGE_HOST") + path
	m["original"] = strings.Split(path, "/")[len(strings.Split(path, "/"))-1]
	m["name"] = strings.Split(path, "/")[len(strings.Split(path, "/"))-1]
	m["type"] = strings.Split(path, ".")[len(strings.Split(path, "."))-1]
	m["state"] = "SUCCESS"
	return json.Marshal(m)
}

func init() {
	f, _ := os.Open("public/ueditor/config.json")
	defer f.Close()
	config, _ = ioutil.ReadAll(f)
	if conf.GET["ueditor"]["uploadpath"] != "" {
		json.Unmarshal([]byte(conf.GET["ueditor"]["uploadpath"]), &uploadpath)
	}
}
