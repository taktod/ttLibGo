package ttLibGoSpeexdsp

/*
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <ttLibC/allocator.h>
#include <ttLibC/resampler/speexdspResampler.h>
#include "../ttLibGo/frame.h"

extern bool ttLibGoSpeexdspFrameCallback(void *ptr, ttLibC_Frame *frame);

typedef struct speexdspResampler_t {
	ttLibC_SpeexdspResampler *resampler;
	ttLibC_Frame             *frame;
} speexdspResampler_t;

speexdspResampler_t *SpeexdspResampler_make(
		uint32_t channelNum,
		uint32_t inSampleRate,
		uint32_t outSampleRate,
		uint32_t quality) {
	if(channelNum == 0 || inSampleRate == 0 || outSampleRate == 0) {
		puts("入力パラメーターが不正です。");
		return NULL;
	}
	speexdspResampler_t *resampler = ttLibC_malloc(sizeof(speexdspResampler_t));
	if(resampler == NULL) {
		puts("failed to allocate speexdspResamplerObject.");
		return NULL;
	}
	resampler->frame = NULL;
	resampler->resampler = ttLibC_SpeexdspResampler_make(channelNum, inSampleRate, outSampleRate, quality);
	if(resampler->resampler == NULL) {
		ttLibC_free(resampler);
		puts("failed to allocate ttLibC_SpeexdspResampler.");
		return NULL;
	}
	return resampler;
}

bool SpeexdspResampler_resample(
		speexdspResampler_t *resampler,
		void *frame,
		uintptr_t ptr) {
	if(resampler == NULL) {
		return false;
	}
	if(frame == NULL) {
		return true;
	}
	ttLibC_Frame *f = (ttLibC_Frame *)frame;
	if(f->type != frameType_pcmS16) {
		puts("pcms16のみ対応します");
		return false;
	}
	ttLibC_PcmS16 *pcm = (ttLibC_PcmS16 *)frame;
	if(pcm->inherit_super.channel_num != resampler->resampler->channel_num) {
		puts("チャンネル数が一致しません");
		return false;
	}
	if(pcm->inherit_super.sample_rate != resampler->resampler->input_sample_rate) {
		puts("入力データのサンプルレートが一致しません");
		return false;
	}
	ttLibC_PcmS16 *p = ttLibC_SpeexdspResampler_resample(resampler->resampler, (ttLibC_PcmS16 *)resampler->frame, pcm);
	if(p == NULL) {
		puts("リサンプル失敗");
		return false;
	}
	resampler->frame = (ttLibC_Frame *)p;
	return ttLibGoSpeexdspFrameCallback((void *)ptr, resampler->frame);
}

void SpeexdspResampler_close(speexdspResampler_t **resampler) {
	speexdspResampler_t *target = *resampler;
	if(target == NULL) {
		return;
	}
	ttLibC_Frame_close(&target->frame);
	ttLibC_SpeexdspResampler_close(&target->resampler);
	ttLibC_free(target);
	*resampler = NULL;
}

*/
import "C"
import (
	"unsafe"

	"github.com/taktod/ttLibGo/ttLibGo"
)

// SpeexdspResampler speexdspによる周波数変換
type SpeexdspResampler struct {
	call       frameCall
	cResampler *C.speexdspResampler_t
}

// Init 初期化
func (speexdsp *SpeexdspResampler) Init(
	channelNum uint32,
	inSampleRate uint32,
	outSampleRate uint32,
	quality uint32) bool {
	if speexdsp.cResampler == nil {
		speexdsp.cResampler = C.SpeexdspResampler_make(
			C.uint32_t(channelNum),
			C.uint32_t(inSampleRate),
			C.uint32_t(outSampleRate),
			C.uint32_t(quality))
	}
	return speexdsp.cResampler != nil
}

// Resample リサンプル実施
func (speexdsp *SpeexdspResampler) Resample(
	frame *ttLibGo.Frame,
	callback ttLibGo.FrameCallback) bool {
	if speexdsp.cResampler == nil {
		return false
	}
	if frame == nil {
		return true
	}
	frame.Restore()
	speexdsp.call.callback = callback
	if !C.SpeexdspResampler_resample(
		speexdsp.cResampler,
		unsafe.Pointer(frame.CFrame),
		C.uintptr_t(uintptr(unsafe.Pointer(speexdsp)))) {
		return false
	}
	return true
}

// Close 解放
func (speexdsp *SpeexdspResampler) Close() {
	C.SpeexdspResampler_close(&speexdsp.cResampler)
	speexdsp.cResampler = nil
}
