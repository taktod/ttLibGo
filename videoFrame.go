package ttLibGo

/*
#include <stdio.h>
#include <stdint.h>
#include <stdlib.h>
#include <stdbool.h>
extern void VideoFrame_getVideoType(void *cFrame, char *buffer, size_t buffer_size);
extern uint32_t VideoFrame_getWidth(void *frame);
extern uint32_t VideoFrame_getHeight(void *frame);
extern void BgrFrame_getBgrType(void *cFrame, char *buffer, size_t buffer_size);
extern uint32_t BgrFrame_getWidthStride(void *frame);
extern void Flv1Frame_getFlv1Type(void *cFrame, char *buffer, size_t buffer_size);
extern void H264Frame_getH264Type(void *cFrame, char *buffer, size_t buffer_size);
extern uint32_t H264Frame_getH264FrameType(void *cFrame);
extern bool H264Frame_isDisposable(void *cFrame);
extern void H265Frame_getH265Type(void *cFrame, char *buffer, size_t buffer_size);
extern uint32_t H265Frame_getH265FrameType(void *cFrame);
extern bool H265Frame_isDisposable(void *cFrame);
extern void TheoraFrame_getTheoraType(void *cFrame, char *buffer, size_t buffer_size);
extern uint64_t TheoraFrame_getGranulePos(void *cFrame);
extern void Yuv420Frame_getYuv420Type(void *cFrame, char *buffer, size_t buffer_size);
extern uint32_t Yuv420Frame_getYStride(void *cFrame);
extern uint32_t Yuv420Frame_getUStride(void *cFrame);
extern uint32_t Yuv420Frame_getVStride(void *cFrame);

extern void *newGoFrame(uint64_t pts, uint64_t dts, uint32_t timebase, uint32_t id,
    uint32_t width, uint32_t height,
    uint32_t sampleRate, uint32_t sampleNum, uint32_t channelNum,
    uint64_t dataPos,  uint32_t widthStride,
    uint64_t yDataPos, uint32_t yStride,
    uint64_t uDataPos, uint32_t uStride,
    uint64_t vDataPos, uint32_t vStride);
extern void deleteGoFrame(void *frame);

extern void *BgrFrame_fromBinary(void *data, size_t data_size,
  uint32_t id, uint64_t pts, uint32_t timebase,
  const char *bgr_type, uint32_t width, uint32_t height, uint32_t width_stride);
extern void *Flv1Frame_fromBinary(void *data, size_t data_size,
  uint32_t id, uint64_t pts, uint32_t timebase);
extern void *H264Frame_fromBinary(void *data, size_t data_size,
  uint32_t id, uint64_t pts, uint32_t timebase);
extern void *H265Frame_fromBinary(void *data, size_t data_size,
  uint32_t id, uint64_t pts, uint32_t timebase);
extern void *JpegFrame_fromBinary(void *data, size_t data_size,
  uint32_t id, uint64_t pts, uint32_t timebase);
extern void *PngFrame_fromBinary(void *data, size_t data_size,
  uint32_t id, uint64_t pts, uint32_t timebase);
extern void *TheoraFrame_fromBinary(void *data, size_t data_size,
	uint32_t id, uint64_t pts, uint32_t timebase,
	uint32_t width, uint32_t height);
extern void *Vp6Frame_fromBinary(void *data, size_t data_size,
	uint32_t id, uint64_t pts, uint32_t timebase,
	uint32_t width, uint32_t height, uint8_t adjustment);
extern void *Vp8Frame_fromBinary(void *data, size_t data_size,
	uint32_t id, uint64_t pts, uint32_t timebase,
	uint32_t width, uint32_t height);
extern void *Vp9Frame_fromBinary(void *data, size_t data_size,
	uint32_t id, uint64_t pts, uint32_t timebase,
	uint32_t width, uint32_t height);
extern void *Yuv420Frame_fromPlaneBinaries(uint32_t id, uint64_t pts, uint32_t timebase,
  const char *yuv_type, uint32_t width, uint32_t height,
  void *y_data, uint32_t y_stride,
  void *u_data, uint32_t u_stride,
	void *v_data, uint32_t v_stride);
*/
import "C"

import (
	"bytes"
	"unsafe"
)

// IsVideo 該当フレームタイプがvideoであるか判定する
func IsVideo(frameType frameType) bool {
	switch frameType {
	case FrameTypes.Bgr,
		FrameTypes.Flv1,
		FrameTypes.H264,
		FrameTypes.H265,
		FrameTypes.Jpeg,
		FrameTypes.Png,
		FrameTypes.Theora,
		FrameTypes.Vp6,
		FrameTypes.Vp8,
		FrameTypes.Vp9,
		FrameTypes.Wmv1,
		FrameTypes.Wmv2,
		FrameTypes.Yuv420:
		return true
	default:
		break
	}
	return false
}

// VideoTypes 映像タイプリスト
var VideoTypes = struct {
	Key   subType
	Inner subType
	Info  subType
}{
	subType{"key"},
	subType{"inner"},
	subType{"info"},
}

// VideoFrame 映像フレーム
type VideoFrame struct {
	Type      frameType
	Pts       uint64
	Dts       uint64
	Timebase  uint32
	ID        uint32
	VideoType subType
	Width     uint32
	Height    uint32
	cFrame    cttLibCFrame
}

// ToFrame 通常のFrameに戻す
func (video *VideoFrame) ToFrame() *Frame {
	frame := new(Frame)
	frame.cFrame = video.cFrame
	frame.ID = video.ID
	frame.Pts = video.Pts
	frame.Dts = video.Dts
	frame.Timebase = video.Timebase
	frame.Type = video.Type
	return frame
}

// newGoRefFrame goの配列データをcの世界に持っていくデータを構築する
func (video *VideoFrame) newGoRefFrame() unsafe.Pointer {
	return C.newGoFrame(C.uint64_t(video.Pts), C.uint64_t(video.Dts), C.uint32_t(video.Timebase), C.uint32_t(video.ID),
		C.uint32_t(video.Width), C.uint32_t(video.Height),
		0, 0, 0,
		0, 0,
		0, 0,
		0, 0,
		0, 0)
}

// refCFrame interface経由でcframe参照
func (video *VideoFrame) refCFrame() cttLibCFrame {
	return video.cFrame
}

// deleteGoRefFrame cの世界に持っていったデータを破棄する
func (video *VideoFrame) deleteGoRefFrame(ptr unsafe.Pointer) {
	C.deleteGoFrame(ptr)
}

// GetBinaryBuffer binaryデータを参照する
func (video *VideoFrame) GetBinaryBuffer(callback func(data []byte) bool) bool {
 	return getFrameBinaryBuffer(video.cFrame, callback)
}

// Video 映像処理
var Video = struct {
	Cast func(frame *Frame) *VideoFrame
}{
	Cast: func(frame *Frame) *VideoFrame {
		if !IsVideo(frame.Type) {
			return nil
		}
		videoFrame := new(VideoFrame)
		videoFrame.ID = frame.ID
		videoFrame.Pts = frame.Pts
		videoFrame.Dts = frame.Dts
		videoFrame.Timebase = frame.Timebase
		videoFrame.Type = frame.Type
		videoFrame.cFrame = frame.cFrame
		buf := make([]byte, 256)
		ptr := unsafe.Pointer(frame.cFrame)
		C.VideoFrame_getVideoType(ptr, (*C.char)(unsafe.Pointer(&buf[0])), C.size_t(256))
		name := string(buf[:bytes.IndexByte(buf, 0)])
		videoFrame.VideoType = subType{name}
		videoFrame.Width = uint32(C.VideoFrame_getWidth(ptr))
		videoFrame.Height = uint32(C.VideoFrame_getHeight(ptr))
		return videoFrame
	},
}

