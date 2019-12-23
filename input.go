package mygoadb

import "image"

type Input interface {
	Click(loc image.Point) (err error)
	SwipeT(p0, p1 image.Point, time int) (err error)

	Touch(loc image.Point, ty int) (err error)
	Key(in string, ty int) (err error)

	Text(in string) (err error)

	Keyevent(in string) (err error)
	KeyHome() error
	KeyBack() error
	KeySwitch() error
	KeyPower() error
}
