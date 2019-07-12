package ttLibGo

/*
#include <stdint.h>
#include <stdlib.h>
extern uint32_t AudioFrame_getSampleRate(void *frame);
extern uint32_t AudioFrame_getSampleNum(void *frame);
extern uint32_t AudioFrame_getChannelNum(void *frame);
extern void AacFrame_getAacType(void *cFrame, char *buffer, size_t buffer_size);
extern void Mp3Frame_getMp3Type(void *cFrame, char *buffer, size_t buffer_size);
extern void OpusFrame_getOpusType(void *cFrame, char *buffer, size_t buffer_size);
extern void PcmF32Frame_getPcmF32Type(void *cFrame, char *buffer, size_t buffer_size);
extern uint32_t PcmF32Frame_getLStride(void *cFrame);
extern uint32_t PcmF32Frame_getLStep(void *cFrame);
extern uint32_t PcmF32Frame_getRStride(void *cFrame);
extern uint32_t PcmF32Frame_getRStep(void *cFrame);
extern void PcmS16Frame_getPcmS16Type(void *cFrame, char *buffer, size_t buffer_size);
extern uint32_t PcmS16Frame_getLStride(void *cFrame);
extern uint32_t PcmS16Frame_getLStep(void *cFrame);
extern uint32_t PcmS16Frame_getRStride(void *cFrame);
extern uint32_t PcmS16Frame_getRStep(void *cFrame);
extern void SpeexFrame_getSpeexType(void *cFrame, char *buffer, size_t buffer_size);
extern void VorbisFrame_getVorbisType(void *cFrame, char *buffer, size_t buffer_size);

extern void *newGoFrame(uint64_t pts, uint64_t dts, uint32_t timebase, uint32_t id,
    uint32_t width, uint32_t height,
    uint32_t sampleRate, uint32_t sampleNum, uint32_t channelNum,
    uint64_t dataPos,  uint32_t widthStride,
    uint64_t yDataPos, uint32_t yStride,
    uint64_t uDataPos, uint32_t uStride,
    uint64_t vDataPos, uint32_t vStride);
extern void deleteGoFrame(void *frame);

extern void *AacFrame_fromBinary(void *data, size_t data_size,
  uint32_t id, uint64_t pts, uint32_t timebase,
  void *dsiFrame);
extern void *AdpcmImaWavFrame_fromBinary(void *data, size_t data_size,
  uint32_t id, uint64_t pts, uint32_t timebase,
  uint32_t sample_rate, uint32_t sample_num, uint32_t channel_num);
extern void *Mp3Frame_fromBinary(void *data, size_t data_size,
  uint32_t id, uint64_t pts, uint32_t timebase);
extern void *NellymoserFrame_fromBinary(void *data, size_t data_size,
  uint32_t id, uint64_t pts, uint32_t timebase,
  uint32_t sample_rate, uint32_t sample_num, uint32_t channel_num);
extern void *OpusFrame_fromBinary(void *data, size_t data_size,
  uint32_t id, uint64_t pts, uint32_t timebase);
extern void *PcmAlawsFrame_fromBinary(void *data, size_t data_size,
  uint32_t id, uint64_t pts, uint32_t timebase,
  uint32_t sample_rate, uint32_t sample_num, uint32_t channel_num);
extern void *PcmF32Frame_fromBinary(void *data, size_t data_size,
  uint32_t id, uint64_t pts, uint32_t timebase,
  const char *pcmF32_type, uint32_t sample_rate, uint32_t sample_num, uint32_t channel_num,
  uint32_t l_index, uint32_t r_index);
extern void *PcmMulawsFrame_fromBinary(void *data, size_t data_size,
  uint32_t id, uint64_t pts, uint32_t timebase,
  uint32_t sample_rate, uint32_t sample_num, uint32_t channel_num);
extern void *PcmS16Frame_fromBinary(void *data, size_t data_size,
  uint32_t id, uint64_t pts, uint32_t timebase,
  const char *pcmS16_type, uint32_t sample_rate, uint32_t sample_num, uint32_t channel_num,
  uint32_t l_index, uint32_t r_index);
extern void *SpeexFrame_fromBinary(void *data, size_t data_size,
  uint32_t id, uint64_t pts, uint32_t timebase,
  uint32_t sample_rate, uint32_t sample_num, uint32_t channel_num);
extern void *VorbisFrame_fromBinary(void *data, size_t data_size,
  uint32_t id, uint64_t pts, uint32_t timebase,
  uint32_t sample_rate, uint32_t sample_num, uint32_t channel_num);
*/
import "C"
import (
	"bytes"
	"reflect"
	"unsafe"
)

// IsAudio 該当フレームタイプがaudioであるか判定する
func IsAudio(frameType frameType) bool {
	switch frameType {
	case FrameTypes.Aac,
		FrameTypes.AdpcmImaWav,
		FrameTypes.Mp3,
		FrameTypes.Nellymoser,
		FrameTypes.Opus,
		FrameTypes.PcmAlaw,
		FrameTypes.PcmF32,
		FrameTypes.PcmMulaw,
		FrameTypes.PcmS16,
		FrameTypes.Speex,
		FrameTypes.Vorbis:
		return true
	default:
		break
	}
	return false
}

// AudioFrame 音声フレーム
type AudioFrame struct {
	Type       frameType
	Pts        uint64
	Dts        uint64
	Timebase   uint32
	ID         uint32
	SampleRate uint32
	SampleNum  uint32
	ChannelNum uint32
	cFrame     cttLibCFrame
}

// ToFrame 通常のFrameに戻す
func (audio *AudioFrame) ToFrame() *Frame {
	frame := new(Frame)
	frame.cFrame = audio.cFrame
	frame.ID = audio.ID
	frame.Pts = audio.Pts
	frame.Dts = audio.Dts
	frame.Timebase = audio.Timebase
	frame.Type = audio.Type
	return frame
}

// newGoRefFrame goの配列データをcの世界に持っていくデータを構築する
func (audio *AudioFrame) newGoRefFrame() unsafe.Pointer {
	return C.newGoFrame(C.uint64_t(audio.Pts), C.uint64_t(audio.Dts), C.uint32_t(audio.Timebase), C.uint32_t(audio.ID),
		0, 0,
		C.uint32_t(audio.SampleRate), C.uint32_t(audio.SampleNum), C.uint32_t(audio.ChannelNum),
		0, 0,
		0, 0,
		0, 0,
		0, 0)
}

// refCFrame interface経由でcframe参照
func (audio *AudioFrame) refCFrame() cttLibCFrame {
	return audio.cFrame
}

// deleteGoRefFrame cの世界に持っていったデータを破棄する
func (audio *AudioFrame) deleteGoRefFrame(ptr unsafe.Pointer) {
	C.deleteGoFrame(ptr)
}

