package ttLibGo

/*
#include <stdint.h>
#include <stdlib.h>
#include <stdbool.h>
extern bool X264Encoder_forceNextFrameType(void *encoder, const char *type);
*/
import "C"
import (
	"unsafe"
)

type x264Encoder encoder

// X264EncoderFrameTypes X264で指定できる、frameTypeリスト
var X264EncoderFrameTypes = struct {
	I        subType
	P        subType
	B        subType
	IDR      subType
	Auto     subType
	Bref     subType
	KeyFrame subType
}{
	subType{"I"},
	subType{"P"},
	subType{"B"},
	subType{"IDR"},
	subType{"Auto"},
	subType{"Bref"},
	subType{"KeyFrame"},
}

// EncodeFrame エンコードを実行
func (x264Encoder *x264Encoder) EncodeFrame(
	frame IFrame,
	callback FrameCallback) bool {
	return encoderEncodeFrame((*encoder)(x264Encoder), frame, callback)
}

// Close 閉じる
func (x264Encoder *x264Encoder) Close() {
	encoderClose((*encoder)(x264Encoder))
}

func (x264Encoder *x264Encoder) ForceNextFrameType(frameType subType) bool {
	ctype := C.CString(frameType.value)
	defer C.free(unsafe.Pointer(ctype))
	return bool(C.X264Encoder_forceNextFrameType(unsafe.Pointer(x264Encoder.cEncoder), ctype))
}
