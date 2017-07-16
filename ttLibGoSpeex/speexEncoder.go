package ttLibGoSpeex

/*
#include <ttLibC/encoder/speexEncoder.h>
#include "../ttLibGo/frame.h"

extern bool ttLibGoSpeexFrameCallback(void *ptr, ttLibC_Frame *frame);

static bool SpeexEncoder_encodeCallback(void *ptr, ttLibC_Speex *speex) {
	return ttLibGoSpeexFrameCallback(ptr, (ttLibC_Frame *)speex);
}

static bool SpeexEncoder_encode(ttLibC_SpeexEncoder *encoder, void *frame, uintptr_t ptr) {
	return ttLibC_SpeexEncoder_encode(encoder, (ttLibC_PcmS16 *)frame, SpeexEncoder_encodeCallback, (void *)ptr);
}
*/
import "C"
import (
	"unsafe"

	"github.com/taktod/ttLibGo/ttLibGo"
)

// SpeexEncoder speexを利用したencode動作
type SpeexEncoder struct {
	call     frameCall
	cEncoder *C.ttLibC_SpeexEncoder
}

// Init 初期化
func (speex *SpeexEncoder) Init(
	sampleRate uint32,
	channelNum uint32,
	quality uint32) bool {
	if speex.cEncoder == nil {
		speex.cEncoder = C.ttLibC_SpeexEncoder_make(
			C.uint32_t(sampleRate),
			C.uint32_t(channelNum),
			C.uint32_t(quality))
	}
	return speex.cEncoder != nil
}

// Encode pcmS16をspeexに変換します
func (speex *SpeexEncoder) Encode(
	frame *ttLibGo.Frame,
	callback ttLibGo.FrameCallback) bool {
	if speex.cEncoder == nil {
		return false
	}
	if frame == nil {
		return true
	}
	frame.Restore()
	speex.call.callback = callback
	if !C.SpeexEncoder_encode(
		speex.cEncoder,
		unsafe.Pointer(frame.CFrame),
		C.uintptr_t(uintptr(unsafe.Pointer(speex)))) {
		return false
	}
	return true
}

// Close 解放
func (speex *SpeexEncoder) Close() {
	C.ttLibC_SpeexEncoder_close(&speex.cEncoder)
	speex.cEncoder = nil
}
