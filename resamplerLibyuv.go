package ttLibGo

/*
#include <stdint.h>
#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>
extern bool LibyuvResampler_resize(
    void *resampler,
    void *dstCFrame, void *dstGoFrame,
    void *srcCFrame, void *srcGoFrame,
    const char *mode,
    const char *sub_mode);

extern bool LibyuvResampler_rotate(
    void *resampler,
    void *dstCFrame, void *dstGoFrame,
    void *srcCFrame, void *srcGoFrame,
    const char *mode);

extern bool LibyuvResampler_toBgr(
    void *resampler,
    void *dstCFrame, void *dstGoFrame,
    void *srcCFrame, void *srcGoFrame);

extern bool LibyuvResampler_toYuv420(
    void *resampler,
    void *dstCFrame, void *dstGoFrame,
    void *srcCFrame, void *srcGoFrame);
*/
import "C"
import (
	"unsafe"
)

var LibyuvModes = struct {
	None     subType
	Linear   subType
	Bilinear subType
	Box      subType
}{
	subType{"None"},
	subType{"Linear"},
	subType{"Bilinear"},
	subType{"Box"},
}

var LibyuvRotateModes = struct {
	Degree0   subType
	Degree90  subType
	Degree180 subType
	Degree270 subType
}{
	subType{"0"},
	subType{"90"},
	subType{"180"},
	subType{"270"},
}

type libyuvResampler resampler

func (libyuvResampler *libyuvResampler) Close() {
	resamplerClose((*resampler)(libyuvResampler))
}

func (libyuvResampler *libyuvResampler) ToBgr(
	dst IFrame,
	src IFrame) bool {
	dstGoFrame := dst.newGoRefFrame()
	srcGoFrame := src.newGoRefFrame()
	defer dst.deleteGoRefFrame(dstGoFrame)
	defer src.deleteGoRefFrame(srcGoFrame)
	return bool(C.LibyuvResampler_toBgr(unsafe.Pointer(libyuvResampler.cResampler),
		unsafe.Pointer(dst.refCFrame()),
		dstGoFrame,
		unsafe.Pointer(src.refCFrame()),
		srcGoFrame))
}

func (libyuvResampler *libyuvResampler) ToYuv(
	dst IFrame,
	src IFrame) bool {
	dstGoFrame := dst.newGoRefFrame()
	srcGoFrame := src.newGoRefFrame()
	defer dst.deleteGoRefFrame(dstGoFrame)
	defer src.deleteGoRefFrame(srcGoFrame)
	return bool(C.LibyuvResampler_toYuv420(unsafe.Pointer(libyuvResampler.cResampler),
		unsafe.Pointer(dst.refCFrame()),
		dstGoFrame,
		unsafe.Pointer(src.refCFrame()),
		srcGoFrame))
}

func (libyuvResampler *libyuvResampler) Resize(
	dst IFrame,
	src IFrame,
	mode subType,
	subMode subType) bool {
	dstGoFrame := dst.newGoRefFrame()
	srcGoFrame := src.newGoRefFrame()
	defer dst.deleteGoRefFrame(dstGoFrame)
	defer src.deleteGoRefFrame(srcGoFrame)
	cMode := C.CString(mode.value)
	defer C.free(unsafe.Pointer(cMode))
	cSubMode := C.CString(subMode.value)
	defer C.free(unsafe.Pointer(cSubMode))
	return bool(C.LibyuvResampler_resize(unsafe.Pointer(libyuvResampler.cResampler),
		unsafe.Pointer(dst.refCFrame()),
		dstGoFrame,
		unsafe.Pointer(src.refCFrame()),
		srcGoFrame,
		cMode,
		cSubMode))
}

func (libyuvResampler *libyuvResampler) Rotate(
	dst IFrame,
	src IFrame,
	mode subType) bool {
	dstGoFrame := dst.newGoRefFrame()
	srcGoFrame := src.newGoRefFrame()
	defer dst.deleteGoRefFrame(dstGoFrame)
	defer src.deleteGoRefFrame(srcGoFrame)
	cMode := C.CString(mode.value)
	defer C.free(unsafe.Pointer(cMode))
	return bool(C.LibyuvResampler_rotate(unsafe.Pointer(libyuvResampler.cResampler),
		unsafe.Pointer(dst.refCFrame()),
		dstGoFrame,
		unsafe.Pointer(src.refCFrame()),
		srcGoFrame,
		cMode))
}
