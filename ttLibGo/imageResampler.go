package ttLibGo

/*
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <ttLibC/allocator.h>
#include <ttLibC/resampler/imageResampler.h>
#include <ttLibC/frame/video/bgr.h>
#include <ttLibC/frame/video/yuv420.h>
#include "frame.h"

extern bool ttLibGoFrameCallback(void *ptr, ttLibC_Frame *frame);

typedef struct imageResampler_t {
	ttLibC_Frame_Type type;
	uint32_t          sub_type;
	ttLibC_Frame     *frame;
} imageResampler_t;

imageResampler_t *ImageResampler_make(
		const char *codecType,
		const char *sub_type) {
	ttLibC_Frame_Type type = Frame_getFrameType(codecType);
	switch(type) {
	case frameType_bgr:
	case frameType_yuv420:
		break;
	default:
		return NULL;
	}
	imageResampler_t *resampler = ttLibC_malloc(sizeof(imageResampler_t));
	if(resampler == NULL) {
		puts("failed to allocate imageResamplerObject.");
		return NULL;
	}
	switch(type) {
	case frameType_bgr:
		{
			resampler->type = type;
			if(strcmp(sub_type, "bgr") == 0) {
				resampler->sub_type = BgrType_bgr;
			}
			else if(strcmp(sub_type, "abgr") == 0) {
				resampler->sub_type = BgrType_abgr;
			}
			else if(strcmp(sub_type, "bgra") == 0) {
				resampler->sub_type = BgrType_bgra;
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
	resampler->frame = NULL;
	return resampler;
}

bool ImageResampler_resample(
		imageResampler_t *resampler,
		void *frame,
		uintptr_t ptr) {
	if(resampler == NULL) {
		return false;
	}
	if(frame == NULL) {
		return true;
	}
	ttLibC_Frame *f = (ttLibC_Frame *)frame;
	// こっちは打ち返しはしない。yuv -> yuv変換等ができないので、あわない場合は、処理しないとする
	switch(f->type) {
	case frameType_bgr:
		{
			if(resampler->type != frameType_yuv420) {
				puts("入力フレームタイプが不正です。");
				return false;
			}
			ttLibC_Yuv420 *y = ttLibC_ImageResampler_makeYuv420FromBgr((ttLibC_Yuv420 *)resampler->frame, resampler->sub_type, (ttLibC_Bgr *)frame);
			if(y == NULL) {
				puts("resample失敗しました。データができてません。");
				return false;
			}
			resampler->frame = (ttLibC_Frame *)y;
		}
		break;
	case frameType_yuv420:
		{
			if(resampler->type != frameType_bgr) {
				puts("入力フレームタイプが不正です。");
				return false;
			}
			ttLibC_Bgr *b = ttLibC_ImageResampler_makeBgrFromYuv420((ttLibC_Bgr *)resampler->frame, resampler->sub_type, (ttLibC_Yuv420 *)frame);
			if(b == NULL) {
				puts("resample失敗しました。データができてません。");
				return false;
			}
			resampler->frame = (ttLibC_Frame *)b;
		}
		break;
	default:
		return false;
	}
	return ttLibGoFrameCallback((void *)ptr, resampler->frame);
}

void ImageResampler_close(imageResampler_t **resampler) {
	imageResampler_t *target = *resampler;
	if(target == NULL) {
		return;
	}
	ttLibC_Frame_close(&target->frame);
	ttLibC_free(target);
	*resampler = NULL;
}
*/
import "C"
import "unsafe"

// ImageResampler bgrとyuvを相互変換するresampler
type ImageResampler struct {
	call       frameCall
	cResampler *C.imageResampler_t
}

// Init 初期化します。 変換後の設定を与えてください
func (image *ImageResampler) Init(
	codecType string,
	subType string) bool {
	if image.cResampler == nil {
		cCodecType := C.CString(codecType)
		defer C.free(unsafe.Pointer(cCodecType))
		cSubType := C.CString(subType)
		defer C.free(unsafe.Pointer(cSubType))
		image.cResampler = C.ImageResampler_make(
			cCodecType,
			cSubType)
	}
	return image.cResampler != nil
}

// Resample 実際に変換を実施します
func (image *ImageResampler) Resample(
	frame *Frame,
	callback FrameCallback) bool {
	if image.cResampler == nil {
		return false
	}
	if frame == nil {
		return true
	}
	frame.Restore()
	image.call.callback = callback
	if !C.ImageResampler_resample(
		image.cResampler,
		unsafe.Pointer(frame.CFrame),
		C.uintptr_t(uintptr(unsafe.Pointer(image)))) {
		return false
	}
	return true
}

// Close 必要なくなったら解放します
func (image *ImageResampler) Close() {
	C.ImageResampler_close(&image.cResampler)
	image.cResampler = nil
}
