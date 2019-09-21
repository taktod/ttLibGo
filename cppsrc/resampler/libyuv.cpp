#include "libyuv.hpp"

#include "../util.hpp"
#include "../frame.hpp"
#include <iostream>

using namespace std;

extern "C" {
typedef bool (* ttLibC_LibyuvResampler_resize_func)(void *, void *, ttLibC_LibyuvFilter_Mode, ttLibC_LibyuvFilter_Mode);
typedef bool (* ttLibC_LibyuvResampler_rotate_func)(void *, void *, ttLibC_LibyuvRotate_Mode);
typedef bool (* ttLibC_Resampler_convert_func)(void *, void *);

extern ttLibC_LibyuvResampler_resize_func ttLibGo_LibyuvResampler_resize;
extern ttLibC_LibyuvResampler_rotate_func ttLibGo_LibyuvResampler_rotate;
extern ttLibC_Resampler_convert_func      ttLibGo_LibyuvResampler_ToBgr;
extern ttLibC_Resampler_convert_func      ttLibGo_LibyuvResampler_ToYuv420;
}

LibyuvResampler::LibyuvResampler(maps *mp) {
}

LibyuvResampler::~LibyuvResampler() {
}

ttLibC_LibyuvFilter_Mode LibyuvResampler::checkMode(string str) {
  if(str == "Linear") {
    return LibyuvFilter_Linear;
  }
  else if(str == "Bilinear") {
    return LibyuvFilter_Bilinear;
  }
  else if(str == "Box") {
    return LibyuvFilter_Box;
  }
  else {
    return LibyuvFilter_None;
  }
}
ttLibC_LibyuvRotate_Mode LibyuvResampler::checkRotate(string str) {
  if(str == "90") {
    return LibyuvRotate_90;
  }
  else if(str == "180") {
    return LibyuvRotate_180;
  }
  else if(str == "270") {
    return LibyuvRotate_270;
  }
  else {
    return LibyuvRotate_0;
  }
}


bool LibyuvResampler::resize(
    ttLibC_Frame *dstCFrame, ttLibGoFrame *dstGoFrame,
    ttLibC_Frame *srcCFrame, ttLibGoFrame *srcGoFrame,
    string strMode,
    string strSubMode) {
  bool result = false;
  if(ttLibGo_LibyuvResampler_resize == nullptr) {
    return false;
  }
  else {
    ttLibC_LibyuvFilter_Mode mode = checkMode(strMode);
    ttLibC_LibyuvFilter_Mode subMode = checkMode(strSubMode);
    update(dstCFrame, dstGoFrame);
    srcfp.update(srcCFrame, srcGoFrame);
    result = (*ttLibGo_LibyuvResampler_resize)(dstCFrame, srcCFrame, mode, subMode);
    srcfp.reset(srcCFrame, srcGoFrame);
    reset(dstCFrame, dstGoFrame);
  }
  return result;
}

bool LibyuvResampler::rotate(
    ttLibC_Frame *dstCFrame, ttLibGoFrame *dstGoFrame,
    ttLibC_Frame *srcCFrame, ttLibGoFrame *srcGoFrame,
    string strMode) {
  bool result = false;
  if(ttLibGo_LibyuvResampler_rotate == nullptr) {
    return false;
  }
  else {
    ttLibC_LibyuvRotate_Mode mode = checkRotate(strMode);
    update(dstCFrame, dstGoFrame);
    srcfp.update(srcCFrame, srcGoFrame);
    result = (*ttLibGo_LibyuvResampler_rotate)(dstCFrame, srcCFrame, mode);
    srcfp.reset(srcCFrame, srcGoFrame);
    reset(dstCFrame, dstGoFrame);
  }
  return result;
}

bool LibyuvResampler::toBgr(
    ttLibC_Frame *dstCFrame, ttLibGoFrame *dstGoFrame,
    ttLibC_Frame *srcCFrame, ttLibGoFrame *srcGoFrame) {
  bool result = false;
  if(ttLibGo_LibyuvResampler_ToBgr == nullptr) {
    return false;
  }
  else {
    update(dstCFrame, dstGoFrame);
    srcfp.update(srcCFrame, srcGoFrame);
    result = (*ttLibGo_LibyuvResampler_ToBgr)(dstCFrame, srcCFrame);
    srcfp.reset(srcCFrame, srcGoFrame);
    reset(dstCFrame, dstGoFrame);
  }
  return result;
}

