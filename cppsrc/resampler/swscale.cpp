#include "swscale.hpp"
#include "../util.hpp"
#include <iostream>

#include <ttLibC/frame/video/bgr.h>
#include <ttLibC/frame/video/yuv420.h>

using namespace std;

extern "C" {
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
#ifdef __ENABLE_SWSCALE__
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
  _resampler = ttLibC_SwscaleResampler_make(
    inType, inSubType, inWidth, inHeight,
    outType, outSubType, outWidth, outHeight,
    scaleMode
  );
#endif
}

SwscaleResampler::~SwscaleResampler() {
#ifdef __ENABLE_SWSCALE__
  ttLibC_SwscaleResampler_close(&_resampler);
#endif
}

bool SwscaleResampler::resampleFrame(ttLibC_Frame *cFrame, ttLibGoFrame *goFrame, void *ptr) {
  bool result = false;
#ifdef __ENABLE_SWSCALE__
  update(cFrame, goFrame);
  result = ttLibC_SwscaleResampler_resample(
    _resampler,
    cFrame,
    ttLibGoFrameCallback,
    ptr);
  reset(cFrame, goFrame);
#endif
  return result;
}