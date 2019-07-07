#include "speexdsp.hpp"
#include "../util.hpp"
#include "../frame.hpp"
#include <iostream>

using namespace std;

extern "C" {
extern bool ttLibGoFrameCallback(void *ptr, ttLibC_Frame *frame);
}

SpeexdspResampler::SpeexdspResampler(maps *mp) {
#ifdef __ENABLE_SPEEXDSP__
  _resampler = ttLibC_SpeexdspResampler_make(
    mp->getUint32("channelNum"),
    mp->getUint32("inSampleRate"),
    mp->getUint32("outSampleRate"),
    mp->getUint32("quality"));
#endif
  _pcm = nullptr;
}

SpeexdspResampler::~SpeexdspResampler() {
#ifdef __ENABLE_SPEEXDSP__
  ttLibC_PcmS16_close(&_pcm);
  ttLibC_SpeexdspResampler_close(&_resampler);
#endif
}
bool SpeexdspResampler::resampleFrame(ttLibC_Frame *cFrame, ttLibGoFrame *goFrame, void *ptr) {
  switch(cFrame->type) {
  case frameType_pcmS16:
    break;
  default:
    return false;
  }
  bool result = false;
#ifdef __ENABLE_SPEEXDSP__
  update(cFrame, goFrame);
  ttLibC_PcmS16 *pcm = (ttLibC_PcmS16 *)cFrame;
  ttLibC_PcmS16 *resampled = ttLibC_SpeexdspResampler_resample(
    _resampler,
    _pcm,
    pcm);
  if(resampled != nullptr) {
    _pcm = resampled;
    result = ttLibGoFrameCallback(ptr, (ttLibC_Frame *)_pcm);
  }
  reset(cFrame, goFrame);
#endif
  return result;
}
