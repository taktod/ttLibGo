#include "speex.hpp"
#include "../util.hpp"
#include <iostream>

using namespace std;

extern "C" {
extern bool ttLibGoFrameCallback(void *ptr, ttLibC_Frame *frame);
}
SpeexDecoder::SpeexDecoder(maps *mp) {
#ifdef __ENABLE_SPEEX__
  _decoder = ttLibC_SpeexDecoder_make(mp->getUint32("sampleRate"), mp->getUint32("channelNum"));
#endif
}
SpeexDecoder::~SpeexDecoder() {
#ifdef __ENABLE_SPEEX__
  ttLibC_SpeexDecoder_close(&_decoder);
#endif
}
bool SpeexDecoder::decodeFrame(ttLibC_Frame *cFrame, ttLibGoFrame *goFrame, void *ptr) {
  bool result = false;
#ifdef __ENABLE_SPEEX__
  update(cFrame, goFrame);
  result = ttLibC_SpeexDecoder_decode(_decoder, (ttLibC_Speex *)cFrame, [](void *ptr, ttLibC_PcmS16 *pcm) -> bool{
    return ttLibGoFrameCallback(ptr, (ttLibC_Frame *)pcm);
  }, ptr);
  reset(cFrame, goFrame);
#endif
  return result;
}
