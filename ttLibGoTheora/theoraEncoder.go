package ttLibGoTheora

/*
#include <ttLibC/encoder/theoraEncoder.h>
#include "../ttLibGo/frame.h"

extern bool ttLibGoTheoraFrameCallback(void *ptr, ttLibC_Frame *frame);

static bool TheoraEncoder_encodeCallback(void *ptr, ttLibC_Theora *theora) {
	return ttLibGoTheoraFrameCallback(ptr, (ttLibC_Frame *)theora);
}

static bool TheoraEncoder_encode(ttLibC_TheoraEncoder *encoder, void *frame, uintptr_t ptr) {
	return ttLibC_TheoraEncoder_encode(encoder, (ttLibC_Yuv420 *)frame, TheoraEncoder_encodeCallback, (void *)ptr);
}
*/
import "C"
import (
	"unsafe"

	"github.com/taktod/ttLibGo/ttLibGo"
)

// TheoraEncoder theoraを利用したencode動作
type TheoraEncoder struct {
	call     frameCall
	cEncoder *C.ttLibC_TheoraEncoder
}

// Init 初期化
func (theora *TheoraEncoder) Init(
	width uint32,
	height uint32,
	quality uint32,
	bitrate uint32,
	keyFrameInterval uint32) bool {
	if theora.cEncoder == nil {
		theora.cEncoder = C.ttLibC_TheoraEncoder_make_ex(
			C.uint32_t(width),
			C.uint32_t(height),
			C.uint32_t(quality),
			C.uint32_t(bitrate),
			C.uint32_t(keyFrameInterval))
	}
	return theora.cEncoder != nil
}

// Encode yuv420をtheoraに変換します
func (theora *TheoraEncoder) Encode(
	frame *ttLibGo.Frame,
	callback ttLibGo.FrameCallback) bool {
	if theora.cEncoder == nil {
		return false
	}
	if frame == nil {
		return true
	}
	frame.Restore()
	theora.call.callback = callback
	if !C.TheoraEncoder_encode(
		theora.cEncoder,
		unsafe.Pointer(frame.CFrame),
		C.uintptr_t(uintptr(unsafe.Pointer(theora)))) {
		return false
	}
	return true
}

// Close 解放
func (theora *TheoraEncoder) Close() {
	C.ttLibC_TheoraEncoder_close(&theora.cEncoder)
	theora.cEncoder = nil
}
