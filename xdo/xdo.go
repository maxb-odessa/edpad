package xdo

// #include <stdlib.h>
// #include <xdo.h>
// #cgo LDFLAGS: -lxdo
import "C"
import "unsafe"

type Xdo struct {
	xdo *C.xdo_t
	win C.Window
}

func New(remote string) *Xdo {

	if remote == "" {
		return nil
	}

	x := new(Xdo)

	x.xdo = C.xdo_new(C.CString(remote))
	if x.xdo == nil {
		x = nil
	} else {
		x.win = C.CURRENTWINDOW
	}

	return x
}

func (x *Xdo) GetActiveWindow() {
	win := C.Window(0)
	C.xdo_get_active_window(x.xdo, &win)
	x.win = win
}

func (x *Xdo) GetWindowName() string {
	var ret *C.uchar
	var retLen C.int
	var dummy C.int

	C.xdo_get_window_name(x.xdo, x.win, &ret, &retLen, &dummy)
	name := C.GoBytes(unsafe.Pointer(ret), retLen)
	C.free(unsafe.Pointer(ret))

	return string(name)
}

func (x *Xdo) KeyDown(key string, udelay int) {
	C.xdo_send_keysequence_window_down(x.xdo, x.win, C.CString(key), C.useconds_t(udelay))
}

func (x *Xdo) KeyUp(key string, udelay int) {
	C.xdo_send_keysequence_window_up(x.xdo, x.win, C.CString(key), C.useconds_t(udelay))
}
