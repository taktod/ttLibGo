package ttLibGoFaac

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

//export ttLibGoFaacFrameCallback
// 本当はttLibGoでの定義を使いまわしたいところだけど、c言語からのcallbackは見えない見えないみたいなので、わざわざ定義する
func ttLibGoFaacFrameCallback(any unsafe.Pointer, cFrame unsafe.Pointer) C.bool {
	call := (*frameCall)(any)
	frame := new(ttLibGo.Frame)
	frame.Init(ttLibGo.CttLibCFrame(cFrame))
	return C.bool(call.callback(frame))
}