// GetBinaryBuffer binaryデータを参照する
func (audio *AudioFrame) GetBinaryBuffer(callback DataCallback) bool {
	return getFrameBinaryBuffer(audio.cFrame, callback)
}

// Audio 音声処理
var Audio = struct {
	Cast func(frame *Frame) *AudioFrame
}{
	Cast: func(frame *Frame) *AudioFrame {
		if !IsAudio(frame.Type) {
			return nil
		}
		audioFrame := new(AudioFrame)
		audioFrame.ID = frame.ID
		audioFrame.Pts = frame.Pts
		audioFrame.Dts = frame.Dts
		audioFrame.Timebase = frame.Timebase
		audioFrame.Type = frame.Type
		audioFrame.cFrame = frame.cFrame
		ptr := unsafe.Pointer(frame.cFrame)
		audioFrame.SampleRate = uint32(C.AudioFrame_getSampleRate(ptr))
		audioFrame.SampleNum = uint32(C.AudioFrame_getSampleNum(ptr))
		audioFrame.ChannelNum = uint32(C.AudioFrame_getChannelNum(ptr))
		return audioFrame
	},
}

// Aac処理 ------------------------------------------------

// AacTypes Aacタイプリスト
var AacTypes = struct {
	Raw  subType
	Adts subType
	Dsi  subType
}{
	subType{"raw"},
	subType{"adts"},
	subType{"dsi"},
}

// AacFrame Aacフレーム
type AacFrame struct {
	Type       frameType
	Pts        uint64
	Dts        uint64
	Timebase   uint32
	ID         uint32
	SampleRate uint32
	SampleNum  uint32
	ChannelNum uint32
	AacType    subType
	cFrame     cttLibCFrame
}

// ToFrame 通常のFrameに戻す
func (aac *AacFrame) ToFrame() *Frame {
	frame := new(Frame)
	frame.cFrame = aac.cFrame
	frame.ID = aac.ID
	frame.Pts = aac.Pts
	frame.Dts = aac.Dts
	frame.Timebase = aac.Timebase
	frame.Type = aac.Type
	return frame
}

// newGoRefFrame goの配列データをcの世界に持っていくデータを構築する
func (aac *AacFrame) newGoRefFrame() unsafe.Pointer {
	return C.newGoFrame(C.uint64_t(aac.Pts), C.uint64_t(aac.Dts), C.uint32_t(aac.Timebase), C.uint32_t(aac.ID),
		0, 0,
		C.uint32_t(aac.SampleRate), C.uint32_t(aac.SampleNum), C.uint32_t(aac.ChannelNum),
		0, 0,
		0, 0,
		0, 0,
		0, 0)
}

// refCFrame interface経由でcframe参照
func (aac *AacFrame) refCFrame() cttLibCFrame {
	return aac.cFrame
}

// deleteGoRefFrame cの世界に持っていったデータを破棄する
func (aac *AacFrame) deleteGoRefFrame(ptr unsafe.Pointer) {
	C.deleteGoFrame(ptr)
}

// GetBinaryBuffer binaryデータを参照する
func (aac *AacFrame) GetBinaryBuffer(callback DataCallback) bool {
	return getFrameBinaryBuffer(aac.cFrame, callback)
}

// Aac aac処理
var Aac = struct {
	Cast       func(frame *Frame) *AacFrame
	FromBinary func(binary []byte, length uint64, id uint32, pts uint64, timebase uint32, dsiFrame IFrame) *Frame
}{
	Cast: func(frame *Frame) *AacFrame {
		audioFrame := Audio.Cast(frame)
		if audioFrame == nil {
			return nil
		}
		aacFrame := new(AacFrame)
		aacFrame.ID = audioFrame.ID
		aacFrame.Pts = audioFrame.Pts
		aacFrame.Dts = audioFrame.Dts
		aacFrame.Timebase = audioFrame.Timebase
		aacFrame.Type = audioFrame.Type
		aacFrame.cFrame = audioFrame.cFrame
		aacFrame.SampleRate = audioFrame.SampleRate
		aacFrame.SampleNum = audioFrame.SampleNum
		aacFrame.ChannelNum = audioFrame.ChannelNum
		// あとは追加データ
		buf := make([]byte, 256)
		ptr := unsafe.Pointer(frame.cFrame)
		C.AacFrame_getAacType(ptr, (*C.char)(unsafe.Pointer(&buf[0])), C.size_t(256))
		name := string(buf[:bytes.IndexByte(buf, 0)])
		aacFrame.AacType = subType{name}
		return aacFrame
	},
	FromBinary: func(binary []byte, length uint64, id uint32, pts uint64, timebase uint32, dsiFrame IFrame) *Frame {
		var ptr unsafe.Pointer = nil
		if !reflect.ValueOf(dsiFrame).IsNil() {
			ptr = unsafe.Pointer(dsiFrame.refCFrame())
		}
		cFrame := C.AacFrame_fromBinary(unsafe.Pointer(&binary[0]), C.size_t(length),
			C.uint32_t(id), C.uint64_t(pts), C.uint32_t(timebase), ptr)
		frame := new(Frame)
		frame.init(cttLibCFrame(cFrame))
		frame.hasBody = true
		return frame
	},
}

// AdpcmImaWav処理 ------------------------------------------------

// AdpcmImaWavFrame AdpcmImaWavフレーム
type AdpcmImaWavFrame AudioFrame

// ToFrame 通常のFrameに戻す
func (adpcmImaWav *AdpcmImaWavFrame) ToFrame() *Frame {
	frame := new(Frame)
	frame.cFrame = adpcmImaWav.cFrame
	frame.ID = adpcmImaWav.ID
	frame.Pts = adpcmImaWav.Pts
	frame.Dts = adpcmImaWav.Dts
	frame.Timebase = adpcmImaWav.Timebase
	frame.Type = adpcmImaWav.Type
	return frame
}

// newGoRefFrame goの配列データをcの世界に持っていくデータを構築する
func (adpcmImaWav *AdpcmImaWavFrame) newGoRefFrame() unsafe.Pointer {
	return C.newGoFrame(C.uint64_t(adpcmImaWav.Pts), C.uint64_t(adpcmImaWav.Dts), C.uint32_t(adpcmImaWav.Timebase), C.uint32_t(adpcmImaWav.ID),
		0, 0,
		C.uint32_t(adpcmImaWav.SampleRate), C.uint32_t(adpcmImaWav.SampleNum), C.uint32_t(adpcmImaWav.ChannelNum),
		0, 0,
		0, 0,
		0, 0,
		0, 0)
}

// refCFrame interface経由でcframe参照
func (adpcmImaWav *AdpcmImaWavFrame) refCFrame() cttLibCFrame {
	return adpcmImaWav.cFrame
}

