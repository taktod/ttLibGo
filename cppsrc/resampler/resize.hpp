#ifndef TTLIBGO_RESAMPLER_RESIZE_HPP
#define TTLIBGO_RESAMPLER_RESIZE_HPP

#include "../resampler.hpp"

class ResizeResampler : public Resampler {
public:
  ResizeResampler(maps *mp);
  ~ResizeResampler();
  bool resize(
    ttLibC_Frame *dstCFrame, ttLibGoFrame *dstGoFrame,
    ttLibC_Frame *srcCFrame, ttLibGoFrame *srcGoFrame,
    bool isQuick);
private:
  FrameProcessor srcfp;
};

#endif
