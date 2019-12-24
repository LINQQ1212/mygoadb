# mygoadb
```
go get -v github.com/LINQQ1212/mygoadb
```
go 调用 adb
链式调用，代码量不多

```
package main

import (
	"fmt"
	"github.com/LINQQ1212/mygoadb"
)

func main() {
	adb := mygoadb.Command("C:/platform-tools/adb.exe")

	// get Devices
	arr := adb.Devices()
	if len(arr) == 0 {
		fmt.Println("not has Devices")
		return
	}
	// You can open Debug to see the executed command
	//adb.Debug(true)

	// set Devices
	adb = adb.Use(arr[0])
	// run command
	adb.ExecOut.Screencap("3.png")
	adb.Shell.Screencap("1.png")
	adb.Shell.Input.Click(506, 838)
	adb.Shell.Input.Swipe(200, 200, 200, 800, 1000)
	adb.Shell.Input.SwipeRandom(200, 200, 200, 800, 2000, 50)
	//adb.Shell.Input.Text("test")
	adb.Shell.Input.KeyEvent("66", false)
	adb.Shell.Input.KeyHome()
}
```
