#include "speex.hpp"
#include "../util.hpp"
#include <iostream>
#include <ttLibC/container/container.h>

using namespace std;

extern "C" {
typedef void *(* ttLibC_SpeexDecoder_make_func)(uint32_t, uint32_t);
typedef void *(* ttLibC_codec_func)(void *, void *, ttLibC_getFrameFunc, void *);
typedef void (* ttLibC_close_func)(void **);

extern ttLibC_SpeexDecoder_make_func ttLibGo_SpeexDecoder_make;
extern ttLibC_codec_func            ttLibGo_SpeexDecoder_decode;
extern ttLibC_close_func             ttLibGo_SpeexDecoder_close;

extern bool ttLibGoFrameCallback(void *ptr, ttLibC_Frame *frame);
}
SpeexDecoder::SpeexDecoder(maps *mp) {
  if(ttLibGo_SpeexDecoder_make != nullptr) {
    _decoder = (*ttLibGo_SpeexDecoder_make)(mp->getUint32("sampleRate"), mp->getUint32("channelNum"));
  }
}
SpeexDecoder::~SpeexDecoder() {
  if(ttLibGo_SpeexDecoder_close != nullptr) {
    (*ttLibGo_SpeexDecoder_close)(&_decoder);
  }
}
bool SpeexDecoder::decodeFrame(ttLibC_Frame *cFrame, ttLibGoFrame *goFrame, void *ptr) {
  bool result = false;
  if(ttLibGo_SpeexDecoder_decode != nullptr) {
    update(cFrame, goFrame);
    result = (*ttLibGo_SpeexDecoder_decode)(_decoder, cFrame, ttLibGoFrameCallback, ptr);
    reset(cFrame, goFrame);
  }
  return result;
}
