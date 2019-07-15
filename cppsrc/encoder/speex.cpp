#include "speex.hpp"
#include "../util.hpp"
#include <iostream>
#include <ttLibC/container/container.h>

using namespace std;

extern "C" {
typedef void *(* ttLibC_SpeexEncoder_make_func)(uint32_t, uint32_t, uint32_t);
typedef void *(* ttLibC_codec_func)(void *, void *, ttLibC_getFrameFunc, void *);
typedef void (* ttLibC_close_func)(void **);

extern ttLibC_SpeexEncoder_make_func ttLibGo_SpeexEncoder_make;
extern ttLibC_codec_func             ttLibGo_SpeexEncoder_encode;
extern ttLibC_close_func             ttLibGo_SpeexEncoder_close;

extern bool ttLibGoFrameCallback(void *ptr, ttLibC_Frame *frame);
}
SpeexEncoder::SpeexEncoder(maps *mp) {
  if(ttLibGo_SpeexEncoder_make != nullptr) {
    _encoder = (*ttLibGo_SpeexEncoder_make)(mp->getUint32("sampleRate"), mp->getUint32("channelNum"), mp->getUint32("quality"));
  }
}

SpeexEncoder::~SpeexEncoder() {
  if(ttLibGo_SpeexEncoder_close != nullptr) {
    (*ttLibGo_SpeexEncoder_close)(&_encoder);
  }
}
bool SpeexEncoder::encodeFrame(ttLibC_Frame *cFrame, ttLibGoFrame *goFrame, void *ptr) {
  bool result = false;
  if(ttLibGo_SpeexEncoder_encode != nullptr) {
    if(cFrame->type != frameType_pcmS16) {
      return result;
    }
    update(cFrame, goFrame);
    result = (*ttLibGo_SpeexEncoder_encode)(_encoder, cFrame, ttLibGoFrameCallback, ptr);
    reset(cFrame, goFrame);
  }
  return result;
}