bool LibyuvResampler::toYuv420(
    ttLibC_Frame *dstCFrame, ttLibGoFrame *dstGoFrame,
    ttLibC_Frame *srcCFrame, ttLibGoFrame *srcGoFrame) {
  bool result = false;
  if(ttLibGo_LibyuvResampler_ToYuv420 == nullptr) {
    return false;
  }
  else {
    update(dstCFrame, dstGoFrame);
    srcfp.update(srcCFrame, srcGoFrame);
    result = (*ttLibGo_LibyuvResampler_ToYuv420)(dstCFrame, srcCFrame);
    srcfp.reset(srcCFrame, srcGoFrame);
    reset(dstCFrame, dstGoFrame);
  }
  return result;
}

extern "C" {
bool LibyuvResampler_resize(
    void *resampler,
    void *dstCFrame, void *dstGoFrame,
    void *srcCFrame, void *srcGoFrame,
    const char *mode,
    const char *sub_mode) {
  if(resampler == nullptr || dstCFrame == nullptr || dstGoFrame == nullptr
  || srcCFrame == nullptr || srcGoFrame == nullptr) {
    return false;
  }
  LibyuvResampler *r = reinterpret_cast<LibyuvResampler *>(resampler);
  ttLibC_Frame *dstCF = reinterpret_cast<ttLibC_Frame *>(dstCFrame);
  ttLibGoFrame *dstGoF = reinterpret_cast<ttLibGoFrame *>(dstGoFrame);
  ttLibC_Frame *srcCF = reinterpret_cast<ttLibC_Frame *>(srcCFrame);
  ttLibGoFrame *srcGoF = reinterpret_cast<ttLibGoFrame *>(srcGoFrame);
  return r->resize(dstCF, dstGoF, srcCF, srcGoF, mode, sub_mode);
}
bool LibyuvResampler_rotate(
    void *resampler,
    void *dstCFrame, void *dstGoFrame,
    void *srcCFrame, void *srcGoFrame,
    const char *mode) {
  if(resampler == nullptr || dstCFrame == nullptr || dstGoFrame == nullptr
  || srcCFrame == nullptr || srcGoFrame == nullptr) {
    return false;
  }
  LibyuvResampler *r = reinterpret_cast<LibyuvResampler *>(resampler);
  ttLibC_Frame *dstCF = reinterpret_cast<ttLibC_Frame *>(dstCFrame);
  ttLibGoFrame *dstGoF = reinterpret_cast<ttLibGoFrame *>(dstGoFrame);
  ttLibC_Frame *srcCF = reinterpret_cast<ttLibC_Frame *>(srcCFrame);
  ttLibGoFrame *srcGoF = reinterpret_cast<ttLibGoFrame *>(srcGoFrame);
  return r->rotate(dstCF, dstGoF, srcCF, srcGoF, mode);
}

bool LibyuvResampler_toBgr(
    void *resampler,
    void *dstCFrame, void *dstGoFrame,
    void *srcCFrame, void *srcGoFrame) {
  if(resampler == nullptr || dstCFrame == nullptr || dstGoFrame == nullptr
  || srcCFrame == nullptr || srcGoFrame == nullptr) {
    return false;
  }
  LibyuvResampler *r = reinterpret_cast<LibyuvResampler *>(resampler);
  ttLibC_Frame *dstCF = reinterpret_cast<ttLibC_Frame *>(dstCFrame);
  ttLibGoFrame *dstGoF = reinterpret_cast<ttLibGoFrame *>(dstGoFrame);
  ttLibC_Frame *srcCF = reinterpret_cast<ttLibC_Frame *>(srcCFrame);
  ttLibGoFrame *srcGoF = reinterpret_cast<ttLibGoFrame *>(srcGoFrame);
  return r->toBgr(dstCF, dstGoF, srcCF, srcGoF);
}

bool LibyuvResampler_toYuv420(
    void *resampler,
    void *dstCFrame, void *dstGoFrame,
    void *srcCFrame, void *srcGoFrame) {
  if(resampler == nullptr || dstCFrame == nullptr || dstGoFrame == nullptr
  || srcCFrame == nullptr || srcGoFrame == nullptr) {
    return false;
  }
  LibyuvResampler *r = reinterpret_cast<LibyuvResampler *>(resampler);
  ttLibC_Frame *dstCF = reinterpret_cast<ttLibC_Frame *>(dstCFrame);
  ttLibGoFrame *dstGoF = reinterpret_cast<ttLibGoFrame *>(dstGoFrame);
  ttLibC_Frame *srcCF = reinterpret_cast<ttLibC_Frame *>(srcCFrame);
  ttLibGoFrame *srcGoF = reinterpret_cast<ttLibGoFrame *>(srcGoFrame);
  return r->toYuv420(dstCF, dstGoF, srcCF, srcGoF);
}
}