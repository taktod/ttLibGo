package ttLibGoFaac

/*
#include <ttLibC/encoder/faacEncoder.h>
#include "../ttLibGo/frame.h"

extern bool ttLibGoFaacFrameCallback(void *ptr, ttLibC_Frame *frame);

static ttLibC_FaacEncoder *FaacEncoder_make(const char *ftype, uint32_t sampleRate, uint32_t channelNum, uint32_t bitrate) {
	if(ftype == NULL || sampleRate == 0 || channelNum == 0 || bitrate == 0) {
		puts("入力パラメーターが不正です");
		return NULL;
	}
	ttLibC_FaacEncoder_Type type = FaacEncoderType_Low;
	if(strcmp(ftype, "Low") == 0) {
	}
	else if(strcmp(ftype, "SSR") == 0) {
		type = FaacEncoderType_SSR;
	}
	else if(strcmp(ftype, "LTP") == 0) {
		type = FaacEncoderType_LTP;
	}
	else if(strcmp(ftype, "Main") == 0) {
		type = FaacEncoderType_Main;
	}
	else {
		return NULL;
	}
	return ttLibC_FaacEncoder_make(type, sampleRate, channelNum, bitrate);
}

static bool FaacEncoder_encodeCallback(void *ptr, ttLibC_Aac *aac) {
	return ttLibGoFaacFrameCallback(ptr, (ttLibC_Frame *)aac);
}

static bool FaacEncoder_encode(ttLibC_FaacEncoder *encoder, void *frame, uintptr_t ptr) {
	return ttLibC_FaacEncoder_encode(encoder, (ttLibC_PcmS16 *)frame, FaacEncoder_encodeCallback, (void *)ptr);
}
*/
import "C"
import (
	"unsafe"

	"github.com/taktod/ttLibGo/ttLibGo"
)

// FaacEncoder libfaacによるaacのencoder
type FaacEncoder struct {
	call     frameCall
	cEncoder *C.ttLibC_FaacEncoder
}

// Init 初期化
func (faac *FaacEncoder) Init(
	aacType string,
	sampleRate uint32,
	channelNum uint32,
	bitrate uint32) bool {
	if faac.cEncoder == nil {
		cType := C.CString(aacType)
		defer C.free(unsafe.Pointer(cType))
		faac.cEncoder = C.FaacEncoder_make(
			cType,
			C.uint32_t(sampleRate),
			C.uint32_t(channelNum),
			C.uint32_t(bitrate))
	}
	return faac.cEncoder != nil
}

// Encode 実際にフレームをencodeします。sourceはpcmS16のみ対応
func (faac *FaacEncoder) Encode(
	frame *ttLibGo.Frame,
	callback ttLibGo.FrameCallback) bool {
	if faac.cEncoder == nil {
		return false
	}
	if frame == nil {
		return true
	}
	frame.Restore()
	faac.call.callback = callback
	if !C.FaacEncoder_encode(
		faac.cEncoder,
		unsafe.Pointer(frame.CFrame),
		C.uintptr_t(uintptr(unsafe.Pointer(faac)))) {
		return false
	}
	return true
}

// Close 解放します
func (faac *FaacEncoder) Close() {
	C.ttLibC_FaacEncoder_close(&faac.cEncoder)
	faac.cEncoder = nil
}
