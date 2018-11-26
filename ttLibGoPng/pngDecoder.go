package ttLibGoPng

/*
#include <ttLibC/decoder/pngDecoder.h>
#include "../ttLibGo/frame.h"

extern bool ttLibGoPngFrameCallback(void *ptr, ttLibC_Frame *frame);

static bool PngDecoder_decodeCallback(void *ptr, ttLibC_Bgr *bgr) {
	return ttLibGoPngFrameCallback(ptr, (ttLibC_Frame *)bgr);
}

static bool PngDecoder_decode(ttLibC_PngDecoder *decoder, void *frame, uintptr_t ptr) {
	return ttLibC_PngDecoder_decode(decoder, (ttLibC_Png *)frame, PngDecoder_decodeCallback, (void *)ptr);
}
*/
import "C"
import (
	"github.com/taktod/ttLibGo/ttLibGo"
)
import "unsafe"

// PngDecoder libpngを利用したpngDecoder
type PngDecoder struct {
	call     frameCall
	cDecoder *C.ttLibC_PngDecoder
}

// Init 初期化します
func (png *PngDecoder) Init() bool {
	if png.cDecoder == nil {
		png.cDecoder = C.ttLibC_PngDecoder_make()
	}
	return png.cDecoder != nil
}

// Decode pngのフレームをbgrにデコードします
func (png *PngDecoder) Decode(
	frame *ttLibGo.Frame,
	callback ttLibGo.FrameCallback) bool {
	if png.cDecoder == nil {
		return false
	}
	if frame == nil {
		return true
	}
	frame.Restore()
	png.call.callback = callback
	if !C.PngDecoder_decode(
		png.cDecoder,
		unsafe.Pointer(frame.CFrame),
		C.uintptr_t(uintptr(unsafe.Pointer(png)))) {
		return false
	}
	return true
}

// Close 解放します
func (png *PngDecoder) Close() {
	C.ttLibC_PngDecoder_close(&png.cDecoder)
	png.cDecoder = nil
}
