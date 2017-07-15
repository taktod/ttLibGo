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

// go側の関数をkickできるようにしておきます。
extern bool ttLibGoDataCallback(void *, void *data, size_t data_size);

// 書き出し側の処理
static ttLibC_ContainerWriter *ContainerWriter_make(const char *format, void *frames, uint32_t frame_num, uint64_t unit_duration) {
	const char **frame_names = (const char **)frames;
 	if(strcmp(format, "flv") == 0) {
		// 音声のみが欲しい場合は"", "aac"と設定して映像のみが欲しい場合は"h264", "aac"とかと入力すればよいとしておく。
		ttLibC_Frame_Type video_frame = Frame_getFrameType(frame_names[0]);
		ttLibC_Frame_Type audio_frame = Frame_getFrameType(frame_names[1]);
		return (ttLibC_ContainerWriter *)ttLibC_FlvWriter_make(video_frame, audio_frame);
 	}
	ttLibC_Frame_Type *types = ttLibC_malloc(sizeof(ttLibC_Frame_Type) * frame_num);
	for(int i = 0;i < frame_num; ++ i) {
		types[i] = Frame_getFrameType(frame_names[i]);
	}
	if(strcmp(format, "mkv") == 0
	|| strcmp(format, "webm") == 0) {
		ttLibC_ContainerWriter *writer = (ttLibC_ContainerWriter *)ttLibC_MkvWriter_make_ex(types, frame_num, unit_duration);
		writer->type = containerType_webm;
		return writer;
	}
	else if(strcmp(format, "mp4") == 0) {
		return (ttLibC_ContainerWriter *)ttLibC_Mp4Writer_make_ex(types, frame_num, unit_duration);
	}
	else if(strcmp(format, "mpegts") == 0) {
		return (ttLibC_ContainerWriter *)ttLibC_MpegtsWriter_make_ex(types, frame_num, unit_duration);
	}
	return NULL;
}

static bool ContainerWriter_write(
		ttLibC_ContainerWriter *writer,
		void *f,
		uintptr_t ptr) {
	if(writer == NULL) {
		return false;
	}
	if(f == NULL) {
		return true;
	}
	ttLibC_Frame *frame = (ttLibC_Frame *)f;
	switch(writer->type) {
	case containerType_flv:
		return ttLibC_FlvWriter_write((ttLibC_FlvWriter *)writer, frame, ttLibGoDataCallback, (void *)ptr);
	case containerType_mkv:
	case containerType_webm:
		return ttLibC_MkvWriter_write((ttLibC_MkvWriter *)writer, frame, ttLibGoDataCallback, (void *)ptr);
	case containerType_mp4:
		return ttLibC_Mp4Writer_write((ttLibC_Mp4Writer *)writer, frame, ttLibGoDataCallback, (void *)ptr);
	case containerType_mpegts:
		return ttLibC_MpegtsWriter_write((ttLibC_MpegtsWriter *)writer, frame, ttLibGoDataCallback, (void *)ptr);
	default:
		break;
	}
	return false;
}

static bool ContainerWriter_writeInfo(
		ttLibC_ContainerWriter *writer,
		uintptr_t ptr) {
	if(writer == NULL) {
		return false;
	}
	if(writer->type != containerType_mpegts) {
		return true;
	}
	return ttLibC_MpegtsWriter_writeInfo((ttLibC_MpegtsWriter *)writer, ttLibGoDataCallback, (void *)ptr);
}
*/
import "C"
import "unsafe"

// Writer ttLibCを利用したコンテナwriterの実装
type Writer struct {
	call    dataCall
	format  string
	cWriter *C.ttLibC_ContainerWriter
	Mode    int
}

// Init 初期化します。
func (writer *Writer) Init(format string, unitDuration uint64, frames ...string) bool {
	if writer.cWriter == nil {
		cFormat := C.CString(format)
		defer C.free(unsafe.Pointer(cFormat))
		frameList := make([](unsafe.Pointer), len(frames))
		for key, value := range frames {
			frameList[key] = unsafe.Pointer(C.CString(value))
			defer C.free(frameList[key])
		}
		writer.cWriter = C.ContainerWriter_make(cFormat, unsafe.Pointer(&frameList[0]), C.uint32_t(len(frames)), 1000)
		if writer.cWriter == nil {
			return false
		}
		writer.format = format
	}
	return writer.cWriter != nil
}

// WriteFrame フレームからbinaryを作成します
func (writer *Writer) WriteFrame(
	frame *Frame,
	callback DataCallback) bool {
	if writer.cWriter == nil {
		return false
	}
	if frame == nil {
		return true
	}
	frame.Restore()
	writer.cWriter.mode = C.uint32_t(writer.Mode)
	// frameの内容をupdateするなにかが必要になる
	writer.call.callback = callback
	if !C.ContainerWriter_write(
		writer.cWriter,
		unsafe.Pointer(frame.CFrame),
		C.uintptr_t(uintptr(unsafe.Pointer(writer)))) {
		return false
	}
	return true
}

// WriteInfo mpegtsのsdt pad pmtというメタデータを再生成するための動作 これをいれないとhlsみたいに分割したときにデータがおかしくなる。
func (writer *Writer) WriteInfo(
	callback DataCallback) bool {
	if writer.cWriter == nil {
		return false
	}
	if writer.format != "mpegts" {
		return true
	}

	writer.call.callback = callback
	if !C.ContainerWriter_writeInfo(
		writer.cWriter,
		C.uintptr_t(uintptr(unsafe.Pointer(writer)))) {
		return false
	}
	return true
}

// Close フレームの書き出しオブジェクトを解放します。
func (writer *Writer) Close() {
	C.ttLibC_ContainerWriter_close(&writer.cWriter)
	writer.cWriter = nil
}
