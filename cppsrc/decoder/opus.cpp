#include "opus.hpp"
#include "../util.hpp"
#include <iostream>
#include <ttLibC/container/container.h>

using namespace std;

extern "C" {
typedef void *(* ttLibC_OpusDecoder_make_func)(uint32_t, uint32_t);
typedef void *(* ttLibC_codec_func)(void *, void *, ttLibC_getFrameFunc, void *);
typedef int (* ttLibC_opus_codecControl_func)(void *, const char *, int);
typedef void (* ttLibC_close_func)(void **);

extern ttLibC_OpusDecoder_make_func  ttLibGo_OpusDecoder_make;
extern ttLibC_codec_func             ttLibGo_OpusDecoder_decode;
extern ttLibC_close_func             ttLibGo_OpusDecoder_close;
extern ttLibC_opus_codecControl_func ttLibGo_OpusDecoder_codecControl;

extern bool ttLibGoFrameCallback(void *ptr, ttLibC_Frame *frame);
}
tOpusDecoder::tOpusDecoder(maps *mp) {
  if(ttLibGo_OpusDecoder_make != nullptr) {
    _decoder = (*ttLibGo_OpusDecoder_make)(mp->getUint32("sampleRate"), mp->getUint32("channelNum"));
  }
}
tOpusDecoder::~tOpusDecoder() {
  if(ttLibGo_OpusDecoder_close != nullptr) {
    (*ttLibGo_OpusDecoder_close)(&_decoder);
  }
}
bool tOpusDecoder::decodeFrame(ttLibC_Frame *cFrame, ttLibGoFrame *goFrame, void *ptr) {
  bool result = false;
  if(ttLibGo_OpusDecoder_decode != nullptr) {
    update(cFrame, goFrame);
    result = (*ttLibGo_OpusDecoder_decode)(_decoder, cFrame, ttLibGoFrameCallback, ptr);
    reset(cFrame, goFrame);
  }
  return result;
}

int tOpusDecoder::codecControl(string control, int value) {
  if(ttLibGo_OpusDecoder_codecControl != nullptr) {
    if(_decoder == NULL) {
      puts("decoderが準備されていません。");
      return -1;
    }
    return (*ttLibGo_OpusDecoder_codecControl)(_decoder, control.c_str(), value);
  }
  return -1;
}

extern "C" {
int OpusDecoder_codecControl(void *decoder, const char *key, int value) {
  tOpusDecoder *d = reinterpret_cast<tOpusDecoder *>(decoder);
  return d->codecControl(key, value);
}
}