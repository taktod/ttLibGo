#include "resize.hpp"
#include "../util.hpp"
#include "../frame.hpp"
#include <iostream>

using namespace std;

extern "C" {
typedef void *(* ttLibC_ImageResizer_resize_func)(void *, void *, bool);

extern ttLibC_ImageResizer_resize_func ttLibGo_ImageResizer_resize;
}

ResizeResampler::ResizeResampler(maps *mp) {
}

ResizeResampler::~ResizeResampler() {
}

bool ResizeResampler::resize(
    ttLibC_Frame *dstCFrame, ttLibGoFrame *dstGoFrame,
    ttLibC_Frame *srcCFrame, ttLibGoFrame *srcGoFrame,
    bool isQuick) {
  bool result = false;
  if(ttLibGo_ImageResizer_resize == nullptr) {
    return false;
  }
  else {
    update(dstCFrame, dstGoFrame);
    srcfp.update(srcCFrame, srcGoFrame);
    result = (*ttLibGo_ImageResizer_resize)(dstCFrame, srcCFrame, isQuick);
    srcfp.reset(srcCFrame, srcGoFrame);
    reset(dstCFrame, dstGoFrame);
  }
  return result;
}

extern "C" {
bool Resizer_resize(
    void *resampler,
    void *dstCFrame, void *dstGoFrame,
    void *srcCFrame, void *srcGoFrame,
    bool is_quick) {
  if(resampler == nullptr || dstCFrame == nullptr || dstGoFrame == nullptr
  || srcCFrame == nullptr || srcGoFrame == nullptr) {
    return false;
  }
  ResizeResampler *r = reinterpret_cast<ResizeResampler *>(resampler);
  ttLibC_Frame *dstCF = reinterpret_cast<ttLibC_Frame *>(dstCFrame);
  ttLibGoFrame *dstGoF = reinterpret_cast<ttLibGoFrame *>(dstGoFrame);
  ttLibC_Frame *srcCF = reinterpret_cast<ttLibC_Frame *>(srcCFrame);
  ttLibGoFrame *srcGoF = reinterpret_cast<ttLibGoFrame *>(srcGoFrame);
  return r->resize(dstCF, dstGoF, srcCF, srcGoF, is_quick);
}
}