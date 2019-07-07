package ttLibGo

/*
#include <stdlib.h>
#include <stdbool.h>

extern void *Encoder_make(void *mp);
extern bool Encoder_encodeFrame(void *encoder, void *frame, void *goFrame, uintptr_t ptr);
extern void Encoder_close(void *encoder);
*/
import "C"
import (
	"fmt"
	"reflect"
	"unsafe"
)

// cttLibCEncoder c言語のencoder
type cttLibCEncoder unsafe.Pointer

// encoderType encoderタイプ
type encoderType struct{ value string }

// IEncoder エンコーダー
type IEncoder interface {
	EncodeFrame(
		frame IFrame,
		callback FrameCallback) bool
	Close()
}

// encoder エンコーダー
type encoder struct {
	Type     encoderType
	cEncoder cttLibCEncoder
	call     frameCall
}

// EncoderTypes Encoderタイプリスト
var EncoderTypes = struct {
	Fdkaac   encoderType
	Jpeg     encoderType
	Openh264 encoderType
	Opus     encoderType
	Speex    encoderType
	Theora   encoderType
	Vorbis   encoderType
	X264     encoderType
	X265     encoderType
}{
	encoderType{"fdkaac"},
	encoderType{"jpeg"},
	encoderType{"openh264"},
	encoderType{"opus"},
	encoderType{"speex"},
	encoderType{"theora"},
	encoderType{"vorbis"},
	encoderType{"x264"},
	encoderType{"x265"},
}