// BGR処理 ------------------------------------------------

// BgrTypes Bgrタイプリスト
var BgrTypes = struct {
	Bgr  subType
	Abgr subType
	Bgra subType
	Rgb  subType
	Argb subType
	Rgba subType
}{
	subType{"bgr"},
	subType{"abgr"},
	subType{"bgra"},
	subType{"rgb"},
	subType{"argb"},
	subType{"rgba"},
}

// BgrFrame Bgrフレーム
type BgrFrame struct {
	Type        frameType
	Pts         uint64
	Dts         uint64
	Timebase    uint32
	ID          uint32
	VideoType   subType
	Width       uint32
	Height      uint32
	BgrType     subType
	Data        uintptr
	WidthStride uint32
	cFrame      cttLibCFrame
}

// ToFrame 通常のFrameに戻す
func (bgr *BgrFrame) ToFrame() *Frame {
	frame := new(Frame)
	frame.cFrame = bgr.cFrame
	frame.ID = bgr.ID
	frame.Pts = bgr.Pts
	frame.Dts = bgr.Dts
	frame.Timebase = bgr.Timebase
	frame.Type = bgr.Type
	return frame
}

// newGoRefFrame goの配列データをcの世界に持っていくデータを構築する
func (bgr *BgrFrame) newGoRefFrame() unsafe.Pointer {
	return C.newGoFrame(C.uint64_t(bgr.Pts), C.uint64_t(bgr.Dts), C.uint32_t(bgr.Timebase), C.uint32_t(bgr.ID),
		C.uint32_t(bgr.Width), C.uint32_t(bgr.Height),
		0, 0, 0,
		C.uint64_t(bgr.Data), C.uint32_t(bgr.WidthStride),
		0, 0,
		0, 0,
		0, 0)
}

// refCFrame interface経由でcframe参照
func (bgr *BgrFrame) refCFrame() cttLibCFrame {
	return bgr.cFrame
}

// deleteGoRefFrame cの世界に持っていったデータを破棄する
func (bgr *BgrFrame) deleteGoRefFrame(ptr unsafe.Pointer) {
	C.deleteGoFrame(ptr)
}

// GetBinaryBuffer binaryデータを参照する
func (bgr *BgrFrame) GetBinaryBuffer(callback func(data []byte) bool) bool {
	return getFrameBinaryBuffer(bgr.cFrame, callback)
}

// Bgr Bgr処理
var Bgr = struct {
	Cast       func(frame *Frame) *BgrFrame // castして、frameを調整する
	FromBinary func(
		binary []byte,
		length uint64,
		id uint32,
		pts uint64,
		timebase uint32,
		bgrType subType,
		width uint32,
		height uint32,
		widthStride uint32) *Frame
}{
	Cast: func(frame *Frame) *BgrFrame {
		videoFrame := Video.Cast(frame)
		if videoFrame == nil {
			return nil
		}
		bgrFrame := new(BgrFrame)
		bgrFrame.ID = videoFrame.ID
		bgrFrame.Pts = videoFrame.Pts
		bgrFrame.Dts = videoFrame.Dts
		bgrFrame.Timebase = videoFrame.Timebase
		bgrFrame.Type = videoFrame.Type
		bgrFrame.VideoType = videoFrame.VideoType
		bgrFrame.cFrame = videoFrame.cFrame
		bgrFrame.Height = videoFrame.Height
		bgrFrame.Width = videoFrame.Width
		// あとは追記のデータだが・・・
		buf := make([]byte, 256)
		ptr := unsafe.Pointer(frame.cFrame)
		C.BgrFrame_getBgrType(ptr, (*C.char)(unsafe.Pointer(&buf[0])), C.size_t(256))
		name := string(buf[:bytes.IndexByte(buf, 0)])
		bgrFrame.BgrType = subType{name}
		bgrFrame.Data = 0
		bgrFrame.WidthStride = uint32(C.BgrFrame_getWidthStride(ptr))
		return bgrFrame
	},
	FromBinary: func(
		binary []byte,
		length uint64,
		id uint32,
		pts uint64,
		timebase uint32,
		bgrType subType,
		width uint32,
		height uint32,
		widthStride uint32) *Frame {
		// とりあえずなんとかして、ttLibC_Frameを作って、
		// そのデータをベースにデータを作ってやればいいか・・・
		// まだ未実装
		cBgrType := C.CString(bgrType.value)
		defer C.free(unsafe.Pointer(cBgrType))
		cFrame := C.BgrFrame_fromBinary(unsafe.Pointer(&binary[0]), C.size_t(length),
			C.uint32_t(id), C.uint64_t(pts), C.uint32_t(timebase),
			cBgrType, C.uint32_t(width), C.uint32_t(height), C.uint32_t(widthStride))
		frame := new(Frame)
		frame.init(cttLibCFrame(cFrame))
		frame.hasBody = true
		return frame
	},
}

// Flv1処理 ------------------------------------------------

// Flv1Types Flvタイプリスト
var Flv1Types = struct {
	Intra           subType
	Inner           subType
	DisposableInner subType
}{
	subType{"intra"},
	subType{"inner"},
	subType{"disposableInner"},
}

// Flv1Frame Flv1フレーム
type Flv1Frame struct {
	Type      frameType
	Pts       uint64
	Dts       uint64
	Timebase  uint32
	ID        uint32
	VideoType subType
	Width     uint32
	Height    uint32
	Flv1Type  subType
	cFrame    cttLibCFrame
}

// ToFrame 通常のFrameに戻す
func (flv1 *Flv1Frame) ToFrame() *Frame {
	frame := new(Frame)
	frame.cFrame = flv1.cFrame
	frame.ID = flv1.ID
	frame.Pts = flv1.Pts
	frame.Dts = flv1.Dts
	frame.Timebase = flv1.Timebase
	frame.Type = flv1.Type
	return frame
}

// newGoRefFrame goの配列データをcの世界に持っていくデータを構築する
func (flv1 *Flv1Frame) newGoRefFrame() unsafe.Pointer {
	return C.newGoFrame(C.uint64_t(flv1.Pts), C.uint64_t(flv1.Dts), C.uint32_t(flv1.Timebase), C.uint32_t(flv1.ID),
		C.uint32_t(flv1.Width), C.uint32_t(flv1.Height),
		0, 0, 0,
		0, 0,
		0, 0,
		0, 0,
		0, 0)
}

// refCFrame interface経由でcframe参照
func (flv1 *Flv1Frame) refCFrame() cttLibCFrame {
	return flv1.cFrame
}

// deleteGoRefFrame cの世界に持っていったデータを破棄する
func (flv1 *Flv1Frame) deleteGoRefFrame(ptr unsafe.Pointer) {
	C.deleteGoFrame(ptr)
}

// GetBinaryBuffer binaryデータを参照する
func (flv1 *Flv1Frame) GetBinaryBuffer(callback func(data []byte) bool) bool {
	return getFrameBinaryBuffer(flv1.cFrame, callback)
}

