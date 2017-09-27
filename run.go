package main

import (
	"log"
	// "path/filepath"
	// "net/http"
	"os"
	"os/exec"
	"strings"
	"time"
)

//守护进程
func main() {
	strun := strings.Join(os.Args[1:], " ")
	cmd := make([][]string, 0)
	for _, j := range strings.Split(strun, "-r") {
		if j != "" {
			stmp := make([]string, 0)
			for _, x := range strings.Split(j, " ") {
				if x != "" {
					stmp = append(stmp, x)
				}
			}
			if len(stmp) > 0 {
				cmd = append(cmd, stmp)
			}
		}
	}
	if len(cmd) <= 0 {
		log.Println("例：go run run.go -r $build -r $server")
		return
	}

	if os.Getppid() != 1 {
		for {
			for _, j := range cmd {
				log.Println(j)
				cmd := exec.Command(j[0], j[1:]...)
				//将其他命令传入生成出的进程
				cmd.Stdin = os.Stdin //给新进程设置文件描述符，可以重定向到文件中
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				cmd.Start() //开始执行新进程，不等待新进程退出
				err := cmd.Wait()
				if err != nil {
					log.Println(err)
					log.Println("sleep 5s")
					time.Sleep(5 * time.Second)
					break
				}
			}
		}
	}
}
