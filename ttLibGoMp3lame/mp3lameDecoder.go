package ttLibGoMp3lame

/*
#include <ttLibC/decoder/mp3lameDecoder.h>
#include "../ttLibGo/frame.h"

extern bool ttLibGoMp3lameFrameCallback(void *ptr, ttLibC_Frame *frame);

static bool Mp3lameDecoder_decodeCallback(void *ptr, ttLibC_PcmS16 *pcm) {
	return ttLibGoMp3lameFrameCallback(ptr, (ttLibC_Frame *)pcm);
}

static bool Mp3lameDecoder_decode(ttLibC_Mp3lameDecoder *decoder, void *frame, uintptr_t ptr) {
	return ttLibC_Mp3lameDecoder_decode(decoder, (ttLibC_Mp3 *)frame, Mp3lameDecoder_decodeCallback, (void *)ptr);
}
*/
import "C"
import (
	"unsafe"

	"github.com/taktod/ttLibGo/ttLibGo"
)

// Mp3lameDecoder mp3lameによるmp3のdecode動作
type Mp3lameDecoder struct {
	call     frameCall
	cDecoder *C.ttLibC_Mp3lameDecoder
}

// Init 初期化
func (mp3lame *Mp3lameDecoder) Init() bool {
	if mp3lame.cDecoder == nil {
		mp3lame.cDecoder = C.ttLibC_Mp3lameDecoder_make()
	}
	return mp3lame.cDecoder != nil
}

// Decode 実際にフレームをdecodeします
func (mp3lame *Mp3lameDecoder) Decode(
	frame *ttLibGo.Frame,
	callback ttLibGo.FrameCallback) bool {
	if mp3lame.cDecoder == nil {
		return false
	}
	if frame == nil {
		return true
	}
	frame.Restore()
	mp3lame.call.callback = callback
	if !C.Mp3lameDecoder_decode(
		mp3lame.cDecoder,
		unsafe.Pointer(frame.CFrame),
		C.uintptr_t(uintptr(unsafe.Pointer(mp3lame)))) {
		return false
	}
	return true
}

// Close 解放
func (mp3lame *Mp3lameDecoder) Close() {
	C.ttLibC_Mp3lameDecoder_close(&mp3lame.cDecoder)
	mp3lame.cDecoder = nil
}
