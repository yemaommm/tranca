package image

import (
	"fmt"
	"net/http"
	// "os"
	"IginServer/framework/API/R"
	"github.com/gin-gonic/gin"
	"strconv"
	// "strings"
	"encoding/json"
	// "github.com/go-martini/martini"
	// "encoding/xml"
	// "io/ioutil"
	"IginServer/conf"
	// "IginServer/lib/Imartini"
	"IginServer/lib/image"
	// "IginServer/lib/redis"
	"IginServer/lib/upload"
)

var imgtoken []string

func uploadfile(c *gin.Context) {
	if c.Request.FormValue("token") == "" || c.Request.FormValue("title") == "" {
		R.Api404(c.Writer)
		return
	}
	token := c.Request.FormValue("token")
	title := c.Request.FormValue("title")
	for _, ktoken := range imgtoken {
		if ktoken == token {
			path, err := upload.Save(c.Request, "upload", token, title)
			if err != nil {
				R.Error(c.Writer, 300, err.Error())
				return
			}
			R.Success(c.Writer, map[string]interface{}{
				"path": conf.GetString("config", "HOST_IginSERVER") + "M/" + path,
			})
			return
		}
	}
	R.Error(c.Writer, 400, "token不正确")
}

func postimg(c *gin.Context) {
	if c.Request.FormValue("token") == "" || c.Request.FormValue("title") == "" {
		R.Api404(c.Writer)
		return
	}
	width, _ := strconv.Atoi(c.Request.FormValue("width"))
	height, _ := strconv.Atoi(c.Request.FormValue("height"))
	if width == 0 && height == 0 {
		R.Api404(c.Writer)
		return
	}
	token := c.Request.FormValue("token")
	title := c.Request.FormValue("title")
	for _, ktoken := range imgtoken {
		if ktoken == token {
			path, err := upload.ImgSave(c.Request, "imagedata", width, height, token, title)

			if err != nil {
				fmt.Println(err)
				R.Error(c.Writer, 300, "图片不正确")
				return
			}
			R.Success(c.Writer, map[string]interface{}{
				"path": conf.GetString("config", "HOST_IginSERVER") + "M/" + path,
			})
			return
		}
	}
	R.Error(c.Writer, 400, "token不正确")
}

func catimg(c *gin.Context) {
	width, _ := strconv.Atoi(c.Request.FormValue("width"))
	height, _ := strconv.Atoi(c.Request.FormValue("height"))

	if width == 0 && height == 0 {
		http.StripPrefix("/M/upload/", http.FileServer(http.Dir(conf.GetString("config", "UPLOAD_PATH")))).ServeHTTP(c.Writer, c.Request)
	} else {
		path := c.Param("path")
		path = conf.GetString("config", "UPLOAD_PATH") + "/" + path

		image.ThumbnailHttp(path, c.Writer, width, height)
	}
}

func init() {
	json.Unmarshal([]byte(conf.GetString("image", "IMAGE_TOKEN")), &imgtoken)
}
