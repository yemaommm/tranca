package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func CopyFile(src, dst string) (w int64, err error) {
	srcFile, err := os.Open(src)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer srcFile.Close()

	dstFile, err := os.Create(dst)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer dstFile.Close()

	return io.Copy(dstFile, srcFile)
}

func walk(dir string) string {
	path := dir
	files := make([]string, 0)
	files = append(files, "package handler")
	dirnum := make(map[string]int)
	dirlist := make([]string, 0)
	filepath.Walk("./src/"+path, func(filename string, fi os.FileInfo, err error) error { //遍历目录
		if err != nil { //忽略错误
			return err
		}

		if !fi.IsDir() && filename[len(filename)-3:len(filename)] == ".go" { //判断是否文件
			f := strings.SplitAfter(filename, path)[1]
			if f != "" {
				names := strings.Split(filename, "/")
				name := strings.Join(names[0:len(names)-1], "/")
				if len(strings.SplitAfter(name, path)) <= 1 {
					return err
				}
				// fmt.Println(name)
				// fmt.Println(filename[len(filename)-3 : len(filename)])
				name = strings.SplitAfter(name, path)[1]
				if dirnum[name] == 0 {
					// dirnum[name] = 0
					dirlist = append(dirlist, name)
				}
				dirnum[name]++
			}
		}

		// if fi.IsDir() { // 忽略目录
		// 	f := strings.SplitAfter(filename, path)[1]
		// 	if f != "" {
		// 		files = append(files, "import _ \""+path+f+"\"")
		// 	}
		// }
		return err
	})
	for _, i := range dirlist {
		if dirnum[i] != 0 {
			files = append(files, "import _ \""+path+i+"\"")
		}
	}
	// fmt.Println(strings.Join(files, "\r\n"))
	return strings.Join(files, "\r\n")
}

var paths = []string{
	"./log",
	"./upload",
	"./config",
	"./public",
	"./templates",
	"./src/IginServer/web/handler",
	"./src/IginServer/web/model",
}

var buildpath = "IginServer/web/handler/"

func main() {
	arg := os.Args

	if len(arg) > 1 {
		switch arg[1] {
		case "bak":
			os.RemoveAll("bak")
			for _, i := range paths {
				os.MkdirAll("bak/"+i, 0666)
				filepath.Walk(i, func(filename string, fi os.FileInfo, err error) error { //遍历目录

					if !fi.IsDir() { // 忽略目录
						CopyFile(filename, "bak/"+filename)
					} else {
						os.MkdirAll("bak/"+filename, 0666)
					}
					return err
				})
			}
		case "clean":
			os.RemoveAll("bak")
			for _, i := range paths {
				os.MkdirAll("bak/"+i, 0666)
				filepath.Walk(i, func(filename string, fi os.FileInfo, err error) error { //遍历目录

					if !fi.IsDir() { // 忽略目录
						CopyFile(filename, "bak/"+filename)
					} else {
						os.MkdirAll("bak/"+filename, 0666)
					}
					return err
				})
				os.RemoveAll(i)
				os.MkdirAll(i, 066)
			}
			CopyFile("bak/config/config.ini", "config/config.ini")
		default:
		}
	} else {
		path := buildpath
		str := walk(path)

		filepath := "./src/" + path + "init.go"
		os.Remove(filepath)
		f, _ := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_RDWR|os.O_SYNC, 0666)
		defer f.Close()
		f.WriteString(str)
	}
}
