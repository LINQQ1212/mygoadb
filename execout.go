package mygoadb

import "io/ioutil"

// NewCmdExecOut  exec-out
func NewCmdExecOut(a *ADB) *CmdExecOut {
	s := &CmdExecOut{
		a: a,
	}
	return s
}

// CmdExecOut shell
type CmdExecOut struct {
	a *ADB
}

// Query 执行 shell cmd
func (s *CmdExecOut) Query(name string, arg ...string) ([]byte, error) {
	args := append([]string{name}, arg...)
	return s.a.Query("exec-out", args...)
}

func (s *CmdExecOut) Screencap(imgpath string) error {
	b, err := s.ScreencapByte()
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(imgpath, b, 0644)
	return err
}

// ScreencapByte shell screencap to byte
func (s *CmdExecOut) ScreencapByte() ([]byte, error) {
	return s.Query("screencap", "-p")
}
