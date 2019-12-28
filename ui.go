package mygoadb

import (
	"bytes"
	"io/ioutil"

	"github.com/PuerkitoBio/goquery"
)

// UII shell ui interface
type UII interface {
	GetUIString() ([]byte, error)
	SaveTo(filePath string) error
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
func (i *CmdUI) SaveTo(filePath string) (err error) {
	var b []byte
	b, err = i.GetUIString()
	if err != nil {
		return
	}
	return ioutil.WriteFile(filePath, b, 0644)
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
