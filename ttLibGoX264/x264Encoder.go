package ttLibGoX264

/*
#include <stdint.h>
#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>
#include <string.h>
#include <x264.h>
#include <ttLibC/encoder/x264Encoder.h>

extern bool ttLibGoX264FrameCallback(void *ptr, ttLibC_Frame *frame);

static bool X264Encoder_encodeCallback(void *ptr, ttLibC_H264 *h264) {
	return ttLibGoX264FrameCallback(ptr, (ttLibC_Frame *)h264);
}

static bool X264Encoder_encode(ttLibC_X264Encoder *encoder, void *frame, uintptr_t ptr) {
	return ttLibC_X264Encoder_encode(encoder, (ttLibC_Yuv420 *)frame, X264Encoder_encodeCallback, (void *)ptr);
}
bool X264Encoder_forceNextFrameType(ttLibC_X264Encoder *encoder, const char *frame_type) {
	if(encoder == NULL) {
		return false;
	}
	ttLibC_X264Encoder_FrameType frameType = X264FrameType_Auto;
			if(strcmp(frame_type, "I") == 0) {
				frameType = X264FrameType_I;
			}
			else if(strcmp(frame_type, "P") == 0) {
				frameType = X264FrameType_P;
			}
			else if(strcmp(frame_type, "B") == 0) {
				frameType = X264FrameType_B;
			}
			else if(strcmp(frame_type, "IDR") == 0) {
				frameType = X264FrameType_IDR;
			}
			else if(strcmp(frame_type, "Auto") == 0) {
				frameType = X264FrameType_Auto;
			}
			else if(strcmp(frame_type, "Bref") == 0) {
				frameType = X264FrameType_Bref;
			}
			else if(strcmp(frame_type, "KeyFrame") == 0) {
				frameType = X264FrameType_KeyFrame;
			}
	else return false;
	return ttLibC_X264Encoder_forceNextFrameType(encoder, frameType);
}
*/
import "C"
import (
	"unsafe"

	"github.com/taktod/ttLibGo/ttLibGo"
)

// X264Encoder x264によるエンコード動作
type X264Encoder struct {
	call     frameCall
	cEncoder *C.ttLibC_X264Encoder
}

// Init 初期化
func (x264 *X264Encoder) Init(
	width uint32,
	height uint32,
	preset string,
	profile string,
	tune string,
	param map[string]string) bool {
	if x264.cEncoder == nil {
		cPreset := C.CString(preset)
		defer C.free(unsafe.Pointer(cPreset))
		cProfile := C.CString(profile)
		defer C.free(unsafe.Pointer(cProfile))
		cTune := C.CString(tune)
		defer C.free(unsafe.Pointer(cTune))
		if width == 0 || height == 0 {
			return false
		}
		var paramT C.struct_x264_param_t
		C.ttLibC_X264Encoder_getDefaultX264ParamTWithPresetTune(
			unsafe.Pointer(&paramT),
			C.uint32_t(width),
			C.uint32_t(height),
			cPreset,
			cTune)

		for key, value := range param {
			cKey := C.CString(key)
			defer C.free(unsafe.Pointer(cKey))
			cValue := C.CString(value)
			defer C.free(unsafe.Pointer(cValue))
			C.x264_param_parse(
				&paramT,
				cKey,
				cValue)
		}
		if C.x264_param_apply_profile(&paramT, cProfile) != 0 {
			return false
		}
		x264.cEncoder = C.ttLibC_X264Encoder_makeWithX264ParamT(unsafe.Pointer(&paramT))
	}
	return x264.cEncoder != nil
}

// Encode Yuv420をh264にエンコードします
func (x264 *X264Encoder) Encode(
	frame *ttLibGo.Frame,
	callback ttLibGo.FrameCallback) bool {
	if x264.cEncoder == nil {
		return false
	}
	if frame == nil {
		return true
	}
	frame.Restore()
	x264.call.callback = callback
	if !C.X264Encoder_encode(
		x264.cEncoder,
		unsafe.Pointer(frame.CFrame),
		C.uintptr_t(uintptr(unsafe.Pointer(x264)))) {
		return false
	}
	return true
}

// Close 解放
func (x264 *X264Encoder) Close() {
	C.ttLibC_X264Encoder_close(&x264.cEncoder)
	x264.cEncoder = nil
}

// ForceNextFrameType 次の変換結果を特定のframeに強制します
func (x264 *X264Encoder) ForceNextFrameType(h26xType string) bool {
	if x264.cEncoder == nil {
		return false
	}
	cH26xType := C.CString(h26xType)
	defer C.free(unsafe.Pointer(cH26xType))
	return bool(C.X264Encoder_forceNextFrameType(x264.cEncoder, cH26xType))
}