// deleteGoRefFrame cの世界に持っていったデータを破棄する
func (adpcmImaWav *AdpcmImaWavFrame) deleteGoRefFrame(ptr unsafe.Pointer) {
	C.deleteGoFrame(ptr)
}

// GetBinaryBuffer binaryデータを参照する
func (adpcmImaWav *AdpcmImaWavFrame) GetBinaryBuffer(callback DataCallback) bool {
	return getFrameBinaryBuffer(adpcmImaWav.cFrame, callback)
}

// AdpcmImaWav AdpcmImaWav処理
var AdpcmImaWav = struct {
	Cast       func(frame *Frame) *AdpcmImaWavFrame
	FromBinary func(binary []byte, length uint64, id uint32, pts uint64, timebase uint32,
		sampleRate uint32, sampleNum uint32, channelNum uint32) *Frame
}{
	Cast: func(frame *Frame) *AdpcmImaWavFrame {
		audioFrame := Audio.Cast(frame)
		if audioFrame == nil {
			return nil
		}
		adpcmImaWavFrame := new(AdpcmImaWavFrame)
		adpcmImaWavFrame.ID = audioFrame.ID
		adpcmImaWavFrame.Pts = audioFrame.Pts
		adpcmImaWavFrame.Dts = audioFrame.Dts
		adpcmImaWavFrame.Timebase = audioFrame.Timebase
		adpcmImaWavFrame.Type = audioFrame.Type
		adpcmImaWavFrame.cFrame = audioFrame.cFrame
		adpcmImaWavFrame.SampleRate = audioFrame.SampleRate
		adpcmImaWavFrame.SampleNum = audioFrame.SampleNum
		adpcmImaWavFrame.ChannelNum = audioFrame.ChannelNum
		return adpcmImaWavFrame
	},
	FromBinary: func(binary []byte, length uint64, id uint32, pts uint64, timebase uint32,
		sampleRate uint32, sampleNum uint32, channelNum uint32) *Frame {
		cFrame := C.AdpcmImaWavFrame_fromBinary(unsafe.Pointer(&binary[0]), C.size_t(length), C.uint32_t(id), C.uint64_t(pts), C.uint32_t(timebase),
			C.uint32_t(sampleRate), C.uint32_t(sampleNum), C.uint32_t(channelNum))
		frame := new(Frame)
		frame.init(cttLibCFrame(cFrame))
		frame.hasBody = true
		return frame
	},
}

// Mp3処理 ------------------------------------------------

// Mp3Types Mp3タイプリスト
var Mp3Types = struct {
	Tag   subType
	ID3   subType
	Frame subType
}{
	subType{"tag"},
	subType{"id3"},
	subType{"frame"},
}

// Mp3Frame Mp3フレーム
type Mp3Frame struct {
	Type       frameType
	Pts        uint64
	Dts        uint64
	Timebase   uint32
	ID         uint32
	SampleRate uint32
	SampleNum  uint32
	ChannelNum uint32
	Mp3Type    subType
	cFrame     cttLibCFrame
}

// ToFrame 通常のFrameに戻す
func (mp3 *Mp3Frame) ToFrame() *Frame {
	frame := new(Frame)
	frame.cFrame = mp3.cFrame
	frame.ID = mp3.ID
	frame.Pts = mp3.Pts
	frame.Dts = mp3.Dts
	frame.Timebase = mp3.Timebase
	frame.Type = mp3.Type
	return frame
}

// newGoRefFrame goの配列データをcの世界に持っていくデータを構築する
func (mp3 *Mp3Frame) newGoRefFrame() unsafe.Pointer {
	return C.newGoFrame(C.uint64_t(mp3.Pts), C.uint64_t(mp3.Dts), C.uint32_t(mp3.Timebase), C.uint32_t(mp3.ID),
		0, 0,
		C.uint32_t(mp3.SampleRate), C.uint32_t(mp3.SampleNum), C.uint32_t(mp3.ChannelNum),
		0, 0,
		0, 0,
		0, 0,
		0, 0)
}

// refCFrame interface経由でcframe参照
func (mp3 *Mp3Frame) refCFrame() cttLibCFrame {
	return mp3.cFrame
}

// deleteGoRefFrame cの世界に持っていったデータを破棄する
func (mp3 *Mp3Frame) deleteGoRefFrame(ptr unsafe.Pointer) {
	C.deleteGoFrame(ptr)
}

// GetBinaryBuffer binaryデータを参照する
func (mp3 *Mp3Frame) GetBinaryBuffer(callback DataCallback) bool {
	return getFrameBinaryBuffer(mp3.cFrame, callback)
}

// Mp3 Mp3処理
var Mp3 = struct {
	Cast       func(frame *Frame) *Mp3Frame
	FromBinary func(binary []byte, length uint64, id uint32, pts uint64, timebase uint32) *Frame
}{
	Cast: func(frame *Frame) *Mp3Frame {
		audioFrame := Audio.Cast(frame)
		if audioFrame == nil {
			return nil
		}
		mp3Frame := new(Mp3Frame)
		mp3Frame.ID = audioFrame.ID
		mp3Frame.Pts = audioFrame.Pts
		mp3Frame.Dts = audioFrame.Dts
		mp3Frame.Timebase = audioFrame.Timebase
		mp3Frame.Type = audioFrame.Type
		mp3Frame.cFrame = audioFrame.cFrame
		mp3Frame.SampleRate = audioFrame.SampleRate
		mp3Frame.SampleNum = audioFrame.SampleNum
		mp3Frame.ChannelNum = audioFrame.ChannelNum
		// あとは追加データ
		buf := make([]byte, 256)
		ptr := unsafe.Pointer(frame.cFrame)
		C.Mp3Frame_getMp3Type(ptr, (*C.char)(unsafe.Pointer(&buf[0])), C.size_t(256))
		name := string(buf[:bytes.IndexByte(buf, 0)])
		mp3Frame.Mp3Type = subType{name}
		return mp3Frame
	},
	FromBinary: func(binary []byte, length uint64, id uint32, pts uint64, timebase uint32) *Frame {
		cFrame := C.Mp3Frame_fromBinary(unsafe.Pointer(&binary[0]), C.size_t(length), C.uint32_t(id), C.uint64_t(pts), C.uint32_t(timebase))
		frame := new(Frame)
		frame.init(cttLibCFrame(cFrame))
		frame.hasBody = true
		return frame
	},
}

// Nellymoser処理 ------------------------------------------------

// NellymoserFrame Nellymoserフレーム
type NellymoserFrame AudioFrame

// ToFrame 通常のFrameに戻す
func (nellymoser *NellymoserFrame) ToFrame() *Frame {
	frame := new(Frame)
	frame.cFrame = nellymoser.cFrame
	frame.ID = nellymoser.ID
	frame.Pts = nellymoser.Pts
	frame.Dts = nellymoser.Dts
	frame.Timebase = nellymoser.Timebase
	frame.Type = nellymoser.Type
	return frame
}

