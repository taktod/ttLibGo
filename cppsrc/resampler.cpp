#include <stdio.h>
#include <string>

#include "resampler.hpp"
#include "resampler/audio.hpp"
#include "resampler/image.hpp"
#include "resampler/resize.hpp"
#include "resampler/soundtouch.hpp"
#include "resampler/speexdsp.hpp"
#include "resampler/swresample.hpp"
#include "resampler/swscale.hpp"
#include <iostream>

using namespace std;

Resampler *Resampler::create(maps *mp) {
  string resamplerType = mp->getString("resampler");
  if(resamplerType == "audio") {
    return new AudioResampler(mp);
  }
  else if(resamplerType == "image") {
    return new ImageResampler(mp);
  }
  else if(resamplerType == "resize") {
    return new ResizeResampler(mp);
  }
  else if(resamplerType == "soundtouch") {
    return new SoundtouchResampler(mp);
  }
  else if(resamplerType == "speexdsp") {
    return new SpeexdspResampler(mp);
  }
  else if(resamplerType == "swresample") {
    return new SwresampleResampler(mp);
  }
  else if(resamplerType == "swscale") {
    return new SwscaleResampler(mp);
  }
  return nullptr;
}

Resampler::Resampler() {
}

Resampler::~Resampler() {
}

extern "C" {

void *Resampler_make(void *mp) {
  maps *m = reinterpret_cast<maps *>(mp);
  return Resampler::create(m);
}
bool Resampler_resampleFrame(void *resampler, void *frame, void *goFrame, uintptr_t ptr) {
  if(resampler == nullptr) {
    return false;
  }
  Resampler *r = reinterpret_cast<Resampler *>(resampler);
  ttLibC_Frame *cF = reinterpret_cast<ttLibC_Frame *>(frame);
  ttLibGoFrame *goF = reinterpret_cast<ttLibGoFrame *>(goFrame);
  return r->resampleFrame(cF, goF, (void *)ptr);
}

void Resampler_close(void *resampler) {
  if(resampler == nullptr) {
    return;
  }
  Resampler *r = reinterpret_cast<Resampler *>(resampler);
  delete r;
}

}
