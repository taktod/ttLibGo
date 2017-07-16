package ttLibGoJpeg

/*
#include <ttLibC/encoder/jpegEncoder.h>
#include "../ttLibGo/frame.h"

extern bool ttLibGoJpegFrameCallback(void *ptr, ttLibC_Frame *frame);

static bool JpegEncoder_encodeCallback(void *ptr, ttLibC_Jpeg *jpeg) {
	return ttLibGoJpegFrameCallback(ptr, (ttLibC_Frame *)jpeg);
}

static bool JpegEncoder_encode(ttLibC_JpegEncoder *encoder, void *frame, uintptr_t ptr) {
	return ttLibC_JpegEncoder_encode(encoder, (ttLibC_Yuv420 *)frame, JpegEncoder_encodeCallback, (void *)ptr);
}
*/
import "C"
import (
	"unsafe"

	"github.com/taktod/ttLibGo/ttLibGo"
)

// JpegEncoder libjpegを利用したencoder
type JpegEncoder struct {
	call     frameCall
	cEncoder *C.ttLibC_JpegEncoder
}

// Init 初期化 qualityは0 - 100 100が一番高品質
func (jpeg *JpegEncoder) Init(
	width uint32,
	height uint32,
	quality uint32) bool {
	if jpeg.cEncoder == nil {
		jpeg.cEncoder = C.ttLibC_JpegEncoder_make(
			C.uint32_t(width),
			C.uint32_t(height),
			C.uint32_t(quality))
	}
	return jpeg.cEncoder != nil
}

// Encode yuv420のフレームをjpegにエンコードします
func (jpeg *JpegEncoder) Encode(
	frame *ttLibGo.Frame,
	callback ttLibGo.FrameCallback) bool {
	if jpeg.cEncoder == nil {
		return false
	}
	if frame == nil {
		return true
	}
	frame.Restore()
	jpeg.call.callback = callback
	if !C.JpegEncoder_encode(
		jpeg.cEncoder,
		unsafe.Pointer(frame.CFrame),
		C.uintptr_t(uintptr(unsafe.Pointer(jpeg)))) {
		return false
	}
	return true
}

// Close 解放します
func (jpeg *JpegEncoder) Close() {
	C.ttLibC_JpegEncoder_close(&jpeg.cEncoder)
	jpeg.cEncoder = nil
}