// newGoRefFrame goの配列データをcの世界に持っていくデータを構築する
func (nellymoser *NellymoserFrame) newGoRefFrame() unsafe.Pointer {
	return C.newGoFrame(C.uint64_t(nellymoser.Pts), C.uint64_t(nellymoser.Dts), C.uint32_t(nellymoser.Timebase), C.uint32_t(nellymoser.ID),
		0, 0,
		C.uint32_t(nellymoser.SampleRate), C.uint32_t(nellymoser.SampleNum), C.uint32_t(nellymoser.ChannelNum),
		0, 0,
		0, 0,
		0, 0,
		0, 0)
}

// refCFrame interface経由でcframe参照
func (nellymoser *NellymoserFrame) refCFrame() cttLibCFrame {
	return nellymoser.cFrame
}

// deleteGoRefFrame cの世界に持っていったデータを破棄する
func (nellymoser *NellymoserFrame) deleteGoRefFrame(ptr unsafe.Pointer) {
	C.deleteGoFrame(ptr)
}

// GetBinaryBuffer binaryデータを参照する
func (nellymoser *NellymoserFrame) GetBinaryBuffer(callback DataCallback) bool {
	return getFrameBinaryBuffer(nellymoser.cFrame, callback)
}

// Nellymoser Nellymoser処理
var Nellymoser = struct {
	Cast       func(frame *Frame) *NellymoserFrame
	FromBinary func(binary []byte, length uint64, id uint32, pts uint64, timebase uint32,
		sampleRate uint32, sampleNum uint32, channelNum uint32) *Frame
}{
	Cast: func(frame *Frame) *NellymoserFrame {
		audioFrame := Audio.Cast(frame)
		if audioFrame == nil {
			return nil
		}
		nellymoserFrame := new(NellymoserFrame)
		nellymoserFrame.ID = audioFrame.ID
		nellymoserFrame.Pts = audioFrame.Pts
		nellymoserFrame.Dts = audioFrame.Dts
		nellymoserFrame.Timebase = audioFrame.Timebase
		nellymoserFrame.Type = audioFrame.Type
		nellymoserFrame.cFrame = audioFrame.cFrame
		nellymoserFrame.SampleRate = audioFrame.SampleRate
		nellymoserFrame.SampleNum = audioFrame.SampleNum
		nellymoserFrame.ChannelNum = audioFrame.ChannelNum
		return nellymoserFrame
	},
	FromBinary: func(binary []byte, length uint64, id uint32, pts uint64, timebase uint32,
		sampleRate uint32, sampleNum uint32, channelNum uint32) *Frame {
		cFrame := C.NellymoserFrame_fromBinary(unsafe.Pointer(&binary[0]), C.size_t(length), C.uint32_t(id), C.uint64_t(pts), C.uint32_t(timebase),
			C.uint32_t(sampleRate), C.uint32_t(sampleNum), C.uint32_t(channelNum))
		frame := new(Frame)
		frame.init(cttLibCFrame(cFrame))
		frame.hasBody = true
		return frame
	},
}

// Opus処理 ------------------------------------------------

// OpusTypes Opusタイプリスト
var OpusTypes = struct {
	Header  subType
	Comment subType
	Frame   subType
}{
	subType{"header"},
	subType{"comment"},
	subType{"frame"},
}

// OpusFrame Opusフレーム
type OpusFrame struct {
	Type       frameType
	Pts        uint64
	Dts        uint64
	Timebase   uint32
	ID         uint32
	SampleRate uint32
	SampleNum  uint32
	ChannelNum uint32
	OpusType   subType
	cFrame     cttLibCFrame
}

// ToFrame 通常のFrameに戻す
func (opus *OpusFrame) ToFrame() *Frame {
	frame := new(Frame)
	frame.cFrame = opus.cFrame
	frame.ID = opus.ID
	frame.Pts = opus.Pts
	frame.Dts = opus.Dts
	frame.Timebase = opus.Timebase
	frame.Type = opus.Type
	return frame
}

// newGoRefFrame goの配列データをcの世界に持っていくデータを構築する
func (opus *OpusFrame) newGoRefFrame() unsafe.Pointer {
	return C.newGoFrame(C.uint64_t(opus.Pts), C.uint64_t(opus.Dts), C.uint32_t(opus.Timebase), C.uint32_t(opus.ID),
		0, 0,
		C.uint32_t(opus.SampleRate), C.uint32_t(opus.SampleNum), C.uint32_t(opus.ChannelNum),
		0, 0,
		0, 0,
		0, 0,
		0, 0)
}

// refCFrame interface経由でcframe参照
func (opus *OpusFrame) refCFrame() cttLibCFrame {
	return opus.cFrame
}

// deleteGoRefFrame cの世界に持っていったデータを破棄する
func (opus *OpusFrame) deleteGoRefFrame(ptr unsafe.Pointer) {
	C.deleteGoFrame(ptr)
}

// GetBinaryBuffer binaryデータを参照する
func (opus *OpusFrame) GetBinaryBuffer(callback DataCallback) bool {
	return getFrameBinaryBuffer(opus.cFrame, callback)
}

// Opus Opus処理
var Opus = struct {
	Cast       func(frame *Frame) *OpusFrame
	FromBinary func(binary []byte, length uint64, id uint32, pts uint64, timebase uint32) *Frame
}{
	Cast: func(frame *Frame) *OpusFrame {
		audioFrame := Audio.Cast(frame)
		if audioFrame == nil {
			return nil
		}
		opusFrame := new(OpusFrame)
		opusFrame.ID = audioFrame.ID
		opusFrame.Pts = audioFrame.Pts
		opusFrame.Dts = audioFrame.Dts
		opusFrame.Timebase = audioFrame.Timebase
		opusFrame.Type = audioFrame.Type
		opusFrame.cFrame = audioFrame.cFrame
		opusFrame.SampleRate = audioFrame.SampleRate
		opusFrame.SampleNum = audioFrame.SampleNum
		opusFrame.ChannelNum = audioFrame.ChannelNum
		// あとは追加データ
		buf := make([]byte, 256)
		ptr := unsafe.Pointer(frame.cFrame)
		C.OpusFrame_getOpusType(ptr, (*C.char)(unsafe.Pointer(&buf[0])), C.size_t(256))
		name := string(buf[:bytes.IndexByte(buf, 0)])
		opusFrame.OpusType = subType{name}
		return opusFrame
	},
	FromBinary: func(binary []byte, length uint64, id uint32, pts uint64, timebase uint32) *Frame {
		cFrame := C.OpusFrame_fromBinary(unsafe.Pointer(&binary[0]), C.size_t(length), C.uint32_t(id), C.uint64_t(pts), C.uint32_t(timebase))
		frame := new(Frame)
		frame.init(cttLibCFrame(cFrame))
		frame.hasBody = true
		return frame
	},
}

