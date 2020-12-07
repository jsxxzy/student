// Author: d1y<chenhonzhou@gmail.com>

package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

var killCmd = "ntsd.exe"

var studentAppname = "studentmain.exe"

var studentPath = filepath.Join("c://", "Program Files (x86)", "Mythware", "极域电子教室软件 v4.0 2016 豪华版", "StudentMain.exe")

func easyKill() error {
	var args = []string{
		"-c", "q", "-pn", studentAppname,
	}
	var cmd = exec.Command(killCmd, args...)
	return cmd.Run()
}

func easyStart() error {
	fmt.Println(studentPath)
	var cacheCmd = exec.Command(studentPath)
	cacheCmd.Run()
	return nil
}

func main() {
	var args = os.Args
	if len(args) <= 1 {
		fmt.Println("use: kill/run")
		return
	}
	var cmd = args[1]
	switch cmd {
	case "run":
		easyStart()
		return
	case "kill":
		var flag = easyKill() == nil
		var msg = "查找进程失败"
		if flag {
			msg = "杀死进程成功"
		}
		fmt.Println(msg)
		return
	}
}

// check file exists
func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// init `bin` data file
func initBinFile() bool {
	flag, err := pathExists(killCmd)
	if err != nil {
		panic(err)
	}
	if !flag {
		return RestoreAsset(".", killCmd) == nil
	}
	return true
}

func init() {
	initBinFile()
}
