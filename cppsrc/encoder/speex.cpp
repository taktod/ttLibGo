#include "speex.hpp"
#include "../util.hpp"
#include <iostream>

using namespace std;

extern "C" {
extern bool ttLibGoFrameCallback(void *ptr, ttLibC_Frame *frame);
}
SpeexEncoder::SpeexEncoder(maps *mp) {
#ifdef __ENABLE_SPEEX__
  _encoder = ttLibC_SpeexEncoder_make(mp->getUint32("sampleRate"), mp->getUint32("channelNum"), mp->getUint32("quality"));
#endif
}

SpeexEncoder::~SpeexEncoder() {
#ifdef __ENABLE_SPEEX__
  ttLibC_SpeexEncoder_close(&_encoder);
#endif
}
bool SpeexEncoder::encodeFrame(ttLibC_Frame *cFrame, ttLibGoFrame *goFrame, void *ptr) {
  bool result = false;
#ifdef __ENABLE_SPEEX__
  if(cFrame->type != frameType_pcmS16) {
    return result;
  }
  update(cFrame, goFrame);
  result = ttLibC_SpeexEncoder_encode(_encoder, (ttLibC_PcmS16 *)cFrame, [](void *ptr, ttLibC_Speex *speex)->bool {
    return ttLibGoFrameCallback(ptr, (ttLibC_Frame *)speex);
  }, ptr);
  reset(cFrame, goFrame);
#endif
  return result;
}
