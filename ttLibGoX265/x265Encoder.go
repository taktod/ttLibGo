package ttLibGoX265

/*
#include <stdint.h>
#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>
#include <string.h>
#include <x265.h>
#include <ttLibC/encoder/x265Encoder.h>

extern bool ttLibGoX265FrameCallback(void *ptr, ttLibC_Frame *frame);

static bool X265Encoder_getDefaultX265ApiAndParam(
		void  *api,
		void  *param,
		const char *preset,
		const char *tune,
		uint32_t width,
		uint32_t height) {
	return ttLibC_X265Encoder_getDefaultX265ApiAndParam(
		(void **)api,
		(void **)param,
		preset,
		tune,
		width,
		height);
}

static bool X265Encoder_encodeCallback(void *ptr, ttLibC_H265 *h265) {
	return ttLibGoX265FrameCallback(ptr, (ttLibC_Frame *)h265);
}

static bool X265Encoder_encode(ttLibC_X265Encoder *encoder, void *frame, uintptr_t ptr) {
	return ttLibC_X265Encoder_encode(encoder, (ttLibC_Yuv420 *)frame, X265Encoder_encodeCallback, (void *)ptr);
}
bool X265Encoder_forceNextFrameType(ttLibC_X265Encoder *encoder, const char *frame_type) {
	if(encoder == NULL) {
		return false;
	}
	ttLibC_X265Encoder_FrameType frameType = X265FrameType_Auto;
	if(strcmp(frame_type, "I") == 0) {
		frameType = X265FrameType_I;
	}
	else if(strcmp(frame_type, "P") == 0) {
		frameType = X265FrameType_P;
	}
	else if(strcmp(frame_type, "B") == 0) {
		frameType = X265FrameType_B;
	}
	else if(strcmp(frame_type, "IDR") == 0) {
		frameType = X265FrameType_IDR;
	}
	else if(strcmp(frame_type, "Auto") == 0) {
		frameType = X265FrameType_Auto;
	}
	else if(strcmp(frame_type, "Bref") == 0) {
		frameType = X265FrameType_Bref;
	}
	else return false;
	return ttLibC_X265Encoder_forceNextFrameType(encoder, frameType);
}
*/
import "C"
import (
	"unsafe"

	"github.com/taktod/ttLibGo/ttLibGo"
)

// X265Encoder x265によるエンコード動作
type X265Encoder struct {
	call     frameCall
	cEncoder *C.ttLibC_X265Encoder
}

// Init 初期化
func (x265 *X265Encoder) Init(
	width uint32,
	height uint32,
	preset string,
	profile string,
	tune string,
	param map[string]string) bool {
	if x265.cEncoder == nil {
		cPreset := C.CString(preset)
		defer C.free(unsafe.Pointer(cPreset))
		cProfile := C.CString(profile)
		defer C.free(unsafe.Pointer(cProfile))
		cTune := C.CString(tune)
		defer C.free(unsafe.Pointer(cTune))
		if width == 0 || height == 0 {
			return false
		}
		var xapi *C.struct_x265_api
		var xparam *C.struct_x265_param

		if !C.X265Encoder_getDefaultX265ApiAndParam(
			unsafe.Pointer(&xapi),
			unsafe.Pointer(&xparam),
			cPreset,
			cTune,
			C.uint32_t(width),
			C.uint32_t(height)) {
			return false
		}
		for key, value := range param {
			cKey := C.CString(key)
			defer C.free(unsafe.Pointer(cKey))
			cValue := C.CString(value)
			defer C.free(unsafe.Pointer(cValue))
			C.x265_param_parse(
				xparam,
				cKey,
				cValue)
		}
		if C.x265_param_apply_profile(xparam, cProfile) != 0 {
			return false
		}
		x265.cEncoder = C.ttLibC_X265Encoder_makeWithX265ApiAndParam(unsafe.Pointer(xapi), unsafe.Pointer(xparam))
	}
	return x265.cEncoder != nil
}

// Encode Yuv420をh265にエンコードします
func (x265 *X265Encoder) Encode(
	frame *ttLibGo.Frame,
	callback ttLibGo.FrameCallback) bool {
	if x265.cEncoder == nil {
		return false
	}
	if frame == nil {
		return true
	}
	frame.Restore()
	x265.call.callback = callback
	if !C.X265Encoder_encode(
		x265.cEncoder,
		unsafe.Pointer(frame.CFrame),
		C.uintptr_t(uintptr(unsafe.Pointer(x265)))) {
		return false
	}
	return true
}

// Close 解放
func (x265 *X265Encoder) Close() {
	C.ttLibC_X265Encoder_close(&x265.cEncoder)
	x265.cEncoder = nil
}

// ForceNextFrameType 次の変換結果を特定のframeに強制します
func (x265 *X265Encoder) ForceNextFrameType(h26xType string) bool {
	if x265.cEncoder == nil {
		return false
	}
	cH26xType := C.CString(h26xType)
	defer C.free(unsafe.Pointer(cH26xType))
	return bool(C.X265Encoder_forceNextFrameType(x265.cEncoder, cH26xType))
}
