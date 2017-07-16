package ttLibGoOpenh264

/*
#include <ttLibC/decoder/openh264Decoder.h>
#include "../ttLibGo/frame.h"

extern bool ttLibGoOpenh264FrameCallback(void *ptr, ttLibC_Frame *frame);

static bool Openh264Decoder_decodeCallback(void *ptr, ttLibC_Yuv420 *yuv) {
	return ttLibGoOpenh264FrameCallback(ptr, (ttLibC_Frame *)yuv);
}

static bool Openh264Decoder_decode(ttLibC_Openh264Decoder *decoder, void *frame, uintptr_t ptr) {
	return ttLibC_Openh264Decoder_decode(decoder, (ttLibC_H264 *)frame, Openh264Decoder_decodeCallback, (void *)ptr);
}
*/
import "C"
import (
	"unsafe"

	"github.com/taktod/ttLibGo/ttLibGo"
)

// Openh264Decoder openh264を利用したh264のdecoder
type Openh264Decoder struct {
	call     frameCall
	cDecoder *C.ttLibC_Openh264Decoder
}

// Init 初期化
func (openh264 *Openh264Decoder) Init() bool {
	if openh264.cDecoder == nil {
		openh264.cDecoder = C.ttLibC_Openh264Decoder_make()
	}
	return openh264.cDecoder != nil
}

// Decode h264フレームをyuv420にデコードします
func (openh264 *Openh264Decoder) Decode(
	frame *ttLibGo.Frame,
	callback ttLibGo.FrameCallback) bool {
	if openh264.cDecoder == nil {
		return false
	}
	if frame == nil {
		return true
	}
	frame.Restore()
	openh264.call.callback = callback
	if !C.Openh264Decoder_decode(
		openh264.cDecoder,
		unsafe.Pointer(frame.CFrame),
		C.uintptr_t(uintptr(unsafe.Pointer(openh264)))) {
		return false
	}
	return true
}

// Close 解放します
func (openh264 *Openh264Decoder) Close() {
	C.ttLibC_Openh264Decoder_close(&openh264.cDecoder)
	openh264.cDecoder = nil
}
