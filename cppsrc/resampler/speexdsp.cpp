#include "speexdsp.hpp"
#include "../util.hpp"
#include "../frame.hpp"
#include <iostream>
#include <ttLibC/frame/audio/pcms16.h>

using namespace std;

extern "C" {
typedef void *(* ttLibC_SpeexdspResampler_make_func)(uint32_t,uint32_t,uint32_t,uint32_t);
typedef void *(* ttLibC_SpeexdspResampler_resample_func)(void *, void *, void *);
typedef void (* ttLibC_close_func)(void **);

extern ttLibC_close_func ttLibGo_Frame_close;
extern ttLibC_SpeexdspResampler_make_func     ttLibGo_SpeexdspResampler_make;
extern ttLibC_close_func                      ttLibGo_SpeexdspResampler_close;
extern ttLibC_SpeexdspResampler_resample_func ttLibGo_SpeexdspResampler_resample;

extern bool ttLibGoFrameCallback(void *ptr, ttLibC_Frame *frame);
}

SpeexdspResampler::SpeexdspResampler(maps *mp) {
  if(ttLibGo_SpeexdspResampler_make != nullptr) {
    _resampler = (*ttLibGo_SpeexdspResampler_make)(
    mp->getUint32("channelNum"),
    mp->getUint32("inSampleRate"),
    mp->getUint32("outSampleRate"),
    mp->getUint32("quality"));
  }
  _pcm = nullptr;
}

SpeexdspResampler::~SpeexdspResampler() {
  if(ttLibGo_Frame_close != nullptr) {
    (*ttLibGo_Frame_close)((void **)&_pcm);
  }
  if(ttLibGo_SpeexdspResampler_close != nullptr) {
    (*ttLibGo_SpeexdspResampler_close)(&_resampler);
  }
}
bool SpeexdspResampler::resampleFrame(ttLibC_Frame *cFrame, ttLibGoFrame *goFrame, void *ptr) {
  switch(cFrame->type) {
  case frameType_pcmS16:
    break;
  default:
    return false;
  }
  bool result = false;
  if(ttLibGo_SpeexdspResampler_resample != nullptr) {
    update(cFrame, goFrame);
    ttLibC_PcmS16 *pcm = (ttLibC_PcmS16 *)cFrame;
    ttLibC_PcmS16 *resampled = (ttLibC_PcmS16 *)(*ttLibGo_SpeexdspResampler_resample)(
      _resampler,
      _pcm,
      pcm);
    if(resampled != nullptr) {
      _pcm = resampled;
      result = ttLibGoFrameCallback(ptr, (ttLibC_Frame *)_pcm);
    }
    reset(cFrame, goFrame);
  }
  return result;
}
