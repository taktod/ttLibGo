package ttLibGo

/*
#include <stdint.h>
#include <stdlib.h>
extern int OpusEncoder_codecControl(void *ptr, const char *key, int value);
*/
import "C"
import (
	"reflect"
	"unsafe"
)

type opusEncoder encoder

// OpusEncoderControls codecControlで設定可能な項目
var OpusEncoderControls = struct {
	SetComplexity             subType
	GetComplexity             subType
	SetBitrate                subType
	GetBitrate                subType
	SetVBR                    subType
	GetVBR                    subType
	SetVBRConstraint          subType
	GetVBRConstraint          subType
	SetForceChannels          subType
	GetForceChannels          subType
	SetMaxBandwidth           subType
	GetMaxBandwidth           subType
	SetBandwidth              subType
	SetSignal                 subType
	GetSignal                 subType
	SetApplication            subType
	GetApplication            subType
	GetLookahead              subType
	SetInbandFEC              subType
	GetInbandFEC              subType
	SetPacketLossPERC         subType
	GetPacketLossPERC         subType
	SetDTX                    subType
	GetDTX                    subType
	SetLSBDepth               subType
	GetLSBDepth               subType
	SetExpertFrameDuration    subType
	GetExpertFrameDuration    subType
	SetPredictionDisabled     subType
	GetPredictionDisabled     subType
	ResetState                subType
	GetFinalRange             subType
	GetBandwidth              subType
	GetSampleRate             subType
	SetPhaseInversionDisabled subType
	GetPhaseInversionDisabled subType
}{
	subType{"OPUS_SET_COMPLEXITY"},
	subType{"OPUS_GET_COMPLEXITY"},
	subType{"OPUS_SET_BITRATE"},
	subType{"OPUS_GET_BITRATE"},
	subType{"OPUS_SET_VBR"},
	subType{"OPUS_GET_VBR"},
	subType{"OPUS_SET_VBR_CONSTRAINT"},
	subType{"OPUS_GET_VBR_CONSTRAINT"},
	subType{"OPUS_SET_FORCE_CHANNELS"},
	subType{"OPUS_GET_FORCE_CHANNELS"},
	subType{"OPUS_SET_MAX_BANDWIDTH"},
	subType{"OPUS_GET_MAX_BANDWIDTH"},
	subType{"OPUS_SET_BANDWIDTH"},
	subType{"OPUS_SET_SIGNAL"},
	subType{"OPUS_GET_SIGNAL"},
	subType{"OPUS_SET_APPLICATION"},
	subType{"OPUS_GET_APPLICATION"},
	subType{"OPUS_GET_LOOKAHEAD"},
	subType{"OPUS_SET_INBAND_FEC"},
	subType{"OPUS_GET_INBAND_FEC"},
	subType{"OPUS_SET_PACKET_LOSS_PERC"},
	subType{"OPUS_GET_PACKET_LOSS_PERC"},
	subType{"OPUS_SET_DTX"},
	subType{"OPUS_GET_DTX"},
	subType{"OPUS_SET_LSB_DEPTH"},
	subType{"OPUS_GET_LSB_DEPTH"},
	subType{"OPUS_SET_EXPERT_FRAME_DURATION"},
	subType{"OPUS_GET_EXPERT_FRAME_DURATION"},
	subType{"OPUS_SET_PREDICTION_DISABLED"},
	subType{"OPUS_GET_PREDICTION_DISABLED"},
	subType{"OPUS_RESET_STATE"},
	subType{"OPUS_GET_FINAL_RANGE"},
	subType{"OPUS_GET_BANDWIDTH"},
	subType{"OPUS_GET_SAMPLE_RATE"},
	subType{"OPUS_SET_PHASE_INVERSION_DISABLED"},
	subType{"OPUS_GET_PHASE_INVERSION_DISABLED"},
}

func (opusEncoder *opusEncoder) EncodeFrame(
	frame IFrame,
	callback FrameCallback) bool {
	return encoderEncodeFrame((*encoder)(opusEncoder), frame, callback)
}

// Close 閉じる
func (opusEncoder *opusEncoder) Close() {
	encoderClose((*encoder)(opusEncoder))
}

// SetCodecControl 追加パラメーター設定
func (opusEncoder *opusEncoder) SetCodecControl(control subType, value interface{}) {
	ckey := C.CString(control.value)
	defer C.free(unsafe.Pointer(ckey))
	// valueについて、reflectでなんとかしなければならない
	ivalue := 0
	switch reflect.TypeOf(value).Kind() {
	case reflect.Uint32:
		ivalue = int(value.(uint32))
	case reflect.Int:
		ivalue = value.(int)
	case reflect.Int32:
		ivalue = int(value.(int32))
	case reflect.Bool:
		if value.(bool) {
			ivalue = 1
		}
	}
	C.OpusEncoder_codecControl(unsafe.Pointer(opusEncoder.cEncoder), ckey, C.int(ivalue))
}

// GetCodecControl 追加パラメーター参照
func (opusEncoder *opusEncoder) GetCodecControl(control subType) int32 {
	ckey := C.CString(control.value)
	defer C.free(unsafe.Pointer(ckey))
	return int32(C.OpusEncoder_codecControl(unsafe.Pointer(opusEncoder.cEncoder), ckey, C.int(0)))
}
