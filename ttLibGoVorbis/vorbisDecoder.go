package ttLibGoVorbis

/*
#include <ttLibC/decoder/vorbisDecoder.h>
#include "../ttLibGo/frame.h"

extern bool ttLibGoVorbisFrameCallback(void *ptr, ttLibC_Frame *frame);

static bool VorbisDecoder_decodeCallback(void *ptr, ttLibC_PcmF32 *pcm) {
	return ttLibGoVorbisFrameCallback(ptr, (ttLibC_Frame *)pcm);
}

static bool VorbisDecoder_decode(ttLibC_VorbisDecoder *decoder, void *frame, uintptr_t ptr) {
	return ttLibC_VorbisDecoder_decode(decoder, (ttLibC_Vorbis *)frame, VorbisDecoder_decodeCallback, (void *)ptr);
}
*/
import "C"
import (
	"unsafe"

	"github.com/taktod/ttLibGo/ttLibGo"
)

// VorbisDecoder vorbisによるデコード動作
type VorbisDecoder struct {
	call     frameCall
	cDecoder *C.ttLibC_VorbisDecoder
}

// Init 初期化
func (vorbis *VorbisDecoder) Init() bool {
	if vorbis.cDecoder == nil {
		vorbis.cDecoder = C.ttLibC_VorbisDecoder_make()
	}
	return vorbis.cDecoder != nil
}

// Decode vorbisをpcmf32にデコードします
func (vorbis *VorbisDecoder) Decode(
	frame *ttLibGo.Frame,
	callback ttLibGo.FrameCallback) bool {
	if vorbis.cDecoder == nil {
		return false
	}
	if frame == nil {
		return true
	}
	frame.Restore()
	vorbis.call.callback = callback
	if !C.VorbisDecoder_decode(
		vorbis.cDecoder,
		unsafe.Pointer(frame.CFrame),
		C.uintptr_t(uintptr(unsafe.Pointer(vorbis)))) {
		return false
	}
	return true
}

// Close 解放
func (vorbis *VorbisDecoder) Close() {
	C.ttLibC_VorbisDecoder_close(&vorbis.cDecoder)
	vorbis.cDecoder = nil
}