// PcmAlaw処理 ------------------------------------------------

// PcmAlawFrame PcmAlawフレーム
type PcmAlawFrame AudioFrame

// ToFrame 通常のFrameに戻す
func (pcmAlaw *PcmAlawFrame) ToFrame() *Frame {
	frame := new(Frame)
	frame.cFrame = pcmAlaw.cFrame
	frame.ID = pcmAlaw.ID
	frame.Pts = pcmAlaw.Pts
	frame.Dts = pcmAlaw.Dts
	frame.Timebase = pcmAlaw.Timebase
	frame.Type = pcmAlaw.Type
	return frame
}

// newGoRefFrame goの配列データをcの世界に持っていくデータを構築する
func (pcmAlaw *PcmAlawFrame) newGoRefFrame() unsafe.Pointer {
	return C.newGoFrame(C.uint64_t(pcmAlaw.Pts), C.uint64_t(pcmAlaw.Dts), C.uint32_t(pcmAlaw.Timebase), C.uint32_t(pcmAlaw.ID),
		0, 0,
		C.uint32_t(pcmAlaw.SampleRate), C.uint32_t(pcmAlaw.SampleNum), C.uint32_t(pcmAlaw.ChannelNum),
		0, 0,
		0, 0,
		0, 0,
		0, 0)
}

// refCFrame interface経由でcframe参照
func (pcmAlaw *PcmAlawFrame) refCFrame() cttLibCFrame {
	return pcmAlaw.cFrame
}

// deleteGoRefFrame cの世界に持っていったデータを破棄する
func (pcmAlaw *PcmAlawFrame) deleteGoRefFrame(ptr unsafe.Pointer) {
	C.deleteGoFrame(ptr)
}

// GetBinaryBuffer binaryデータを参照する
func (pcmAlaw *PcmAlawFrame) GetBinaryBuffer(callback DataCallback) bool {
	return getFrameBinaryBuffer(pcmAlaw.cFrame, callback)
}

// PcmAlaw PcmAlaw処理
var PcmAlaw = struct {
	Cast       func(frame *Frame) *PcmAlawFrame
	FromBinary func(binary []byte, length uint64, id uint32, pts uint64, timebase uint32,
		sampleRate uint32, sampleNum uint32, channelNum uint32) *Frame
}{
	Cast: func(frame *Frame) *PcmAlawFrame {
		audioFrame := Audio.Cast(frame)
		if audioFrame == nil {
			return nil
		}
		pcmAlawFrame := new(PcmAlawFrame)
		pcmAlawFrame.ID = audioFrame.ID
		pcmAlawFrame.Pts = audioFrame.Pts
		pcmAlawFrame.Dts = audioFrame.Dts
		pcmAlawFrame.Timebase = audioFrame.Timebase
		pcmAlawFrame.Type = audioFrame.Type
		pcmAlawFrame.cFrame = audioFrame.cFrame
		pcmAlawFrame.SampleRate = audioFrame.SampleRate
		pcmAlawFrame.SampleNum = audioFrame.SampleNum
		pcmAlawFrame.ChannelNum = audioFrame.ChannelNum
		return pcmAlawFrame
	},
	FromBinary: func(binary []byte, length uint64, id uint32, pts uint64, timebase uint32,
		sampleRate uint32, sampleNum uint32, channelNum uint32) *Frame {
		cFrame := C.PcmAlawsFrame_fromBinary(unsafe.Pointer(&binary[0]), C.size_t(length), C.uint32_t(id), C.uint64_t(pts), C.uint32_t(timebase),
			C.uint32_t(sampleRate), C.uint32_t(sampleNum), C.uint32_t(channelNum))
		frame := new(Frame)
		frame.init(cttLibCFrame(cFrame))
		frame.hasBody = true
		return frame
	},
}

// PcmF32処理 ------------------------------------------------

// PcmF32Types PcmF32タイプリスト
var PcmF32Types = struct {
	Interleave subType
	Planar     subType
}{
	subType{"interleave"},
	subType{"planar"},
}

// PcmF32Frame PcmF32フレーム
type PcmF32Frame struct {
	Type       frameType
	Pts        uint64
	Dts        uint64
	Timebase   uint32
	ID         uint32
	SampleRate uint32
	SampleNum  uint32
	ChannelNum uint32
	PcmF32Type subType
	LStride    uint32 // pcmについては、binaryを復元して、取得したデータを書き換えて、pcmを作り直すというのが編集で使えそう
	LStep      uint32
	RStride    uint32
	RStep      uint32
	cFrame     cttLibCFrame
}

// ToFrame 通常のFrameに戻す
func (pcmF32 *PcmF32Frame) ToFrame() *Frame {
	frame := new(Frame)
	frame.cFrame = pcmF32.cFrame
	frame.ID = pcmF32.ID
	frame.Pts = pcmF32.Pts
	frame.Dts = pcmF32.Dts
	frame.Timebase = pcmF32.Timebase
	frame.Type = pcmF32.Type
	return frame
}

// newGoRefFrame goの配列データをcの世界に持っていくデータを構築する
func (pcmF32 *PcmF32Frame) newGoRefFrame() unsafe.Pointer {
	return C.newGoFrame(C.uint64_t(pcmF32.Pts), C.uint64_t(pcmF32.Dts), C.uint32_t(pcmF32.Timebase), C.uint32_t(pcmF32.ID),
		0, 0,
		C.uint32_t(pcmF32.SampleRate), C.uint32_t(pcmF32.SampleNum), C.uint32_t(pcmF32.ChannelNum),
		0, 0,
		0, 0,
		0, 0,
		0, 0)
}

// refCFrame interface経由でcframe参照
func (pcmF32 *PcmF32Frame) refCFrame() cttLibCFrame {
	return pcmF32.cFrame
}

// deleteGoRefFrame cの世界に持っていったデータを破棄する
func (pcmF32 *PcmF32Frame) deleteGoRefFrame(ptr unsafe.Pointer) {
	C.deleteGoFrame(ptr)
}

// GetBinaryBuffer binaryデータを参照する
func (pcmF32 *PcmF32Frame) GetBinaryBuffer(callback DataCallback) bool {
	return getFrameBinaryBuffer(pcmF32.cFrame, callback)
}

