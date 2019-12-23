package mygoadb

import (
	"bytes"
	"log"
	"os/exec"
	"sync"
)

const (
	KEY_UP   = "up"
	KEY_MV   = "move"
	KEY_DOWN = "down"
)

type ADB struct {
	sync.RWMutex
	cmd   *exec.Cmd
	Path  string
	Args  []string
	debug bool
}

func Command(name string) (adb *ADB) {
	adb = &ADB{
		debug: false,
		Path:  name,
		cmd:   exec.Command(name),
	}
	adb.Args = append([]string{}, adb.cmd.Args...)
	return
}

// Debug 调试切换
func (a *ADB) Debug(f bool) {
	a.debug = f
}

// Use use device with given serial
func (a *ADB) Use(SERIAL string) *ADB {
	adb := &ADB{
		Path: a.Path,
		cmd:  a.cmd,
	}
	adb.Args = []string{a.Path, "-s", SERIAL}
	return adb
}

// Cmd 获取cmd
func (a *ADB) Cmd() *exec.Cmd {
	return a.cmd
}

// Devices
func (a *ADB) Devices() []string {
	b, err := a.Query("devices")
	if a.checkErr(err) {
		return []string{}
	}
	dcb := []byte("device")
	arr := []string{}
	barr := bytes.Split(b, []byte("\n"))
	for i, sb := range barr {
		if i == 0 {
			continue
		}
		if bytes.Contains(sb, dcb) {
			sb = bytes.TrimSpace(bytes.ReplaceAll(sb, dcb, []byte{}))
			arr = append(arr, string(sb))
		}
	}
	return arr
}

// Shell shell
func (a *ADB) ShellQuety(arg ...string) ([]byte, error) {
	return a.Query("shell", arg...)
}

// Query 执行
func (a *ADB) Query(name string, arg ...string) ([]byte, error) {
	args := append([]string{}, a.Args...)
	args = append(args, name)
	if len(arg) > 0 {
		args = append(args, arg...)
	}
	a.cmd.Args = args
	if a.debug {
		log.Println("mygoadb debug:", a.cmd.String())
		return []byte(""), nil
	}
	//a.cmd.Run()
	return a.cmd.Output()
}

func (a *ADB) checkErr(err error) bool {
	if err != nil {
		log.Println(err)
		return true
	}
	return false
}
