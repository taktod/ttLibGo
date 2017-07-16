package ttLibGoOpus

/*
#include <ttLibC/decoder/opusDecoder.h>
#include "../ttLibGo/frame.h"

extern bool ttLibGoOpusFrameCallback(void *ptr, ttLibC_Frame *frame);

static bool OpusDecoder_decodeCallback(void *ptr, ttLibC_PcmS16 *pcm) {
	return ttLibGoOpusFrameCallback(ptr, (ttLibC_Frame *)pcm);
}

static bool OpusDecoder_decode(ttLibC_OpusDecoder *decoder, void *frame, uintptr_t ptr) {
	return ttLibC_OpusDecoder_decode(decoder, (ttLibC_Opus *)frame, OpusDecoder_decodeCallback, (void *)ptr);
}
*/
import "C"
import (
	"unsafe"

	"github.com/taktod/ttLibGo/ttLibGo"
)

// OpusDecoder opusを利用したdecoder
type OpusDecoder struct {
	call     frameCall
	cDecoder *C.ttLibC_OpusDecoder
}

// Init 初期化
func (opus *OpusDecoder) Init(
	sampleRate uint32,
	channelNum uint32) bool {
	if opus.cDecoder == nil {
		opus.cDecoder = C.ttLibC_OpusDecoder_make(
			C.uint32_t(sampleRate),
			C.uint32_t(channelNum))
	}
	return opus.cDecoder != nil
}

// Decode opusをpcmS16にデコードします
func (opus *OpusDecoder) Decode(
	frame *ttLibGo.Frame,
	callback ttLibGo.FrameCallback) bool {
	if opus.cDecoder == nil {
		return false
	}
	if frame == nil {
		return true
	}
	frame.Restore()
	opus.call.callback = callback
	if !C.OpusDecoder_decode(
		opus.cDecoder,
		unsafe.Pointer(frame.CFrame),
		C.uintptr_t(uintptr(unsafe.Pointer(opus)))) {
		return false
	}
	return true
}

// Close 解放
func (opus *OpusDecoder) Close() {
	C.ttLibC_OpusDecoder_close(&opus.cDecoder)
	opus.cDecoder = nil
}
