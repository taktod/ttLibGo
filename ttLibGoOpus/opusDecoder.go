package ttLibGoOpus

/*
#include <ttLibC/decoder/opusDecoder.h>
#include "../ttLibGo/frame.h"
#include <opus/opus.h>

extern bool ttLibGoOpusFrameCallback(void *ptr, ttLibC_Frame *frame);

static bool OpusDecoder_decodeCallback(void *ptr, ttLibC_PcmS16 *pcm) {
	return ttLibGoOpusFrameCallback(ptr, (ttLibC_Frame *)pcm);
}

static bool OpusDecoder_decode(ttLibC_OpusDecoder *decoder, void *frame, uintptr_t ptr) {
	return ttLibC_OpusDecoder_decode(decoder, (ttLibC_Opus *)frame, OpusDecoder_decodeCallback, (void *)ptr);
}
static int64_t OpusDecoder_OPUS_SET_PHASE_INVERSION_DISABLED(ttLibC_OpusDecoder *decoder, int value) {
	return opus_decoder_ctl(ttLibC_OpusDecoder_refNativeDecoder(decoder), OPUS_SET_PHASE_INVERSION_DISABLED(value));
}
static int64_t OpusDecoder_OPUS_SET_GAIN(ttLibC_OpusDecoder *decoder, int value) {
	return opus_decoder_ctl(ttLibC_OpusDecoder_refNativeDecoder(decoder), OPUS_SET_GAIN(value));
}
static int64_t OpusDecoder_OPUS_RESET_STATE(ttLibC_OpusDecoder *decoder) {
	return opus_decoder_ctl(ttLibC_OpusDecoder_refNativeDecoder(decoder), OPUS_RESET_STATE);
}
static int64_t OpusDecoder_OPUS_GET_FINAL_RANGE(ttLibC_OpusDecoder *decoder) {
	int res;
	uint32_t result;
	res = opus_decoder_ctl(ttLibC_OpusDecoder_refNativeDecoder(decoder), OPUS_GET_FINAL_RANGE(&result));
	if(res != OPUS_OK) {
		return res;
	}
	return result;
}
static int64_t OpusDecoder_OPUS_GET_BANDWIDTH(ttLibC_OpusDecoder *decoder) {
	int res, result;
	res = opus_decoder_ctl(ttLibC_OpusDecoder_refNativeDecoder(decoder), OPUS_GET_BANDWIDTH(&result));
	if(res != OPUS_OK) {
		return res;
	}
	return result;
}
static int64_t OpusDecoder_OPUS_GET_SAMPLE_RATE(ttLibC_OpusDecoder *decoder) {
	int res, result;
	res = opus_decoder_ctl(ttLibC_OpusDecoder_refNativeDecoder(decoder), OPUS_GET_SAMPLE_RATE(&result));
	if(res != OPUS_OK) {
		return res;
	}
	return result;
}
static int64_t OpusDecoder_OPUS_GET_PHASE_INVERSION_DISABLED(ttLibC_OpusDecoder *decoder) {
	int res, result;
	res = opus_decoder_ctl(ttLibC_OpusDecoder_refNativeDecoder(decoder), OPUS_GET_PHASE_INVERSION_DISABLED(&result));
	if(res != OPUS_OK) {
		return res;
	}
	return result;
}
static int64_t OpusDecoder_OPUS_GET_GAIN(ttLibC_OpusDecoder *decoder) {
	int res, result;
	res = opus_decoder_ctl(ttLibC_OpusDecoder_refNativeDecoder(decoder), OPUS_GET_GAIN(&result));
	if(res != OPUS_OK) {
		return res;
	}
	return result;
}
static int64_t OpusDecoder_OPUS_GET_LAST_PACKET_DURATION(ttLibC_OpusDecoder *decoder) {
	int res, result;
	res = opus_decoder_ctl(ttLibC_OpusDecoder_refNativeDecoder(decoder), OPUS_GET_LAST_PACKET_DURATION(&result));
	if(res != OPUS_OK) {
		return res;
	}
	return result;
}
static int64_t OpusDecoder_OPUS_GET_PITCH(ttLibC_OpusDecoder *decoder) {
	int res, result;
	res = opus_decoder_ctl(ttLibC_OpusDecoder_refNativeDecoder(decoder), OPUS_GET_PITCH(&result));
	if(res != OPUS_OK) {
		return res;
	}
	return result;
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

// SetCodecControl opus_decoder_ctlで特殊設定を実施します。
func (opus *OpusDecoder) SetCodecControl(control string, value int64) int64 {
	if opus.cDecoder == nil {
		return -1
	}
	switch control {
	case "OPUS_SET_PHASE_INVERSION_DISABLED":
		return int64(C.OpusDecoder_OPUS_SET_PHASE_INVERSION_DISABLED(opus.cDecoder, C.int(value)))
	case "OPUS_SET_GAIN":
		return int64(C.OpusDecoder_OPUS_SET_GAIN(opus.cDecoder, C.int(value)))
	default:
		return -1
	}
}

// GetCodecControl opus_decoder_ctlで特殊設定参照を実施します
func (opus *OpusDecoder) GetCodecControl(control string) int64 {
	if opus.cDecoder == nil {
		return -1
	}
	switch control {
	case "OPUS_RESET_STATE":
		return int64(C.OpusDecoder_OPUS_RESET_STATE(opus.cDecoder))
	case "OPUS_GET_FINAL_RANGE":
		return int64(C.OpusDecoder_OPUS_GET_FINAL_RANGE(opus.cDecoder))
	case "OPUS_GET_BANDWIDTH":
		return int64(C.OpusDecoder_OPUS_GET_BANDWIDTH(opus.cDecoder))
	case "OPUS_GET_SAMPLE_RATE":
		return int64(C.OpusDecoder_OPUS_GET_SAMPLE_RATE(opus.cDecoder))
	case "OPUS_GET_PHASE_INVERSION_DISABLED":
		return int64(C.OpusDecoder_OPUS_GET_PHASE_INVERSION_DISABLED(opus.cDecoder))
	case "OPUS_GET_GAIN":
		return int64(C.OpusDecoder_OPUS_GET_GAIN(opus.cDecoder))
	case "OPUS_GET_LAST_PACKET_DURATION":
		return int64(C.OpusDecoder_OPUS_GET_LAST_PACKET_DURATION(opus.cDecoder))
	case "OPUS_GET_PITCH":
		return int64(C.OpusDecoder_OPUS_GET_PITCH(opus.cDecoder))
	default:
		return -1
	}
}