// Flv1 Flv1処理
var Flv1 = struct {
	Cast       func(frame *Frame) *Flv1Frame
	FromBinary func(binary []byte, length uint64, id uint32, pts uint64, timebase uint32) *Frame
}{
	Cast: func(frame *Frame) *Flv1Frame {
		videoFrame := Video.Cast(frame)
		if videoFrame == nil {
			return nil
		}
		flv1Frame := new(Flv1Frame)
		flv1Frame.ID = videoFrame.ID
		flv1Frame.Pts = videoFrame.Pts
		flv1Frame.Dts = videoFrame.Dts
		flv1Frame.Timebase = videoFrame.Timebase
		flv1Frame.Type = videoFrame.Type
		flv1Frame.VideoType = videoFrame.VideoType
		flv1Frame.cFrame = videoFrame.cFrame
		flv1Frame.Height = videoFrame.Height
		flv1Frame.Width = videoFrame.Width
		// あとは追加データ
		buf := make([]byte, 256)
		ptr := unsafe.Pointer(frame.cFrame)
		C.Flv1Frame_getFlv1Type(ptr, (*C.char)(unsafe.Pointer(&buf[0])), C.size_t(256))
		name := string(buf[:bytes.IndexByte(buf, 0)])
		flv1Frame.Flv1Type = subType{name}
		return flv1Frame
	},
	FromBinary: func(binary []byte, length uint64, id uint32, pts uint64, timebase uint32) *Frame {
		cFrame := C.Flv1Frame_fromBinary(unsafe.Pointer(&binary[0]), C.size_t(length),
			C.uint32_t(id), C.uint64_t(pts), C.uint32_t(timebase))
		frame := new(Frame)
		frame.init(cttLibCFrame(cFrame))
		frame.hasBody = true
		return frame
	},
}

// H264処理 ------------------------------------------------

// H264Types H264タイプリスト
var H264Types = struct {
	ConfigData subType
	SliceIDR   subType
	Slice      subType
	Unknown    subType
}{
	subType{"configData"},
	subType{"sliceIDR"},
	subType{"slice"},
	subType{"unknown"},
}

type h264FrameType int

// H264FrameTypes H264フレームタイプリスト
var H264FrameTypes = struct {
	P  h264FrameType
	B  h264FrameType
	I  h264FrameType
	SP h264FrameType
	SI h264FrameType
}{
	0,
	1,
	2,
	3,
	4,
}

// H264Frame H264フレーム
type H264Frame struct {
	Type          frameType
	Pts           uint64
	Dts           uint64
	Timebase      uint32
	ID            uint32
	VideoType     subType
	Width         uint32
	Height        uint32
	H264Type      subType
	H264FrameType h264FrameType
	IsDisposable  bool
	cFrame        cttLibCFrame
}

// ToFrame 通常のFrameに戻す
func (h264 *H264Frame) ToFrame() *Frame {
	frame := new(Frame)
	frame.cFrame = h264.cFrame
	frame.ID = h264.ID
	frame.Pts = h264.Pts
	frame.Dts = h264.Dts
	frame.Timebase = h264.Timebase
	frame.Type = h264.Type
	return frame
}

// newGoRefFrame goの配列データをcの世界に持っていくデータを構築する
func (h264 *H264Frame) newGoRefFrame() unsafe.Pointer {
	return C.newGoFrame(C.uint64_t(h264.Pts), C.uint64_t(h264.Dts), C.uint32_t(h264.Timebase), C.uint32_t(h264.ID),
		C.uint32_t(h264.Width), C.uint32_t(h264.Height),
		0, 0, 0,
		0, 0,
		0, 0,
		0, 0,
		0, 0)
}

// refCFrame interface経由でcframe参照
func (h264 *H264Frame) refCFrame() cttLibCFrame {
	return h264.cFrame
}

// deleteGoRefFrame cの世界に持っていったデータを破棄する
func (h264 *H264Frame) deleteGoRefFrame(ptr unsafe.Pointer) {
	C.deleteGoFrame(ptr)
}

// GetBinaryBuffer binaryデータを参照する
func (h264 *H264Frame) GetBinaryBuffer(callback func(data []byte) bool) bool {
	return getFrameBinaryBuffer(h264.cFrame, callback)
}

// H264 H264処理
var H264 = struct {
	Cast       func(frame *Frame) *H264Frame
	FromBinary func(binary []byte, length uint64, id uint32, pts uint64, timebase uint32) *Frame
}{
	Cast: func(frame *Frame) *H264Frame {
		videoFrame := Video.Cast(frame)
		if videoFrame == nil {
			return nil
		}
		h264Frame := new(H264Frame)
		h264Frame.ID = videoFrame.ID
		h264Frame.Pts = videoFrame.Pts
		h264Frame.Dts = videoFrame.Dts
		h264Frame.Timebase = videoFrame.Timebase
		h264Frame.Type = videoFrame.Type
		h264Frame.VideoType = videoFrame.VideoType
		h264Frame.cFrame = videoFrame.cFrame
		h264Frame.Height = videoFrame.Height
		h264Frame.Width = videoFrame.Width
		// あとは追加データ
		buf := make([]byte, 256)
		ptr := unsafe.Pointer(frame.cFrame)
		C.H264Frame_getH264Type(ptr, (*C.char)(unsafe.Pointer(&buf[0])), C.size_t(256))
		name := string(buf[:bytes.IndexByte(buf, 0)])
		h264Frame.H264Type = subType{name}
		h264Frame.H264FrameType = h264FrameType(C.H264Frame_getH264FrameType(ptr))
		h264Frame.IsDisposable = bool(C.H264Frame_isDisposable(ptr))
		return h264Frame
	},
	FromBinary: func(binary []byte, length uint64, id uint32, pts uint64, timebase uint32) *Frame {
		cFrame := C.H264Frame_fromBinary(unsafe.Pointer(&binary[0]), C.size_t(length),
			C.uint32_t(id), C.uint64_t(pts), C.uint32_t(timebase))
		frame := new(Frame)
		frame.init(cttLibCFrame(cFrame))
		frame.hasBody = true
		return frame
	},
}

// H265処理 ------------------------------------------------

// H265Types H265タイプリスト
var H265Types = struct {
	ConfigData subType
	SliceIDR   subType
	Slice      subType
	Unknown    subType
}{
	subType{"configData"},
	subType{"sliceIDR"},
	subType{"slice"},
	subType{"unknown"},
}

type h265FrameType int

// H265FrameTypes H265フレームタイプリスト
var H265FrameTypes = struct {
	B h265FrameType
	P h265FrameType
	I h265FrameType
}{
	0,
	1,
	2,
}

// H265Frame H265フレーム
type H265Frame struct {
	Type          frameType
	Pts           uint64
	Dts           uint64
	Timebase      uint32
	ID            uint32
	VideoType     subType
	Width         uint32
	Height        uint32
	H265Type      subType
	H265FrameType h265FrameType
	IsDisposable  bool
	cFrame        cttLibCFrame
}

// ToFrame 通常のFrameに戻す
func (h265 *H265Frame) ToFrame() *Frame {
	frame := new(Frame)
	frame.cFrame = h265.cFrame
	frame.ID = h265.ID
	frame.Pts = h265.Pts
	frame.Dts = h265.Dts
	frame.Timebase = h265.Timebase
	frame.Type = h265.Type
	return frame
}

