#ifndef TTLIBGO_RESAMPLER_SPEEXDSP_HPP
#define TTLIBGO_RESAMPLER_SPEEXDSP_HPP

#include "../resampler.hpp"
#include <ttLibC/resampler/speexdspResampler.h>

class SpeexdspResampler : public Resampler {
public:
  SpeexdspResampler(maps *mp);
  ~SpeexdspResampler();
  bool resampleFrame(ttLibC_Frame *cFrame, ttLibGoFrame *goFrame, void *ptr);
private:
  ttLibC_PcmS16 *_pcm;
  ttLibC_SpeexdspResampler *_resampler;
};

#endif
