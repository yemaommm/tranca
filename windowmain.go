package main

import (
	"IginServer/conf"
	"IginServer/lib/Imartini"
	_ "IginServer/web/handler"
	"flag"
	"fmt"
	// "github.com/gin-gonic/gin"
	"github.com/howeyc/fsnotify"
	// "github.com/tabalt/gracehttp"
	"net/http"
	"os"
	// "os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

var mode, host *string
var DEBUG = conf.GetBool("config", "DEBUG")
var iswatch = true
var basepath = "src/IginServer"

func watch(p string) {
	path := p
	files := make([]string, 0)
	filepath.Walk(path+basepath, func(filename string, fi os.FileInfo, err error) error { //遍历目录
		if fi.IsDir() { // 忽略目录
			files = append(files, filename)
		}
		return err
	})
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		fmt.Println(err)
	}

	go func() {
		for iswatch {
			select {
			case name := <-watcher.Event:
				fmt.Println("Chanage:", name)
				os.Exit(0)
			case err := <-watcher.Error:
				fmt.Println(err)
			}
		}
	}()
	// watcher.Watch(p)
	for _, i := range files {
		watcher.Watch(i)
	}
}

func watchconfig(p, exe string) {
	path := p
	files := make([]string, 0)
	filepath.Walk(path, func(filename string, fi os.FileInfo, err error) error { //遍历目录
		names := strings.Split(filename, ".")
		if fi.IsDir() || names[len(names)-1] == exe { // 忽略目录
			files = append(files, filename)
		}
		return err
	})
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		fmt.Println(err)
	}

	go func() {
		for {
			select {
			case name := <-watcher.Event:
				fmt.Println("Chanage:", name)
				if exe == "ini" {
					conf.Update()
				} else if exe == "xml" {
					conf.UpdateXML()
				}
			case err := <-watcher.Error:
				fmt.Println(err)
			}
		}
	}()
	// watcher.Watch(p)
	for _, i := range files {
		watcher.Watch(i)
	}
}

func main() {
	host = flag.String("host", ":8005", "启动服务器的端口号")
	mode = flag.String("mode", "EXE", "SET MODE EXE OR SCRIPT")
	flag.Parse()

	var MULTICORE int = runtime.NumCPU() * 4 //number of core
	runtime.GOMAXPROCS(MULTICORE)            //running in multicore

	router := Imartini.ServerInit()
	// gin.DefaultWriter = os.Stdout
	// gin.DefaultErrorWriter = os.Stderr
	// router.Use(gin.Logger(), gin.Recovery())

	if DEBUG {
		watch("./")
		watchconfig("./config", "ini")
		watchconfig("./XML", "xml")
	}

	//监听端口
	s := &http.Server{
		Addr:           *host,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