// newGoRefFrame goの配列データをcの世界に持っていくデータを構築する
func (h265 *H265Frame) newGoRefFrame() unsafe.Pointer {
	return C.newGoFrame(C.uint64_t(h265.Pts), C.uint64_t(h265.Dts), C.uint32_t(h265.Timebase), C.uint32_t(h265.ID),
		C.uint32_t(h265.Width), C.uint32_t(h265.Height),
		0, 0, 0,
		0, 0,
		0, 0,
		0, 0,
		0, 0)
}

// refCFrame interface経由でcframe参照
func (h265 *H265Frame) refCFrame() cttLibCFrame {
	return h265.cFrame
}

// deleteGoRefFrame cの世界に持っていったデータを破棄する
func (h265 *H265Frame) deleteGoRefFrame(ptr unsafe.Pointer) {
	C.deleteGoFrame(ptr)
}

// GetBinaryBuffer binaryデータを参照する
func (h265 *H265Frame) GetBinaryBuffer(callback func(data []byte) bool) bool {
	return getFrameBinaryBuffer(h265.cFrame, callback)
}

// H265 H265処理
var H265 = struct {
	Cast       func(frame *Frame) *H265Frame
	FromBinary func(binary []byte, length uint64, id uint32, pts uint64, timebase uint32) *Frame
}{
	Cast: func(frame *Frame) *H265Frame {
		videoFrame := Video.Cast(frame)
		if videoFrame == nil {
			return nil
		}
		h265Frame := new(H265Frame)
		h265Frame.ID = videoFrame.ID
		h265Frame.Pts = videoFrame.Pts
		h265Frame.Dts = videoFrame.Dts
		h265Frame.Timebase = videoFrame.Timebase
		h265Frame.Type = videoFrame.Type
		h265Frame.VideoType = videoFrame.VideoType
		h265Frame.cFrame = videoFrame.cFrame
		h265Frame.Height = videoFrame.Height
		h265Frame.Width = videoFrame.Width
		// あとは追加データ
		buf := make([]byte, 256)
		ptr := unsafe.Pointer(frame.cFrame)
		C.H265Frame_getH265Type(ptr, (*C.char)(unsafe.Pointer(&buf[0])), C.size_t(256))
		name := string(buf[:bytes.IndexByte(buf, 0)])
		h265Frame.H265Type = subType{name}
		h265Frame.H265FrameType = h265FrameType(C.H265Frame_getH265FrameType(ptr))
		h265Frame.IsDisposable = bool(C.H265Frame_isDisposable(ptr))
		return h265Frame
	},
	FromBinary: func(binary []byte, length uint64, id uint32, pts uint64, timebase uint32) *Frame {
		cFrame := C.H265Frame_fromBinary(unsafe.Pointer(&binary[0]), C.size_t(length),
			C.uint32_t(id), C.uint64_t(pts), C.uint32_t(timebase))
		frame := new(Frame)
		frame.init(cttLibCFrame(cFrame))
		frame.hasBody = true
		return frame
	},
}

// Jpeg処理 ------------------------------------------------

// JpegFrame Jpegフレーム
type JpegFrame VideoFrame

// ToFrame 通常のFrameに戻す
func (jpeg *JpegFrame) ToFrame() *Frame {
	frame := new(Frame)
	frame.cFrame = jpeg.cFrame
	frame.ID = jpeg.ID
	frame.Pts = jpeg.Pts
	frame.Dts = jpeg.Dts
	frame.Timebase = jpeg.Timebase
	frame.Type = jpeg.Type
	return frame
}

// newGoRefFrame goの配列データをcの世界に持っていくデータを構築する
func (jpeg *JpegFrame) newGoRefFrame() unsafe.Pointer {
	return C.newGoFrame(C.uint64_t(jpeg.Pts), C.uint64_t(jpeg.Dts), C.uint32_t(jpeg.Timebase), C.uint32_t(jpeg.ID),
		C.uint32_t(jpeg.Width), C.uint32_t(jpeg.Height),
		0, 0, 0,
		0, 0,
		0, 0,
		0, 0,
		0, 0)
}

// refCFrame interface経由でcframe参照
func (jpeg *JpegFrame) refCFrame() cttLibCFrame {
	return jpeg.cFrame
}

// deleteGoRefFrame cの世界に持っていったデータを破棄する
func (jpeg *JpegFrame) deleteGoRefFrame(ptr unsafe.Pointer) {
	C.deleteGoFrame(ptr)
}

// GetBinaryBuffer binaryデータを参照する
func (jpeg *JpegFrame) GetBinaryBuffer(callback func(data []byte) bool) bool {
	return getFrameBinaryBuffer(jpeg.cFrame, callback)
}

// Jpeg Jpeg処理
var Jpeg = struct {
	Cast       func(frame *Frame) *JpegFrame
	FromBinary func(binary []byte, length uint64, id uint32, pts uint64, timebase uint32) *Frame
}{
	Cast: func(frame *Frame) *JpegFrame {
		videoFrame := Video.Cast(frame)
		if videoFrame == nil {
			return nil
		}
		jpegFrame := new(JpegFrame)
		jpegFrame.ID = videoFrame.ID
		jpegFrame.Pts = videoFrame.Pts
		jpegFrame.Dts = videoFrame.Dts
		jpegFrame.Timebase = videoFrame.Timebase
		jpegFrame.Type = videoFrame.Type
		jpegFrame.VideoType = videoFrame.VideoType
		jpegFrame.cFrame = videoFrame.cFrame
		jpegFrame.Height = videoFrame.Height
		jpegFrame.Width = videoFrame.Width
		return jpegFrame
	},
	FromBinary: func(binary []byte, length uint64, id uint32, pts uint64, timebase uint32) *Frame {
		cFrame := C.JpegFrame_fromBinary(unsafe.Pointer(&binary[0]), C.size_t(length),
			C.uint32_t(id), C.uint64_t(pts), C.uint32_t(timebase))
		frame := new(Frame)
		frame.init(cttLibCFrame(cFrame))
		frame.hasBody = true
		return frame
	},
}

// Png処理 ------------------------------------------------

// PngFrame Pngフレーム
type PngFrame VideoFrame

// ToFrame 通常のFrameに戻す
func (png *PngFrame) ToFrame() *Frame {
	frame := new(Frame)
	frame.cFrame = png.cFrame
	frame.ID = png.ID
	frame.Pts = png.Pts
	frame.Dts = png.Dts
	frame.Timebase = png.Timebase
	frame.Type = png.Type
	return frame
}

// newGoRefFrame goの配列データをcの世界に持っていくデータを構築する
func (png *PngFrame) newGoRefFrame() unsafe.Pointer {
	return C.newGoFrame(C.uint64_t(png.Pts), C.uint64_t(png.Dts), C.uint32_t(png.Timebase), C.uint32_t(png.ID),
		C.uint32_t(png.Width), C.uint32_t(png.Height),
		0, 0, 0,
		0, 0,
		0, 0,
		0, 0,
		0, 0)
}

// refCFrame interface経由でcframe参照
func (png *PngFrame) refCFrame() cttLibCFrame {
	return png.cFrame
}