// PcmF32 PcmF32処理
var PcmF32 = struct {
	Cast       func(frame *Frame) *PcmF32Frame
	FromBinary func(binary []byte, length uint64, id uint32, pts uint64, timebase uint32,
		pcmF32Type subType, sampleRate uint32, sampleNum uint32, channelNum uint32,
		lIndex uint32, rIndex uint32) *Frame
}{
	Cast: func(frame *Frame) *PcmF32Frame {
		audioFrame := Audio.Cast(frame)
		if audioFrame == nil {
			return nil
		}
		pcmF32Frame := new(PcmF32Frame)
		pcmF32Frame.ID = audioFrame.ID
		pcmF32Frame.Pts = audioFrame.Pts
		pcmF32Frame.Dts = audioFrame.Dts
		pcmF32Frame.Timebase = audioFrame.Timebase
		pcmF32Frame.Type = audioFrame.Type
		pcmF32Frame.cFrame = audioFrame.cFrame
		pcmF32Frame.SampleRate = audioFrame.SampleRate
		pcmF32Frame.SampleNum = audioFrame.SampleNum
		pcmF32Frame.ChannelNum = audioFrame.ChannelNum
		// あとは追加データ
		buf := make([]byte, 256)
		ptr := unsafe.Pointer(frame.cFrame)
		C.PcmF32Frame_getPcmF32Type(ptr, (*C.char)(unsafe.Pointer(&buf[0])), C.size_t(256))
		name := string(buf[:bytes.IndexByte(buf, 0)])
		pcmF32Frame.PcmF32Type = subType{name}
		pcmF32Frame.LStride = uint32(C.PcmF32Frame_getLStride(ptr))
		pcmF32Frame.LStep = uint32(C.PcmF32Frame_getLStep(ptr))
		pcmF32Frame.RStride = uint32(C.PcmF32Frame_getRStride(ptr))
		pcmF32Frame.RStep = uint32(C.PcmF32Frame_getRStep(ptr))
		return pcmF32Frame
	},
	FromBinary: func(binary []byte, length uint64, id uint32, pts uint64, timebase uint32,
		pcmF32Type subType, sampleRate uint32, sampleNum uint32, channelNum uint32,
		lIndex uint32, rIndex uint32) *Frame {
		cPcmF32Type := C.CString(pcmF32Type.value)
		defer C.free(unsafe.Pointer(cPcmF32Type))
		cFrame := C.PcmF32Frame_fromBinary(unsafe.Pointer(&binary[0]), C.size_t(length), C.uint32_t(id), C.uint64_t(pts), C.uint32_t(timebase),
			cPcmF32Type, C.uint32_t(sampleRate), C.uint32_t(sampleNum), C.uint32_t(channelNum),
			C.uint32_t(lIndex), C.uint32_t(rIndex))
		frame := new(Frame)
		frame.init(cttLibCFrame(cFrame))
		frame.hasBody = true
		return frame
	},
}

// PcmMulaw処理 ------------------------------------------------

// PcmMulawFrame PcmMulawフレーム
type PcmMulawFrame AudioFrame

// ToFrame 通常のFrameに戻す
func (pcmMulaw *PcmMulawFrame) ToFrame() *Frame {
	frame := new(Frame)
	frame.cFrame = pcmMulaw.cFrame
	frame.ID = pcmMulaw.ID
	frame.Pts = pcmMulaw.Pts
	frame.Dts = pcmMulaw.Dts
	frame.Timebase = pcmMulaw.Timebase
	frame.Type = pcmMulaw.Type
	return frame
}

// newGoRefFrame goの配列データをcの世界に持っていくデータを構築する
func (pcmMulaw *PcmMulawFrame) newGoRefFrame() unsafe.Pointer {
	return C.newGoFrame(C.uint64_t(pcmMulaw.Pts), C.uint64_t(pcmMulaw.Dts), C.uint32_t(pcmMulaw.Timebase), C.uint32_t(pcmMulaw.ID),
		0, 0,
		C.uint32_t(pcmMulaw.SampleRate), C.uint32_t(pcmMulaw.SampleNum), C.uint32_t(pcmMulaw.ChannelNum),
		0, 0,
		0, 0,
		0, 0,
		0, 0)
}

// refCFrame interface経由でcframe参照
func (pcmMulaw *PcmMulawFrame) refCFrame() cttLibCFrame {
	return pcmMulaw.cFrame
}

// deleteGoRefFrame cの世界に持っていったデータを破棄する
func (pcmMulaw *PcmMulawFrame) deleteGoRefFrame(ptr unsafe.Pointer) {
	C.deleteGoFrame(ptr)
}

// GetBinaryBuffer binaryデータを参照する
func (pcmMulaw *PcmMulawFrame) GetBinaryBuffer(callback DataCallback) bool {
	return getFrameBinaryBuffer(pcmMulaw.cFrame, callback)
}

// PcmMulaw PcmMulaw処理
var PcmMulaw = struct {
	Cast       func(frame *Frame) *PcmMulawFrame
	FromBinary func(binary []byte, length uint64, id uint32, pts uint64, timebase uint32,
		sampleRate uint32, sampleNum uint32, channelNum uint32) *Frame
}{
	Cast: func(frame *Frame) *PcmMulawFrame {
		audioFrame := Audio.Cast(frame)
		if audioFrame == nil {
			return nil
		}
		pcmMulawFrame := new(PcmMulawFrame)
		pcmMulawFrame.ID = audioFrame.ID
		pcmMulawFrame.Pts = audioFrame.Pts
		pcmMulawFrame.Dts = audioFrame.Dts
		pcmMulawFrame.Timebase = audioFrame.Timebase
		pcmMulawFrame.Type = audioFrame.Type
		pcmMulawFrame.cFrame = audioFrame.cFrame
		pcmMulawFrame.SampleRate = audioFrame.SampleRate
		pcmMulawFrame.SampleNum = audioFrame.SampleNum
		pcmMulawFrame.ChannelNum = audioFrame.ChannelNum
		return pcmMulawFrame
	},
	FromBinary: func(binary []byte, length uint64, id uint32, pts uint64, timebase uint32,
		sampleRate uint32, sampleNum uint32, channelNum uint32) *Frame {
		cFrame := C.PcmMulawsFrame_fromBinary(unsafe.Pointer(&binary[0]), C.size_t(length), C.uint32_t(id), C.uint64_t(pts), C.uint32_t(timebase),
			C.uint32_t(sampleRate), C.uint32_t(sampleNum), C.uint32_t(channelNum))
		frame := new(Frame)
		frame.init(cttLibCFrame(cFrame))
		frame.hasBody = true
		return frame
	},
}

// PcmS16処理 ------------------------------------------------

// PcmS16Types PcmS16タイプリスト
var PcmS16Types = struct {
	BigEndian          subType
	BigEndianPlanar    subType
	LittleEndian       subType
	LittleEndianPlanar subType
}{
	subType{"bigEndian"},
	subType{"bigEndianPlanar"},
	subType{"littleEndian"},
	subType{"littleEndianPlanar"},
}

