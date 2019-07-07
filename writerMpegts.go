package ttLibGo

/*
#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>

extern bool ContainerWriter_writeInfo(void *writer, uintptr_t ptr);
*/
import "C"
import "unsafe"

type mpegtsWriter writer

// WriteFrame Frame書き出し
func (mpegtsWriter *mpegtsWriter) WriteFrame(
	frame IFrame,
	callback DataCallback) bool {
	return writerWriteFrame((*writer)(mpegtsWriter), frame, callback)
}

// Close 閉じる
func (mpegtsWriter *mpegtsWriter) Close() {
	writerClose((*writer)(mpegtsWriter))
}

// WriteInfo 情報タグ書き出し
func (mpegtsWriter *mpegtsWriter) WriteInfo(
	callback DataCallback) bool {
	if mpegtsWriter.cWriter == nil {
		return false
	}
	mpegtsWriter.call.callback = callback
	result := C.ContainerWriter_writeInfo(
		unsafe.Pointer(mpegtsWriter.cWriter),
		C.uintptr_t(uintptr(unsafe.Pointer(&mpegtsWriter.call))))
	return bool(result)
}
