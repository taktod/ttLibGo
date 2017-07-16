package ttLibGoLibyuv

/*
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <ttLibC/resampler/libyuvResampler.h>
#include <ttLibC/allocator.h>

extern bool ttLibGoLibyuvFrameCallback(void *ptr, ttLibC_Frame *frame);

typedef struct libyuvRotateResampler_t {
	ttLibC_LibyuvRotate_Mode degree;
	ttLibC_Yuv420 *frame;
} libyuvRotateResampler_t;

libyuvRotateResampler_t *LibyuvRotateResampler_make(
		uint32_t degree) {
	libyuvRotateResampler_t *resampler = ttLibC_malloc(sizeof(libyuvRotateResampler_t));
	if(resampler == NULL) {
		puts("failed to allocate libyuvRotateResampler");
		return NULL;
	}
	switch(degree) {
	case 0:
		{
			resampler->degree = LibyuvRotate_0;
		}
		break;
	case 90:
		{
			resampler->degree = LibyuvRotate_90;
		}
		break;
	case 180:
		{
			resampler->degree = LibyuvRotate_180;
		}
		break;
	case 270:
		{
			resampler->degree = LibyuvRotate_270;
		}
		break;
	default:
		{
			ttLibC_free(resampler);
			puts("degreeが不正です");
		}
		return NULL;
	}
	resampler->frame = NULL;
	return resampler;
}

bool LibyuvRotateResampler_resample(
		libyuvRotateResampler_t *resampler,
		void *frame,
		uintptr_t ptr) {
	if(resampler == NULL) {
		return false;
	}
	if(frame == NULL) {
		return true;
	}
	ttLibC_Frame *f = (ttLibC_Frame *)frame;
	if(f->type != frameType_yuv420) {
		puts("入力フレームはyuv420のみ対応です。");
		return false;
	}
	ttLibC_Yuv420 *yuv = (ttLibC_Yuv420 *)frame;
	if(yuv->type != Yuv420Type_planar) {
		puts("サブタイプはyuv420_Planarのみ対応です");
		return false;
	}
	ttLibC_Yuv420 *y = ttLibC_LibyuvResampler_rotate(resampler->frame, yuv, resampler->degree);
	if(y == NULL) {
		puts("resample失敗しました。");
		return false;
	}
	resampler->frame = y;
	return ttLibGoLibyuvFrameCallback((void *)ptr, (ttLibC_Frame *)resampler->frame);
}

void LibyuvRotateResampler_close(libyuvRotateResampler_t **resampler) {
	libyuvRotateResampler_t *target = *resampler;
	if(target == NULL) {
		return;
	}
	ttLibC_Yuv420_close(&target->frame);
	ttLibC_free(target);
	*resampler = NULL;
}
*/
import "C"
import (
	"unsafe"

	"github.com/taktod/ttLibGo/ttLibGo"
)

// LibyuvRotateResampler libyuvを利用した回転処理
type LibyuvRotateResampler struct {
	call       frameCall
	cResampler *C.libyuvRotateResampler_t
}

// Init 初期化 degreeは90 180 270 0を設定できます
func (libyuvRotate *LibyuvRotateResampler) Init(
	degree uint32) bool {
	if libyuvRotate.cResampler == nil {
		libyuvRotate.cResampler = C.LibyuvRotateResampler_make(C.uint32_t(degree))
	}
	return libyuvRotate.cResampler != nil
}

// Resample リサンプル実行
func (libyuvRotate *LibyuvRotateResampler) Resample(
	frame *ttLibGo.Frame,
	callback ttLibGo.FrameCallback) bool {
	if libyuvRotate.cResampler == nil {
		return false
	}
	if frame == nil {
		return true
	}
	frame.Restore()
	libyuvRotate.call.callback = callback
	if !C.LibyuvRotateResampler_resample(
		libyuvRotate.cResampler,
		unsafe.Pointer(frame.CFrame),
		C.uintptr_t(uintptr(unsafe.Pointer(libyuvRotate)))) {
		return false
	}
	return true
}

// Close 解放
func (libyuvRotate *LibyuvRotateResampler) Close() {
	C.LibyuvRotateResampler_close(&libyuvRotate.cResampler)
	libyuvRotate.cResampler = nil
}
