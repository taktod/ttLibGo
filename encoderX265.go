package ttLibGo

/*
#include <stdint.h>
#include <stdlib.h>
#include <stdbool.h>
extern bool X265Encoder_forceNextFrameType(void *encoder, const char *type);
*/
import "C"
import (
	"unsafe"
)

type x265Encoder encoder

// X265EncoderFrameTypes X264で指定できる、frameTypeリスト
var X265EncoderFrameTypes = struct {
	I    subType
	P    subType
	B    subType
	IDR  subType
	Auto subType
	Bref subType
}{
	subType{"I"},
	subType{"P"},
	subType{"B"},
	subType{"IDR"},
	subType{"Auto"},
	subType{"Bref"},
}

// EncodeFrame エンコードを実行
func (x265Encoder *x265Encoder) EncodeFrame(
	frame IFrame,
	callback FrameCallback) bool {
	return encoderEncodeFrame((*encoder)(x265Encoder), frame, callback)
}

// Close 閉じる
func (x265Encoder *x265Encoder) Close() {
	encoderClose((*encoder)(x265Encoder))
}

func (x265Encoder *x265Encoder) ForceNextFrameType(frameType subType) bool {
	ctype := C.CString(frameType.value)
	defer C.free(unsafe.Pointer(ctype))
	return bool(C.X265Encoder_forceNextFrameType(unsafe.Pointer(x265Encoder.cEncoder), ctype))
}
