package ttLibGo

/*
#include <stdint.h>
#include <stdlib.h>
#include <stdbool.h>

extern void *Resampler_make(void *mp);
extern bool Resampler_resampleFrame(void *resampler, void *frame, void *goFrame, uintptr_t ptr);
extern void Resampler_close(void *resampler);
*/
import "C"
import (
	"unsafe"
)

// cttlibCResampler c言語のresampler
type cttlibCResampler unsafe.Pointer

// resamplerType resamplerタイプ
type resamplerType struct{ value string }

// IResampler リサンプラ
type IResampler interface {
	ResampleFrame(
		frame IFrame,
		callback FrameCallback) bool
	Close()
}

// resampler リサンプラ
type resampler struct {
	Type       resamplerType
	cResampler cttlibCResampler
	call       frameCall
}

// ResamplerTypes Resamplerタイプリスト
var ResamplerTypes = struct {
	Audio      resamplerType
	Image      resamplerType
	Resize     resamplerType
	Soundtouch resamplerType
	Speexdsp   resamplerType
	Swresample resamplerType
	Swscale    resamplerType
}{
	resamplerType{"audio"},
	resamplerType{"image"},
	resamplerType{"resize"},
	resamplerType{"soundtouch"},
	resamplerType{"speexdsp"},
	resamplerType{"swresample"},
	resamplerType{"swscale"},
}

// Resamplers Resamplers処理
var Resamplers = struct {
	Audio      func(frameType frameType, subType subType) *resampler
	Image      func(frameType frameType, subType subType) *resampler
	Resize     func(width uint32, height uint32, isQuick bool) *resampler
	Soundtouch func(sampleRate uint32, channelNum uint32) *soundtouchResampler
	Speexdsp   func(inSampleRate uint32, outSampleRate uint32, channelNum uint32, quality uint32) *resampler
	Swresample func(inType frameType, inSubType subType, inSampleRate uint32, inChannelNum uint32,
		outType frameType, outSubtype subType, outSampleRate uint32, outChannelNum uint32) *resampler
	Swscale func(inType frameType, inSubType subType, inWidth uint32, inHeight uint32,
		outType frameType, outSubType subType, outWidth uint32, outHeight uint32,
		mode subType) *resampler
}{
	Audio: func(frameType frameType, subType subType) *resampler {
		resampler := new(resampler)
		resampler.Type = ResamplerTypes.Audio
		params := map[string]interface{}{
			"resampler": resampler.Type.value,
			"frameType": frameType.value,
			"subType":   subType.value,
		}
		v := mapUtil.fromMap(params)
		resampler.cResampler = cttlibCResampler(C.Resampler_make(v))
		mapUtil.close(v)
		return resampler
	},
	Image: func(frameType frameType, subType subType) *resampler {
		resampler := new(resampler)
		resampler.Type = ResamplerTypes.Image
		params := map[string]interface{}{
			"resampler": resampler.Type.value,
			"frameType": frameType.value,
			"subType":   subType.value,
		}
		v := mapUtil.fromMap(params)
		resampler.cResampler = cttlibCResampler(C.Resampler_make(v))
		mapUtil.close(v)
		return resampler
	},
	Resize: func(width uint32, height uint32, isQuick bool) *resampler {
		resampler := new(resampler)
		resampler.Type = ResamplerTypes.Resize
		_isQuick := uint32(0)
		if isQuick {
			_isQuick = uint32(1)
		}
		params := map[string]interface{}{
			"resampler": resampler.Type.value,
			"width":     width,
			"height":    height,
			"isQuick":   _isQuick,
		}
		v := mapUtil.fromMap(params)
		resampler.cResampler = cttlibCResampler(C.Resampler_make(v))
		mapUtil.close(v)
		return resampler
	},
	Soundtouch: func(sampleRate uint32, channelNum uint32) *soundtouchResampler {
		resampler := new(soundtouchResampler)
		resampler.Type = ResamplerTypes.Soundtouch
		params := map[string]interface{}{
			"resampler":  resampler.Type.value,
			"sampleRate": sampleRate,
			"channelNum": channelNum,
		}
		v := mapUtil.fromMap(params)
		resampler.cResampler = cttlibCResampler(C.Resampler_make(v))
		mapUtil.close(v)
		return resampler
	},
	Speexdsp: func(inSampleRate uint32, outSampleRate uint32, channelNum uint32, quality uint32) *resampler {
		resampler := new(resampler)
		resampler.Type = ResamplerTypes.Speexdsp
		params := map[string]interface{}{
			"resampler":     resampler.Type.value,
			"inSampleRate":  inSampleRate,
			"outSampleRate": outSampleRate,
			"channelNum":    channelNum,
			"quality":       quality,
		}
		v := mapUtil.fromMap(params)
		resampler.cResampler = cttlibCResampler(C.Resampler_make(v))
		mapUtil.close(v)
		return resampler
	},
	Swresample: func(inType frameType, inSubType subType, inSampleRate uint32, inChannelNum uint32,
		outType frameType, outSubType subType, outSampleRate uint32, outChannelNum uint32) *resampler {
		resampler := new(resampler)
		resampler.Type = ResamplerTypes.Swresample
		params := map[string]interface{}{
			"resampler":     resampler.Type.value,
			"inType":        inType.value,
			"inSubType":     inSubType.value,
			"inSampleRate":  inSampleRate,
			"inChannelNum":  inChannelNum,
			"outType":       outType.value,
			"outSubType":    outSubType.value,
			"outSampleRate": outSampleRate,
			"outChannelNum": outChannelNum,
		}
		v := mapUtil.fromMap(params)
		resampler.cResampler = cttlibCResampler(C.Resampler_make(v))
		mapUtil.close(v)
		return resampler
	},
	Swscale: func(inType frameType, inSubType subType, inWidth uint32, inHeight uint32,
		outType frameType, outSubType subType, outWidth uint32, outHeight uint32,
		mode subType) *resampler {
		resampler := new(resampler)
		resampler.Type = ResamplerTypes.Swscale
		params := map[string]interface{}{
			"resampler":  resampler.Type.value,
			"inType":     inType.value,
			"inSubType":  inSubType.value,
			"inWidth":    inWidth,
			"inHeight":   inHeight,
			"outType":    outType.value,
			"outSubType": outSubType.value,
			"outWidth":   outWidth,
			"outHeight":  outHeight,
			"mode":       mode.value,
		}
		v := mapUtil.fromMap(params)
		resampler.cResampler = cttlibCResampler(C.Resampler_make(v))
		mapUtil.close(v)
		return resampler
	},
}

func resamplerResampleFrame(
	resampler *resampler,
	frame IFrame,
	callback FrameCallback) bool {
	if resampler.cResampler == nil {
		return false
	}
	if frame == nil {
		return true
	}
	cframe := frame.newGoRefFrame()
	resampler.call.callback = callback
	result := C.Resampler_resampleFrame(
		unsafe.Pointer(resampler.cResampler),
		unsafe.Pointer(frame.refCFrame()),
		cframe,
		C.uintptr_t(uintptr(unsafe.Pointer(&resampler.call))))
	frame.deleteGoRefFrame(cframe)
	return bool(result)
}

func resamplerClose(resampler *resampler) {
	C.Resampler_close(unsafe.Pointer(resampler.cResampler))
}

// ResampleFrame リサンプルを実行
func (resampler *resampler) ResampleFrame(
	frame IFrame,
	callback FrameCallback) bool {
	return resamplerResampleFrame(resampler, frame, callback)
}

// Close 閉じる
func (resampler *resampler) Close() {
	resamplerClose(resampler)
}
