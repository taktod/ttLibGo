#include "swscale.hpp"
#include "../util.hpp"
#include <iostream>

#include "ttLibC/ttLibC/frame/video/bgr.h"
#include "ttLibC/ttLibC/frame/video/yuv420.h"
#include "ttLibC/ttLibC/resampler/swscaleResampler.h"

using namespace std;

extern "C" {
typedef void *(* ttLibC_SwscaleResampler_make_func)(ttLibC_Frame_Type, uint32_t, uint32_t, uint32_t, ttLibC_Frame_Type, uint32_t, uint32_t, uint32_t, ttLibC_SwscaleResampler_Mode);
typedef void *(* ttLibC_SwscaleResampler_resample_func)(void *, void *, void *);
typedef void (* ttLibC_close_func)(void **);

extern ttLibC_SwscaleResampler_make_func     ttLibGo_SwscaleResampler_make;
extern ttLibC_SwscaleResampler_resample_func ttLibGo_SwscaleResampler_resample;
extern ttLibC_close_func                     ttLibGo_SwscaleResampler_close;

extern bool ttLibGoFrameCallback(void *ptr, ttLibC_Frame *frame);
}

uint32_t SwscaleResampler::getSubType(ttLibC_Frame_Type type, string name) {
  switch(type) {
  case frameType_bgr:
    if(name == "bgr") {return BgrType_bgr;}
    if(name == "rgb") {return BgrType_rgb;}
    if(name == "abgr") {return BgrType_abgr;}
    if(name == "argb") {return BgrType_argb;}
    if(name == "bgra") {return BgrType_bgra;}
    if(name == "rgba") {return BgrType_rgba;}
    break;
  case frameType_yuv420:
    if(name == "yuv420") {return Yuv420Type_planar;}
    if(name == "yuv420SemiPlanar") {return Yuv420Type_semiPlanar;}
    if(name == "yvu420") {return Yvu420Type_planar;}
    if(name == "yvu420SemiPlanar") {return Yvu420Type_semiPlanar;}
    break;
  default:
    break;
  }
  return 99;
}

SwscaleResampler::SwscaleResampler(maps *mp) {
  if(ttLibGo_SwscaleResampler_make != nullptr) {
    ttLibC_Frame_Type inType = Frame_getFrameTypeFromString(mp->getString("inType"));
    uint32_t inSubType = getSubType(inType, mp->getString("inSubType"));
    uint32_t inWidth = mp->getUint32("inWidth");
    uint32_t inHeight = mp->getUint32("inHeight");
    ttLibC_Frame_Type outType = Frame_getFrameTypeFromString(mp->getString("outType"));
    uint32_t outSubType = getSubType(outType, mp->getString("outSubType"));
    uint32_t outWidth = mp->getUint32("outWidth");
    uint32_t outHeight = mp->getUint32("outHeight");
    string mode = mp->getString("mode");
    ttLibC_SwscaleResampler_Mode scaleMode = SwscaleResampler_FastBiLinear;
    if(mode == "X") {
      scaleMode = SwscaleResampler_X;
    }
    else if(mode == "Area") {
      scaleMode = SwscaleResampler_Area;
    }
    else if(mode == "Sinc") {
      scaleMode = SwscaleResampler_Sinc;
    }
    else if(mode == "Point") {
      scaleMode = SwscaleResampler_Point;
    }
    else if(mode == "Gauss") {
      scaleMode = SwscaleResampler_Gauss;
    }
    else if(mode == "Spline") {
      scaleMode = SwscaleResampler_Spline;
    }
    else if(mode == "Bicubic") {
      scaleMode = SwscaleResampler_Bicubic;
    }
    else if(mode == "Lanczos") {
      scaleMode = SwscaleResampler_Lanczos;
    }
    else if(mode == "Bilinear") {
      scaleMode = SwscaleResampler_Bilinear;
    }
    else if(mode == "Bicublin") {
      scaleMode = SwscaleResampler_Bicublin;
    }
    else if(mode == "FastBilinear") {
      scaleMode = SwscaleResampler_FastBiLinear;
    }
    _resampler = (*ttLibGo_SwscaleResampler_make)(
      inType, inSubType, inWidth, inHeight,
      outType, outSubType, outWidth, outHeight,
      scaleMode
    );
  }
}

SwscaleResampler::~SwscaleResampler() {
  if(ttLibGo_SwscaleResampler_close != nullptr) {
    (* ttLibGo_SwscaleResampler_close)(&_resampler);
  }
}
bool SwscaleResampler::resample(
    ttLibC_Frame *dstCFrame, ttLibGoFrame *dstGoFrame,
    ttLibC_Frame *srcCFrame, ttLibGoFrame *srcGoFrame) {
  bool result = false;
  if(ttLibGo_SwscaleResampler_resample == nullptr) {
    return false;
  }
  else {
    update(dstCFrame, dstGoFrame);
    srcfp.update(srcCFrame, srcGoFrame);
    result = (*ttLibGo_SwscaleResampler_resample)(_resampler, dstCFrame, srcCFrame);
    srcfp.reset(srcCFrame, srcGoFrame);
    reset(dstCFrame, dstGoFrame);
  }
  return result;
}

extern "C" {
bool SwscaleResampler_resample(void *resampler,
    void *dstCFrame, void *dstGoFrame,
    void *srcCFrame, void *srcGoFrame) {
  if(resampler == nullptr || dstCFrame == nullptr || dstGoFrame == nullptr
  || srcCFrame == nullptr || srcGoFrame == nullptr) {
    return false;
  }
  SwscaleResampler *r = reinterpret_cast<SwscaleResampler *>(resampler);
  ttLibC_Frame *dstCF = reinterpret_cast<ttLibC_Frame *>(dstCFrame);
  ttLibGoFrame *dstGoF = reinterpret_cast<ttLibGoFrame *>(dstGoFrame);
  ttLibC_Frame *srcCF = reinterpret_cast<ttLibC_Frame *>(srcCFrame);
  ttLibGoFrame *srcGoF = reinterpret_cast<ttLibGoFrame *>(srcGoFrame);
  return r->resample(dstCF, dstGoF, srcCF, srcGoF);
}
}
