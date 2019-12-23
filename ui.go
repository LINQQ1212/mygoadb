package mygoadb

import (
	"bytes"

	"github.com/PuerkitoBio/goquery"
)

// UII shell ui interface
type UII interface {
	GetUIString() ([]byte, error)
	SaveTo(dirname string) ([]byte, error)
	GetUIToDoc() (*goquery.Document, error)
}

// NewUI shell
func NewUI(s *CmdShell) *CmdUI {
	u := &CmdUI{
		s:       s,
		tmpPath: s.a.TmpDir + "/ui.xml",
	}
	return u
}

/*
adb -s Y15MFBP8236XF shell uiautomator dump /sdcard/ui.xml
adb -s Y15MFBP8236XF pull /sdcard/ui.xml C:/Users/Administrator/Desktop
adb -s Y15MFBP8236XF shell cat /sdcard/ui.xml
*/
// CmdUI shell
type CmdUI struct {
	s       *CmdShell
	tmpPath string
}

func (i *CmdUI) dumpui() {
	i.s.Query("uiautomator", "dump", i.tmpPath)
}

// GetUIString get UI xml to string
func (i *CmdUI) GetUIString() ([]byte, error) {
	i.dumpui()
	return i.s.Query("cat", i.tmpPath)
}

// SaveTo pull file to you computer dir
func (i *CmdUI) SaveTo(dirname string) ([]byte, error) {
	i.dumpui()
	return i.s.Query("pull", i.tmpPath, dirname)
}

// GetUIToDoc get ui to goquery.Document
func (i *CmdUI) GetUIToDoc() (*goquery.Document, error) {
	i.dumpui()
	b, err := i.s.Query("cat", i.tmpPath)
	if err != nil {
		return nil, err
	}
	return goquery.NewDocumentFromReader(bytes.NewReader(b))
}
