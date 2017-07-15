package ttLibGoFfmpeg

/*
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <ttLibC/resampler/swresampleResampler.h>
#include "../ttLibGo/frame.h"

extern bool ttLibGoFfmpegFrameCallback(void *ptr, ttLibC_Frame *frame);

static uint32_t SwresampleResampler_getSubType(
		ttLibC_Frame_Type type,
		const char *subType) {
	switch(type) {
	case frameType_pcmS16:
		{
			if(strcmp(subType, "bigEndian") == 0) {
				return PcmS16Type_bigEndian;
			}
			else if(strcmp(subType, "littleEndian") == 0) {
				return PcmS16Type_littleEndian;
			}
			else if(strcmp(subType, "bigEndianPlanar") == 0) {
				return PcmS16Type_bigEndian_planar;
			}
			else if(strcmp(subType, "littleEndianPlanar") == 0) {
				return PcmS16Type_littleEndian_planar;
			}
			else {
				puts("unexpected pcmS16 subType");
			}
		}
		break;
	case frameType_pcmF32:
		{
			if(strcmp(subType, "planar") == 0) {
				return PcmF32Type_planar;
			}
			else if(strcmp(subType, "interleave") == 0) {
				return PcmF32Type_interleave;
			}
			else {
				puts("unexpected pcmF32 subType");
			}
		}
		break;
	default:
		puts("unexpected frametype.");
	}
	return 99;
}
static ttLibC_SwresampleResampler *SwresampleResampler_make(
		const char *inCodecType,
		const char *inSubType,
		uint32_t inSampleRate,
		uint32_t inChannelNum,
		const char *outCodecType,
		const char *outSubType,
		uint32_t outSampleRate,
		uint32_t outChannelNum) {
	ttLibC_Frame_Type inType = Frame_getFrameType(inCodecType);
	ttLibC_Frame_Type outType = Frame_getFrameType(outCodecType);
	return ttLibC_SwresampleResampler_make(
		inType,
		SwresampleResampler_getSubType(inType, inSubType),
		inSampleRate,
		inChannelNum,
		outType,
		SwresampleResampler_getSubType(outType, outSubType),
		outSampleRate,
		outChannelNum);
}
static bool SwresampleResampler_resampleCallback(void *ptr, ttLibC_Frame *audio) {
	return ttLibGoFfmpegFrameCallback(ptr, audio);
}
static bool SwresampleResampler_resample(
		ttLibC_SwresampleResampler *resampler,
		void *frame,
		uintptr_t ptr) {
	return ttLibC_SwresampleResampler_resample(resampler, (ttLibC_Frame *)frame, SwresampleResampler_resampleCallback, (void *)ptr);
}
*/
import "C"
import (
	"unsafe"

	"github.com/taktod/ttLibGo/ttLibGo"
)

// SwresampleResampler libswresampleを利用した、pcmデータのリサンプラ
type SwresampleResampler struct {
	call       frameCall
	cResampler *C.ttLibC_SwresampleResampler
}

// Init 初期化
func (swresample *SwresampleResampler) Init(
	inCodecType string,
	inSubType string,
	inSampleRate uint32,
	inChannelNum uint32,
	outCodecType string,
	outSubType string,
	outSampleRate uint32,
	outChannelNum uint32) bool {
	if swresample.cResampler == nil {
		cInCodecType := C.CString(inCodecType)
		defer C.free(unsafe.Pointer(cInCodecType))
		cInSubType := C.CString(inSubType)
		defer C.free(unsafe.Pointer(cInSubType))
		cOutCodecType := C.CString(outCodecType)
		defer C.free(unsafe.Pointer(cOutCodecType))
		cOutSubType := C.CString(outSubType)
		defer C.free(unsafe.Pointer(cOutSubType))
		swresample.cResampler = C.SwresampleResampler_make(
			cInCodecType,
			cInSubType,
			C.uint32_t(inSampleRate),
			C.uint32_t(inChannelNum),
			cOutCodecType,
			cOutSubType,
			C.uint32_t(outSampleRate),
			C.uint32_t(outChannelNum))
	}
	return swresample.cResampler != nil
}

// Resample pcmデータのリサンプルを実行します
func (swresample *SwresampleResampler) Resample(
	frame *ttLibGo.Frame,
	callback ttLibGo.FrameCallback) bool {
	if swresample.cResampler == nil {
		return false
	}
	if frame == nil {
		return true
	}
	frame.Restore()
	swresample.call.callback = callback
	if !C.SwresampleResampler_resample(
		swresample.cResampler,
		unsafe.Pointer(frame.CFrame),
		C.uintptr_t(uintptr(unsafe.Pointer(swresample)))) {
		return false
	}
	return true
}

// Close swresampleResamplerを閉じます
func (swresample *SwresampleResampler) Close() {
	C.ttLibC_SwresampleResampler_close(&swresample.cResampler)
	swresample.cResampler = nil
}
