package mygoadb

import (
	"bytes"
	"log"
	"os/exec"
	"strings"
	"sync"
)

// ADB  mygoadb
type ADB struct {
	sync.RWMutex
	cmd     *exec.Cmd
	Path    string
	Args    []string
	debug   bool
	Shell   *CmdShell
	ExecOut *CmdExecOut
	// TmpDir android tmp dir
	TmpDir string
}

// Command 调用 mygoadb
func Command(name string) (adb *ADB) {
	adb = &ADB{
		debug:  false,
		Path:   name,
		cmd:    exec.Command(name),
		TmpDir: "/data/local/tmp",
	}
	adb.Shell = NewCmdShell(adb)
	adb.ExecOut = NewCmdExecOut(adb)
	adb.Args = append([]string{}, adb.cmd.Args...)
	return
}

// SetTmpDir set android tmp dir
func (a *ADB) SetTmpDir(dir string) {
	a.TmpDir = dir
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

// Devices 查询 Devices
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

// Query 执行
func (a *ADB) Query(parts string, arg ...string) ([]byte, error) {
	a.Lock()
	defer a.Unlock()
	args := append([]string{}, a.Args...)
	args = append(args, parts)
	if len(arg) > 0 {
		args = append(args, arg...)
	}
	a.cmd.Args = args
	if a.debug {
		log.Println("mygoadb debug:", strings.Join(a.cmd.Args, " "))
		return []byte(""), nil
	}
	//a.cmd.Run()
	return a.cmd.Output()
}

//UnInstallApp UnInstall App
func (a *ADB) UnInstallApp(packages string) ([]byte, error) {
	return a.Query("uninstall", packages)
}

//InstallApp Install App  If the part up is true, use -r
func (a *ADB) InstallApp(appPath string, up bool) ([]byte, error) {
	if up {
		return a.Query("install", "-r", appPath)
	}
	return a.Query("install", appPath)
}

func (a *ADB) checkErr(err error) bool {
	if err != nil {
		log.Println(err)
		return true
	}
	return false
}