// deleteGoRefFrame cの世界に持っていったデータを破棄する
func (png *PngFrame) deleteGoRefFrame(ptr unsafe.Pointer) {
	C.deleteGoFrame(ptr)
}

// GetBinaryBuffer binaryデータを参照する
func (png *PngFrame) GetBinaryBuffer(callback func(data []byte) bool) bool {
	return getFrameBinaryBuffer(png.cFrame, callback)
}

// Png Png処理
var Png = struct {
	Cast       func(frame *Frame) *PngFrame
	FromBinary func(binary []byte, length uint64, id uint32, pts uint64, timebase uint32) *Frame
}{
	Cast: func(frame *Frame) *PngFrame {
		videoFrame := Video.Cast(frame)
		if videoFrame == nil {
			return nil
		}
		pngFrame := new(PngFrame)
		pngFrame.ID = videoFrame.ID
		pngFrame.Pts = videoFrame.Pts
		pngFrame.Dts = videoFrame.Dts
		pngFrame.Timebase = videoFrame.Timebase
		pngFrame.Type = videoFrame.Type
		pngFrame.VideoType = videoFrame.VideoType
		pngFrame.cFrame = videoFrame.cFrame
		pngFrame.Height = videoFrame.Height
		pngFrame.Width = videoFrame.Width
		return pngFrame
	},
	FromBinary: func(binary []byte, length uint64, id uint32, pts uint64, timebase uint32) *Frame {
		cFrame := C.PngFrame_fromBinary(unsafe.Pointer(&binary[0]), C.size_t(length),
			C.uint32_t(id), C.uint64_t(pts), C.uint32_t(timebase))
		frame := new(Frame)
		frame.init(cttLibCFrame(cFrame))
		frame.hasBody = true
		return frame
	},
}

// Theora処理 ------------------------------------------------

// TheoraTypes Theoraタイプリスト
var TheoraTypes = struct {
	IdentificationHeader subType
	CommenntHeader       subType
	SetupHeader          subType
	Intra                subType
	Inner                subType
}{
	subType{"identificationHeader"},
	subType{"commentHeader"},
	subType{"setupHeader"},
	subType{"intra"},
	subType{"inner"},
}

// TheoraFrame Theoraフレーム
type TheoraFrame struct {
	Type       frameType
	Pts        uint64
	Dts        uint64
	Timebase   uint32
	ID         uint32
	VideoType  subType
	Width      uint32
	Height     uint32
	TheoraType subType
	GranulePos uint64
	cFrame     cttLibCFrame
}

// ToFrame 通常のFrameに戻す
func (theora *TheoraFrame) ToFrame() *Frame {
	frame := new(Frame)
	frame.cFrame = theora.cFrame
	frame.ID = theora.ID
	frame.Pts = theora.Pts
	frame.Dts = theora.Dts
	frame.Timebase = theora.Timebase
	frame.Type = theora.Type
	return frame
}

// newGoRefFrame goの配列データをcの世界に持っていくデータを構築する
func (theora *TheoraFrame) newGoRefFrame() unsafe.Pointer {
	return C.newGoFrame(C.uint64_t(theora.Pts), C.uint64_t(theora.Dts), C.uint32_t(theora.Timebase), C.uint32_t(theora.ID),
		C.uint32_t(theora.Width), C.uint32_t(theora.Height),
		0, 0, 0,
		0, 0,
		0, 0,
		0, 0,
		0, 0)
}

// refCFrame interface経由でcframe参照
func (theora *TheoraFrame) refCFrame() cttLibCFrame {
	return theora.cFrame
}

// deleteGoRefFrame cの世界に持っていったデータを破棄する
func (theora *TheoraFrame) deleteGoRefFrame(ptr unsafe.Pointer) {
	C.deleteGoFrame(ptr)
}

// GetBinaryBuffer binaryデータを参照する
func (theora *TheoraFrame) GetBinaryBuffer(callback func(data []byte) bool) bool {
	return getFrameBinaryBuffer(theora.cFrame, callback)
}

// Theora Theora処理
var Theora = struct {
	Cast       func(frame *Frame) *TheoraFrame
	FromBinary func(binary []byte, length uint64, id uint32, pts uint64, timebase uint32, width uint32, height uint32) *Frame
}{
	Cast: func(frame *Frame) *TheoraFrame {
		videoFrame := Video.Cast(frame)
		if videoFrame == nil {
			return nil
		}
		theoraFrame := new(TheoraFrame)
		theoraFrame.ID = videoFrame.ID
		theoraFrame.Pts = videoFrame.Pts
		theoraFrame.Dts = videoFrame.Dts
		theoraFrame.Timebase = videoFrame.Timebase
		theoraFrame.Type = videoFrame.Type
		theoraFrame.VideoType = videoFrame.VideoType
		theoraFrame.cFrame = videoFrame.cFrame
		theoraFrame.Height = videoFrame.Height
		theoraFrame.Width = videoFrame.Width
		// あとは追加データ
		buf := make([]byte, 256)
		ptr := unsafe.Pointer(frame.cFrame)
		C.TheoraFrame_getTheoraType(ptr, (*C.char)(unsafe.Pointer(&buf[0])), C.size_t(256))
		name := string(buf[:bytes.IndexByte(buf, 0)])
		theoraFrame.TheoraType = subType{name}
		theoraFrame.GranulePos = uint64(C.TheoraFrame_getGranulePos(ptr))
		return theoraFrame
	},
	FromBinary: func(binary []byte, length uint64, id uint32, pts uint64, timebase uint32, width uint32, height uint32) *Frame {
		cFrame := C.TheoraFrame_fromBinary(unsafe.Pointer(&binary[0]), C.size_t(length),
			C.uint32_t(id), C.uint64_t(pts), C.uint32_t(timebase),
			C.uint32_t(width), C.uint32_t(height))
		frame := new(Frame)
		frame.init(cttLibCFrame(cFrame))
		frame.hasBody = true
		return frame
	},
}

// Vp6処理 ------------------------------------------------

// Vp6Frame Vp6フレーム
type Vp6Frame VideoFrame

// ToFrame 通常のFrameに戻す
func (vp6 *Vp6Frame) ToFrame() *Frame {
	frame := new(Frame)
	frame.cFrame = vp6.cFrame
	frame.ID = vp6.ID
	frame.Pts = vp6.Pts
	frame.Dts = vp6.Dts
	frame.Timebase = vp6.Timebase
	frame.Type = vp6.Type
	return frame
}

// newGoRefFrame goの配列データをcの世界に持っていくデータを構築する
func (vp6 *Vp6Frame) newGoRefFrame() unsafe.Pointer {
	return C.newGoFrame(C.uint64_t(vp6.Pts), C.uint64_t(vp6.Dts), C.uint32_t(vp6.Timebase), C.uint32_t(vp6.ID),
		C.uint32_t(vp6.Width), C.uint32_t(vp6.Height),
		0, 0, 0,
		0, 0,
		0, 0,
		0, 0,
		0, 0)
}

// refCFrame interface経由でcframe参照
func (vp6 *Vp6Frame) refCFrame() cttLibCFrame {
	return vp6.cFrame
}

// deleteGoRefFrame cの世界に持っていったデータを破棄する
func (vp6 *Vp6Frame) deleteGoRefFrame(ptr unsafe.Pointer) {
	C.deleteGoFrame(ptr)
}

