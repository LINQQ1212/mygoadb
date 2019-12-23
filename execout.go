package mygoadb

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
	return s.a.Query("exec-out", arg...)
}

// Screencap shell screencap
func (s *CmdExecOut) Screencap(imgpath string) error {
	_, err := s.Query("screencap", "-p", ">", imgpath)
	return err
}

// ScreencapByte shell screencap to byte
func (s *CmdExecOut) ScreencapByte() ([]byte, error) {
	return s.Query("screencap", "-p")
}
