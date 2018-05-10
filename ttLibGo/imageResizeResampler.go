package ttLibGo

/*
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <ttLibC/allocator.h>
#include <ttLibC/resampler/imageResizer.h>
#include <ttLibC/frame/video/bgr.h>
#include <ttLibC/frame/video/yuv420.h>
#include "frame.h"

extern bool ttLibGoFrameCallback(void *ptr, ttLibC_Frame *frame);

typedef struct imageResizerResampler_t {
	ttLibC_Frame_Type type;
	uint32_t          sub_type;
	uint32_t          width;
	uint32_t          height;
	bool              isQuick;
	ttLibC_Frame     *frame;
} imageResizerResampler_t;

imageResizerResampler_t *ImageResizerResampler_make(
		const char *codecType,
		const char *sub_type,
		uint32_t width,
		uint32_t height,
		bool is_quick) {
	ttLibC_Frame_Type type = Frame_getFrameType(codecType);
	switch(type) {
	case frameType_bgr:
	case frameType_yuv420:
		break;
	default:
		return NULL;
	}
	imageResizerResampler_t *resampler = ttLibC_malloc(sizeof(imageResizerResampler_t));
	if(resampler == NULL) {
		puts("failed to allocate imageResizerResampler Object.");
		return NULL;
	}
	switch(type) {
	case frameType_bgr:
		{
			resampler->type = type;
			if(strcmp(sub_type, "bgr") == 0) {
				resampler->sub_type = BgrType_bgr;
			}
			else if(strcmp(sub_type, "rgb") == 0) {
				resampler->sub_type = BgrType_rgb;
			}
			else if(strcmp(sub_type, "bgra") == 0) {
				resampler->sub_type = BgrType_bgra;
			}
			else if(strcmp(sub_type, "abgr") == 0) {
				resampler->sub_type = BgrType_abgr;
			}
			else if(strcmp(sub_type, "argb") == 0) {
				resampler->sub_type = BgrType_argb;
			}
			else if(strcmp(sub_type, "rgba") == 0) {
				resampler->sub_type = BgrType_rgba;
			}
			else {
				puts("subTypeが不正です。");
				ttLibC_free(resampler);
				return NULL;
			}
		}
		break;
	case frameType_yuv420:
		{
			resampler->type = type;
			if(strcmp(sub_type, "planar") == 0) {
				resampler->sub_type = Yuv420Type_planar;
			}
			else if(strcmp(sub_type, "yvuPlanar") == 0) {
				resampler->sub_type = Yvu420Type_planar;
			}
			else if(strcmp(sub_type, "semiPlanar") == 0) {
				resampler->sub_type = Yuv420Type_semiPlanar;
			}
			else if(strcmp(sub_type, "yvuSemiPlanar") == 0) {
				resampler->sub_type = Yvu420Type_semiPlanar;
			}
			else {
				puts("subTypeが不正です。");
				ttLibC_free(resampler);
				return NULL;
			}
		}
		break;
	default:
		return NULL;
	}
	resampler->frame  = NULL;
	resampler->width  = width;
	resampler->height = height;
	resampler->isQuick = is_quick;
	return resampler;
}

bool ImageResizerResampler_resample(
		imageResizerResampler_t *resampler,
		void *frame,
		uintptr_t ptr) {
	if(resampler == NULL) {
		return false;
	}
	if(frame == NULL) {
		return true;
	}
	ttLibC_Frame *f = (ttLibC_Frame *)frame;
	switch(f->type) {
	case frameType_bgr:
		{
			if(resampler->type != frameType_bgr) {
				puts("入力フレームタイプが不正です。");
				return false;
			}
			ttLibC_Bgr *b = ttLibC_ImageResizer_resizeBgr((ttLibC_Bgr *)resampler->frame, resampler->sub_type, resampler->width, resampler->height, (ttLibC_Bgr *)frame);
			if(b == NULL) {
				puts("resample失敗しました。データができてません。");
				return false;
			}
			resampler->frame = (ttLibC_Frame *)b;
		}
		break;
	case frameType_yuv420:
		{
			if(resampler->type != frameType_yuv420) {
				puts("入力フレームタイプが不正です。");
				return false;
			}
			ttLibC_Yuv420 *y = ttLibC_ImageResizer_resizeYuv420((ttLibC_Yuv420 *)resampler->frame, resampler->sub_type, resampler->width, resampler->height, (ttLibC_Yuv420 *)frame, resampler->isQuick);
			if(y == NULL) {
				puts("resample失敗しました。データができてません。");
				return false;
			}
			resampler->frame = (ttLibC_Frame *)y;
		}
		break;
	default:
		return false;
	}
	return ttLibGoFrameCallback((void *)ptr, resampler->frame);
}

void ImageResizerResampler_close(imageResizerResampler_t **resampler) {
	imageResizerResampler_t *target = *resampler;
	if(target == NULL) {
		return;
	}
	ttLibC_Frame_close(&target->frame);
	ttLibC_free(target);
	*resampler = NULL;
}
*/
import "C"
import (
	"fmt"
	"unsafe"
)

// ImageResizeResampler 画像データのサイズを変更するresampler
type ImageResizeResampler struct {
	call       frameCall
	cResampler *C.imageResizerResampler_t
}

// Init 初期化
func (image *ImageResizeResampler) Init(
	codecType string,
	subType string,
	width uint32,
	height uint32,
	isQuick bool) bool {
	if image.cResampler == nil {
		cCodecType := C.CString(codecType)
		defer C.free(unsafe.Pointer(cCodecType))
		cSubType := C.CString(subType)
		defer C.free(unsafe.Pointer(cSubType))
		image.cResampler = C.ImageResizerResampler_make(
			cCodecType,
			cSubType,
			C.uint32_t(width),
			C.uint32_t(height),
			C.bool(isQuick))
	}
	return image.cResampler != nil
}

// Resample リサンプル実行
func (image *ImageResizeResampler) Resample(
	frame *Frame,
	callback FrameCallback) bool {
	if image.cResampler == nil {
		//		fmt.Println("resamplerがない")
		return false
	}
	if frame == nil {
		fmt.Println("frameがない")
		return true
	}
	frame.Restore()
	image.call.callback = callback
	if !C.ImageResizerResampler_resample(
		image.cResampler,
		unsafe.Pointer(frame.CFrame),
		C.uintptr_t(uintptr(unsafe.Pointer(image)))) {
		return false
	}
	return true
}

// Close 破棄
func (image *ImageResizeResampler) Close() {
	C.ImageResizerResampler_close(&image.cResampler)
	image.cResampler = nil
}
