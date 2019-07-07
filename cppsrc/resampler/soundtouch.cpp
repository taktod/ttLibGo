#include "soundtouch.hpp"
#include "../util.hpp"
#include <iostream>

using namespace std;

extern "C" {
extern bool ttLibGoFrameCallback(void *ptr, ttLibC_Frame *frame);
}

SoundtouchResampler::SoundtouchResampler(maps *mp) {
#ifdef __ENABLE_SOUNDTOUCH__
  _resampler = ttLibC_Soundtouch_make(mp->getUint32("sampleRate"), mp->getUint32("channelNum"));
#endif
}

SoundtouchResampler::~SoundtouchResampler() {
#ifdef __ENABLE_SOUNDTOUCH__
  ttLibC_Soundtouch_close(&_resampler);
#endif
}

bool SoundtouchResampler::resampleFrame(ttLibC_Frame *cFrame, ttLibGoFrame *goFrame, void *ptr) {
  bool result = false;
#ifdef __ENABLE_SOUNDTOUCH__
  update(cFrame, goFrame);
  result = ttLibC_Soundtouch_resample(
    _resampler,
    (ttLibC_Audio *)cFrame,
    [](void *ptr, ttLibC_Audio *audio) -> bool {
      return ttLibGoFrameCallback(ptr, (ttLibC_Frame *)audio);
    },
    ptr);
  reset(cFrame, goFrame);
#endif
  return result;
}

void SoundtouchResampler::setRate(double newRate) {
#ifdef __ENABLE_SOUNDTOUCH__
  ttLibC_Soundtouch_setRate(_resampler, newRate);
#endif
}
void SoundtouchResampler::setTempo(double newTempo) {
#ifdef __ENABLE_SOUNDTOUCH__
  ttLibC_Soundtouch_setTempo(_resampler, newTempo);
#endif
}
void SoundtouchResampler::setRateChange(double newRate) {
#ifdef __ENABLE_SOUNDTOUCH__
  ttLibC_Soundtouch_setRateChange(_resampler, newRate);
#endif
}
void SoundtouchResampler::setTempoChange(double newTempo) {
#ifdef __ENABLE_SOUNDTOUCH__
  ttLibC_Soundtouch_setTempoChange(_resampler, newTempo);
#endif
}
void SoundtouchResampler::setPitch(double newPitch) {
#ifdef __ENABLE_SOUNDTOUCH__
  ttLibC_Soundtouch_setPitch(_resampler, newPitch);
#endif
}
void SoundtouchResampler::setPitchOctaves(double newPitch) {
#ifdef __ENABLE_SOUNDTOUCH__
  ttLibC_Soundtouch_setPitchOctaves(_resampler, newPitch);
#endif
}
void SoundtouchResampler::setPitchSemiTones(double newPitch) {
#ifdef __ENABLE_SOUNDTOUCH__
  ttLibC_Soundtouch_setPitchSemiTones(_resampler, newPitch);
#endif
}

extern "C" {
void SoundtouchResampler_setRate(void *resampler, float newRate) {
  SoundtouchResampler *s = reinterpret_cast<SoundtouchResampler *>(resampler);
  s->setRate((double)newRate);
}

void SoundtouchResampler_setTempo(void *resampler, float newTempo) {
  SoundtouchResampler *s = reinterpret_cast<SoundtouchResampler *>(resampler);
  s->setTempo((double)newTempo);
}

void SoundtouchResampler_setRateChange(void *resampler, float newRate) {
  SoundtouchResampler *s = reinterpret_cast<SoundtouchResampler *>(resampler);
  s->setRateChange((double)newRate);
}

void SoundtouchResampler_setTempoChange(void *resampler, float newTempo) {
  SoundtouchResampler *s = reinterpret_cast<SoundtouchResampler *>(resampler);
  s->setTempoChange((double)newTempo);
}

void SoundtouchResampler_setPitch(void *resampler, float newPitch) {
  SoundtouchResampler *s = reinterpret_cast<SoundtouchResampler *>(resampler);
  s->setPitch((double)newPitch);
}

void SoundtouchResampler_setPitchOctaves(void *resampler, float newPitch) {
  SoundtouchResampler *s = reinterpret_cast<SoundtouchResampler *>(resampler);
  s->setPitchOctaves((double)newPitch);
}

void SoundtouchResampler_setPitchSemiTones(void *resampler, float newPitch) {
  SoundtouchResampler *s = reinterpret_cast<SoundtouchResampler *>(resampler);
  s->setPitchSemiTones((double)newPitch);
}

}