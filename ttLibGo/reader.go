package ttLibGo

/*
#include <stdlib.h>
#include <ttLibC/allocator.h>
#include <ttLibC/container/flv.h>
#include <ttLibC/container/mkv.h>
#include <ttLibC/container/mp4.h>
#include <ttLibC/container/mpegts.h>

#include <string.h>
#include <stdbool.h>
#include "frame.h"

// 読み込み側の処理
// go側の関数をkickできるようにしておきます。
extern bool ttLibGoFrameCallback(void *ptr, ttLibC_Frame *frame);

// コンテナリーダーをつくっておく
static ttLibC_ContainerReader *ContainerReader_make(const char *format) {
	// formatの指定文字列でなにをつくるか決定します。
	if(strcmp(format, "flv") == 0) {
		return (ttLibC_ContainerReader *)ttLibC_FlvReader_make();
	}
	else if(strcmp(format, "mkv") == 0
	|| strcmp(format, "webm") == 0) {
		return (ttLibC_ContainerReader *)ttLibC_MkvReader_make();
	}
	else if(strcmp(format, "mp4") == 0) {
		return (ttLibC_ContainerReader *)ttLibC_Mp4Reader_make();
	}
	else if(strcmp(format, "mpegts") == 0) {
		return (ttLibC_ContainerReader *)ttLibC_MpegtsReader_make();
	}
	return NULL;
}

// flvのcallback
static bool ContainerReader_flvCallback(void *ptr, ttLibC_Flv *flv) {
	return ttLibC_Flv_getFrame(flv, ttLibGoFrameCallback, ptr);
}

// mkvのcallback
static bool ContainerReader_mkvCallback(void *ptr, ttLibC_Mkv *mkv) {
	return ttLibC_Mkv_getFrame(mkv, ttLibGoFrameCallback, ptr);
}

// mp4のcallback
static bool ContainerReader_mp4Callback(void *ptr, ttLibC_Mp4 *mp4) {
	return ttLibC_Mp4_getFrame(mp4, ttLibGoFrameCallback, ptr);
}

// mpegtsのcallback
static bool ContainerReader_mpegtsCallback(void *ptr, ttLibC_Mpegts *mpegts) {
	return ttLibC_Mpegts_getFrame(mpegts, ttLibGoFrameCallback, ptr);
}

// 読み込み実施
static bool ContainerReader_read(
		ttLibC_ContainerReader *reader,
		void *data,
		size_t data_size,
		uintptr_t ptr) {
	if(reader == NULL) {
		return false;
	}
	switch(reader->type) {
	case containerType_flv:
		return ttLibC_FlvReader_read((ttLibC_FlvReader *)reader, data, data_size, ContainerReader_flvCallback, (void *)ptr);
	case containerType_mkv:
	case containerType_webm:
		return ttLibC_MkvReader_read((ttLibC_MkvReader *)reader, data, data_size, ContainerReader_mkvCallback, (void *)ptr);
	case containerType_mp4:
		return ttLibC_Mp4Reader_read((ttLibC_Mp4Reader *)reader, data, data_size, ContainerReader_mp4Callback, (void *)ptr);
	case containerType_mpegts:
		return ttLibC_MpegtsReader_read((ttLibC_MpegtsReader *)reader, data, data_size, ContainerReader_mpegtsCallback, (void *)ptr);
	default:
		break;
	}
	return false;
}
*/
import "C"
import (
	"unsafe"
)

// Reader ttLibCを利用したコンテナReaderの実装
type Reader struct {
	call    frameCall
	format  string
	cReader *C.ttLibC_ContainerReader
}

// Init 初期化
func (reader *Reader) Init(format string) bool {
	if reader.cReader == nil {
		cFormat := C.CString(format)
		defer C.free(unsafe.Pointer(cFormat))
		reader.cReader = C.ContainerReader_make(cFormat)
		if reader.cReader == nil {
			return false
		}
		reader.format = format
	}
	return reader.cReader != nil
}

// ReadFrame 実際のフレームをbinaryから読み込みます。
func (reader *Reader) ReadFrame(
	binary []byte,
	length uint64,
	callback FrameCallback) bool {
	// readerができてなかったら、アウト
	if reader.cReader == nil {
		return false
	}
	reader.call.callback = callback
	if !C.ContainerReader_read(
		reader.cReader,
		unsafe.Pointer(&binary[0]),
		C.size_t(length),
		C.uintptr_t(uintptr(unsafe.Pointer(reader)))) {
		// このunsafe.Pointerを無理やりuintptrにキャストしてから、cのポインタ化するというやり方は少々あやういけど
		// 現状の動作では、ContainerReader_readの動作が別スレッドではないので、比較的安全だと思われ
		return false
	}
	return true
}

// Close フレーム読み込みオブジェクトを解放します
func (reader *Reader) Close() {
	C.ttLibC_ContainerReader_close(&reader.cReader)
	reader.cReader = nil
}
