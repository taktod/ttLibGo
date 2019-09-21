#ifndef TTLIBGO_RESAMPLER_LIBYUV_HPP
#define TTLIBGO_RESAMPLER_LIBYUV_HPP

#include "../resampler.hpp"
#include "ttLibC/ttLibC/resampler/libyuvResampler.h"

class LibyuvResampler : public Resampler {
public:
  LibyuvResampler(maps *mp);
  ~LibyuvResampler();
  bool resize(
    ttLibC_Frame *dstCFrame, ttLibGoFrame *dstGoFrame,
    ttLibC_Frame *srcCFrame, ttLibGoFrame *srcGoFrame,
    string mode,
    string subMode);
  bool rotate(
    ttLibC_Frame *dstCFrame, ttLibGoFrame *dstGoFrame,
    ttLibC_Frame *srcCFrame, ttLibGoFrame *srcGoFrame,
    string mode);
  bool toBgr(
    ttLibC_Frame *dstCFrame, ttLibGoFrame *dstGoFrame,
    ttLibC_Frame *srcCFrame, ttLibGoFrame *srcGoFrame);
  bool toYuv420(
    ttLibC_Frame *dstCFrame, ttLibGoFrame *dstGoFrame,
    ttLibC_Frame *srcCFrame, ttLibGoFrame *srcGoFrame);
private:
  ttLibC_LibyuvFilter_Mode checkMode(string str);
  ttLibC_LibyuvRotate_Mode checkRotate(string str);
  FrameProcessor srcfp;
};

#endif
