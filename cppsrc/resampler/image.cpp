#include "image.hpp"
#include "../util.hpp"
#include "../frame.hpp"
#include <iostream>

using namespace std;

extern "C" {
extern bool ttLibGoFrameCallback(void *ptr, ttLibC_Frame *frame);
}

uint32_t ImageResampler::getSubType(ttLibC_Frame_Type frameType, string name) {
  switch(frameType) {
  case frameType_bgr:
    if(name == "bgr") {
      return BgrType_bgr;
    }
    else if(name == "abgr") {
      return BgrType_abgr;
    }
    else if(name == "bgra") {
      return BgrType_bgra;
    }
    else if(name == "rgb") {
      return BgrType_rgb;
    }
    else if(name == "argb") {
      return BgrType_argb;
    }
    else if(name == "rgba") {
      return BgrType_rgba;
    }
    break;
  case frameType_yuv420:
    if(name == "yuv420") {
      return Yuv420Type_planar;
    }
    else if(name == "yuv420SemiPlanar") {
      return Yuv420Type_semiPlanar;
    }
    else if(name == "yvu420") {
      return Yvu420Type_planar;
    }
    else if(name == "yvu420SemiPlanar") {
      return Yvu420Type_semiPlanar;
    }
  default:
    break;
  }
  return 99;
}

ImageResampler::ImageResampler(maps *mp) {
  _targetType = Frame_getFrameTypeFromString(mp->getString("frameType"));
  _subType = getSubType(_targetType, mp->getString("subType"));
  _image = nullptr;
  switch(_targetType) {
  case frameType_bgr:
    _sourceType = frameType_yuv420;
    break;
  case frameType_yuv420:
    _sourceType = frameType_bgr;
    break;
  default:
    break;
  }
}
ImageResampler::~ImageResampler() {
  ttLibC_Video_close(&_image);
}

bool ImageResampler::resampleFrame(ttLibC_Frame *cFrame, ttLibGoFrame *goFrame, void *ptr) {
  if(cFrame->type != _sourceType) {
    return false;
  }
  bool result = false;
  update(cFrame, goFrame);
  switch(cFrame->type) {
  case frameType_bgr:
    {
      if(_image != nullptr && _image->inherit_super.type != frameType_yuv420) {
        ttLibC_Video_close(&_image);
      }
      ttLibC_Yuv420 *yuv = ttLibC_ImageResampler_makeYuv420FromBgr(
        (ttLibC_Yuv420 *)_image,
        (ttLibC_Yuv420_Type)_subType,
        (ttLibC_Bgr *)cFrame);
      if(yuv != nullptr) {
        _image = (ttLibC_Video *)yuv;
        result = ttLibGoFrameCallback(ptr, (ttLibC_Frame *)_image);
      }
    }
    break;
  case frameType_yuv420:
    {
      if(_image != nullptr && _image->inherit_super.type != frameType_bgr) {
        ttLibC_Video_close(&_image);
      }
      ttLibC_Bgr *bgr = ttLibC_ImageResampler_makeBgrFromYuv420(
        (ttLibC_Bgr *)_image,
        (ttLibC_Bgr_Type)_subType,
        (ttLibC_Yuv420 *)cFrame);
      if(bgr != nullptr) {
        _image = (ttLibC_Video *)bgr;
        result = ttLibGoFrameCallback(ptr, (ttLibC_Frame *)_image);
      }
    }
    break;
  default:
    break;
  }
  reset(cFrame, goFrame);
  return result;
}