// Encoders Encoder作成
var Encoders = struct {
	Fdkaac   func(aacType subType, sampleRate uint32, channelNum uint32, bitrate uint32) *fdkaacEncoder
	Jpeg     func(width uint32, height uint32, quality uint32) *jpegEncoder
	Openh264 func(width uint32, height uint32, param map[interface{}]interface{}, spatialParamArray [](map[interface{}]interface{})) *openh264Encoder
	Opus     func(sampleRate uint32, channelNum uint32, unitSampleNum uint32) *opusEncoder
	Speex    func(sampleRate uint32, channelNum uint32, quality uint32) *encoder
	Theora   func(width uint32, height uint32, quality uint32, bitrate uint32, keyFrameInterval uint32) *encoder
	Vorbis   func(sampleRate uint32, channelNum uint32) *encoder
	X264     func(width uint32, height uint32, preset string, tune string, profile string, param map[string]string) *x264Encoder
	X265     func(width uint32, height uint32, preset string, tune string, profile string, param map[string]string) *x265Encoder
}{
	Fdkaac: func(aacType subType, sampleRate uint32, channelNum uint32, bitrate uint32) *fdkaacEncoder {
		encoder := new(fdkaacEncoder)
		encoder.Type = EncoderTypes.Fdkaac
		params := map[string]interface{}{
			"encoder":    encoder.Type.value,
			"aacType":    aacType.value,
			"sampleRate": sampleRate,
			"channelNum": channelNum,
			"bitrate":    bitrate,
		}
		v := mapUtil.fromMap(params)
		encoder.cEncoder = cttLibCEncoder(C.Encoder_make(v))
		mapUtil.close(v)
		return encoder
	},
	Jpeg: func(width uint32, height uint32, quality uint32) *jpegEncoder {
		encoder := new(jpegEncoder)
		encoder.Type = EncoderTypes.Jpeg
		params := map[string]interface{}{
			"encoder": encoder.Type.value,
			"width":   width,
			"height":  height,
			"quality": quality,
		}
		v := mapUtil.fromMap(params)
		encoder.cEncoder = cttLibCEncoder(C.Encoder_make(v))
		mapUtil.close(v)
		return encoder
	},
	Openh264: func(width uint32, height uint32, param map[interface{}]interface{}, spatialParamArray [](map[interface{}]interface{})) *openh264Encoder {
		encoder := new(openh264Encoder)
		encoder.Type = EncoderTypes.Openh264
		var paramSlices []string
		if param != nil {
			for i, v := range param {
				ii := reflect.ValueOf(i)
				paramSlices = append(paramSlices, fmt.Sprintf("%v:%v", ii.FieldByName("value"), v))
			}
		}
		var spatialParamArrays []string
		if spatialParamArray != nil {
			for i, v := range spatialParamArray {
				for ii, vv := range v {
					iii := reflect.ValueOf(ii)
					spatialParamArrays = append(spatialParamArrays, fmt.Sprintf("%d:%v:%v", i, iii.FieldByName("value"), vv))
				}
			}
		}
		params := map[string]interface{}{
			"encoder":           encoder.Type.value,
			"width":             width,
			"height":            height,
			"param":             paramSlices,
			"spatialParamArray": spatialParamArrays,
		}
		v := mapUtil.fromMap(params)
		encoder.cEncoder = cttLibCEncoder(C.Encoder_make(v))
		mapUtil.close(v)
		return encoder
	},
	Opus: func(sampleRate uint32, channelNum uint32, unitSampleNum uint32) *opusEncoder {
		encoder := new(opusEncoder)
		encoder.Type = EncoderTypes.Opus
		params := map[string]interface{}{
			"encoder":       encoder.Type.value,
			"sampleRate":    sampleRate,
			"channelNum":    channelNum,
			"unitSampleNum": unitSampleNum,
		}
		v := mapUtil.fromMap(params)
		encoder.cEncoder = cttLibCEncoder(C.Encoder_make(v))
		mapUtil.close(v)
		return encoder
	},
	Speex: func(sampleRate uint32, channelNum uint32, quality uint32) *encoder {
		encoder := new(encoder)
		encoder.Type = EncoderTypes.Speex
		params := map[string]interface{}{
			"encoder":    encoder.Type.value,
			"sampleRate": sampleRate,
			"channelNum": channelNum,
			"quality":    quality,
		}
		v := mapUtil.fromMap(params)
		encoder.cEncoder = cttLibCEncoder(C.Encoder_make(v))
		mapUtil.close(v)
		return encoder
	},
	Theora: func(width uint32, height uint32, quality uint32, bitrate uint32, keyFrameInterval uint32) *encoder {
		encoder := new(encoder)
		encoder.Type = EncoderTypes.Theora
		params := map[string]interface{}{
			"encoder":          encoder.Type.value,
			"width":            width,
			"height":           height,
			"quality":          quality,
			"bitrate":          bitrate,
			"keyFrameInterval": keyFrameInterval,
		}
		v := mapUtil.fromMap(params)
		encoder.cEncoder = cttLibCEncoder(C.Encoder_make(v))
		mapUtil.close(v)
		return encoder
	},
	Vorbis: func(sampleRate uint32, channelNum uint32) *encoder {
		encoder := new(encoder)
		encoder.Type = EncoderTypes.Vorbis
		params := map[string]interface{}{
			"encoder":    encoder.Type.value,
			"sampleRate": sampleRate,
			"channelNum": channelNum,
		}
		v := mapUtil.fromMap(params)
		encoder.cEncoder = cttLibCEncoder(C.Encoder_make(v))
		mapUtil.close(v)
		return encoder
	},
	X264: func(width uint32, height uint32, preset string, tune string, profile string, param map[string]string) *x264Encoder {
		encoder := new(x264Encoder)
		encoder.Type = EncoderTypes.X264
		var params []string
		if param != nil {
			for i, v := range param {
				params = append(params, fmt.Sprintf("%s:%s", i, v))
			}
		}
		if profile == "" {
			profile = "baseline"
		}
		vv := map[string]interface{}{
			"encoder": encoder.Type.value,
			"width":   width,
			"height":  height,
			"preset":  preset,
			"tune":    tune,
			"profile": profile,
			"param":   params,
		}
		v := mapUtil.fromMap(vv)
		encoder.cEncoder = cttLibCEncoder(C.Encoder_make(v))
		mapUtil.close(v)
		return encoder
	},
	X265: func(width uint32, height uint32, preset string, tune string, profile string, param map[string]string) *x265Encoder {
		encoder := new(x265Encoder)
		encoder.Type = EncoderTypes.X265
		var params []string
		if param != nil {
			for i, v := range param {
				params = append(params, fmt.Sprintf("%s:%s", i, v))
			}
		}
		if profile == "" {
			profile = "main"
		}
		vv := map[string]interface{}{
			"encoder": encoder.Type.value,
			"width":   width,
			"height":  height,
			"preset":  preset,
			"tune":    tune,
			"profile": profile,
			"param":   params,
		}
		v := mapUtil.fromMap(vv)
		encoder.cEncoder = cttLibCEncoder(C.Encoder_make(v))
		mapUtil.close(v)
		return encoder
	},
}

func encoderEncodeFrame(
	encoder *encoder,
	frame IFrame,
	callback FrameCallback) bool {
	if encoder.cEncoder == nil {
		return false
	}
	if frame == nil {
		return true
	}
	cframe := frame.newGoRefFrame()
	encoder.call.callback = callback
	result := C.Encoder_encodeFrame(
		unsafe.Pointer(encoder.cEncoder),
		unsafe.Pointer(frame.refCFrame()),
		cframe,
		C.uintptr_t(uintptr(unsafe.Pointer(&encoder.call))))
	frame.deleteGoRefFrame(cframe)
	return bool(result)
}

func encoderClose(encoder *encoder) {
	C.Encoder_close(unsafe.Pointer(encoder.cEncoder))
}

// EncodeFrame エンコードを実行
func (encoder *encoder) EncodeFrame(
	frame IFrame,
	callback FrameCallback) bool {
	return encoderEncodeFrame(encoder, frame, callback)
}

// Close 閉じる
func (encoder *encoder) Close() {
	encoderClose(encoder)
}