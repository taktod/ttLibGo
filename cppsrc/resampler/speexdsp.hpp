#ifndef TTLIBGO_RESAMPLER_SPEEXDSP_HPP
#define TTLIBGO_RESAMPLER_SPEEXDSP_HPP

#include "../resampler.hpp"

class SpeexdspResampler : public Resampler {
public:
  SpeexdspResampler(maps *mp);
  ~SpeexdspResampler();
  bool resampleFrame(ttLibC_Frame *cFrame, ttLibGoFrame *goFrame, void *ptr);
private:
  ttLibC_PcmS16 *_pcm;
  void *_resampler;
};

#endif
