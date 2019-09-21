package ttLibGo

/*
#include <stdint.h>
#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>

extern void *ContainerWriter_make(void *mp);
extern bool ContainerWriter_writeFrame(void *writer, void *frame, void *goFrame, uint32_t mode, uint32_t unitDuration, uintptr_t ptr);
extern void ContainerWriter_close(void *writer);
extern uint64_t ContainerWriter_getPts(void *writer);
extern uint32_t ContainerWriter_getTimebase(void *writer);
*/
import "C"
import (
	"unsafe"
)

type cttlibCWriter unsafe.Pointer

// writerType writerタイプ
type writerType struct{ value string }

// IWriter メディアデータの書き出し
type IWriter interface {
	WriteFrame(
		frame IFrame,
		callback func(data []byte) bool) bool
	Close()
}

// writer メディアデータ書き出し
type writer struct {
	Type         writerType
	Pts          uint64
	Timebase     uint32
	Mode         uint32
	UnitDuration uint32
	cWriter      cttlibCWriter
	call         dataCall
}

// WriterModes Writerのモードリスト or で結合可能
var WriterModes = struct {
	EnableDts        uint32
	SplitAllKeyFrame uint32
	//SplitKeyFrame         uint32
	SplitInnerFrame       uint32
	SplitPFrame           uint32
	SplitDisposableBFrame uint32
	SplitBFrame           uint32
}{
	EnableDts:             0x20,
	SplitAllKeyFrame:      0x10,
	SplitInnerFrame:       0x01,
	SplitPFrame:           0x02,
	SplitDisposableBFrame: 0x04,
	SplitBFrame:           0x08,
}

// WriterTypes Writerタイプリスト
var WriterTypes = struct {
	Flv    writerType
	Mp4    writerType
	Mkv    writerType
	Mpegts writerType
	Webm   writerType
}{
	Flv:    writerType{"flv"},
	Mp4:    writerType{"mp4"},
	Mkv:    writerType{"mkv"},
	Mpegts: writerType{"mpegts"},
	Webm:   writerType{"webm"},
}

func writerInit(writer *writer, params map[string]interface{}) *writer {
	v := mapUtil.fromMap(params)
	writer.cWriter = cttlibCWriter(C.ContainerWriter_make(v))
	writer.Timebase = uint32(C.ContainerWriter_getTimebase(unsafe.Pointer(writer.cWriter)))
	mapUtil.close(v)
	return writer
}

// Writers Writers処理
var Writers = struct {
	Flv    func(videoType frameType, audioType frameType) *writer
	Mkv    func(types ...frameType) *writer
	Webm   func(types ...frameType) *writer
	Mp4    func(types ...frameType) *writer
	Mpegts func(types ...frameType) *mpegtsWriter
}{
	Flv: func(videoType frameType, audioType frameType) *writer {
		writer := writerInit(new(writer), map[string]interface{}{
			"type":      WriterTypes.Flv.value,
			"videoType": videoType.value,
			"audioType": audioType.value,
		})
		writer.Type = WriterTypes.Flv
		writer.UnitDuration = 1000
		return writer
	},
	Mkv: func(types ...frameType) *writer {
		var frameNames []string
		for _, v := range types {
			frameNames = append(frameNames, v.value)
		}
		writer := writerInit(new(writer), map[string]interface{}{
			"type":       WriterTypes.Mkv.value,
			"frameTypes": frameNames,
		})
		writer.Type = WriterTypes.Mkv
		writer.UnitDuration = 1000
		return writer
	},
	Webm: func(types ...frameType) *writer {
		var frameNames []string
		for _, v := range types {
			frameNames = append(frameNames, v.value)
		}
		writer := writerInit(new(writer), map[string]interface{}{
			"type":       WriterTypes.Webm.value,
			"frameTypes": frameNames,
		})
		writer.Type = WriterTypes.Webm
		writer.UnitDuration = 1000
		return writer
	},
	Mp4: func(types ...frameType) *writer {
		var frameNames []string
		for _, v := range types {
			frameNames = append(frameNames, v.value)
		}
		writer := writerInit(new(writer), map[string]interface{}{
			"type":       WriterTypes.Mp4.value,
			"frameTypes": frameNames,
		})
		writer.Type = WriterTypes.Mp4
		writer.UnitDuration = 1000
		return writer
	},
	Mpegts: func(types ...frameType) *mpegtsWriter {
		var frameNames []string
		for _, v := range types {
			frameNames = append(frameNames, v.value)
		}
		writer := writerInit((*writer)(new(mpegtsWriter)), map[string]interface{}{
			"type":       WriterTypes.Mpegts.value,
			"frameTypes": frameNames,
		})
		writer.Type = WriterTypes.Mpegts
		writer.UnitDuration = 180000
		return (*mpegtsWriter)(writer)
	},
}

func writerWriteFrame(
	writer *writer,
	frame IFrame,
	callback func(data []byte) bool) bool {
	if writer.cWriter == nil {
		return false
	}
	if frame == nil {
		return true
	}
	cframe := frame.newGoRefFrame()
	writer.call.callback = func(data []byte) bool {
		writer.Pts = uint64(C.ContainerWriter_getPts(unsafe.Pointer(writer.cWriter)))
		writer.Timebase = uint32(C.ContainerWriter_getTimebase(unsafe.Pointer(writer.cWriter)))
		return callback(data)
	}
	result := C.ContainerWriter_writeFrame(
		unsafe.Pointer(writer.cWriter),
		unsafe.Pointer(frame.refCFrame()),
		cframe,
		C.uint32_t(writer.Mode),
		C.uint32_t(writer.UnitDuration),
		C.uintptr_t(uintptr(unsafe.Pointer(&writer.call))))
	frame.deleteGoRefFrame(cframe)
	return bool(result)
}

func writerClose(writer *writer) {
	C.ContainerWriter_close(unsafe.Pointer(writer.cWriter))
}

// WriteFrame Frame書き出し
func (writer *writer) WriteFrame(
	frame IFrame,
	callback func(data []byte) bool) bool {
	return writerWriteFrame(writer, frame, callback)
}

// Close 閉じる
func (writer *writer) Close() {
	writerClose(writer)
}
