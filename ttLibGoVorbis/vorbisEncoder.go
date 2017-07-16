package ttLibGoVorbis

/*
#include <ttLibC/encoder/vorbisEncoder.h>
#include "../ttLibGo/frame.h"

extern bool ttLibGoVorbisFrameCallback(void *ptr, ttLibC_Frame *frame);

static bool VorbisEncoder_encodeCallback(void *ptr, ttLibC_Vorbis *vorbis) {
	return ttLibGoVorbisFrameCallback(ptr, (ttLibC_Frame *)vorbis);
}

static bool VorbisEncoder_encode(ttLibC_VorbisEncoder *encoder, void *frame, uintptr_t ptr) {
	return ttLibC_VorbisEncoder_encode(encoder, (ttLibC_Audio *)frame, VorbisEncoder_encodeCallback, (void *)ptr);
}
*/
import "C"
import (
	"unsafe"

	"github.com/taktod/ttLibGo/ttLibGo"
)

// VorbisEncoder vorbisを利用したencode動作
type VorbisEncoder struct {
	call     frameCall
	cEncoder *C.ttLibC_VorbisEncoder
}

// Init 初期化
func (vorbis *VorbisEncoder) Init(
	sampleRate uint32,
	channelNum uint32) bool {
	if vorbis.cEncoder == nil {
		vorbis.cEncoder = C.ttLibC_VorbisEncoder_make(
			C.uint32_t(sampleRate),
			C.uint32_t(channelNum))
	}
	return vorbis.cEncoder != nil
}

// Encode pcmS16をvorbisに変換します
func (vorbis *VorbisEncoder) Encode(
	frame *ttLibGo.Frame,
	callback ttLibGo.FrameCallback) bool {
	if vorbis.cEncoder == nil {
		return false
	}
	if frame == nil {
		return true
	}
	frame.Restore()
	vorbis.call.callback = callback
	if !C.VorbisEncoder_encode(
		vorbis.cEncoder,
		unsafe.Pointer(frame.CFrame),
		C.uintptr_t(uintptr(unsafe.Pointer(vorbis)))) {
		return false
	}
	return true
}

// Close 解放
func (vorbis *VorbisEncoder) Close() {
	C.ttLibC_VorbisEncoder_close(&vorbis.cEncoder)
	vorbis.cEncoder = nil
}
