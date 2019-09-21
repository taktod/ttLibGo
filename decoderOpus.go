package ttLibGo

/*
#include <stdint.h>
#include <stdlib.h>
extern int OpusDecoder_codecControl(void *ptr, const char *key, int value);
*/
import "C"
import (
	"reflect"
	"unsafe"
)

type opusDecoder decoder

// OpusDecoderControls codecControlで設定可能な項目
var OpusDecoderControls = struct {
	ResetState                subType
	GetFinalRange             subType
	GetBandwidth              subType
	GetSampleRate             subType
	SetPhaseInversionDisabled subType
	GetPhaseInversionDisabled subType
	SetGain                   subType
	GetGain                   subType
	GetLastPacketDuration     subType
	GetPitch                  subType
}{
	subType{"OPUS_RESET_STATE"},
	subType{"OPUS_GET_FINAL_RANGE"},
	subType{"OPUS_GET_BANDWIDTH"},
	subType{"OPUS_GET_SAMPLE_RATE"},
	subType{"OPUS_SET_PHASE_INVERSION_DISABLED"},
	subType{"OPUS_GET_PHASE_INVERSION_DISABLED"},
	subType{"OPUS_SET_GAIN"},
	subType{"OPUS_GET_GAIN"},
	subType{"OPUS_GET_LAST_PACKET_DURATION"},
	subType{"OPUS_GET_PITCH"},
}

// DecodeFrame デコードを実行
func (opusDecoder *opusDecoder) DecodeFrame(
	frame IFrame,
	callback func(frame *Frame) bool) bool {
	return decoderDecodeFrame((*decoder)(opusDecoder), frame, callback)
}

// Close 閉じる
func (opusDecoder *opusDecoder) Close() {
	decoderClose((*decoder)(opusDecoder))
}

// SetCodecControl 追加パラメーター設定
func (opusDecoder *opusDecoder) SetCodecControl(control subType, value interface{}) {
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
	C.OpusDecoder_codecControl(unsafe.Pointer(opusDecoder.cDecoder), ckey, C.int(ivalue))
}

// GetCodecControl 追加パラメーター参照
func (opusDecoder *opusDecoder) GetCodecControl(control subType) int32 {
	ckey := C.CString(control.value)
	defer C.free(unsafe.Pointer(ckey))
	return int32(C.OpusDecoder_codecControl(unsafe.Pointer(opusDecoder.cDecoder), ckey, C.int(0)))
}
