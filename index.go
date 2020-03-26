package mygoadb

import (
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"os/exec"
	"strings"
	"sync"
	"syscall"
)

// ADB  mygoadb
type ADB struct {
	sync.RWMutex
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
		TmpDir: "/data/local/tmp",
		Args:   []string{},
	}
	adb.Shell = NewCmdShell(adb)
	adb.ExecOut = NewCmdExecOut(adb)

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

// IsDebug  is Debug
func (a *ADB) IsDebug() bool {
	return a.debug
}

// Use use device with given serial
func (a *ADB) Use(SERIAL string) *ADB {
	adb := &ADB{
		debug:  a.debug,
		TmpDir: a.TmpDir,
		Path:   a.Path,
		Args:   append([]string{}, "-s", SERIAL),
	}
	adb.Shell = NewCmdShell(adb)
	adb.ExecOut = NewCmdExecOut(adb)

	return adb
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

func (a *ADB) Query(parts string, arg ...string) (b []byte, err error) {
	a.Lock()
	defer a.Unlock()
	args := append([]string{}, a.Args...)
	args = append(args, parts)
	if len(arg) > 0 {
		args = append(args, arg...)
	}

	if a.debug {
		cmdstr := strings.Join(args, " ")
		log.Println("mygoadb debug:", cmdstr)
		return []byte(cmdstr), errors.New(cmdstr)
	}
	adbExec := exec.Command(a.Path, args...)
	adbExec.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	in, _ := adbExec.StdinPipe()
	errorOut, _ := adbExec.StderrPipe()
	out, _ := adbExec.StdoutPipe()
	defer closeIO(in)
	defer closeIO(errorOut)
	defer closeIO(out)

	if err := adbExec.Start(); err != nil {
		return []byte(""), errors.New("start adb process error")
	}

	outData, _ := ioutil.ReadAll(out)
	errorData, _ := ioutil.ReadAll(errorOut)

	var adbError error = nil

	if err := adbExec.Wait(); err != nil {
		if _, ok := err.(*exec.ExitError); ok {
			adbError = errors.New("adb return error")
			outData = errorData
		} else {
			return []byte(""), errors.New("start adb process error")
		}
	}
	return outData, adbError
}

// close a stream
func closeIO(c io.Closer) {
	if c != nil {
		c.Close()
	}
}

// Query 执行
func (a *ADB) Query1(parts string, arg ...string) (b []byte, err error) {
	a.Lock()
	defer a.Unlock()
	args := append([]string{}, a.Args...)
	args = append(args, parts)
	if len(arg) > 0 {
		args = append(args, arg...)
	}

	if a.debug {
		cmdstr := strings.Join(args, " ")
		log.Println("mygoadb debug:", cmdstr)
		return []byte(cmdstr), errors.New(cmdstr)
	}
	cmd := exec.Command(a.Path, args...)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	return cmd.Output()
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
