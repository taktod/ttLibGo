#ifndef TTLIBGO_RESAMPLER_SWSCALE_HPP
#define TTLIBGO_RESAMPLER_SWSCALE_HPP

#include "../resampler.hpp"
#include <ttLibC/resampler/swscaleResampler.h>

class SwscaleResampler : public Resampler {
public:
  SwscaleResampler(maps *mp);
  ~SwscaleResampler();
  bool resampleFrame(ttLibC_Frame *cFrame, ttLibGoFrame *goFrame, void *ptr);
private:
  uint32_t getSubType(ttLibC_Frame_Type type, string name);
  ttLibC_SwscaleResampler *_resampler;
};

#endif
