package main

import (
	"IginServer/conf"
	_ "IginServer/framework"
	"IginServer/lib/Imartini"
	_ "IginServer/web/handler"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/howeyc/fsnotify"
	"github.com/tabalt/gracehttp"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

var mode, host *string
var Server *gracehttp.Server
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
				if strings.Index(name.Name, "web/handler/init.go") == -1 {
					iswatch = false
					fmt.Println("Chanage:", name)
					iswatch = !UpdateSystem()
					// os.Exit(0)
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

	router.GET("/my/UpdateSystem", func(c *gin.Context) {
		if strings.Split(c.Request.RemoteAddr, ":")[0] == "127.0.0.1" {
			if UpdateSystem() {
				c.String(200, "TRUE")
				return
			}
		}
		c.String(200, "FALSE")
	})

	//监听端口
	s := &http.Server{
		Addr:           *host,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	Server = gracehttp.NewHttpServer(s, DEBUG)
	Server.ListenAndServe()
}

func UpdateSystem() bool {
	if *mode == "SCRIPT" {
		var cmdrun []string
		fmt.Println("Start build file....")

		cmdrun = []string{"go", "run", "build.go"}
		if !RunCmd(cmdrun[0], cmdrun[1:]...) {
			fmt.Println("Start build FAIL")
			return false
		}

		cmdrun = []string{"go", "build", "main.go"}
		if !RunCmd(cmdrun[0], cmdrun[1:]...) {
			fmt.Println("Start build FAIL")
			return false
		}

		fmt.Println("Start build OK")
	}
	Server.StartNewProcess()
	Server.Close()
	return true
}

func RunCmd(name string, m ...string) bool {
	cmd := exec.Command(name, m...)
	//将其他命令传入生成出的进程
	cmd.Stdin = os.Stdin //给新进程设置文件描述符，可以重定向到文件中
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Start() //开始执行新进程，不等待新进程退出
	err := cmd.Wait()
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