// GetBinaryBuffer binaryデータを参照する
func (vp6 *Vp6Frame) GetBinaryBuffer(callback func(data []byte) bool) bool {
	return getFrameBinaryBuffer(vp6.cFrame, callback)
}

// Vp6 Vp6処理
var Vp6 = struct {
	Cast       func(frame *Frame) *Vp6Frame
	FromBinary func(binary []byte, length uint64, id uint32, pts uint64, timebase uint32, width uint32, height uint32, adjustment uint8) *Frame
}{
	Cast: func(frame *Frame) *Vp6Frame {
		videoFrame := Video.Cast(frame)
		if videoFrame == nil {
			return nil
		}
		vp6Frame := new(Vp6Frame)
		vp6Frame.ID = videoFrame.ID
		vp6Frame.Pts = videoFrame.Pts
		vp6Frame.Dts = videoFrame.Dts
		vp6Frame.Timebase = videoFrame.Timebase
		vp6Frame.Type = videoFrame.Type
		vp6Frame.VideoType = videoFrame.VideoType
		vp6Frame.cFrame = videoFrame.cFrame
		vp6Frame.Height = videoFrame.Height
		vp6Frame.Width = videoFrame.Width
		return vp6Frame
	},
	FromBinary: func(binary []byte, length uint64, id uint32, pts uint64, timebase uint32, width uint32, height uint32, adjustment uint8) *Frame {
		cFrame := C.Vp6Frame_fromBinary(unsafe.Pointer(&binary[0]), C.size_t(length),
			C.uint32_t(id), C.uint64_t(pts), C.uint32_t(timebase),
			C.uint32_t(width), C.uint32_t(height), C.uint8_t(adjustment))
		frame := new(Frame)
		frame.init(cttLibCFrame(cFrame))
		frame.hasBody = true
		return frame
	},
}

// Vp8処理 ------------------------------------------------

// Vp8Frame Vp8フレーム
type Vp8Frame VideoFrame

// ToFrame 通常のFrameに戻す
func (vp8 *Vp8Frame) ToFrame() *Frame {
	frame := new(Frame)
	frame.cFrame = vp8.cFrame
	frame.ID = vp8.ID
	frame.Pts = vp8.Pts
	frame.Dts = vp8.Dts
	frame.Timebase = vp8.Timebase
	frame.Type = vp8.Type
	return frame
}

// newGoRefFrame goの配列データをcの世界に持っていくデータを構築する
func (vp8 *Vp8Frame) newGoRefFrame() unsafe.Pointer {
	return C.newGoFrame(C.uint64_t(vp8.Pts), C.uint64_t(vp8.Dts), C.uint32_t(vp8.Timebase), C.uint32_t(vp8.ID),
		C.uint32_t(vp8.Width), C.uint32_t(vp8.Height),
		0, 0, 0,
		0, 0,
		0, 0,
		0, 0,
		0, 0)
}

// refCFrame interface経由でcframe参照
func (vp8 *Vp8Frame) refCFrame() cttLibCFrame {
	return vp8.cFrame
}

// deleteGoRefFrame cの世界に持っていったデータを破棄する
func (vp8 *Vp8Frame) deleteGoRefFrame(ptr unsafe.Pointer) {
	C.deleteGoFrame(ptr)
}

// GetBinaryBuffer binaryデータを参照する
func (vp8 *Vp8Frame) GetBinaryBuffer(callback func(data []byte) bool) bool {
	return getFrameBinaryBuffer(vp8.cFrame, callback)
}

// Vp8 Vp8処理
var Vp8 = struct {
	Cast       func(frame *Frame) *Vp8Frame
	FromBinary func(binary []byte, length uint64, id uint32, pts uint64, timebase uint32, width uint32, height uint32) *Frame
}{
	Cast: func(frame *Frame) *Vp8Frame {
		videoFrame := Video.Cast(frame)
		if videoFrame == nil {
			return nil
		}
		vp8Frame := new(Vp8Frame)
		vp8Frame.ID = videoFrame.ID
		vp8Frame.Pts = videoFrame.Pts
		vp8Frame.Dts = videoFrame.Dts
		vp8Frame.Timebase = videoFrame.Timebase
		vp8Frame.Type = videoFrame.Type
		vp8Frame.VideoType = videoFrame.VideoType
		vp8Frame.cFrame = videoFrame.cFrame
		vp8Frame.Height = videoFrame.Height
		vp8Frame.Width = videoFrame.Width
		return vp8Frame
	},
	FromBinary: func(binary []byte, length uint64, id uint32, pts uint64, timebase uint32, width uint32, height uint32) *Frame {
		cFrame := C.Vp8Frame_fromBinary(unsafe.Pointer(&binary[0]), C.size_t(length),
			C.uint32_t(id), C.uint64_t(pts), C.uint32_t(timebase),
			C.uint32_t(width), C.uint32_t(height))
		frame := new(Frame)
		frame.init(cttLibCFrame(cFrame))
		frame.hasBody = true
		return frame
	},
}

// Vp9処理 ------------------------------------------------

// Vp9Frame Vp9フレーム
type Vp9Frame VideoFrame

// ToFrame 通常のFrameに戻す
func (vp9 *Vp9Frame) ToFrame() *Frame {
	frame := new(Frame)
	frame.cFrame = vp9.cFrame
	frame.ID = vp9.ID
	frame.Pts = vp9.Pts
	frame.Dts = vp9.Dts
	frame.Timebase = vp9.Timebase
	frame.Type = vp9.Type
	return frame
}

// newGoRefFrame goの配列データをcの世界に持っていくデータを構築する
func (vp9 *Vp9Frame) newGoRefFrame() unsafe.Pointer {
	return C.newGoFrame(C.uint64_t(vp9.Pts), C.uint64_t(vp9.Dts), C.uint32_t(vp9.Timebase), C.uint32_t(vp9.ID),
		C.uint32_t(vp9.Width), C.uint32_t(vp9.Height),
		0, 0, 0,
		0, 0,
		0, 0,
		0, 0,
		0, 0)
}

// refCFrame interface経由でcframe参照
func (vp9 *Vp9Frame) refCFrame() cttLibCFrame {
	return vp9.cFrame
}

// deleteGoRefFrame cの世界に持っていったデータを破棄する
func (vp9 *Vp9Frame) deleteGoRefFrame(ptr unsafe.Pointer) {
	C.deleteGoFrame(ptr)
}

// GetBinaryBuffer binaryデータを参照する
func (vp9 *Vp9Frame) GetBinaryBuffer(callback func(data []byte) bool) bool {
	return getFrameBinaryBuffer(vp9.cFrame, callback)
}

