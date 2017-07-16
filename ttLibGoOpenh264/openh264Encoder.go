package ttLibGoOpenh264

/*
#include <ttLibC/encoder/openh264Encoder.h>
#include "../ttLibGo/frame.h"
#include <wels/codec_api.h>

extern bool ttLibGoOpenh264FrameCallback(void *ptr, ttLibC_Frame *frame);

typedef struct openh264Param_t {
	SEncParamExt paramExt;
} openh264Param_t;

static bool Openh264Encoder_encodeCallback(void *ptr, ttLibC_H264 *h264) {
	return ttLibGoOpenh264FrameCallback(ptr, (ttLibC_Frame *)h264);
}

static bool Openh264Encoder_encode(ttLibC_Openh264Encoder *encoder, void *frame, uintptr_t ptr) {
	return ttLibC_Openh264Encoder_encode(encoder, (ttLibC_Yuv420 *)frame, Openh264Encoder_encodeCallback, (void *)ptr);
}

bool Openh264Encoder_setRCMode(ttLibC_Openh264Encoder *encoder, const char *mode) {
	if(encoder == NULL) {
		return false;
	}
	ttLibC_Openh264Encoder_RCType type;
	if(!strcmp(mode, "QualityMode")) {
		type = Openh264EncoderRCType_QualityMode;
	}
	else if(!strcmp(mode, "BitrateMode")) {
		type = Openh264EncoderRCType_BitrateMode;
	}
	else if(!strcmp(mode, "BufferbasedMode")) {
		type = Openh264EncoderRCType_BufferbasedMode;
	}
	else if(!strcmp(mode, "TimestampMode")) {
		type = Openh264EncoderRCType_TimestampMode;
	}
	else if(!strcmp(mode, "BitrateModePostSkip")) {
		type = Openh264EncoderRCType_BitrateModePostSkip;
	}
	else if(!strcmp(mode, "OffMode")) {
		type = Openh264EncoderRCType_OffMode;
	}
	else {
		return false;
	}
	return ttLibC_Openh264Encoder_setRCMode(encoder, type);
}
*/
import "C"
import (
	"unsafe"

	"github.com/taktod/ttLibGo/ttLibGo"
)

// Openh264Encoder openh264を利用したh264のencoder
type Openh264Encoder struct {
	call     frameCall
	cEncoder *C.ttLibC_Openh264Encoder
}

// Init 初期化
func (openh264 *Openh264Encoder) Init(
	width uint32,
	height uint32,
	param map[string]string,
	sLayers [](map[string]string)) bool {
	if openh264.cEncoder == nil {
		if width == 0 || height == 0 {
			return false
		}
		var oParamT C.struct_openh264Param_t
		C.ttLibC_Openh264Encoder_getDefaultSEncParamExt(
			unsafe.Pointer(&oParamT.paramExt),
			C.uint32_t(width),
			C.uint32_t(height))
		for key, value := range param {
			// stringが決定しているわけか・・
			cKey := C.CString(key)
			defer C.free(unsafe.Pointer(cKey))
			cValue := C.CString(value)
			defer C.free(unsafe.Pointer(cValue))
			C.ttLibC_Openh264Encoder_paramParse(
				unsafe.Pointer(&oParamT.paramExt),
				cKey,
				cValue)
		}
		// spatialLayer
		for i := 0; i < len(sLayers); i++ {
			for key, value := range sLayers[i] {
				cKey := C.CString(key)
				defer C.free(unsafe.Pointer(cKey))
				cValue := C.CString(value)
				defer C.free(unsafe.Pointer(cValue))
				C.ttLibC_Openh264Encoder_spatialParamParse(
					unsafe.Pointer(&oParamT.paramExt),
					C.uint32_t(i),
					cKey,
					cValue)
			}
		}
		openh264.cEncoder = C.ttLibC_Openh264Encoder_makeWithSEncParamExt(unsafe.Pointer(&oParamT.paramExt))
	}
	return openh264.cEncoder != nil
}

// Encode yuv420のイメージをh264に変換します
func (openh264 *Openh264Encoder) Encode(
	frame *ttLibGo.Frame,
	callback ttLibGo.FrameCallback) bool {
	if openh264.cEncoder == nil {
		return false
	}
	if frame == nil {
		return true
	}
	frame.Restore()
	openh264.call.callback = callback
	if !C.Openh264Encoder_encode(
		openh264.cEncoder,
		unsafe.Pointer(frame.CFrame),
		C.uintptr_t(uintptr(unsafe.Pointer(openh264)))) {
		return false
	}
	return true
}

// Close 解放
func (openh264 *Openh264Encoder) Close() {
	C.ttLibC_Openh264Encoder_close(&openh264.cEncoder)
	openh264.cEncoder = nil
}

// SetRCMode RCModeを変更します
func (openh264 *Openh264Encoder) SetRCMode(mode string) bool {
	if openh264.cEncoder == nil {
		return false
	}
	cMode := C.CString(mode)
	defer C.free(unsafe.Pointer(cMode))
	return bool(C.Openh264Encoder_setRCMode(openh264.cEncoder, cMode))
}

// SetIDRInterval keyFrameに相当するsliceIDRの間隔を変更します
func (openh264 *Openh264Encoder) SetIDRInterval(value int32) bool {
	if openh264.cEncoder == nil {
		return false
	}
	return bool(C.ttLibC_Openh264Encoder_setIDRInterval(openh264.cEncoder, C.int32_t(value)))
}

// ForceNextKeyFrame 次のエンコード結果を強制的にkeyFrameにします
func (openh264 *Openh264Encoder) ForceNextKeyFrame() bool {
	if openh264.cEncoder == nil {
		return false
	}
	return bool(C.ttLibC_Openh264Encoder_forceNextKeyFrame(openh264.cEncoder))
}
