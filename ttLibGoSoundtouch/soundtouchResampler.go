package ttLibGoSoundtouch

/*
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <ttLibC/allocator.h>
#include <ttLibC/resampler/soundtouchResampler.h>
#include <ttLibC/frame/audio/pcms16.h>
#include <ttLibC/frame/audio/pcmf32.h>
#include "../ttLibGo/frame.h"

extern bool ttLibGoSoundtouchFrameCallback(void *ptr, ttLibC_Frame *frame);

static bool SoundtouchResampler_resampleCallback(void *ptr, ttLibC_Audio *audio) {
	return ttLibGoSoundtouchFrameCallback(ptr, (ttLibC_Frame *)audio);
}

static bool SoundtouchResampler_resample(
		ttLibC_Soundtouch *resampler,
		void *frame,
		uintptr_t ptr) {
	return ttLibC_Soundtouch_resample(resampler, (ttLibC_Audio *)frame, SoundtouchResampler_resampleCallback, (void *)ptr);
}
*/
import "C"
import (
	"unsafe"

	"github.com/taktod/ttLibGo/ttLibGo"
)

// SoundtouchResampler soundtouchを利用した音声の加工動作
type SoundtouchResampler struct {
	call       frameCall
	cResampler *C.ttLibC_Soundtouch
}

// Init 初期化
func (soundtouch *SoundtouchResampler) Init(
	sampleRate uint32,
	channelNum uint32) bool {
	if soundtouch.cResampler == nil {
		soundtouch.cResampler = C.ttLibC_Soundtouch_make(
			C.uint32_t(sampleRate),
			C.uint32_t(channelNum))
	}
	return soundtouch.cResampler != nil
}

// Resample pcmをリサンプルします
func (soundtouch *SoundtouchResampler) Resample(
	frame *ttLibGo.Frame,
	callback ttLibGo.FrameCallback) bool {
	if soundtouch.cResampler == nil {
		return false
	}
	if frame == nil {
		return true
	}
	frame.Restore()
	soundtouch.call.callback = callback
	if !C.SoundtouchResampler_resample(
		soundtouch.cResampler,
		unsafe.Pointer(frame.CFrame),
		C.uintptr_t(uintptr(unsafe.Pointer(soundtouch)))) {
		return false
	}
	return true
}

// Close 解放
func (soundtouch *SoundtouchResampler) Close() {
	C.ttLibC_Soundtouch_close(&soundtouch.cResampler)
	soundtouch.cResampler = nil
}

// SetRate 動作レートを変更します 1.0が通常の速さ
func (soundtouch *SoundtouchResampler) SetRate(value float32) bool {
	if soundtouch.cResampler == nil {
		return false
	}
	C.ttLibC_Soundtouch_setRate(soundtouch.cResampler, C.double(value))
	return true
}

// SetTempo テンポを変更します
func (soundtouch *SoundtouchResampler) SetTempo(value float32) bool {
	if soundtouch.cResampler == nil {
		return false
	}
	C.ttLibC_Soundtouch_setTempo(soundtouch.cResampler, C.double(value))
	return true
}

// SetRateChange rateChangeします
func (soundtouch *SoundtouchResampler) SetRateChange(value float32) bool {
	if soundtouch.cResampler == nil {
		return false
	}
	C.ttLibC_Soundtouch_setRateChange(soundtouch.cResampler, C.double(value))
	return true
}

// SetTempoChange tempoChangeします
func (soundtouch *SoundtouchResampler) SetTempoChange(value float32) bool {
	if soundtouch.cResampler == nil {
		return false
	}
	C.ttLibC_Soundtouch_setTempoChange(soundtouch.cResampler, C.double(value))
	return true
}

// SetPitch ピッチを変更します
func (soundtouch *SoundtouchResampler) SetPitch(value float32) bool {
	if soundtouch.cResampler == nil {
		return false
	}
	C.ttLibC_Soundtouch_setPitch(soundtouch.cResampler, C.double(value))
	return true
}

// SetPitchOctaves ピッチを変更します
func (soundtouch *SoundtouchResampler) SetPitchOctaves(value float32) bool {
	if soundtouch.cResampler == nil {
		return false
	}
	C.ttLibC_Soundtouch_setPitchOctaves(soundtouch.cResampler, C.double(value))
	return true
}

// SetPitchSemiTones ピッチを変更します
func (soundtouch *SoundtouchResampler) SetPitchSemiTones(value float32) bool {
	if soundtouch.cResampler == nil {
		return false
	}
	C.ttLibC_Soundtouch_setPitchSemiTones(soundtouch.cResampler, C.double(value))
	return true
}
