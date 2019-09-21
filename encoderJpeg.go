package ttLibGo

/*
#include <stdint.h>
#include <stdlib.h>
#include <stdbool.h>
extern bool JpegEncoder_setQuality(void *encoder, uint32_t quality);
*/
import "C"
import (
	"unsafe"
)

type jpegEncoder encoder

// EncodeFrame エンコードを実行
func (jpegEncoder *jpegEncoder) EncodeFrame(
	frame IFrame,
	callback func(frame *Frame) bool) bool {
	return encoderEncodeFrame((*encoder)(jpegEncoder), frame, callback)
}

// Close 閉じる
func (jpegEncoder *jpegEncoder) Close() {
	encoderClose((*encoder)(jpegEncoder))
}

// SetQuality jpegEncodeの動作qualityを設定する 0 - 100
func (jpegEncoder *jpegEncoder) SetQuality(quality uint32) bool {
	return bool(C.JpegEncoder_setQuality(unsafe.Pointer(jpegEncoder.cEncoder), C.uint32_t(quality)))
}
