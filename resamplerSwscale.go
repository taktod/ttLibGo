package ttLibGo

/*
#include <stdint.h>
#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>
extern bool SwscaleResampler_resample(void *resampler,
    void *dstCFrame, void *dstGoFrame,
    void *srcCFrame, void *srcGoFrame);
*/
import "C"
import (
	"unsafe"
)

type swscaleResampler resampler

// SwscaleModes Swscaleのモードリスト
var SwscaleModes = struct {
	X            subType
	Area         subType
	Sinc         subType
	Point        subType
	Gauss        subType
	Spline       subType
	Bicubic      subType
	Lanczos      subType
	Bilinear     subType
	Bicublin     subType
	FastBilinear subType
}{
	subType{"X"},
	subType{"Area"},
	subType{"Sinc"},
	subType{"Point"},
	subType{"Gauss"},
	subType{"Spline"},
	subType{"Bicubic"},
	subType{"Lanczos"},
	subType{"Bilinear"},
	subType{"Bicublin"},
	subType{"FastBilinear"},
}

func (swscaleResampler *swscaleResampler) Close() {
	resamplerClose((*resampler)(swscaleResampler))
}

func (swscaleResampler *swscaleResampler) Resample(dst IFrame, src IFrame) bool {
	dstGoFrame := dst.newGoRefFrame()
	srcGoFrame := src.newGoRefFrame()
	result := C.SwscaleResampler_resample(
		unsafe.Pointer(swscaleResampler.cResampler),
		unsafe.Pointer(dst.refCFrame()),
		dstGoFrame,
		unsafe.Pointer(src.refCFrame()),
		srcGoFrame)
	dst.deleteGoRefFrame(dstGoFrame)
	src.deleteGoRefFrame(srcGoFrame)
	return bool(result)
}
