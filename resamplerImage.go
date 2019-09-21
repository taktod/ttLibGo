package ttLibGo

/*
#include <stdint.h>
#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>
extern bool ImageResampler_toBgr(
    void *resampler,
    void *dstCFrame, void *dstGoFrame,
    void *srcCFrame, void *srcGoFrame);

extern bool ImageResampler_toYuv420(
    void *resampler,
    void *dstCFrame, void *dstGoFrame,
    void *srcCFrame, void *srcGoFrame);
*/
import "C"
import (
	"unsafe"
)

type imageResampler resampler

func (imageResampler *imageResampler) Close() {
	resamplerClose((*resampler)(imageResampler))
}

func (imageResampler *imageResampler) ToBgr(
	dst IFrame,
	src IFrame) bool {
	dstGoFrame := dst.newGoRefFrame()
	srcGoFrame := src.newGoRefFrame()
	defer dst.deleteGoRefFrame(dstGoFrame)
	defer src.deleteGoRefFrame(srcGoFrame)
	return bool(C.ImageResampler_toBgr(unsafe.Pointer(imageResampler.cResampler),
		unsafe.Pointer(dst.refCFrame()),
		dstGoFrame,
		unsafe.Pointer(src.refCFrame()),
		srcGoFrame))
}

func (imageResampler *imageResampler) ToYuv(
	dst IFrame,
	src IFrame) bool {
	dstGoFrame := dst.newGoRefFrame()
	srcGoFrame := src.newGoRefFrame()
	defer dst.deleteGoRefFrame(dstGoFrame)
	defer src.deleteGoRefFrame(srcGoFrame)
	return bool(C.ImageResampler_toYuv420(unsafe.Pointer(imageResampler.cResampler),
		unsafe.Pointer(dst.refCFrame()),
		dstGoFrame,
		unsafe.Pointer(src.refCFrame()),
		srcGoFrame))
}