// Vp9 Vp9処理
var Vp9 = struct {
	Cast       func(frame *Frame) *Vp9Frame
	FromBinary func(binary []byte, length uint64, id uint32, pts uint64, timebase uint32, width uint32, height uint32) *Frame
}{
	Cast: func(frame *Frame) *Vp9Frame {
		videoFrame := Video.Cast(frame)
		if videoFrame == nil {
			return nil
		}
		vp9Frame := new(Vp9Frame)
		vp9Frame.ID = videoFrame.ID
		vp9Frame.Pts = videoFrame.Pts
		vp9Frame.Dts = videoFrame.Dts
		vp9Frame.Timebase = videoFrame.Timebase
		vp9Frame.Type = videoFrame.Type
		vp9Frame.VideoType = videoFrame.VideoType
		vp9Frame.cFrame = videoFrame.cFrame
		vp9Frame.Height = videoFrame.Height
		vp9Frame.Width = videoFrame.Width
		return vp9Frame
	},
	FromBinary: func(binary []byte, length uint64, id uint32, pts uint64, timebase uint32, width uint32, height uint32) *Frame {
		cFrame := C.Vp9Frame_fromBinary(unsafe.Pointer(&binary[0]), C.size_t(length),
			C.uint32_t(id), C.uint64_t(pts), C.uint32_t(timebase),
			C.uint32_t(width), C.uint32_t(height))
		frame := new(Frame)
		frame.init(cttLibCFrame(cFrame))
		frame.hasBody = true
		return frame
	},
}

// Wmv1処理 ------------------------------------------------

// Wmv1Frame Wmv1フレーム
type Wmv1Frame VideoFrame

// ToFrame 通常のFrameに戻す
func (wmv1 *Wmv1Frame) ToFrame() *Frame {
	frame := new(Frame)
	frame.cFrame = wmv1.cFrame
	frame.ID = wmv1.ID
	frame.Pts = wmv1.Pts
	frame.Dts = wmv1.Dts
	frame.Timebase = wmv1.Timebase
	frame.Type = wmv1.Type
	return frame
}

// newGoRefFrame goの配列データをcの世界に持っていくデータを構築する
func (wmv1 *Wmv1Frame) newGoRefFrame() unsafe.Pointer {
	return C.newGoFrame(C.uint64_t(wmv1.Pts), C.uint64_t(wmv1.Dts), C.uint32_t(wmv1.Timebase), C.uint32_t(wmv1.ID),
		C.uint32_t(wmv1.Width), C.uint32_t(wmv1.Height),
		0, 0, 0,
		0, 0,
		0, 0,
		0, 0,
		0, 0)
}

// refCFrame interface経由でcframe参照
func (wmv1 *Wmv1Frame) refCFrame() cttLibCFrame {
	return wmv1.cFrame
}

// deleteGoRefFrame cの世界に持っていったデータを破棄する
func (wmv1 *Wmv1Frame) deleteGoRefFrame(ptr unsafe.Pointer) {
	C.deleteGoFrame(ptr)
}

// GetBinaryBuffer binaryデータを参照する
func (wmv1 *Wmv1Frame) GetBinaryBuffer(callback func(data []byte) bool) bool {
	return getFrameBinaryBuffer(wmv1.cFrame, callback)
}

// Wmv1 Wmv1処理
var Wmv1 = struct {
	Cast       func(frame *Frame) *Wmv1Frame
	FromBinary func(binary []byte, id uint32, pts uint64, timebase uint32) *Frame
}{
	Cast: func(frame *Frame) *Wmv1Frame {
		videoFrame := Video.Cast(frame)
		if videoFrame == nil {
			return nil
		}
		wmv1Frame := new(Wmv1Frame)
		wmv1Frame.ID = videoFrame.ID
		wmv1Frame.Pts = videoFrame.Pts
		wmv1Frame.Dts = videoFrame.Dts
		wmv1Frame.Timebase = videoFrame.Timebase
		wmv1Frame.Type = videoFrame.Type
		wmv1Frame.VideoType = videoFrame.VideoType
		wmv1Frame.cFrame = videoFrame.cFrame
		wmv1Frame.Height = videoFrame.Height
		wmv1Frame.Width = videoFrame.Width
		return wmv1Frame
	},
	FromBinary: func(binary []byte, id uint32, pts uint64, timebase uint32) *Frame {
		panic("not implemented for wmv1")
	},
}

// Wmv2処理 ------------------------------------------------

// Wmv2Frame Wmv2フレーム
type Wmv2Frame VideoFrame

// ToFrame 通常のFrameに戻す
func (wmv2 *Wmv2Frame) ToFrame() *Frame {
	frame := new(Frame)
	frame.cFrame = wmv2.cFrame
	frame.ID = wmv2.ID
	frame.Pts = wmv2.Pts
	frame.Dts = wmv2.Dts
	frame.Timebase = wmv2.Timebase
	frame.Type = wmv2.Type
	return frame
}

// newGoRefFrame goの配列データをcの世界に持っていくデータを構築する
func (wmv2 *Wmv2Frame) newGoRefFrame() unsafe.Pointer {
	return C.newGoFrame(C.uint64_t(wmv2.Pts), C.uint64_t(wmv2.Dts), C.uint32_t(wmv2.Timebase), C.uint32_t(wmv2.ID),
		C.uint32_t(wmv2.Width), C.uint32_t(wmv2.Height),
		0, 0, 0,
		0, 0,
		0, 0,
		0, 0,
		0, 0)
}

// refCFrame interface経由でcframe参照
func (wmv2 *Wmv2Frame) refCFrame() cttLibCFrame {
	return wmv2.cFrame
}

// deleteGoRefFrame cの世界に持っていったデータを破棄する
func (wmv2 *Wmv2Frame) deleteGoRefFrame(ptr unsafe.Pointer) {
	C.deleteGoFrame(ptr)
}

// GetBinaryBuffer binaryデータを参照する
func (wmv2 *Wmv2Frame) GetBinaryBuffer(callback func(data []byte) bool) bool {
	return getFrameBinaryBuffer(wmv2.cFrame, callback)
}

// Wmv2 Wmv2処理
var Wmv2 = struct {
	Cast       func(frame *Frame) *Wmv2Frame
	FromBinary func(binary []byte, id uint32, pts uint64, timebase uint32) *Frame
}{
	Cast: func(frame *Frame) *Wmv2Frame {
		videoFrame := Video.Cast(frame)
		if videoFrame == nil {
			return nil
		}
		wmv2Frame := new(Wmv2Frame)
		wmv2Frame.ID = videoFrame.ID
		wmv2Frame.Pts = videoFrame.Pts
		wmv2Frame.Dts = videoFrame.Dts
		wmv2Frame.Timebase = videoFrame.Timebase
		wmv2Frame.Type = videoFrame.Type
		wmv2Frame.VideoType = videoFrame.VideoType
		wmv2Frame.cFrame = videoFrame.cFrame
		wmv2Frame.Height = videoFrame.Height
		wmv2Frame.Width = videoFrame.Width
		return wmv2Frame
	},
	FromBinary: func(binary []byte, id uint32, pts uint64, timebase uint32) *Frame {
		panic("not implemented for wmv2")
	},
}

// Yuv420処理 ------------------------------------------------

// Yuv420Types Yuv420タイプリスト
var Yuv420Types = struct {
	Yuv420           subType
	Yuv420SemiPlanar subType
	Yvu420           subType
	Yvu420SemiPlanar subType
}{
	subType{"yuv420"},
	subType{"yuv420SemiPlanar"},
	subType{"yvu420"},
	subType{"yvu420SemiPlanar"},
}

// Yuv420Frame Yuv420フレーム
type Yuv420Frame struct {
	Type       frameType
	Pts        uint64
	Dts        uint64
	Timebase   uint32
	ID         uint32
	VideoType  subType
	Width      uint32
	Height     uint32
	Yuv420Type subType
	YData      uintptr
	YStride    uint32
	UData      uintptr
	UStride    uint32
	VData      uintptr
	VStride    uint32
	cFrame     cttLibCFrame
}

