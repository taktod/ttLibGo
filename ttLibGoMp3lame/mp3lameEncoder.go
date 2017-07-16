package ttLibGoMp3lame

/*
#include <ttLibC/encoder/mp3lameEncoder.h>
#include "../ttLibGo/frame.h"

extern bool ttLibGoMp3lameFrameCallback(void *ptr, ttLibC_Frame *frame);

static bool Mp3lameEncoder_encodeCallback(void *ptr, ttLibC_Mp3 *mp3) {
	return ttLibGoMp3lameFrameCallback(ptr, (ttLibC_Frame *)mp3);
}

static bool Mp3lameEncoder_encode(ttLibC_Mp3lameEncoder *encoder, void *frame, uintptr_t ptr) {
	return ttLibC_Mp3lameEncoder_encode(encoder, (ttLibC_PcmS16 *)frame, Mp3lameEncoder_encodeCallback, (void *)ptr);
}
*/
import "C"
import (
	"unsafe"

	"github.com/taktod/ttLibGo/ttLibGo"
)

// Mp3lameEncoder mp3lameによるmp3のencode動作
type Mp3lameEncoder struct {
	call     frameCall
	cEncoder *C.ttLibC_Mp3lameEncoder
}

// Init 初期化
func (mp3lame *Mp3lameEncoder) Init(
	sampleRate uint32,
	channelNum uint32,
	quality uint32) bool {
	if mp3lame.cEncoder == nil {
		mp3lame.cEncoder = C.ttLibC_Mp3lameEncoder_make(
			C.uint32_t(sampleRate),
			C.uint32_t(channelNum),
			C.uint32_t(quality))
	}
	return mp3lame.cEncoder != nil
}

// Encode pcmS16をmp3にエンコードします
func (mp3lame *Mp3lameEncoder) Encode(
	frame *ttLibGo.Frame,
	callback ttLibGo.FrameCallback) bool {
	if mp3lame.cEncoder == nil {
		return false
	}
	if frame == nil {
		return true
	}
	frame.Restore()
	mp3lame.call.callback = callback
	if !C.Mp3lameEncoder_encode(
		mp3lame.cEncoder,
		unsafe.Pointer(frame.CFrame),
		C.uintptr_t(uintptr(unsafe.Pointer(mp3lame)))) {
		return false
	}
	return true
}

// Close 解放
func (mp3lame *Mp3lameEncoder) Close() {
	C.ttLibC_Mp3lameEncoder_close(&mp3lame.cEncoder)
	mp3lame.cEncoder = nil
}
