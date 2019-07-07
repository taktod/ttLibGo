#include "jpeg.hpp"
#include "../util.hpp"
#include <iostream>

using namespace std;

extern "C" {
extern bool ttLibGoFrameCallback(void *ptr, ttLibC_Frame *frame);
}
JpegEncoder::JpegEncoder(maps *mp) {
#ifdef __ENABLE_JPEG__
  _encoder = ttLibC_JpegEncoder_make(mp->getUint32("width"), mp->getUint32("height"), mp->getUint32("quality"));
#endif
}

JpegEncoder::~JpegEncoder() {
#ifdef __ENABLE_JPEG__
  ttLibC_JpegEncoder_close(&_encoder);
#endif
}
bool JpegEncoder::encodeFrame(ttLibC_Frame *cFrame, ttLibGoFrame *goFrame, void *ptr) {
  bool result = false;
#ifdef __ENABLE_JPEG__
  if(cFrame->type != frameType_yuv420) {
    return result;
  }
  update(cFrame, goFrame);
  result = ttLibC_JpegEncoder_encode(_encoder, (ttLibC_Yuv420 *)cFrame, [](void *ptr, ttLibC_Jpeg *jpeg)->bool {
    return ttLibGoFrameCallback(ptr, (ttLibC_Frame *)jpeg);
  }, ptr);
  reset(cFrame, goFrame);
#endif
  return result;
}

bool JpegEncoder::setQuality(uint32_t quality) {
  bool result = false;
#ifdef __ENABLE_JPEG__
  result = ttLibC_JpegEncoder_setQuality(_encoder, quality);
#endif
  return result;
}

extern "C" {
bool JpegEncoder_setQuality(void *encoder, uint32_t quality) {
  JpegEncoder *j = reinterpret_cast<JpegEncoder *>(encoder);
  return j->setQuality(quality);
}  
}