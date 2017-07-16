package ttLibGo

import (
	"unsafe"
)

/*
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <ttLibC/allocator.h>
#include <ttLibC/resampler/audioResampler.h>
#include <ttLibC/frame/audio/pcms16.h>
#include <ttLibC/frame/audio/pcmf32.h>
#include "frame.h"
extern bool ttLibGoFrameCallback(void *ptr, ttLibC_Frame *frame);

typedef struct audioResampler_t {
	ttLibC_Frame_Type type;
	uint32_t sub_type;
	uint32_t channel_num;
	ttLibC_Frame *frame;
} audioResampler_t;

audioResampler_t *AudioResampler_make(
		const char *codecType,
		const char *subType,
		uint32_t channel_num) {
	ttLibC_Frame_Type type = Frame_getFrameType(codecType);
	switch(type) {
	case frameType_pcmS16:
	case frameType_pcmF32:
		break;
	default:
		return NULL;
	}
	audioResampler_t *resampler = ttLibC_malloc(sizeof(audioResampler_t));
	if(resampler == NULL) {
		puts("audioResamplerオブジェクト作成できませんでした");
		return NULL;
	}
	switch(type) {
	case frameType_pcmS16:
		{
			resampler->type = frameType_pcmS16;
			if(strcmp(subType, "bigEndian") == 0) {
				resampler->sub_type = PcmS16Type_bigEndian;
			}
			else if(strcmp(subType, "littleEndian") == 0) {
				resampler->sub_type = PcmS16Type_littleEndian;
			}
			else if(strcmp(subType, "bigEndianPlanar") == 0) {
				resampler->sub_type = PcmS16Type_bigEndian_planar;
			}
			else if(strcmp(subType, "littleEndianPlanar") == 0) {
				resampler->sub_type = PcmS16Type_littleEndian_planar;
			}
			else {
				puts("subTypeが不正です。");
				ttLibC_free(resampler);
				return NULL;
			}
		}
		break;
	case frameType_pcmF32:
		{
			resampler->type = frameType_pcmF32;
			if(strcmp(subType, "planar") == 0) {
				resampler->sub_type = PcmF32Type_planar;
			}
			else if(strcmp(subType, "interleave") == 0) {
				resampler->sub_type = PcmF32Type_interleave;
			}
			else {
				puts("subTypeが不正です。");
				ttLibC_free(resampler);
				return NULL;
			}
		}
		break;
	default:
		ttLibC_free(resampler);
		return NULL;
	}
	resampler->frame = NULL;
	resampler->channel_num = channel_num;
	return resampler;
}

bool AudioResampler_resample(
		audioResampler_t *resampler,
		void *frame,
		uintptr_t ptr) {
	if(resampler == NULL) {
		return false;
	}
	if(frame == NULL) {
		return true;
	}
	ttLibC_Frame *f = (ttLibC_Frame *)frame;
	// 入力したフレームと、出力データが一致する場合はそのまま打ち返す
	switch(f->type) {
	case frameType_pcmS16:
		{
			ttLibC_PcmS16 *pcm = (ttLibC_PcmS16 *)frame;
			if(resampler->type == f->type && resampler->sub_type == pcm->type) {
				// タイプが一致する場合はそのまま打ち返す
				return ttLibGoFrameCallback((void *)ptr, frame);
			}
		}
		break;
	case frameType_pcmF32:
		{
			ttLibC_PcmF32 *pcm = (ttLibC_PcmF32 *)frame;
			if(resampler->type == f->type && resampler->sub_type == pcm->type) {
				// タイプが一致する場合はそのまま打ち返す
				return ttLibGoFrameCallback((void *)ptr, frame);
			}
		}
		break;
	default:
		return false;
	}
	ttLibC_Audio *audio = (ttLibC_Audio *)frame;
	uint32_t channel_num = resampler->channel_num;
	if(channel_num == 0) {
		channel_num = audio->channel_num;
	}
	ttLibC_Audio *resampled = ttLibC_AudioResampler_convertFormat(
		(ttLibC_Audio *)resampler->frame,
		resampler->type,
		resampler->sub_type,
		channel_num,
		audio);
	if(resampled == NULL) {
		puts("リサンプル失敗しました。");
		return false;
	}
	resampler->frame = (ttLibC_Frame *)resampled;
	return ttLibGoFrameCallback((void *)ptr, resampler->frame);
}

void AudioResampler_close(audioResampler_t **resampler) {
	audioResampler_t *target = *resampler;
	if(target == NULL) {
		return;
	}
	ttLibC_Frame_close(&target->frame);
	ttLibC_free(target);
	*resampler = NULL;
}
*/
import "C"

// AudioResampler 音声をpcmS16とpcmF32で相互変換します
type AudioResampler struct {
	call       frameCall
	cResampler *C.audioResampler_t
}

// Init 変換後のデータタイプを指定します
func (audio *AudioResampler) Init(
	codecType string,
	subType string,
	channelNum uint32) bool {
	if audio.cResampler == nil {
		cCodecType := C.CString(codecType)
		defer C.free(unsafe.Pointer(cCodecType))
		cSubType := C.CString(subType)
		defer C.free(unsafe.Pointer(cSubType))
		audio.cResampler = C.AudioResampler_make(
			cCodecType,
			cSubType,
			C.uint32_t(channelNum))
	}
	return audio.cResampler != nil
}

// Resample 変換を実施します
func (audio *AudioResampler) Resample(
	frame *Frame,
	callback FrameCallback) bool {
	if audio.cResampler == nil {
		return false
	}
	if frame == nil {
		return true
	}
	frame.Restore()
	audio.call.callback = callback
	if !C.AudioResampler_resample(
		audio.cResampler,
		unsafe.Pointer(frame.CFrame),
		C.uintptr_t(uintptr(unsafe.Pointer(audio)))) {
		return false
	}
	return true
}

// Close 必要なくなった場合に破棄します
func (audio *AudioResampler) Close() {
	C.AudioResampler_close(&audio.cResampler)
	audio.cResampler = nil
}
