#ifndef TTLIBGO_RESAMPLER_IMAGE_HPP
#define TTLIBGO_RESAMPLER_IMAGE_HPP

#include "../resampler.hpp"

class ImageResampler : public Resampler {
public:
  ImageResampler(maps *mp);
  ~ImageResampler();
//  bool resampleFrame(ttLibC_Frame *cFrame, ttLibGoFrame *goFrame, void *ptr);
  bool toBgr(
    ttLibC_Frame *dstCFrame, ttLibGoFrame *dstGoFrame,
    ttLibC_Frame *srcCFrame, ttLibGoFrame *srcGoFrame);
  bool toYuv420(
    ttLibC_Frame *dstCFrame, ttLibGoFrame *dstGoFrame,
    ttLibC_Frame *srcCFrame, ttLibGoFrame *srcGoFrame);
private:
/*  uint32_t getSubType(ttLibC_Frame_Type frameType, string name);
  ttLibC_Frame_Type _sourceType;

  ttLibC_Frame_Type _targetType;
  uint32_t _subType;
  ttLibC_Video *_image;*/
  FrameProcessor srcfp;
};

#endif
