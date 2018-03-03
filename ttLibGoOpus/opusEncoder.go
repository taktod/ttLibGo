package ttLibGoOpus

/*
#include <ttLibC/encoder/opusEncoder.h>
#include "../ttLibGo/frame.h"
#include <opus/opus.h>

extern bool ttLibGoOpusFrameCallback(void *ptr, ttLibC_Frame *frame);

static bool OpusEncoder_encodeCallback(void *ptr, ttLibC_Opus *opus) {
	return ttLibGoOpusFrameCallback(ptr, (ttLibC_Frame *)opus);
}

static bool OpusEncoder_encode(ttLibC_OpusEncoder *encoder, void *frame, uintptr_t ptr) {
	return ttLibC_OpusEncoder_encode(encoder, (ttLibC_PcmS16 *)frame, OpusEncoder_encodeCallback, (void *)ptr);
}

static int64_t OpusEncoder_OPUS_SET_COMPLEXITY(ttLibC_OpusEncoder *encoder, int value) {
	return opus_encoder_ctl(ttLibC_OpusEncoder_refNativeEncoder(encoder), OPUS_SET_COMPLEXITY(value));
}
static int64_t OpusEncoder_OPUS_SET_BITRATE(ttLibC_OpusEncoder *encoder, int value) {
	return opus_encoder_ctl(ttLibC_OpusEncoder_refNativeEncoder(encoder), OPUS_SET_BITRATE(value));
}
static int64_t OpusEncoder_OPUS_SET_VBR(ttLibC_OpusEncoder *encoder, int value) {
	return opus_encoder_ctl(ttLibC_OpusEncoder_refNativeEncoder(encoder), OPUS_SET_VBR(value));
}
static int64_t OpusEncoder_OPUS_SET_VBR_CONSTRAINT(ttLibC_OpusEncoder *encoder, int value) {
	return opus_encoder_ctl(ttLibC_OpusEncoder_refNativeEncoder(encoder), OPUS_SET_VBR_CONSTRAINT(value));
}
static int64_t OpusEncoder_OPUS_SET_FORCE_CHANNELS(ttLibC_OpusEncoder *encoder, int value) {
	return opus_encoder_ctl(ttLibC_OpusEncoder_refNativeEncoder(encoder), OPUS_SET_FORCE_CHANNELS(value));
}
static int64_t OpusEncoder_OPUS_SET_MAX_BANDWIDTH(ttLibC_OpusEncoder *encoder, int value) {
	return opus_encoder_ctl(ttLibC_OpusEncoder_refNativeEncoder(encoder), OPUS_SET_MAX_BANDWIDTH(value));
}
static int64_t OpusEncoder_OPUS_SET_BANDWIDTH(ttLibC_OpusEncoder *encoder, int value) {
	return opus_encoder_ctl(ttLibC_OpusEncoder_refNativeEncoder(encoder), OPUS_SET_BANDWIDTH(value));
}
static int64_t OpusEncoder_OPUS_SET_SIGNAL(ttLibC_OpusEncoder *encoder, int value) {
	return opus_encoder_ctl(ttLibC_OpusEncoder_refNativeEncoder(encoder), OPUS_SET_SIGNAL(value));
}
static int64_t OpusEncoder_OPUS_SET_APPLICATION(ttLibC_OpusEncoder *encoder, int value) {
	return opus_encoder_ctl(ttLibC_OpusEncoder_refNativeEncoder(encoder), OPUS_SET_APPLICATION(value));
}
static int64_t OpusEncoder_OPUS_SET_INBAND_FEC(ttLibC_OpusEncoder *encoder, int value) {
	return opus_encoder_ctl(ttLibC_OpusEncoder_refNativeEncoder(encoder), OPUS_SET_INBAND_FEC(value));
}
static int64_t OpusEncoder_OPUS_SET_PACKET_LOSS_PERC(ttLibC_OpusEncoder *encoder, int value) {
	return opus_encoder_ctl(ttLibC_OpusEncoder_refNativeEncoder(encoder), OPUS_SET_PACKET_LOSS_PERC(value));
}
static int64_t OpusEncoder_OPUS_SET_DTX(ttLibC_OpusEncoder *encoder, int value) {
	return opus_encoder_ctl(ttLibC_OpusEncoder_refNativeEncoder(encoder), OPUS_SET_DTX(value));
}
static int64_t OpusEncoder_OPUS_SET_LSB_DEPTH(ttLibC_OpusEncoder *encoder, int value) {
	return opus_encoder_ctl(ttLibC_OpusEncoder_refNativeEncoder(encoder), OPUS_SET_LSB_DEPTH(value));
}
static int64_t OpusEncoder_OPUS_SET_EXPERT_FRAME_DURATION(ttLibC_OpusEncoder *encoder, int value) {
	return opus_encoder_ctl(ttLibC_OpusEncoder_refNativeEncoder(encoder), OPUS_SET_EXPERT_FRAME_DURATION(value));
}
static int64_t OpusEncoder_OPUS_SET_PREDICTION_DISABLED(ttLibC_OpusEncoder *encoder, int value) {
	return opus_encoder_ctl(ttLibC_OpusEncoder_refNativeEncoder(encoder), OPUS_SET_PREDICTION_DISABLED(value));
}
static int64_t OpusEncoder_OPUS_SET_PHASE_INVERSION_DISABLED(ttLibC_OpusEncoder *encoder, int value) {
	return opus_encoder_ctl(ttLibC_OpusEncoder_refNativeEncoder(encoder), OPUS_SET_PHASE_INVERSION_DISABLED(value));
}
static int64_t OpusEncoder_OPUS_GET_COMPLEXITY(ttLibC_OpusEncoder *encoder) {
	int res, result;
	res = opus_encoder_ctl(ttLibC_OpusEncoder_refNativeEncoder(encoder), OPUS_GET_COMPLEXITY(&result));
	if(res != OPUS_OK) {
		return res;
	}
	return result;
}
static int64_t OpusEncoder_OPUS_GET_BITRATE(ttLibC_OpusEncoder *encoder) {
	int res, result;
	res = opus_encoder_ctl(ttLibC_OpusEncoder_refNativeEncoder(encoder), OPUS_GET_BITRATE(&result));
	if(res != OPUS_OK) {
		return res;
	}
	return result;
}
static int64_t OpusEncoder_OPUS_GET_VBR(ttLibC_OpusEncoder *encoder) {
	int res, result;
	res = opus_encoder_ctl(ttLibC_OpusEncoder_refNativeEncoder(encoder), OPUS_GET_VBR(&result));
	if(res != OPUS_OK) {
		return res;
	}
	return result;
}
static int64_t OpusEncoder_OPUS_GET_VBR_CONSTRAINT(ttLibC_OpusEncoder *encoder) {
	int res, result;
	res = opus_encoder_ctl(ttLibC_OpusEncoder_refNativeEncoder(encoder), OPUS_GET_VBR_CONSTRAINT(&result));
	if(res != OPUS_OK) {
		return res;
	}
	return result;
}
static int64_t OpusEncoder_OPUS_GET_FORCE_CHANNELS(ttLibC_OpusEncoder *encoder) {
	int res, result;
	res = opus_encoder_ctl(ttLibC_OpusEncoder_refNativeEncoder(encoder), OPUS_GET_FORCE_CHANNELS(&result));
	if(res != OPUS_OK) {
		return res;
	}
	return result;
}
static int64_t OpusEncoder_OPUS_GET_MAX_BANDWIDTH(ttLibC_OpusEncoder *encoder) {
	int res, result;
	res = opus_encoder_ctl(ttLibC_OpusEncoder_refNativeEncoder(encoder), OPUS_GET_MAX_BANDWIDTH(&result));
	if(res != OPUS_OK) {
		return res;
	}
	return result;
}
static int64_t OpusEncoder_OPUS_GET_SIGNAL(ttLibC_OpusEncoder *encoder) {
	int res, result;
	res = opus_encoder_ctl(ttLibC_OpusEncoder_refNativeEncoder(encoder), OPUS_GET_SIGNAL(&result));
	if(res != OPUS_OK) {
		return res;
	}
	return result;
}
static int64_t OpusEncoder_OPUS_GET_APPLICATION(ttLibC_OpusEncoder *encoder) {
	int res, result;
	res = opus_encoder_ctl(ttLibC_OpusEncoder_refNativeEncoder(encoder), OPUS_GET_APPLICATION(&result));
	if(res != OPUS_OK) {
		return res;
	}
	return result;
}
static int64_t OpusEncoder_OPUS_GET_LOOKAHEAD(ttLibC_OpusEncoder *encoder) {
	int res, result;
	res = opus_encoder_ctl(ttLibC_OpusEncoder_refNativeEncoder(encoder), OPUS_GET_LOOKAHEAD(&result));
	if(res != OPUS_OK) {
		return res;
	}
	return result;
}
static int64_t OpusEncoder_OPUS_GET_INBAND_FEC(ttLibC_OpusEncoder *encoder) {
	int res, result;
	res = opus_encoder_ctl(ttLibC_OpusEncoder_refNativeEncoder(encoder), OPUS_GET_INBAND_FEC(&result));
	if(res != OPUS_OK) {
		return res;
	}
	return result;
}
static int64_t OpusEncoder_OPUS_GET_PACKET_LOSS_PERC(ttLibC_OpusEncoder *encoder) {
	int res, result;
	res = opus_encoder_ctl(ttLibC_OpusEncoder_refNativeEncoder(encoder), OPUS_GET_PACKET_LOSS_PERC(&result));
	if(res != OPUS_OK) {
		return res;
	}
	return result;
}
static int64_t OpusEncoder_OPUS_GET_DTX(ttLibC_OpusEncoder *encoder) {
	int res, result;
	res = opus_encoder_ctl(ttLibC_OpusEncoder_refNativeEncoder(encoder), OPUS_GET_DTX(&result));
	if(res != OPUS_OK) {
		return res;
	}
	return result;
}
static int64_t OpusEncoder_OPUS_GET_LSB_DEPTH(ttLibC_OpusEncoder *encoder) {
	int res, result;
	res = opus_encoder_ctl(ttLibC_OpusEncoder_refNativeEncoder(encoder), OPUS_GET_LSB_DEPTH(&result));
	if(res != OPUS_OK) {
		return res;
	}
	return result;
}
static int64_t OpusEncoder_OPUS_GET_EXPERT_FRAME_DURATION(ttLibC_OpusEncoder *encoder) {
	int res, result;
	res = opus_encoder_ctl(ttLibC_OpusEncoder_refNativeEncoder(encoder), OPUS_GET_EXPERT_FRAME_DURATION(&result));
	if(res != OPUS_OK) {
		return res;
	}
	return result;
}
static int64_t OpusEncoder_OPUS_GET_PREDICTION_DISABLED(ttLibC_OpusEncoder *encoder) {
	int res, result;
	res = opus_encoder_ctl(ttLibC_OpusEncoder_refNativeEncoder(encoder), OPUS_GET_PREDICTION_DISABLED(&result));
	if(res != OPUS_OK) {
		return res;
	}
	return result;
}
static int64_t OpusEncoder_OPUS_RESET_STATE(ttLibC_OpusEncoder *encoder) {
	return opus_encoder_ctl(ttLibC_OpusEncoder_refNativeEncoder(encoder), OPUS_RESET_STATE);
}
static int64_t OpusEncoder_OPUS_GET_FINAL_RANGE(ttLibC_OpusEncoder *encoder) {
	int res;
	uint32_t result;
	res = opus_encoder_ctl(ttLibC_OpusEncoder_refNativeEncoder(encoder), OPUS_GET_FINAL_RANGE(&result));
	if(res != OPUS_OK) {
		return res;
	}
	return result;
}
static int64_t OpusEncoder_OPUS_GET_BANDWIDTH(ttLibC_OpusEncoder *encoder) {
	int res, result;
	res = opus_encoder_ctl(ttLibC_OpusEncoder_refNativeEncoder(encoder), OPUS_GET_BANDWIDTH(&result));
	if(res != OPUS_OK) {
		return res;
	}
	return result;
}
static int64_t OpusEncoder_OPUS_GET_SAMPLE_RATE(ttLibC_OpusEncoder *encoder) {
	int res, result;
	res = opus_encoder_ctl(ttLibC_OpusEncoder_refNativeEncoder(encoder), OPUS_GET_SAMPLE_RATE(&result));
	if(res != OPUS_OK) {
		return res;
	}
	return result;
}
static int64_t OpusEncoder_OPUS_GET_PHASE_INVERSION_DISABLED(ttLibC_OpusEncoder *encoder) {
	int res, result;
	res = opus_encoder_ctl(ttLibC_OpusEncoder_refNativeEncoder(encoder), OPUS_GET_PHASE_INVERSION_DISABLED(&result));
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

// SetCodecControl opus_encoder_ctlで特殊設定を実施します
func (opus *OpusEncoder) SetCodecControl(control string, value int64) int64 {
	if opus.cEncoder == nil {
		return -1
	}
	switch control {
	case "OPUS_SET_COMPLEXITY":
		return int64(C.OpusEncoder_OPUS_SET_COMPLEXITY(opus.cEncoder, C.int(value)))
	case "OPUS_SET_BITRATE":
		return int64(C.OpusEncoder_OPUS_SET_BITRATE(opus.cEncoder, C.int(value)))
	case "OPUS_SET_VBR":
		return int64(C.OpusEncoder_OPUS_SET_VBR(opus.cEncoder, C.int(value)))
	case "OPUS_SET_VBR_CONSTRAINT":
		return int64(C.OpusEncoder_OPUS_SET_VBR_CONSTRAINT(opus.cEncoder, C.int(value)))
	case "OPUS_SET_FORCE_CHANNELS":
		return int64(C.OpusEncoder_OPUS_SET_FORCE_CHANNELS(opus.cEncoder, C.int(value)))
	case "OPUS_SET_MAX_BANDWIDTH":
		return int64(C.OpusEncoder_OPUS_SET_MAX_BANDWIDTH(opus.cEncoder, C.int(value)))
	case "OPUS_SET_BANDWIDTH":
		return int64(C.OpusEncoder_OPUS_SET_BANDWIDTH(opus.cEncoder, C.int(value)))
	case "OPUS_SET_SIGNAL":
		return int64(C.OpusEncoder_OPUS_SET_SIGNAL(opus.cEncoder, C.int(value)))
	case "OPUS_SET_APPLICATION":
		return int64(C.OpusEncoder_OPUS_SET_APPLICATION(opus.cEncoder, C.int(value)))
	case "OPUS_SET_INBAND_FEC":
		return int64(C.OpusEncoder_OPUS_SET_INBAND_FEC(opus.cEncoder, C.int(value)))
	case "OPUS_SET_PACKET_LOSS_PERC":
		return int64(C.OpusEncoder_OPUS_SET_PACKET_LOSS_PERC(opus.cEncoder, C.int(value)))
	case "OPUS_SET_DTX":
		return int64(C.OpusEncoder_OPUS_SET_DTX(opus.cEncoder, C.int(value)))
	case "OPUS_SET_LSB_DEPTH":
		return int64(C.OpusEncoder_OPUS_SET_LSB_DEPTH(opus.cEncoder, C.int(value)))
	case "OPUS_SET_EXPERT_FRAME_DURATION":
		return int64(C.OpusEncoder_OPUS_SET_EXPERT_FRAME_DURATION(opus.cEncoder, C.int(value)))
	case "OPUS_SET_PREDICTION_DISABLED":
		return int64(C.OpusEncoder_OPUS_SET_PREDICTION_DISABLED(opus.cEncoder, C.int(value)))
	case "OPUS_SET_PHASE_INVERSION_DISABLED":
		return int64(C.OpusEncoder_OPUS_SET_PHASE_INVERSION_DISABLED(opus.cEncoder, C.int(value)))
	default:
		return -1
	}
}

// GetCodecControl opus_encoder_ctlで特殊設定参照を実施します
func (opus *OpusEncoder) GetCodecControl(control string) int64 {
	if opus.cEncoder == nil {
		return -1
	}
	switch control {
	case "OPUS_GET_COMPLEXITY":
		return int64(C.OpusEncoder_OPUS_GET_COMPLEXITY(opus.cEncoder))
	case "OPUS_GET_BITRATE":
		return int64(C.OpusEncoder_OPUS_GET_BITRATE(opus.cEncoder))
	case "OPUS_GET_VBR":
		return int64(C.OpusEncoder_OPUS_GET_VBR(opus.cEncoder))
	case "OPUS_GET_VBR_CONSTRAINT":
		return int64(C.OpusEncoder_OPUS_GET_VBR_CONSTRAINT(opus.cEncoder))
	case "OPUS_GET_FORCE_CHANNELS":
		return int64(C.OpusEncoder_OPUS_GET_FORCE_CHANNELS(opus.cEncoder))
	case "OPUS_GET_MAX_BANDWIDTH":
		return int64(C.OpusEncoder_OPUS_GET_MAX_BANDWIDTH(opus.cEncoder))
	case "OPUS_GET_SIGNAL":
		return int64(C.OpusEncoder_OPUS_GET_SIGNAL(opus.cEncoder))
	case "OPUS_GET_APPLICATION":
		return int64(C.OpusEncoder_OPUS_GET_APPLICATION(opus.cEncoder))
	case "OPUS_GET_LOOKAHEAD":
		return int64(C.OpusEncoder_OPUS_GET_LOOKAHEAD(opus.cEncoder))
	case "OPUS_GET_INBAND_FEC":
		return int64(C.OpusEncoder_OPUS_GET_INBAND_FEC(opus.cEncoder))
	case "OPUS_GET_PACKET_LOSS_PERC":
		return int64(C.OpusEncoder_OPUS_GET_PACKET_LOSS_PERC(opus.cEncoder))
	case "OPUS_GET_DTX":
		return int64(C.OpusEncoder_OPUS_GET_DTX(opus.cEncoder))
	case "OPUS_GET_LSB_DEPTH":
		return int64(C.OpusEncoder_OPUS_GET_LSB_DEPTH(opus.cEncoder))
	case "OPUS_GET_EXPERT_FRAME_DURATION":
		return int64(C.OpusEncoder_OPUS_GET_EXPERT_FRAME_DURATION(opus.cEncoder))
	case "OPUS_GET_PREDICTION_DISABLED":
		return int64(C.OpusEncoder_OPUS_GET_PREDICTION_DISABLED(opus.cEncoder))
	case "OPUS_RESET_STATE":
		return int64(C.OpusEncoder_OPUS_RESET_STATE(opus.cEncoder))
	case "OPUS_GET_FINAL_RANGE":
		return int64(C.OpusEncoder_OPUS_GET_FINAL_RANGE(opus.cEncoder))
	case "OPUS_GET_BANDWIDTH":
		return int64(C.OpusEncoder_OPUS_GET_BANDWIDTH(opus.cEncoder))
	case "OPUS_GET_SAMPLE_RATE":
		return int64(C.OpusEncoder_OPUS_GET_SAMPLE_RATE(opus.cEncoder))
	case "OPUS_GET_PHASE_INVERSION_DISABLED":
		return int64(C.OpusEncoder_OPUS_GET_PHASE_INVERSION_DISABLED(opus.cEncoder))
	default:
		return -1
	}
	return 0
}
