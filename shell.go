package mygoadb

import (
	"bytes"
	"io/ioutil"
)

// NewCmdShell new CmdShell
func NewCmdShell(a *ADB) *CmdShell {
	s := &CmdShell{
		a: a,
	}
	s.Input = NewCmdInput(s)
	s.UI = NewUI(s)
	return s
}

// CmdShell shell
type CmdShell struct {
	a     *ADB
	Input InputI
	UI    UII
}

// Query 执行 shell cmd
func (s *CmdShell) Query(name string, arg ...string) ([]byte, error) {
	args := append([]string{name}, arg...)
	return s.a.Query("shell", args...)
}

// Screencap shell screencap
func (s *CmdShell) Screencap(imgpath string) error {
	b, err := s.ScreencapByte()
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(imgpath, b, 0644)
	//_, err = s.Query("'screencap -p > " + imgpath + "'")
	return err
}

// ScreencapByte shell screencap to byte
func (s *CmdShell) ScreencapByte() (b []byte, err error) {
	b, err = s.Query("screencap", "-p")
	b = bytes.ReplaceAll(b, []byte("\r\r\n"), []byte("\n"))
	return
}

// StartApp start app
func (s *CmdShell) StartApp(app string) (err error) {
	_, err = s.Query("monkey -p " + app + " -c android.intent.category.LAUNCHER 1")
	return
}

// KillApp kill app
func (s *CmdShell) KillApp(app string) (err error) {
	_, err = s.Query("am force-stop " + app)
	return
}
