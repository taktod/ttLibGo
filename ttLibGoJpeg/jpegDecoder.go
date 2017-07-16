package ttLibGoJpeg

/*
#include <ttLibC/decoder/jpegDecoder.h>
#include "../ttLibGo/frame.h"

extern bool ttLibGoJpegFrameCallback(void *ptr, ttLibC_Frame *frame);

static bool JpegDecoder_decodeCallback(void *ptr, ttLibC_Yuv420 *yuv) {
	return ttLibGoJpegFrameCallback(ptr, (ttLibC_Frame *)yuv);
}

static bool JpegDecoder_decode(ttLibC_JpegDecoder *decoder, void *frame, uintptr_t ptr) {
	return ttLibC_JpegDecoder_decode(decoder, (ttLibC_Jpeg *)frame, JpegDecoder_decodeCallback, (void *)ptr);
}
*/
import "C"
import (
	"unsafe"

	"github.com/taktod/ttLibGo/ttLibGo"
)

// JpegDecoder libjpegを利用したjpegDecoder
type JpegDecoder struct {
	call     frameCall
	cDecoder *C.ttLibC_JpegDecoder
}

// Init 初期化します
func (jpeg *JpegDecoder) Init() bool {
	if jpeg.cDecoder == nil {
		jpeg.cDecoder = C.ttLibC_JpegDecoder_make()
	}
	return jpeg.cDecoder != nil
}

// Decode jpegのフレームをyuvにデコードします
func (jpeg *JpegDecoder) Decode(
	frame *ttLibGo.Frame,
	callback ttLibGo.FrameCallback) bool {
	if jpeg.cDecoder == nil {
		return false
	}
	if frame == nil {
		return true
	}
	frame.Restore()
	jpeg.call.callback = callback
	if !C.JpegDecoder_decode(
		jpeg.cDecoder,
		unsafe.Pointer(frame.CFrame),
		C.uintptr_t(uintptr(unsafe.Pointer(jpeg)))) {
		return false
	}
	return true
}

// Close 解放します
func (jpeg *JpegDecoder) Close() {
	C.ttLibC_JpegDecoder_close(&jpeg.cDecoder)
	jpeg.cDecoder = nil
}
