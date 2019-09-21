#include "image.hpp"
#include "../util.hpp"
#include "../frame.hpp"
#include <iostream>

using namespace std;

extern "C" {
typedef bool (* ttLibC_Resampler_convert_func)(void *, void *);
extern ttLibC_Resampler_convert_func ttLibGo_ImageResampler_ToBgr;
extern ttLibC_Resampler_convert_func ttLibGo_ImageResampler_ToYuv420;
}

ImageResampler::ImageResampler(maps *mp) {
}
ImageResampler::~ImageResampler() {
}

bool ImageResampler::toBgr(
    ttLibC_Frame *dstCFrame, ttLibGoFrame *dstGoFrame,
    ttLibC_Frame *srcCFrame, ttLibGoFrame *srcGoFrame) {
  bool result = false;
  if(ttLibGo_ImageResampler_ToBgr == nullptr) {
    return false;
  }
  else {
    update(dstCFrame, dstGoFrame);
    srcfp.update(srcCFrame, srcGoFrame);
    result = (*ttLibGo_ImageResampler_ToBgr)(dstCFrame, srcCFrame);
    srcfp.reset(srcCFrame, srcGoFrame);
    reset(dstCFrame, dstGoFrame);
  }
  return result;
}

bool ImageResampler::toYuv420(
    ttLibC_Frame *dstCFrame, ttLibGoFrame *dstGoFrame,
    ttLibC_Frame *srcCFrame, ttLibGoFrame *srcGoFrame) {
  bool result = false;
  if(ttLibGo_ImageResampler_ToYuv420 == nullptr) {
    return false;
  }
  else {
    update(dstCFrame, dstGoFrame);
    srcfp.update(srcCFrame, srcGoFrame);
    result = (*ttLibGo_ImageResampler_ToYuv420)(dstCFrame, srcCFrame);
    srcfp.reset(srcCFrame, srcGoFrame);
    reset(dstCFrame, dstGoFrame);
  }
  return result;
}

extern "C" {
bool ImageResampler_toBgr(
    void *resampler,
    void *dstCFrame, void *dstGoFrame,
    void *srcCFrame, void *srcGoFrame) {
  if(resampler == nullptr || dstCFrame == nullptr || dstGoFrame == nullptr
  || srcCFrame == nullptr || srcGoFrame == nullptr) {
    return false;
  }
  ImageResampler *r = reinterpret_cast<ImageResampler *>(resampler);
  ttLibC_Frame *dstCF = reinterpret_cast<ttLibC_Frame *>(dstCFrame);
  ttLibGoFrame *dstGoF = reinterpret_cast<ttLibGoFrame *>(dstGoFrame);
  ttLibC_Frame *srcCF = reinterpret_cast<ttLibC_Frame *>(srcCFrame);
  ttLibGoFrame *srcGoF = reinterpret_cast<ttLibGoFrame *>(srcGoFrame);
  return r->toBgr(dstCF, dstGoF, srcCF, srcGoF);
}

bool ImageResampler_toYuv420(
    void *resampler,
    void *dstCFrame, void *dstGoFrame,
    void *srcCFrame, void *srcGoFrame) {
  if(resampler == nullptr || dstCFrame == nullptr || dstGoFrame == nullptr
  || srcCFrame == nullptr || srcGoFrame == nullptr) {
    return false;
  }
  ImageResampler *r = reinterpret_cast<ImageResampler *>(resampler);
  ttLibC_Frame *dstCF = reinterpret_cast<ttLibC_Frame *>(dstCFrame);
  ttLibGoFrame *dstGoF = reinterpret_cast<ttLibGoFrame *>(dstGoFrame);
  ttLibC_Frame *srcCF = reinterpret_cast<ttLibC_Frame *>(srcCFrame);
  ttLibGoFrame *srcGoF = reinterpret_cast<ttLibGoFrame *>(srcGoFrame);
  return r->toYuv420(dstCF, dstGoF, srcCF, srcGoF);
}

}