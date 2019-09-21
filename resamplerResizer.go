package ttLibGo

/*
#include <stdint.h>
#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>
extern bool Resizer_resize(
    void *resampler,
    void *dstCFrame, void *dstGoFrame,
    void *srcCFrame, void *srcGoFrame,
		bool is_quick);
*/
import "C"
import (
	"unsafe"
)

type resizeResampler resampler

func (resizeResampler *resizeResampler) Close() {
	resamplerClose((*resampler)(resizeResampler))
}

func (resizeResampler *resizeResampler) Resize(
	dst IFrame,
	src IFrame,
	isQuick bool) bool {
	dstGoFrame := dst.newGoRefFrame()
	srcGoFrame := src.newGoRefFrame()
	defer dst.deleteGoRefFrame(dstGoFrame)
	defer src.deleteGoRefFrame(srcGoFrame)
	return bool(C.Resizer_resize(unsafe.Pointer(resizeResampler.cResampler),
		unsafe.Pointer(dst.refCFrame()),
		dstGoFrame,
		unsafe.Pointer(src.refCFrame()),
		srcGoFrame,
		C.bool(isQuick)))
}
