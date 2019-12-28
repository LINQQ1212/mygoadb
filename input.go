package mygoadb

import (
	"fmt"
	"math/rand"
	"time"
)

// KEYTYPE 类型
type KEYTYPE string

const (
	// KeyUp key up
	KeyUp KEYTYPE = "up"
	// KeyMv key move
	KeyMv KEYTYPE = "move"
	// KeyDown  key down
	KeyDown KEYTYPE = "down"
)

// InputI input interface
type InputI interface {
	Query(arg ...string) ([]byte, error)
	Click(x, y int) (err error)
	Swipe(x, y, x1, y1, dtime int) (err error)
	SwipeRandom(x, y, x1, y1, dtime, r int) (err error)
	Text(str string) (err error)
	KeyEvent(key string, longpress bool) (err error)

	KeyHome() error
	KeyBack() error
	KeySwitch() error
	KeyPower() error
	KeyMenu() error
}

// NewCmdInput new CmdInput
func NewCmdInput(s *CmdShell) *CmdInput {
	i := CmdInput{
		s: s,
	}
	return &i
}

// CmdInput shell input bind
type CmdInput struct {
	s *CmdShell
}

// Query Input Query
func (i *CmdInput) Query(arg ...string) ([]byte, error) {
	return i.s.Query("input", arg...)
}

// Query Input Query
func (i *CmdInput) query2(arg ...string) (err error) {
	_, err = i.s.Query("input", arg...)
	return
}

// Click input tap
func (i *CmdInput) Click(x, y int) error {
	return i.query2("tap", fmt.Sprintf("%d %d", x, y))
}

// Swipe input Swipe
func (i *CmdInput) Swipe(x, y, x1, y1, dtime int) error {
	if dtime < 0 {
		dtime = 300
	}
	dur := time.Duration(dtime) * time.Millisecond

	return i.query2("swipe", fmt.Sprintf("%d %d %d %d %d", x, y, x1, y1, dur.Milliseconds()))
}

// SwipeRandom input SwipeRandom
func (i *CmdInput) SwipeRandom(x, y, x1, y1, dtime, r int) error {

	return i.Swipe(i.random(x, r), i.random(y, r), i.random(x1, r), i.random(y1, r), dtime)
	//return i.query2("swipe", fmt.Sprintf("%d %d %d %d %d", i.random(x, r), i.random(y, r), i.random(x1, r), i.random(y1, r), dur.Milliseconds()))
}

func (i *CmdInput) random(max, min int) int {
	rand.Seed(time.Now().UnixNano())
	o := max - min
	if o > 0 {
		time.Sleep(100 * time.Microsecond)
		return rand.Intn(o) + min
	}
	return max + min
}

// Text input text
func (i *CmdInput) Text(str string) error {
	return i.query2("text", str)
}

// KeyEvent input keyevent
// key KEYCODE
//longpress 是否长按
func (i *CmdInput) KeyEvent(key string, longpress bool) error {
	if longpress {
		return i.query2("keyevent", "--longpress", key)
	}
	return i.query2("keyevent", key)
}

// KeyHome input keyevent KEYCODE_HOME
func (i *CmdInput) KeyHome() error {
	return i.KeyEvent("KEYCODE_HOME", false)
}

// KeyBack input keyevent KEYCODE_BACK
func (i *CmdInput) KeyBack() error {
	return i.KeyEvent("KEYCODE_BACK", false)
}

// KeySwitch input keyevent KEYCODE_BACK
func (i *CmdInput) KeySwitch() error {
	return i.KeyEvent("KEYCODE_APP_SWITCH", false)
}

// KeyPower input keyevent KEYCODE_POWER
func (i *CmdInput) KeyPower() error {
	return i.KeyEvent("KEYCODE_POWER", false)
}

// KeyMenu input keyevent KEYCODE_POWER
func (i *CmdInput) KeyMenu() error {
	return i.KeyEvent("KEYCODE_MENU", false)
}