// PcmS16Frame PcmS16フレーム
type PcmS16Frame struct {
	Type       frameType
	Pts        uint64
	Dts        uint64
	Timebase   uint32
	ID         uint32
	SampleRate uint32
	SampleNum  uint32
	ChannelNum uint32
	PcmS16Type subType
	LStride    uint32 // pcmについては、binaryを復元して、取得したデータを書き換えて、pcmを作り直すというのが編集で使えそう
	LStep      uint32
	RStride    uint32
	RStep      uint32
	cFrame     cttLibCFrame
}

// ToFrame 通常のFrameに戻す
func (pcmS16 *PcmS16Frame) ToFrame() *Frame {
	frame := new(Frame)
	frame.cFrame = pcmS16.cFrame
	frame.ID = pcmS16.ID
	frame.Pts = pcmS16.Pts
	frame.Dts = pcmS16.Dts
	frame.Timebase = pcmS16.Timebase
	frame.Type = pcmS16.Type
	return frame
}

// newGoRefFrame goの配列データをcの世界に持っていくデータを構築する
func (pcmS16 *PcmS16Frame) newGoRefFrame() unsafe.Pointer {
	return C.newGoFrame(C.uint64_t(pcmS16.Pts), C.uint64_t(pcmS16.Dts), C.uint32_t(pcmS16.Timebase), C.uint32_t(pcmS16.ID),
		0, 0,
		C.uint32_t(pcmS16.SampleRate), C.uint32_t(pcmS16.SampleNum), C.uint32_t(pcmS16.ChannelNum),
		0, 0,
		0, 0,
		0, 0,
		0, 0)
}

// refCFrame interface経由でcframe参照
func (pcmS16 *PcmS16Frame) refCFrame() cttLibCFrame {
	return pcmS16.cFrame
}

// deleteGoRefFrame cの世界に持っていったデータを破棄する
func (pcmS16 *PcmS16Frame) deleteGoRefFrame(ptr unsafe.Pointer) {
	C.deleteGoFrame(ptr)
}

// GetBinaryBuffer binaryデータを参照する
func (pcmS16 *PcmS16Frame) GetBinaryBuffer(callback DataCallback) bool {
	return getFrameBinaryBuffer(pcmS16.cFrame, callback)
}

// PcmS16 PcmS16処理
var PcmS16 = struct {
	Cast       func(frame *Frame) *PcmS16Frame
	FromBinary func(binary []byte, length uint64, id uint32, pts uint64, timebase uint32,
		pcmS16Type subType, sampleRate uint32, sampleNum uint32, channelNum uint32,
		lIndex uint32, rIndex uint32) *Frame
}{
	Cast: func(frame *Frame) *PcmS16Frame {
		audioFrame := Audio.Cast(frame)
		if audioFrame == nil {
			return nil
		}
		pcmS16Frame := new(PcmS16Frame)
		pcmS16Frame.ID = audioFrame.ID
		pcmS16Frame.Pts = audioFrame.Pts
		pcmS16Frame.Dts = audioFrame.Dts
		pcmS16Frame.Timebase = audioFrame.Timebase
		pcmS16Frame.Type = audioFrame.Type
		pcmS16Frame.cFrame = audioFrame.cFrame
		pcmS16Frame.SampleRate = audioFrame.SampleRate
		pcmS16Frame.SampleNum = audioFrame.SampleNum
		pcmS16Frame.ChannelNum = audioFrame.ChannelNum
		// あとは追加データ
		buf := make([]byte, 256)
		ptr := unsafe.Pointer(frame.cFrame)
		C.PcmS16Frame_getPcmS16Type(ptr, (*C.char)(unsafe.Pointer(&buf[0])), C.size_t(256))
		name := string(buf[:bytes.IndexByte(buf, 0)])
		pcmS16Frame.PcmS16Type = subType{name}
		pcmS16Frame.LStride = uint32(C.PcmS16Frame_getLStride(ptr))
		pcmS16Frame.LStep = uint32(C.PcmS16Frame_getLStep(ptr))
		pcmS16Frame.RStride = uint32(C.PcmS16Frame_getRStride(ptr))
		pcmS16Frame.RStep = uint32(C.PcmS16Frame_getRStep(ptr))
		return pcmS16Frame
	},
	FromBinary: func(binary []byte, length uint64, id uint32, pts uint64, timebase uint32,
		pcmS16Type subType, sampleRate uint32, sampleNum uint32, channelNum uint32,
		lIndex uint32, rIndex uint32) *Frame {
		cPcmS16Type := C.CString(pcmS16Type.value)
		defer C.free(unsafe.Pointer(cPcmS16Type))
		// で、ステレオでplanarでない場合は、r_dataは未設定にしなければならないと・・・
		cFrame := C.PcmS16Frame_fromBinary(unsafe.Pointer(&binary[0]), C.size_t(length), C.uint32_t(id), C.uint64_t(pts), C.uint32_t(timebase),
			cPcmS16Type, C.uint32_t(sampleRate), C.uint32_t(sampleNum), C.uint32_t(channelNum),
			C.uint32_t(lIndex), C.uint32_t(rIndex))
		frame := new(Frame)
		frame.init(cttLibCFrame(cFrame))
		frame.hasBody = true
		return frame
	},
}

// Speex処理 ------------------------------------------------

// SpeexTypes Speexタイプリスト
var SpeexTypes = struct {
	Comment subType
	Header  subType
	Frame   subType
}{
	subType{"comment"},
	subType{"header"},
	subType{"frame"},
}

// SpeexFrame Speexフレーム
type SpeexFrame struct {
	Type       frameType
	Pts        uint64
	Dts        uint64
	Timebase   uint32
	ID         uint32
	SampleRate uint32
	SampleNum  uint32
	ChannelNum uint32
	SpeexType  subType
	cFrame     cttLibCFrame
}

// ToFrame 通常のFrameに戻す
func (speex *SpeexFrame) ToFrame() *Frame {
	frame := new(Frame)
	frame.cFrame = speex.cFrame
	frame.ID = speex.ID
	frame.Pts = speex.Pts
	frame.Dts = speex.Dts
	frame.Timebase = speex.Timebase
	frame.Type = speex.Type
	return frame
}

// newGoRefFrame goの配列データをcの世界に持っていくデータを構築する
func (speex *SpeexFrame) newGoRefFrame() unsafe.Pointer {
	return C.newGoFrame(C.uint64_t(speex.Pts), C.uint64_t(speex.Dts), C.uint32_t(speex.Timebase), C.uint32_t(speex.ID),
		0, 0,
		C.uint32_t(speex.SampleRate), C.uint32_t(speex.SampleNum), C.uint32_t(speex.ChannelNum),
		0, 0,
		0, 0,
		0, 0,
		0, 0)
}

// refCFrame interface経由でcframe参照
func (speex *SpeexFrame) refCFrame() cttLibCFrame {
	return speex.cFrame
}

