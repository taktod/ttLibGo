package ttLibGoSoundtouch

/*
#include <ttLibC/frame/frame.h>
*/
import "C"
import (
	"unsafe"

	"github.com/taktod/ttLibGo/ttLibGo"
)

type frameCall struct {
	callback ttLibGo.FrameCallback
}

//export ttLibGoSoundtouchFrameCallback
func ttLibGoSoundtouchFrameCallback(any unsafe.Pointer, cFrame unsafe.Pointer) C.bool {
	call := (*frameCall)(any)
	frame := new(ttLibGo.Frame)
	frame.Init(ttLibGo.CttLibCFrame(cFrame))
	return C.bool(call.callback(frame))
}
