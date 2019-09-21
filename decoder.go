package ttLibGo

/*
#include <stdint.h>
#include <stdlib.h>
#include <stdbool.h>

extern void *Decoder_make(void *mp);
extern bool Decoder_decodeFrame(void *decoder, void *frame, void *classFrame, uintptr_t ptr);
extern void Decoder_close(void *decoder);
*/
import "C"
import (
	"unsafe"
)

// cttLibCDecoder c言語のdecoder
type cttLibCDecoder unsafe.Pointer

// decoderType decoderタイプ
type decoderType struct{ value string }

// IDecoder デコーダー
type IDecoder interface {
	DecodeFrame(
		frame IFrame,
		callback func(frame *Frame) bool) bool
	Close()
}

// dsecoder デコーダー
type decoder struct {
	Type     decoderType
	cDecoder cttLibCDecoder
	call     frameCall
}

// DecoderTypes Decoderタイプリスト
var DecoderTypes = struct {
	AvcodecVideo decoderType
	AvcodecAudio decoderType
	Jpeg         decoderType
	Openh264     decoderType
	Opus         decoderType
	Png          decoderType
	Speex        decoderType
	Theora       decoderType
	Vorbis       decoderType
}{
	decoderType{"avcodecVideo"},
	decoderType{"avcodecAudio"},
	decoderType{"jpeg"},
	decoderType{"openh264"},
	decoderType{"opus"},
	decoderType{"png"},
	decoderType{"speex"},
	decoderType{"theora"},
	decoderType{"vorbis"},
}

// Decoders Decoder作成
var Decoders = struct {
	AvcodecVideo func(frameType frameType, width uint32, height uint32) *decoder
	AvcodecAudio func(frameType frameType, sampleRate uint32, channelNum uint32) *decoder
	Jpeg         func() *decoder
	Openh264     func() *decoder
	Opus         func(sampleRate uint32, channelNum uint32) *opusDecoder
	Png          func() *decoder
	Speex        func(sampleRate uint32, channelNum uint32) *decoder
	Theora       func() *decoder
	Vorbis       func() *decoder
}{
	AvcodecVideo: func(frameType frameType, width uint32, height uint32) *decoder {
		decoder := new(decoder)
		decoder.Type = DecoderTypes.AvcodecVideo
		params := map[string]interface{}{
			"decoder":   decoder.Type.value,
			"frameType": frameType.value,
			"width":     width,
			"height":    height,
		}
		v := mapUtil.fromMap(params)
		decoder.cDecoder = cttLibCDecoder(C.Decoder_make(v))
		mapUtil.close(v)
		return decoder
	},
	AvcodecAudio: func(frameType frameType, sampleRate uint32, channelNum uint32) *decoder {
		decoder := new(decoder)
		decoder.Type = DecoderTypes.AvcodecAudio
		params := map[string]interface{}{
			"decoder":    decoder.Type.value,
			"frameType":  frameType.value,
			"sampleRate": sampleRate,
			"channelNum": channelNum,
		}
		v := mapUtil.fromMap(params)
		decoder.cDecoder = cttLibCDecoder(C.Decoder_make(v))
		mapUtil.close(v)
		return decoder
	},
	Jpeg: func() *decoder {
		decoder := new(decoder)
		decoder.Type = DecoderTypes.Jpeg
		params := map[string]interface{}{
			"decoder": decoder.Type.value,
		}
		v := mapUtil.fromMap(params)
		decoder.cDecoder = cttLibCDecoder(C.Decoder_make(v))
		mapUtil.close(v)
		return decoder
	},
	Openh264: func() *decoder {
		decoder := new(decoder)
		decoder.Type = DecoderTypes.Openh264
		params := map[string]interface{}{
			"decoder": decoder.Type.value,
		}
		v := mapUtil.fromMap(params)
		decoder.cDecoder = cttLibCDecoder(C.Decoder_make(v))
		mapUtil.close(v)
		return decoder
	},
	Opus: func(sampleRate uint32, channelNum uint32) *opusDecoder {
		decoder := new(opusDecoder)
		decoder.Type = DecoderTypes.Opus
		params := map[string]interface{}{
			"decoder":    decoder.Type.value,
			"sampleRate": sampleRate,
			"channelNum": channelNum,
		}
		v := mapUtil.fromMap(params)
		decoder.cDecoder = cttLibCDecoder(C.Decoder_make(v))
		mapUtil.close(v)
		return decoder
	},
	Png: func() *decoder {
		decoder := new(decoder)
		decoder.Type = DecoderTypes.Png
		params := map[string]interface{}{
			"decoder": decoder.Type.value,
		}
		v := mapUtil.fromMap(params)
		decoder.cDecoder = cttLibCDecoder(C.Decoder_make(v))
		mapUtil.close(v)
		return decoder
	},
	Speex: func(sampleRate uint32, channelNum uint32) *decoder {
		decoder := new(decoder)
		decoder.Type = DecoderTypes.Speex
		params := map[string]interface{}{
			"decoder":    decoder.Type.value,
			"sampleRate": sampleRate,
			"channelNum": channelNum,
		}
		v := mapUtil.fromMap(params)
		decoder.cDecoder = cttLibCDecoder(C.Decoder_make(v))
		mapUtil.close(v)
		return decoder
	},
	Theora: func() *decoder {
		decoder := new(decoder)
		decoder.Type = DecoderTypes.Theora
		params := map[string]interface{}{
			"decoder": decoder.Type.value,
		}
		v := mapUtil.fromMap(params)
		decoder.cDecoder = cttLibCDecoder(C.Decoder_make(v))
		mapUtil.close(v)
		return decoder
	},
	Vorbis: func() *decoder {
		decoder := new(decoder)
		decoder.Type = DecoderTypes.Vorbis
		params := map[string]interface{}{
			"decoder": decoder.Type.value,
		}
		v := mapUtil.fromMap(params)
		decoder.cDecoder = cttLibCDecoder(C.Decoder_make(v))
		mapUtil.close(v)
		return decoder
	},
}

func decoderDecodeFrame(
	decoder *decoder,
	frame IFrame,
	callback func(frame *Frame) bool) bool {
	if decoder.cDecoder == nil {
		return false
	}
	if frame == nil {
		return true
	}
	cframe := frame.newGoRefFrame()
	decoder.call.callback = callback
	result := C.Decoder_decodeFrame(
		unsafe.Pointer(decoder.cDecoder),
		unsafe.Pointer(frame.refCFrame()),
		cframe,
		C.uintptr_t(uintptr(unsafe.Pointer(&decoder.call))))
	frame.deleteGoRefFrame(cframe)
	return bool(result)
}

func decoderClose(decoder *decoder) {
	C.Decoder_close(unsafe.Pointer(decoder.cDecoder))
}

// DecodeFrame デコードを実行
func (decoder *decoder) DecodeFrame(
	frame IFrame,
	callback func(frame *Frame) bool) bool {
	return decoderDecodeFrame(decoder, frame, callback)
}

// Close 閉じる
func (decoder *decoder) Close() {
	decoderClose(decoder)
}
