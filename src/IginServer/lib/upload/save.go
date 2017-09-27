package upload

import (
	"IginServer/conf"
	"IginServer/lib/image"
	"crypto/md5"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

var img_type = conf.GET["config"]["IMAGE_TYPE"]

func Getpath() string {
	d := time.Now()
	filename := strconv.Itoa(int(d.UnixNano()))
	filename = fmt.Sprintf("%x", md5.Sum([]byte(filename)))

	y := strconv.Itoa(d.Year())
	m := d.Month().String()
	da := strconv.Itoa(d.Day())
	path := y + "/" + m + "/" + da + "/" + filename
	return path
}

func Makedir(path string) error {
	p := strings.Join(strings.Split(path, "/")[0:len(strings.Split(path, "/"))-1], "/")
	err := os.MkdirAll(p, 0777)
	return err
}

func SaveFile(b []byte, param ...string) (string, error) {
	var path string
	if len(param) == 1 {
		path = conf.GET["config"]["UPLOAD_PATH"] + "/" + param[0] + "/" + Getpath()
	} else if len(param) > 1 {
		path = conf.GET["config"]["UPLOAD_PATH"] + "/" + param[0] + "/" + Getpath() + param[len(param)-1]
	} else {
		path = conf.GET["config"]["UPLOAD_PATH"] + "/" + Getpath()
	}
	err := Makedir(path)
	if err != nil {
		fmt.Printf("%s", err)
	}
	err = ioutil.WriteFile(path, b, 0666)
	return path, err
}

func Save(req *http.Request, key string, param ...string) (string, error) {
	f, fh, err := req.FormFile(key)
	if err != nil {
		return "", err
	}
	stmp := strings.Split(fh.Filename, ".")
	fi := ""
	if len(stmp) > 1 {
		fi = "." + stmp[len(stmp)-1]
	}
	u := &Upload{f, fh}

	var path string
	if len(param) > 0 {
		path = conf.GET["config"]["UPLOAD_PATH"] + "/" + strings.Join(param, "/") + "/" + Getpath() + fi
	} else {
		path = conf.GET["config"]["UPLOAD_PATH"] + "/" + Getpath() + fi
	}
	u.save(path)

	return path, err
}

func ImgSave(req *http.Request, key string, w, h int, param ...string) (string, error) {
	_, fh, err := req.FormFile(key)
	if err != nil {
		return "", err
	}
	stmp := strings.Split(fh.Filename, ".")
	ty := stmp[len(stmp)-1]
	if len(stmp) > 1 && strings.Index(img_type, "|"+strings.ToLower(ty)+"|") == -1 {
		return "", errors.New("-1")
	}

	path, err := Save(req, key, param...)
	if err != nil {
		return path, err
	}
	if w != 0 && h != 0 {
		err = image.Thumbnail(path, path, w, h)
	}
	return path, err
}

func ImgBase64Save(req *http.Request, w, h int, param ...string) (string, error) {
	img, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return "", err
	}
	limg := strings.Split(string(img), ";")
	ty := ""
	for _, j := range limg[0 : len(limg)-1] {
		s := strings.Split(j, ":")
		if s[0] == "data" {
			ty = s[1]
		}
	}
	data := strings.Split(limg[len(limg)-1], "base64,")[1]
	b, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return "", err
	}
	if ty == "image/jpeg" || ty == "image/jpg" {
		param = append(param, ".jpg")
	} else if ty == "image/png" {
		param = append(param, ".png")
	} else if ty == "image/gif" {
		param = append(param, ".gif")
	} else if ty == "image/bmp" {
		param = append(param, ".bmp")
	}
	if strings.Index(img_type, "|"+strings.Split(ty, "/")[1]+"|") == -1 {
		return "", errors.New("-1")
	}
	path, err := SaveFile(b, param...)
	if err != nil {
		return path, err
	}
	err = image.Thumbnail(path, path, w, h)
	return path, err
}

type Upload struct {
	File       multipart.File
	FileHeader *multipart.FileHeader
}

func (u *Upload) save(path string) {
	err := Makedir(path)
	if err != nil {
		fmt.Printf("%s", err)
	}
	fw, err2 := os.Create(path)
	if err != nil {
		fmt.Println(err2)
	}
	_, err = io.Copy(fw, u.File)
	if err != nil {
		fmt.Println(err)
	}

}

func init() {
	// err := json.Unmarshal([]byte(conf.GET["config"]["IMAGE_TYPE"]), &img_type)
	// if err != nil {
	// 	fmt.Println(err)
	// }
}
