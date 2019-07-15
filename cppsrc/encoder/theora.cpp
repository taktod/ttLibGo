#include "theora.hpp"
#include "../util.hpp"
#include <iostream>

using namespace std;

extern "C" {
typedef void *(* ttLibC_TheoraEncoder_make_ex_func)(uint32_t, uint32_t, uint32_t, uint32_t, uint32_t);
typedef void *(* ttLibC_codec_func)(void *, void *, ttLibC_getFrameFunc, void *);
typedef void (* ttLibC_close_func)(void **);

extern ttLibC_TheoraEncoder_make_ex_func ttLibGo_TheoraEncoder_make_ex;
extern ttLibC_codec_func                 ttLibGo_TheoraEncoder_encode;
extern ttLibC_close_func                 ttLibGo_TheoraEncoder_close;

extern bool ttLibGoFrameCallback(void *ptr, ttLibC_Frame *frame);
}
TheoraEncoder::TheoraEncoder(maps *mp) {
  if(ttLibGo_TheoraEncoder_make_ex != nullptr) {
    _encoder = (*ttLibGo_TheoraEncoder_make_ex)(
      mp->getUint32("width"),
      mp->getUint32("height"),
      mp->getUint32("quality"),
      mp->getUint32("bitrate"),
      mp->getUint32("keyFrameInterval"));
  }
}

TheoraEncoder::~TheoraEncoder() {
  if(ttLibGo_TheoraEncoder_close != nullptr) {
    (*ttLibGo_TheoraEncoder_close)(&_encoder);
  }
}
bool TheoraEncoder::encodeFrame(ttLibC_Frame *cFrame, ttLibGoFrame *goFrame, void *ptr) {
  bool result = false;
  if(ttLibGo_TheoraEncoder_encode != nullptr) {
    if(cFrame->type != frameType_yuv420) {
      return result;
    }
    update(cFrame, goFrame);
    result = (*ttLibGo_TheoraEncoder_encode)(_encoder, cFrame, ttLibGoFrameCallback, ptr);
    reset(cFrame, goFrame);
  }
  return result;
}
