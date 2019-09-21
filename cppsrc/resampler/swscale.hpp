#ifndef TTLIBGO_RESAMPLER_SWSCALE_HPP
#define TTLIBGO_RESAMPLER_SWSCALE_HPP

#include "../resampler.hpp"

class SwscaleResampler : public Resampler {
public:
  SwscaleResampler(maps *mp);
  ~SwscaleResampler();
  bool resample(
    ttLibC_Frame *dstCFrame, ttLibGoFrame *dstGoFrame,
    ttLibC_Frame *srcCFrame, ttLibGoFrame *srcGoFrame);
private:
  uint32_t getSubType(ttLibC_Frame_Type type, string name);
  void *_resampler;
  FrameProcessor srcfp;
};

#endif
