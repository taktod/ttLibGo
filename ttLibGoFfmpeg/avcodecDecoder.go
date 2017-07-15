package ttLibGoFfmpeg

/*
#include <ttLibC/decoder/avcodecDecoder.h>
#include "../ttLibGo/frame.h"
extern bool ttLibGoFfmpegFrameCallback(void *ptr, ttLibC_Frame *frame);

static ttLibC_AvcodecDecoder *AvcodecDecoder_make(
		const char *codec_type,
		uint32_t width,
		uint32_t height,
		uint32_t sample_rate,
		uint32_t channel_num) {
	ttLibC_Frame_Type type = Frame_getFrameType(codec_type);
	switch(type) {
	case frameType_aac:
	case frameType_adpcm_ima_wav:
	case frameType_mp3:
	case frameType_nellymoser:
	case frameType_opus:
	case frameType_pcm_alaw:
	case frameType_pcm_mulaw:
	case frameType_speex:
	case frameType_vorbis:
		return ttLibC_AvcodecAudioDecoder_make(type, sample_rate, channel_num);
	case frameType_flv1:
	case frameType_h264:
	case frameType_h265:
	case frameType_jpeg:
	case frameType_theora:
	case frameType_vp6:
	case frameType_vp8:
	case frameType_vp9:
	case frameType_wmv1:
	case frameType_wmv2:
		return ttLibC_AvcodecVideoDecoder_make(type, width, height);
	case frameType_pcmF32:
	case frameType_pcmS16:
	case frameType_bgr:
	case frameType_yuv420:
		puts("rawデータはdecodeできません");
		return NULL;
	default:
		puts("不明なコーデックでした");
		return NULL;
	}
}

static bool AvcodecDecoder_decode(ttLibC_AvcodecDecoder *decoder, void *frame, uintptr_t ptr) {
	return ttLibC_AvcodecDecoder_decode(decoder, (ttLibC_Frame *)frame, ttLibGoFfmpegFrameCallback, (void *)ptr);
}
*/
import "C"
import (
	"unsafe"

	"github.com/taktod/ttLibGo/ttLibGo"
)

// AvcodecDecoder ffmpegのcodecであるlibavcodecを利用して動作するデコーダー
type AvcodecDecoder struct {
	call     frameCall
	cDecoder *C.ttLibC_AvcodecDecoder
}

// InitVideo 映像のデコーダーとして初期化する
func (avcodec *AvcodecDecoder) InitVideo(codecType string, width uint32, height uint32) bool {
	if avcodec.cDecoder == nil {
		cCodecType := C.CString(codecType)
		defer C.free(unsafe.Pointer(cCodecType))
		avcodec.cDecoder = C.AvcodecDecoder_make(
			cCodecType,
			C.uint32_t(width),
			C.uint32_t(height),
			0,
			0)
	}
	return avcodec.cDecoder != nil
}

// InitAudio 音声のdecoderとして初期化します
func (avcodec *AvcodecDecoder) InitAudio(codecType string, sampleRate uint32, channelNum uint32) bool {
	if avcodec.cDecoder == nil {
		cCodecType := C.CString(codecType)
		defer C.free(unsafe.Pointer(cCodecType))
		avcodec.cDecoder = C.AvcodecDecoder_make(
			cCodecType,
			0,
			0,
			C.uint32_t(sampleRate),
			C.uint32_t(channelNum))
	}
	return avcodec.cDecoder != nil
}

// Decode 実際にフレームをデコードします
func (avcodec *AvcodecDecoder) Decode(
	frame *ttLibGo.Frame,
	callback ttLibGo.FrameCallback) bool {
	if avcodec.cDecoder == nil {
		return false
	}
	if frame == nil {
		return true
	}
	frame.Restore()
	avcodec.call.callback = callback
	if !C.AvcodecDecoder_decode(
		avcodec.cDecoder,
		unsafe.Pointer(frame.CFrame), // cFrameが見えてないのが、こまるのか・・・
		C.uintptr_t(uintptr(unsafe.Pointer(avcodec)))) {
		return false
	}
	return true
}

// Close デコーダーをとじます
func (avcodec *AvcodecDecoder) Close() {
	C.ttLibC_AvcodecDecoder_close(&avcodec.cDecoder)
	avcodec.cDecoder = nil
}
