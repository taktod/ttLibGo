package ttLibGoSpeex

/*
#include <ttLibC/decoder/speexDecoder.h>
#include "../ttLibGo/frame.h"

extern bool ttLibGoSpeexFrameCallback(void *ptr, ttLibC_Frame *frame);

static bool SpeexDecoder_decodeCallback(void *ptr, ttLibC_PcmS16 *pcm) {
	return ttLibGoSpeexFrameCallback(ptr, (ttLibC_Frame *)pcm);
}

static bool SpeexDecoder_decode(ttLibC_SpeexDecoder *decoder, void *frame, uintptr_t ptr) {
	return ttLibC_SpeexDecoder_decode(decoder, (ttLibC_Speex *)frame, SpeexDecoder_decodeCallback, (void *)ptr);
}
*/
import "C"
import (
	"unsafe"

	"github.com/taktod/ttLibGo/ttLibGo"
)

// SpeexDecoder speexによるデコード動作
type SpeexDecoder struct {
	call     frameCall
	cDecoder *C.ttLibC_SpeexDecoder
}

// Init 初期化
func (speex *SpeexDecoder) Init(
	sampleRate uint32,
	channelNum uint32) bool {
	if speex.cDecoder == nil {
		speex.cDecoder = C.ttLibC_SpeexDecoder_make(
			C.uint32_t(sampleRate),
			C.uint32_t(channelNum))
	}
	return speex.cDecoder != nil
}

// Decode speexをpcms16にデコードします
func (speex *SpeexDecoder) Decode(
	frame *ttLibGo.Frame,
	callback ttLibGo.FrameCallback) bool {
	if speex.cDecoder == nil {
		return false
	}
	if frame == nil {
		return true
	}
	frame.Restore()
	speex.call.callback = callback
	if !C.SpeexDecoder_decode(
		speex.cDecoder,
		unsafe.Pointer(frame.CFrame),
		C.uintptr_t(uintptr(unsafe.Pointer(speex)))) {
		return false
	}
	return true
}

// Close 解放
func (speex *SpeexDecoder) Close() {
	C.ttLibC_SpeexDecoder_close(&speex.cDecoder)
	speex.cDecoder = nil
}
