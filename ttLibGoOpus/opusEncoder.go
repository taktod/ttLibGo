package ttLibGoOpus

/*
#include <ttLibC/encoder/opusEncoder.h>
#include "../ttLibGo/frame.h"

extern bool ttLibGoOpusFrameCallback(void *ptr, ttLibC_Frame *frame);

static bool OpusEncoder_encodeCallback(void *ptr, ttLibC_Opus *opus) {
	return ttLibGoOpusFrameCallback(ptr, (ttLibC_Frame *)opus);
}

static bool OpusEncoder_encode(ttLibC_OpusEncoder *encoder, void *frame, uintptr_t ptr) {
	return ttLibC_OpusEncoder_encode(encoder, (ttLibC_PcmS16 *)frame, OpusEncoder_encodeCallback, (void *)ptr);
}
*/
import "C"
import (
	"unsafe"

	"github.com/taktod/ttLibGo/ttLibGo"
)

// OpusEncoder opusを利用したエンコーダー
type OpusEncoder struct {
	call     frameCall
	cEncoder *C.ttLibC_OpusEncoder
}

// Init 初期化
func (opus *OpusEncoder) Init(
	sampleRate uint32,
	channelNum uint32,
	unitSampleNum uint32) bool {
	if opus.cEncoder == nil {
		opus.cEncoder = C.ttLibC_OpusEncoder_make(
			C.uint32_t(sampleRate),
			C.uint32_t(channelNum),
			C.uint32_t(unitSampleNum))
	}
	return opus.cEncoder != nil
}

// Encode pcmS16をopusにエンコードする
func (opus *OpusEncoder) Encode(
	frame *ttLibGo.Frame,
	callback ttLibGo.FrameCallback) bool {
	if opus.cEncoder == nil {
		return false
	}
	if frame == nil {
		return true
	}
	frame.Restore()
	opus.call.callback = callback
	if !C.OpusEncoder_encode(
		opus.cEncoder,
		unsafe.Pointer(frame.CFrame),
		C.uintptr_t(uintptr(unsafe.Pointer(opus)))) {
		return false
	}
	return true
}

// Close 解放
func (opus *OpusEncoder) Close() {
	C.ttLibC_OpusEncoder_close(&opus.cEncoder)
	opus.cEncoder = nil
}

// SetBitrate 出力bitrateを再設定します
func (opus *OpusEncoder) SetBitrate(value uint32) bool {
	if opus.cEncoder == nil {
		return false
	}
	return bool(C.ttLibC_OpusEncoder_setBitrate(opus.cEncoder, C.uint32_t(value)))
}

// SetComplexity 処理のcomplexityを設定します
func (opus *OpusEncoder) SetComplexity(value uint32) bool {
	if opus.cEncoder == nil {
		return false
	}
	return bool(C.ttLibC_OpusEncoder_setComplexity(opus.cEncoder, C.uint32_t(value)))
}
