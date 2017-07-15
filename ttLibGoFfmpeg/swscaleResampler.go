package ttLibGoFfmpeg

/*
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <ttLibC/allocator.h>
#include <ttLibC/resampler/swscaleResampler.h>
#include "log.h"
#include "../ttLibGo/frame.h"

extern bool ttLibGoFfmpegFrameCallback(void *ptr, ttLibC_Frame *frame);

static uint32_t SwscaleResampler_getSubType(ttLibC_Frame_Type type, const char *subType) {
	switch(type) {
	case frameType_yuv420:
		{
			if(strcmp(subType, "planar") == 0) {
				return Yuv420Type_planar;
			}
			else if(strcmp(subType, "yvuPlanar") == 0) {
				return Yvu420Type_planar;
			}
			else if(strcmp(subType, "semiPlanar") == 0) {
				return Yuv420Type_semiPlanar;
			}
			else if(strcmp(subType, "yvuSemiPlanar") == 0) {
				return Yvu420Type_semiPlanar;
			}
		}
		break;
	case frameType_bgr:
		{
			if(strcmp(subType, "bgr") == 0) {
				return BgrType_bgr;
			}
			else if(strcmp(subType, "bgra") == 0) {
				return BgrType_bgra;
			}
			else if(strcmp(subType, "abgr") == 0) {
				return BgrType_abgr;
			}
		}
		break;
	default:
		break;
	}
	return 99;
}
static ttLibC_SwscaleResampler *SwscaleResampler_make(
		const char *inCodecType,
		const char *inSubType,
		uint32_t inWidth,
		uint32_t inHeight,
		const char *outCodecType,
		const char *outSubType,
		uint32_t outWidth,
		uint32_t outHeight,
		const char *mode) {
	ttLibC_SwscaleResampler_Mode scaleMode = SwscaleResampler_FastBiLinear;
	if(strcmp(mode, "X") == 0){
		scaleMode = SwscaleResampler_X;
	}
	else if(strcmp(mode, "Area") == 0){
		scaleMode = SwscaleResampler_Area;
	}
	else if(strcmp(mode, "Sinc") == 0){
		scaleMode = SwscaleResampler_Sinc;
	}
	else if(strcmp(mode, "Point") == 0){
		scaleMode = SwscaleResampler_Point;
	}
	else if(strcmp(mode, "Gauss") == 0){
		scaleMode = SwscaleResampler_Gauss;
	}
	else if(strcmp(mode, "Spline") == 0){
		scaleMode = SwscaleResampler_Spline;
	}
	else if(strcmp(mode, "Bicubic") == 0){
		scaleMode = SwscaleResampler_Bicubic;
	}
	else if(strcmp(mode, "Lanczos") == 0){
		scaleMode = SwscaleResampler_Lanczos;
	}
	else if(strcmp(mode, "Bilinear") == 0){
		scaleMode = SwscaleResampler_Bilinear;
	}
	else if(strcmp(mode, "Bicublin") == 0){
		scaleMode = SwscaleResampler_Bicublin;
	}
	else if(strcmp(mode, "FastBilinear") == 0) {
		scaleMode = SwscaleResampler_FastBiLinear;
	}
	else {
		return NULL;
	}
	ttLibC_Frame_Type inType = Frame_getFrameType(inCodecType);
	ttLibC_Frame_Type outType = Frame_getFrameType(outCodecType);
	return ttLibC_SwscaleResampler_make(
		inType,
		SwscaleResampler_getSubType(inType, inSubType),
		inWidth,
		inHeight,
		outType,
		SwscaleResampler_getSubType(outType, outSubType),
		outWidth,
		outHeight,
		scaleMode);
}
static bool SwscaleResampler_resampleCallback(void *ptr, ttLibC_Frame *image) {
	return ttLibGoFfmpegFrameCallback(ptr, image);
}
static bool SwscaleResampler_resample(
		ttLibC_SwscaleResampler *resampler,
		void *frame,
		uintptr_t ptr) {
	return ttLibC_SwscaleResampler_resample(resampler, (ttLibC_Frame *)frame, SwscaleResampler_resampleCallback, (void *)ptr);
}

*/
import "C"
import (
	"unsafe"

	"github.com/taktod/ttLibGo/ttLibGo"
)

// SwscaleResampler libswscaleを利用したyuv bgrデータのリサンプル処理
type SwscaleResampler struct {
	call       frameCall
	cResampler *C.ttLibC_SwscaleResampler
}

// Init 初期化
func (swscale *SwscaleResampler) Init(
	inCodecType string,
	inSubType string,
	inWidth uint32,
	inHeight uint32,
	outCodecType string,
	outSubType string,
	outWidth uint32,
	outHeight uint32,
	mode string) bool {
	if swscale.cResampler == nil {
		cInCodecType := C.CString(inCodecType)
		defer C.free(unsafe.Pointer(cInCodecType))
		cInSubType := C.CString(inSubType)
		defer C.free(unsafe.Pointer(cInSubType))
		cOutCodecType := C.CString(outCodecType)
		defer C.free(unsafe.Pointer(cOutCodecType))
		cOutSubType := C.CString(outSubType)
		defer C.free(unsafe.Pointer(cOutSubType))
		cMode := C.CString(mode)
		defer C.free(unsafe.Pointer(cMode))
		swscale.cResampler = C.SwscaleResampler_make(
			cInCodecType,
			cInSubType,
			C.uint32_t(inWidth),
			C.uint32_t(inHeight),
			cOutCodecType,
			cOutSubType,
			C.uint32_t(outWidth),
			C.uint32_t(outHeight),
			cMode)
	}
	return swscale.cResampler != nil
}

// Resample データをリサンプルします
func (swscale *SwscaleResampler) Resample(
	frame *ttLibGo.Frame,
	callback ttLibGo.FrameCallback) bool {
	if swscale.cResampler == nil {
		return false
	}
	if frame == nil {
		return true
	}
	frame.Restore()
	swscale.call.callback = callback
	if !C.SwscaleResampler_resample(
		swscale.cResampler,
		unsafe.Pointer(frame.CFrame),
		C.uintptr_t(uintptr(unsafe.Pointer(swscale)))) {
		return false
	}
	return true
}

// Close 解放します
func (swscale *SwscaleResampler) Close() {
	C.ttLibC_SwscaleResampler_close(&swscale.cResampler)
	swscale.cResampler = nil
}
