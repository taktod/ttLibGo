#include "theora.hpp"
#include "../util.hpp"
#include <iostream>

using namespace std;

extern "C" {
extern bool ttLibGoFrameCallback(void *ptr, ttLibC_Frame *frame);
}
TheoraEncoder::TheoraEncoder(maps *mp) {
#ifdef __ENABLE_THEORA__
  _encoder = ttLibC_TheoraEncoder_make_ex(
    mp->getUint32("width"),
    mp->getUint32("height"),
    mp->getUint32("quality"),
    mp->getUint32("bitrate"),
    mp->getUint32("keyFrameInterval"));
#endif
}

TheoraEncoder::~TheoraEncoder() {
#ifdef __ENABLE_THEORA__
  ttLibC_TheoraEncoder_close(&_encoder);
#endif
}
bool TheoraEncoder::encodeFrame(ttLibC_Frame *cFrame, ttLibGoFrame *goFrame, void *ptr) {
  bool result = false;
#ifdef __ENABLE_THEORA__
  if(cFrame->type != frameType_yuv420) {
    return result;
  }
  update(cFrame, goFrame);
  result = ttLibC_TheoraEncoder_encode(_encoder, (ttLibC_Yuv420 *)cFrame, [](void *ptr, ttLibC_Theora *theora)->bool {
    return ttLibGoFrameCallback(ptr, (ttLibC_Frame *)theora);
  }, ptr);
  reset(cFrame, goFrame);
#endif
  return result;
}
