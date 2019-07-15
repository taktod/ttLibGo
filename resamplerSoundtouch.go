package ttLibGo

/*
#include <stdint.h>
#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>
extern void SoundtouchResampler_setRate(void *resampler, float newRate);
extern void SoundtouchResampler_setTempo(void *resampler, float newTempo);
extern void SoundtouchResampler_setRateChange(void *resampler, float newRate);
extern void SoundtouchResampler_setTempoChange(void *resampler, float newTempo);
extern void SoundtouchResampler_setPitch(void *resampler, float newPitch);
extern void SoundtouchResampler_setPitchOctaves(void *resampler, float newPitch);
extern void SoundtouchResampler_setPitchSemiTones(void *resampler, float newPitch);
*/
import "C"
import (
	"unsafe"
)

type soundtouchResampler resampler

// ResampleFrame リサンプルを実行
func (soundtouchResampler *soundtouchResampler) ResampleFrame(
	frame IFrame,
	callback FrameCallback) bool {
	return resamplerResampleFrame((*resampler)(soundtouchResampler), frame, callback)
}

// Close 閉じる
func (soundtouchResampler *soundtouchResampler) Close() {
	resamplerClose((*resampler)(soundtouchResampler))
}

func (soundtouchResampler *soundtouchResampler) SetRate(newRate float32) {
	C.SoundtouchResampler_setRate(unsafe.Pointer(soundtouchResampler.cResampler), C.float(newRate))
}

func (soundtouchResampler *soundtouchResampler) SetTempo(newTempo float32) {
	C.SoundtouchResampler_setTempo(unsafe.Pointer(soundtouchResampler.cResampler), C.float(newTempo))
}

func (soundtouchResampler *soundtouchResampler) SetRateChange(newRate float32) {
	C.SoundtouchResampler_setRateChange(unsafe.Pointer(soundtouchResampler.cResampler), C.float(newRate))
}

func (soundtouchResampler *soundtouchResampler) SetTempoChange(newTempo float32) {
	C.SoundtouchResampler_setTempoChange(unsafe.Pointer(soundtouchResampler.cResampler), C.float(newTempo))
}

func (soundtouchResampler *soundtouchResampler) SetPitch(newPitch float32) {
	C.SoundtouchResampler_setPitch(unsafe.Pointer(soundtouchResampler.cResampler), C.float(newPitch))
}

func (soundtouchResampler *soundtouchResampler) SetPitchOctaves(newPitch float32) {
	C.SoundtouchResampler_setPitchOctaves(unsafe.Pointer(soundtouchResampler.cResampler), C.float(newPitch))
}

func (soundtouchResampler *soundtouchResampler) SetPitchSemiTones(newPitch float32) {
	C.SoundtouchResampler_setPitchSemiTones(unsafe.Pointer(soundtouchResampler.cResampler), C.float(newPitch))
}
