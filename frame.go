package ttLibGo

/*
#include <stdio.h>
#include <stdlib.h>
#include <stdint.h>
#include <ttLibC/frame/frame.h>
extern void Frame_getFrameType(void *cFrame, char *buffer, size_t buffer_size);
extern uint64_t Frame_getPts(void *frame);
extern uint64_t Frame_getDts(void *frame);
extern uint32_t Frame_getTimebase(void *frame);
extern uint32_t Frame_getID(void *frame);
extern void Frame_close(void *frame);
extern ttLibC_Frame *Frame_clone(void *frame);
extern void *newGoFrame(uint64_t pts, uint64_t dts, uint32_t timebase, uint32_t id,
    uint32_t width, uint32_t height,
    uint32_t sampleRate, uint32_t sampleNum, uint32_t channelNum,
    uint64_t dataPos, uint32_t widthStride,
    uint64_t yDataPos, uint32_t yStride,
    uint64_t uDataPos, uint32_t uStride,
    uint64_t vDataPos, uint32_t vStride);
extern void deleteGoFrame(void *frame);
extern bool Frame_getBinaryBuffer(uintptr_t ptr, void *frame);
*/
import "C"
import (
	"bytes"
	"unsafe"
)

// cttLibCFrame c言語のframeデータ
type cttLibCFrame unsafe.Pointer

// IFrame Frameのインターフェイス
type IFrame interface {
	newGoRefFrame() unsafe.Pointer
	refCFrame() cttLibCFrame
	deleteGoRefFrame(ptr unsafe.Pointer)
}

// frameType frameタイプ
type frameType struct{ value string }

// subType subタイプ
type subType struct{ value string }

// FrameTypes Frameタイプリスト
var FrameTypes = struct {
	Bgr         frameType
	Flv1        frameType
	H264        frameType
	H265        frameType
	Jpeg        frameType
	Png         frameType
	Theora      frameType
	Vp6         frameType
	Vp8         frameType
	Vp9         frameType
	Wmv1        frameType
	Wmv2        frameType
	Yuv420      frameType
	Aac         frameType
	AdpcmImaWav frameType
	Mp3         frameType
	Nellymoser  frameType
	Opus        frameType
	PcmAlaw     frameType
	PcmF32      frameType
	PcmMulaw    frameType
	PcmS16      frameType
	Speex       frameType
	Vorbis      frameType
	Unknown     frameType
}{
	frameType{"bgr"},
	frameType{"flv1"},
	frameType{"h264"},
	frameType{"h265"},
	frameType{"jpeg"},
	frameType{"png"},
	frameType{"theora"},
	frameType{"vp6"},
	frameType{"vp8"},
	frameType{"vp9"},
	frameType{"wmv1"},
	frameType{"wmv2"},
	frameType{"yuv420"},
	frameType{"aac"},
	frameType{"adpcmImaWav"},
	frameType{"mp3"},
	frameType{"nellymoser"},
	frameType{"opus"},
	frameType{"pcmAlaw"},
	frameType{"pcmF32"},
	frameType{"pcmMulaw"},
	frameType{"pcmS16"},
	frameType{"speex"},
	frameType{"vorbis"},
	frameType{""},
}

// Frame フレーム
type Frame struct {
	// Type 対ーぷ
	Type     frameType
	Pts      uint64
	Dts      uint64
	Timebase uint32
	ID       uint32

	cFrame  cttLibCFrame
	hasBody bool // このデータが解放する時に、closeすべきかのフラグ
}

// init データを使って初期化します。
func (frame *Frame) init(cframe cttLibCFrame) {
	// cのframeデータを元にどういうデータであるかを取得しなければならない。
	if cframe == nil {
		panic("nullなデータのframe初期化を実施しようとしました。")
	}
	frame.cFrame = cframe
	buf := make([]byte, 256)
	ptr := unsafe.Pointer(cframe)
	C.Frame_getFrameType(ptr, (*C.char)(unsafe.Pointer(&buf[0])), C.size_t(256))
	name := string(buf[:bytes.IndexByte(buf, 0)])
	frame.Type = frameType{name}
	frame.ID = uint32(C.Frame_getID(ptr))
	frame.Pts = uint64(C.Frame_getPts(ptr))
	frame.Dts = uint64(C.Frame_getDts(ptr))
	frame.Timebase = uint32(C.Frame_getTimebase(ptr))
}

// Close フレームを解放します
func (frame *Frame) Close() {
	if frame.hasBody {
		C.Frame_close(unsafe.Pointer(frame.cFrame))
	}
	frame.cFrame = nil
}

// Clone フレームのcloneを作成し、応答します
func (frame *Frame) Clone() *Frame {
	clonedCFrame := C.Frame_clone(unsafe.Pointer(frame.cFrame))
	if clonedCFrame == nil {
		return nil
	}
	frame.init(cttLibCFrame(clonedCFrame))
	frame.hasBody = true
	return frame
}

// newGoRefFrame goの配列データをcの世界に持っていくデータを構築する
func (frame *Frame) newGoRefFrame() unsafe.Pointer {
	return C.newGoFrame(C.uint64_t(frame.Pts), C.uint64_t(frame.Dts), C.uint32_t(frame.Timebase), C.uint32_t(frame.ID),
		0, 0,
		0, 0, 0,
		0, 0,
		0, 0,
		0, 0,
		0, 0)
}

// refCFrame interface経由でcframe参照
func (frame *Frame) refCFrame() cttLibCFrame {
	return frame.cFrame
}

// deleteGoRefFrame cの世界に持っていったデータを破棄する
func (frame *Frame) deleteGoRefFrame(ptr unsafe.Pointer) {
	C.deleteGoFrame(ptr)
}

// getFrameBinaryBuffer frameのbyteデータを参照します
func getFrameBinaryBuffer(cFrame cttLibCFrame, callback DataCallback) bool {
	call := new(dataCall)
	call.callback = callback
	return bool(C.Frame_getBinaryBuffer(C.uintptr_t(uintptr(unsafe.Pointer(call))), unsafe.Pointer(cFrame)))
}

// GetBinaryBuffer binaryデータを参照する
func (frame *Frame) GetBinaryBuffer(callback DataCallback) bool {
	return getFrameBinaryBuffer(frame.cFrame, callback)
}