// deleteGoRefFrame cの世界に持っていったデータを破棄する
func (speex *SpeexFrame) deleteGoRefFrame(ptr unsafe.Pointer) {
	C.deleteGoFrame(ptr)
}

// GetBinaryBuffer binaryデータを参照する
func (speex *SpeexFrame) GetBinaryBuffer(callback DataCallback) bool {
	return getFrameBinaryBuffer(speex.cFrame, callback)
}

// Speex Speex処理
var Speex = struct {
	Cast       func(frame *Frame) *SpeexFrame
	FromBinary func(binary []byte, length uint64, id uint32, pts uint64, timebase uint32,
		sampleRate uint32, sampleNum uint32, channelNum uint32) *Frame
}{
	Cast: func(frame *Frame) *SpeexFrame {
		audioFrame := Audio.Cast(frame)
		if audioFrame == nil {
			return nil
		}
		speexFrame := new(SpeexFrame)
		speexFrame.ID = audioFrame.ID
		speexFrame.Pts = audioFrame.Pts
		speexFrame.Dts = audioFrame.Dts
		speexFrame.Timebase = audioFrame.Timebase
		speexFrame.Type = audioFrame.Type
		speexFrame.cFrame = audioFrame.cFrame
		speexFrame.SampleRate = audioFrame.SampleRate
		speexFrame.SampleNum = audioFrame.SampleNum
		speexFrame.ChannelNum = audioFrame.ChannelNum
		// あとは追加データ
		buf := make([]byte, 256)
		ptr := unsafe.Pointer(frame.cFrame)
		C.SpeexFrame_getSpeexType(ptr, (*C.char)(unsafe.Pointer(&buf[0])), C.size_t(256))
		name := string(buf[:bytes.IndexByte(buf, 0)])
		speexFrame.SpeexType = subType{name}
		return speexFrame
	},
	FromBinary: func(binary []byte, length uint64, id uint32, pts uint64, timebase uint32,
		sampleRate uint32, sampleNum uint32, channelNum uint32) *Frame {
		cFrame := C.SpeexFrame_fromBinary(unsafe.Pointer(&binary[0]), C.size_t(length), C.uint32_t(id), C.uint64_t(pts), C.uint32_t(timebase),
			C.uint32_t(sampleRate), C.uint32_t(sampleNum), C.uint32_t(channelNum))
		frame := new(Frame)
		frame.init(cttLibCFrame(cFrame))
		frame.hasBody = true
		return frame
	},
}

// Vorbis処理 ------------------------------------------------

// VorbisTypes Vorbisタイプリスト
var VorbisTypes = struct {
	Comment        subType
	Setup          subType
	Identification subType
	Frame          subType
}{
	subType{"comment"},
	subType{"setup"},
	subType{"identification"},
	subType{"frame"},
}

// VorbisFrame Vorbisフレーム
type VorbisFrame struct {
	Type       frameType
	Pts        uint64
	Dts        uint64
	Timebase   uint32
	ID         uint32
	SampleRate uint32
	SampleNum  uint32
	ChannelNum uint32
	VorbisType subType
	cFrame     cttLibCFrame
}

// ToFrame 通常のFrameに戻す
func (vorbis *VorbisFrame) ToFrame() *Frame {
	frame := new(Frame)
	frame.cFrame = vorbis.cFrame
	frame.ID = vorbis.ID
	frame.Pts = vorbis.Pts
	frame.Dts = vorbis.Dts
	frame.Timebase = vorbis.Timebase
	frame.Type = vorbis.Type
	return frame
}

// newGoRefFrame goの配列データをcの世界に持っていくデータを構築する
func (vorbis *VorbisFrame) newGoRefFrame() unsafe.Pointer {
	return C.newGoFrame(C.uint64_t(vorbis.Pts), C.uint64_t(vorbis.Dts), C.uint32_t(vorbis.Timebase), C.uint32_t(vorbis.ID),
		0, 0,
		C.uint32_t(vorbis.SampleRate), C.uint32_t(vorbis.SampleNum), C.uint32_t(vorbis.ChannelNum),
		0, 0,
		0, 0,
		0, 0,
		0, 0)
}

// refCFrame interface経由でcframe参照
func (vorbis *VorbisFrame) refCFrame() cttLibCFrame {
	return vorbis.cFrame
}

// deleteGoRefFrame cの世界に持っていったデータを破棄する
func (vorbis *VorbisFrame) deleteGoRefFrame(ptr unsafe.Pointer) {
	C.deleteGoFrame(ptr)
}

// GetBinaryBuffer binaryデータを参照する
func (vorbis *VorbisFrame) GetBinaryBuffer(callback DataCallback) bool {
	return getFrameBinaryBuffer(vorbis.cFrame, callback)
}

// Vorbis Vorbis処理
var Vorbis = struct {
	Cast       func(frame *Frame) *VorbisFrame
	FromBinary func(binary []byte, length uint64, id uint32, pts uint64, timebase uint32,
		sampleRate uint32, sampleNum uint32, channelNum uint32) *Frame
}{
	Cast: func(frame *Frame) *VorbisFrame {
		audioFrame := Audio.Cast(frame)
		if audioFrame == nil {
			return nil
		}
		vorbisFrame := new(VorbisFrame)
		vorbisFrame.ID = audioFrame.ID
		vorbisFrame.Pts = audioFrame.Pts
		vorbisFrame.Dts = audioFrame.Dts
		vorbisFrame.Timebase = audioFrame.Timebase
		vorbisFrame.Type = audioFrame.Type
		vorbisFrame.cFrame = audioFrame.cFrame
		vorbisFrame.SampleRate = audioFrame.SampleRate
		vorbisFrame.SampleNum = audioFrame.SampleNum
		vorbisFrame.ChannelNum = audioFrame.ChannelNum
		// あとは追加データ
		buf := make([]byte, 256)
		ptr := unsafe.Pointer(frame.cFrame)
		C.VorbisFrame_getVorbisType(ptr, (*C.char)(unsafe.Pointer(&buf[0])), C.size_t(256))
		name := string(buf[:bytes.IndexByte(buf, 0)])
		vorbisFrame.VorbisType = subType{name}
		return vorbisFrame
	},
	FromBinary: func(binary []byte, length uint64, id uint32, pts uint64, timebase uint32,
		sampleRate uint32, sampleNum uint32, channelNum uint32) *Frame {
		cFrame := C.VorbisFrame_fromBinary(unsafe.Pointer(&binary[0]), C.size_t(length), C.uint32_t(id), C.uint64_t(pts), C.uint32_t(timebase),
			C.uint32_t(sampleRate), C.uint32_t(sampleNum), C.uint32_t(channelNum))
		frame := new(Frame)
		frame.init(cttLibCFrame(cFrame))
		frame.hasBody = true
		return frame
	},
}
