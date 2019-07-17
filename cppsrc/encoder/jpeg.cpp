#include "jpeg.hpp"
#include "../util.hpp"
#include <iostream>
#include <ttLibC/container/container.h>

using namespace std;

extern "C" {
typedef void *(* ttLibC_JpegEncoder_make_func)(uint32_t, uint32_t, uint32_t);
typedef void *(* ttLibC_codec_func)(void *, void *, ttLibC_getFrameFunc, void *);
typedef void *(* ttLibC_JpegEncoder_setQuality_func)(void *, uint32_t);
typedef void (* ttLibC_close_func)(void **);

extern ttLibC_JpegEncoder_make_func       ttLibGo_JpegEncoder_make;
extern ttLibC_codec_func                  ttLibGo_JpegEncoder_encode;
extern ttLibC_JpegEncoder_setQuality_func ttLibGo_JpegEncoder_setQuality;
extern ttLibC_close_func                  ttLibGo_JpegEncoder_close;

extern bool ttLibGoFrameCallback(void *ptr, ttLibC_Frame *frame);
}
JpegEncoder::JpegEncoder(maps *mp) {
  if(ttLibGo_JpegEncoder_make != nullptr) {
    _encoder = (*ttLibGo_JpegEncoder_make)(mp->getUint32("width"), mp->getUint32("height"), mp->getUint32("quality"));
  }
}

JpegEncoder::~JpegEncoder() {
  if(ttLibGo_JpegEncoder_close != nullptr) {
    (*ttLibGo_JpegEncoder_close)(&_encoder);
  }
}
bool JpegEncoder::encodeFrame(ttLibC_Frame *cFrame, ttLibGoFrame *goFrame, void *ptr) {
  bool result = false;
  if(ttLibGo_JpegEncoder_encode != nullptr) {
    if(cFrame->type != frameType_yuv420) {
      return result;
    }
    update(cFrame, goFrame);
    result = (*ttLibGo_JpegEncoder_encode)(_encoder, cFrame, ttLibGoFrameCallback, ptr);
    reset(cFrame, goFrame);
  }
  return result;
}

bool JpegEncoder::setQuality(uint32_t quality) {
  bool result = false;
  if(ttLibGo_JpegEncoder_setQuality != nullptr) {
    result = (*ttLibGo_JpegEncoder_setQuality)(_encoder, quality);
  }
  return result;
}

extern "C" {
bool JpegEncoder_setQuality(void *encoder, uint32_t quality) {
  JpegEncoder *j = reinterpret_cast<JpegEncoder *>(encoder);
  return j->setQuality(quality);
}  
}
