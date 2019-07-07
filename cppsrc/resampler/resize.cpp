#include "resize.hpp"
#include "../util.hpp"
#include "../frame.hpp"
#include <iostream>

using namespace std;

extern "C" {
extern bool ttLibGoFrameCallback(void *ptr, ttLibC_Frame *frame);
}

ResizeResampler::ResizeResampler(maps *mp) {
  _width = mp->getUint32("width");
  _height = mp->getUint32("height");
  _isQuick = mp->getUint32("isQuick") == 1;
  _image = nullptr;
}

ResizeResampler::~ResizeResampler() {
  ttLibC_Video_close(&_image);
}

bool ResizeResampler::resampleFrame(ttLibC_Frame *cFrame, ttLibGoFrame *goFrame, void *ptr) {
  if(_image != nullptr && _image->inherit_super.type != cFrame->type) {
    ttLibC_Video_close(&_image);
  }
  switch(cFrame->type) {
  case frameType_bgr:
    {
      bool result = false;
      update(cFrame, goFrame);
      ttLibC_Bgr *source = (ttLibC_Bgr *)cFrame;
      ttLibC_Bgr *bgr = ttLibC_ImageResizer_resizeBgr(
        (ttLibC_Bgr *)_image,
		    source->type,
		    _width,
		    _height,
		    source);
      if(bgr != nullptr) {
        _image = (ttLibC_Video *)bgr;
        result = ttLibGoFrameCallback(ptr, (ttLibC_Frame *)_image);
      }
      reset(cFrame, goFrame);
      return result;
    }
    break;
  case frameType_yuv420:
    {
      bool result = false;
      update(cFrame, goFrame);
      ttLibC_Yuv420 *source = (ttLibC_Yuv420 *)cFrame;
      ttLibC_Yuv420 *yuv = ttLibC_ImageResizer_resizeYuv420(
		    (ttLibC_Yuv420 *)_image,
		    source->type,
		    _width,
	      _height,
	  	  source,
	    	_isQuick);
      if(yuv != nullptr) {
        _image = (ttLibC_Video *)yuv;
        result = ttLibGoFrameCallback(ptr, (ttLibC_Frame *)_image);
      }
      reset(cFrame, goFrame);
      return result;
    }
    break;
  default:
    return false;
  }
}
