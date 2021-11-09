package display

// #include <stdlib.h>
// #include <xdo.h>
// #cgo LDFLAGS: -lxdo
import "C"
import "unsafe"

type xdo struct {
	xdo *C.xdo_t
	win C.Window
}

func xdoNew(remote string) *xdo {
	x := new(xdo)
	x.xdo = C.xdo_new(C.CString(remote))
	x.win = C.CURRENTWINDOW
	return x
}

func (x *xdo) getActiveWindow() {
	win := C.Window(0)
	C.xdo_get_active_window(x.xdo, &win)
	x.win = win
}

func (x *xdo) getWindowName() string {
	var ret *C.uchar
	var retLen C.int
	var dummy C.int

	C.xdo_get_window_name(x.xdo, x.win, &ret, &retLen, &dummy)
	name := C.GoBytes(unsafe.Pointer(ret), retLen)
	C.free(unsafe.Pointer(ret))

	return string(name)
}

func (x *xdo) keyDown(keys string, udelay int) {
	C.xdo_send_keysequence_window_down(x.xdo, x.win, C.CString(keys), C.useconds_t(udelay))
}

func (x *xdo) keyUp(keys string, udelay int) {
	C.xdo_send_keysequence_window_up(x.xdo, x.win, C.CString(keys), C.useconds_t(udelay))
}
