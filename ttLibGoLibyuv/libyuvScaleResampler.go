package ttLibGoLibyuv

/*
#include <stdbool.h>
#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include <ttLibC/resampler/libyuvResampler.h>
#include <ttLibC/allocator.h>

extern bool ttLibGoLibyuvFrameCallback(void *ptr, ttLibC_Frame *frame);

typedef struct libyuvScaleResampler_t {
	uint32_t width;
	uint32_t height;
	ttLibC_LibyuvFilter_Mode y_filter;
	ttLibC_LibyuvFilter_Mode u_filter;
	ttLibC_LibyuvFilter_Mode v_filter;
	ttLibC_Yuv420 *frame;
} libyuvScaleResampler_t;

static ttLibC_LibyuvFilter_Mode libyuvScaleResampler_getFilter(const char *filter) {
	if(strcmp(filter, "box") == 0) {
		return LibyuvFilter_Box;
	}
	else if(strcmp(filter, "linear") == 0) {
		return LibyuvFilter_Linear;
	}
	else if(strcmp(filter, "bilinear") == 0) {
		return LibyuvFilter_Bilinear;
	}
	return LibyuvFilter_None;
}

libyuvScaleResampler_t *LibyuvScaleResampler_make(
		uint32_t width,
		uint32_t height,
		const char *yFilter,
		const char *uFilter,
		const char *vFilter) {
	libyuvScaleResampler_t *resampler = ttLibC_malloc(sizeof(libyuvScaleResampler_t));
	resampler->width = width;
	resampler->height = height;
	resampler->y_filter = libyuvScaleResampler_getFilter(yFilter);
	resampler->u_filter = libyuvScaleResampler_getFilter(uFilter);
	resampler->v_filter = libyuvScaleResampler_getFilter(vFilter);
	resampler->frame = NULL;
	return resampler;
}

bool LibyuvScaleResampler_resample(
		libyuvScaleResampler_t *resampler,
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
	ttLibC_Yuv420 *y = ttLibC_LibyuvResampler_resize(
		resampler->frame,
		resampler->width,
		resampler->height,
		yuv,
		resampler->y_filter,
		resampler->u_filter,
		resampler->v_filter);
	if(y == NULL) {
		puts("resample失敗しました。");
		return false;
	}
	resampler->frame = y;
	return ttLibGoLibyuvFrameCallback((void *)ptr, (ttLibC_Frame *)resampler->frame);
}

void LibyuvScaleResampler_close(libyuvScaleResampler_t **resampler) {
	libyuvScaleResampler_t *target = *resampler;
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

// LibyuvScaleResampler libyuvを利用した拡大縮小処理
type LibyuvScaleResampler struct {
	call       frameCall
	cResampler *C.libyuvScaleResampler_t
}

// Init 初期化
func (libyuvScale *LibyuvScaleResampler) Init(
	width uint32,
	height uint32,
	yFilter string,
	uFilter string,
	vFilter string) bool {
	if libyuvScale.cResampler == nil {
		cYFilter := C.CString(yFilter)
		defer C.free(unsafe.Pointer(cYFilter))
		cUFilter := C.CString(uFilter)
		defer C.free(unsafe.Pointer(cUFilter))
		cVFilter := C.CString(vFilter)
		defer C.free(unsafe.Pointer(cVFilter))
		libyuvScale.cResampler = C.LibyuvScaleResampler_make(
			C.uint32_t(width),
			C.uint32_t(height),
			cYFilter,
			cUFilter,
			cVFilter)
	}
	return libyuvScale.cResampler != nil
}

// Resample リサンプル実行
func (libyuvScale *LibyuvScaleResampler) Resample(
	frame *ttLibGo.Frame,
	callback ttLibGo.FrameCallback) bool {
	if libyuvScale.cResampler == nil {
		return false
	}
	if frame == nil {
		return true
	}
	frame.Restore()
	libyuvScale.call.callback = callback
	if !C.LibyuvScaleResampler_resample(
		libyuvScale.cResampler,
		unsafe.Pointer(frame.CFrame),
		C.uintptr_t(uintptr(unsafe.Pointer(libyuvScale)))) {
		return false
	}
	return true
}

// Close 解放
func (libyuvScale *LibyuvScaleResampler) Close() {
	C.LibyuvScaleResampler_close(&libyuvScale.cResampler)
	libyuvScale.cResampler = nil
}
