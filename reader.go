package ttLibGo

/*
#include <stdint.h>
#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>
#include <ttLibC/container/container.h>

extern ttLibC_ContainerReader *ContainerReader_make(const char *format);

// 読み込み実施
extern bool ContainerReader_read(
		ttLibC_ContainerReader *reader,
		void *data,
		size_t data_size,
		uintptr_t ptr);
extern void ContainerReader_close(ttLibC_ContainerReader *reader);
*/
import "C"
import "unsafe"

// readerType readerタイプ
type readerType struct{ value string }

type IReader interface {
	ReadFrame(
		binary []byte,
		length uint64,
		callback func(frame *Frame) bool) bool
	Close()
}

// reader メディアデータ読み込み
type reader struct {
	Type    readerType
	cReader *C.ttLibC_ContainerReader
	call    frameCall
}

// ReaderTypes Readerタイプリスト
var ReaderTypes = struct {
	Flv    readerType
	Mp4    readerType
	Mkv    readerType
	Webm   readerType
	Mpegts readerType
}{
	readerType{"flv"},
	readerType{"mp4"},
	readerType{"mkv"},
	readerType{"webm"},
	readerType{"mpegts"},
}

// readerInit 初期化共通関数
func readerInit(readerType readerType) *reader {
	reader := new(reader)
	reader.Type = readerType
	cFormat := C.CString(reader.Type.value)
	defer C.free(unsafe.Pointer(cFormat))
	reader.cReader = C.ContainerReader_make(cFormat)
	if reader.cReader == nil {
		return nil
	}
	return reader
}

// Readers Readers処理
var Readers = struct {
	// Flv FlvReaderを作ります
	Flv    func() *reader
	Mkv    func() *reader
	Webm   func() *reader
	Mp4    func() *reader
	Mpegts func() *reader
}{
	Flv: func() *reader {
		return readerInit(ReaderTypes.Flv)
	},
	Mkv: func() *reader {
		return readerInit(ReaderTypes.Mkv)
	},
	Webm: func() *reader {
		return readerInit(ReaderTypes.Webm)
	},
	Mp4: func() *reader {
		return readerInit(ReaderTypes.Mp4)
	},
	Mpegts: func() *reader {
		return readerInit(ReaderTypes.Mpegts)
	},
}

// ReadFrame Frame読み込み
func (reader *reader) ReadFrame(
	binary []byte,
	length uint64,
	callback func(frame *Frame) bool) bool {
	if reader.cReader == nil {
		return false
	}
	reader.call.callback = callback
	if !C.ContainerReader_read(
		reader.cReader,
		unsafe.Pointer(&binary[0]),
		C.size_t(length),
		C.uintptr_t(uintptr(unsafe.Pointer(&reader.call)))) {
		return false
	}
	return true
}

// Close 閉じる
func (reader *reader) Close() {
	C.ContainerReader_close(reader.cReader)
	reader.cReader = nil
}