// ToFrame 通常のFrameに戻す
func (yuv420 *Yuv420Frame) ToFrame() *Frame {
	frame := new(Frame)
	frame.cFrame = yuv420.cFrame
	frame.ID = yuv420.ID
	frame.Pts = yuv420.Pts
	frame.Dts = yuv420.Dts
	frame.Timebase = yuv420.Timebase
	frame.Type = yuv420.Type
	return frame
}

// newGoRefFrame goの配列データをcの世界に持っていくデータを構築する
func (yuv420 *Yuv420Frame) newGoRefFrame() unsafe.Pointer {
	return C.newGoFrame(C.uint64_t(yuv420.Pts), C.uint64_t(yuv420.Dts), C.uint32_t(yuv420.Timebase), C.uint32_t(yuv420.ID),
		C.uint32_t(yuv420.Width), C.uint32_t(yuv420.Height),
		0, 0, 0,
		0, 0,
		C.uint64_t(yuv420.YData), C.uint32_t(yuv420.YStride),
		C.uint64_t(yuv420.UData), C.uint32_t(yuv420.UStride),
		C.uint64_t(yuv420.VData), C.uint32_t(yuv420.VStride))
}

// refCFrame interface経由でcframe参照
func (yuv420 *Yuv420Frame) refCFrame() cttLibCFrame {
	return yuv420.cFrame
}

// deleteGoRefFrame cの世界に持っていったデータを破棄する
func (yuv420 *Yuv420Frame) deleteGoRefFrame(ptr unsafe.Pointer) {
	C.deleteGoFrame(ptr)
}

// GetBinaryBuffer binaryデータを参照する
func (yuv420 *Yuv420Frame) GetBinaryBuffer(callback func(data []byte) bool) bool {
	return getFrameBinaryBuffer(yuv420.cFrame, callback)
}

// Yuv420 Yuv420処理
var Yuv420 = struct {
	Cast       func(frame *Frame) *Yuv420Frame
	FromBinary func(
		binary []byte,
		length uint64,
		id uint32,
		pts uint64,
		timebase uint32,
		yuv420Type subType,
		width uint32,
		height uint32) *Frame
	FromPlaneBinaries func(
		id uint32,
		pts uint64,
		timebase uint32,
		yuv420Type subType,
		width uint32,
		height uint32,
		yBinary []byte,
		yStride uint32,
		uBinary []byte,
		uStride uint32,
		vBinary []byte,
		vStride uint32) *Frame
}{
	Cast: func(frame *Frame) *Yuv420Frame {
		videoFrame := Video.Cast(frame)
		if videoFrame == nil {
			return nil
		}
		yuv420Frame := new(Yuv420Frame)
		yuv420Frame.ID = videoFrame.ID
		yuv420Frame.Pts = videoFrame.Pts
		yuv420Frame.Dts = videoFrame.Dts
		yuv420Frame.Timebase = videoFrame.Timebase
		yuv420Frame.Type = videoFrame.Type
		yuv420Frame.VideoType = videoFrame.VideoType
		yuv420Frame.cFrame = videoFrame.cFrame
		yuv420Frame.Height = videoFrame.Height
		yuv420Frame.Width = videoFrame.Width
		// あとは追記のデータ
		buf := make([]byte, 256)
		ptr := unsafe.Pointer(frame.cFrame)
		C.Yuv420Frame_getYuv420Type(ptr, (*C.char)(unsafe.Pointer(&buf[0])), C.size_t(256))
		name := string(buf[:bytes.IndexByte(buf, 0)])
		yuv420Frame.Yuv420Type = subType{name}
		yuv420Frame.YData = 0
		yuv420Frame.YStride = uint32(C.Yuv420Frame_getYStride(ptr))
		yuv420Frame.UData = 0
		yuv420Frame.UStride = uint32(C.Yuv420Frame_getUStride(ptr))
		yuv420Frame.VData = 0
		yuv420Frame.VStride = uint32(C.Yuv420Frame_getVStride(ptr))
		return yuv420Frame
	},
	FromBinary: func(
		binary []byte,
		length uint64,
		id uint32,
		pts uint64,
		timebase uint32,
		yuv420Type subType,
		width uint32,
		height uint32) *Frame {
		//strideを自力で計算して・・・とすれば行けるか・・・
		halfWidth := (width + 1) >> 1
		halfHeight := (height + 1) >> 1
		yStride := width
		uStride := halfWidth
		vStride := halfWidth
		yBinary := binary
		var uBinary []byte
		var vBinary []byte
		switch yuv420Type {
		case Yuv420Types.Yuv420:
			uBinary = binary[width*height:]
			vBinary = uBinary[halfWidth*halfHeight:]
			return yuv420_fromPlaneBinaries(
				id,
				pts,
				timebase,
				yuv420Type,
				width,
				height,
				yBinary,
				yStride,
				uBinary,
				uStride,
				vBinary,
				vStride)
		case Yuv420Types.Yuv420SemiPlanar:
			uStride := (halfWidth << 1)
			vStride := (halfWidth << 1)
			uBinary = binary[width*height:]
			vBinary = uBinary[1:]
			return yuv420_fromPlaneBinaries(
				id,
				pts,
				timebase,
				yuv420Type,
				width,
				height,
				yBinary,
				yStride,
				uBinary,
				uStride,
				vBinary,
				vStride)
		case Yuv420Types.Yvu420:
			vBinary = binary[width*height:]
			uBinary = vBinary[halfWidth*halfHeight:]
			return yuv420_fromPlaneBinaries(
				id,
				pts,
				timebase,
				yuv420Type,
				width,
				height,
				yBinary,
				yStride,
				uBinary,
				uStride,
				vBinary,
				vStride)
		case Yuv420Types.Yvu420SemiPlanar:
			uStride := (halfWidth << 1)
			vStride := (halfWidth << 1)
			vBinary = binary[width*height:]
			uBinary = vBinary[1:]
			return yuv420_fromPlaneBinaries(
				id,
				pts,
				timebase,
				yuv420Type,
				width,
				height,
				yBinary,
				yStride,
				uBinary,
				uStride,
				vBinary,
				vStride)
		default:
			return nil
		}
		return nil
	},
	FromPlaneBinaries: yuv420_fromPlaneBinaries,
}

func yuv420_fromPlaneBinaries(
	id uint32,
	pts uint64,
	timebase uint32,
	yuv420Type subType,
	width uint32,
	height uint32,
	yBinary []byte,
	yStride uint32,
	uBinary []byte,
	uStride uint32,
	vBinary []byte,
	vStride uint32) *Frame {

	cYuv420Type := C.CString(yuv420Type.value)
	defer C.free(unsafe.Pointer(cYuv420Type))
	cFrame := C.Yuv420Frame_fromPlaneBinaries(
		C.uint32_t(id), C.uint64_t(pts), C.uint32_t(timebase),
		cYuv420Type, C.uint32_t(width), C.uint32_t(height),
		unsafe.Pointer(&yBinary[0]), C.uint32_t(yStride),
		unsafe.Pointer(&uBinary[0]), C.uint32_t(uStride),
		unsafe.Pointer(&vBinary[0]), C.uint32_t(vStride))
	frame := new(Frame)
	frame.init(cttLibCFrame(cFrame))
	frame.hasBody = true
	return frame
}
