#include "soundtouch.hpp"
#include "../util.hpp"
#include <iostream>

using namespace std;

extern "C" {
typedef void *(* ttLibC_Soundtouch_make_func)(uint32_t, uint32_t);
typedef void *(* ttLibC_codec_func)(void *, void *, ttLibC_getFrameFunc, void *);
typedef void (* ttLibC_soundtouch_set_func)(void *, double);
typedef void (* ttLibC_close_func)(void **);

extern ttLibC_Soundtouch_make_func ttLibGo_Soundtouch_make;
extern ttLibC_codec_func           ttLibGo_Soundtouch_resample;
extern ttLibC_close_func           ttLibGo_Soundtouch_close;
extern ttLibC_soundtouch_set_func  ttLibGo_Soundtouch_setRate;
extern ttLibC_soundtouch_set_func  ttLibGo_Soundtouch_setTempo;
extern ttLibC_soundtouch_set_func  ttLibGo_Soundtouch_setRateChange;
extern ttLibC_soundtouch_set_func  ttLibGo_Soundtouch_setTempoChange;
extern ttLibC_soundtouch_set_func  ttLibGo_Soundtouch_setPitch;
extern ttLibC_soundtouch_set_func  ttLibGo_Soundtouch_setPitchOctaves;
extern ttLibC_soundtouch_set_func  ttLibGo_Soundtouch_setPitchSemiTones;

extern bool ttLibGoFrameCallback(void *ptr, ttLibC_Frame *frame);
}

SoundtouchResampler::SoundtouchResampler(maps *mp) {
  if(ttLibGo_Soundtouch_make != nullptr) {
    _resampler = (* ttLibGo_Soundtouch_make)(mp->getUint32("sampleRate"), mp->getUint32("channelNum"));
  }
}

SoundtouchResampler::~SoundtouchResampler() {
  if(ttLibGo_Soundtouch_close != nullptr) {
    (*ttLibGo_Soundtouch_close)(&_resampler);
  }
}

bool SoundtouchResampler::resampleFrame(ttLibC_Frame *cFrame, ttLibGoFrame *goFrame, void *ptr) {
  bool result = false;
  if(ttLibGo_Soundtouch_resample != nullptr) {
    update(cFrame, goFrame);
    result = (*ttLibGo_Soundtouch_resample)(_resampler, cFrame, ttLibGoFrameCallback, ptr);
    reset(cFrame, goFrame);
  }
  return result;
}

void SoundtouchResampler::setRate(double newRate) {
  if(ttLibGo_Soundtouch_setRate != nullptr) {
    (*ttLibGo_Soundtouch_setRate)(_resampler, newRate);
  }
}
void SoundtouchResampler::setTempo(double newTempo) {
  if(ttLibGo_Soundtouch_setTempo != nullptr) {
    (*ttLibGo_Soundtouch_setTempo)(_resampler, newTempo);
  }
}
void SoundtouchResampler::setRateChange(double newRate) {
  if(ttLibGo_Soundtouch_setRateChange != nullptr) {
    (*ttLibGo_Soundtouch_setRateChange)(_resampler, newRate);
  }
}
void SoundtouchResampler::setTempoChange(double newTempo) {
  if(ttLibGo_Soundtouch_setTempoChange != nullptr) {
    (*ttLibGo_Soundtouch_setTempoChange)(_resampler, newTempo);
  }
}
void SoundtouchResampler::setPitch(double newPitch) {
  if(ttLibGo_Soundtouch_setPitch != nullptr) {
    (*ttLibGo_Soundtouch_setPitch)(_resampler, newPitch);
  }
}
void SoundtouchResampler::setPitchOctaves(double newPitch) {
  if(ttLibGo_Soundtouch_setPitchOctaves != nullptr) {
    (*ttLibGo_Soundtouch_setPitchOctaves)(_resampler, newPitch);
  }
}
void SoundtouchResampler::setPitchSemiTones(double newPitch) {
  if(ttLibGo_Soundtouch_setPitchSemiTones != nullptr) {
    (*ttLibGo_Soundtouch_setPitchSemiTones)(_resampler, newPitch);
  }
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