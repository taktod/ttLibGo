package ttLibGoTheora

/*
#include <ttLibC/decoder/theoraDecoder.h>
#include "../ttLibGo/frame.h"

extern bool ttLibGoTheoraFrameCallback(void *ptr, ttLibC_Frame *frame);

static bool TheoraDecoder_decodeCallback(void *ptr, ttLibC_Yuv420 *yuv) {
	return ttLibGoTheoraFrameCallback(ptr, (ttLibC_Frame *)yuv);
}

static bool TheoraDecoder_decode(ttLibC_TheoraDecoder *decoder, void *frame, uintptr_t ptr) {
	return ttLibC_TheoraDecoder_decode(decoder, (ttLibC_Theora *)frame, TheoraDecoder_decodeCallback, (void *)ptr);
}
*/
import "C"
import (
	"unsafe"

	"github.com/taktod/ttLibGo/ttLibGo"
)

// TheoraDecoder theoraによるデコード動作
type TheoraDecoder struct {
	call     frameCall
	cDecoder *C.ttLibC_TheoraDecoder
}

// Init 初期化
func (theora *TheoraDecoder) Init() bool {
	if theora.cDecoder == nil {
		theora.cDecoder = C.ttLibC_TheoraDecoder_make()
	}
	return theora.cDecoder != nil
}

// Decode theoraをyuv420にデコードします
func (theora *TheoraDecoder) Decode(
	frame *ttLibGo.Frame,
	callback ttLibGo.FrameCallback) bool {
	if theora.cDecoder == nil {
		return false
	}
	if frame == nil {
		return true
	}
	frame.Restore()
	theora.call.callback = callback
	if !C.TheoraDecoder_decode(
		theora.cDecoder,
		unsafe.Pointer(frame.CFrame),
		C.uintptr_t(uintptr(unsafe.Pointer(theora)))) {
		return false
	}
	return true
}

// Close 解放
func (theora *TheoraDecoder) Close() {
	C.ttLibC_TheoraDecoder_close(&theora.cDecoder)
	theora.cDecoder = nil
}
