#include "opus.hpp"
#include "../util.hpp"
#include <iostream>
#include <ttLibC/container/container.h>

using namespace std;

extern "C" {
typedef void *(* ttLibC_OpusEncoder_make_func)(uint32_t, uint32_t, uint32_t);
typedef void *(* ttLibC_codec_func)(void *, void *, ttLibC_getFrameFunc, void *);
typedef int (* ttLibC_opus_codecControl_func)(void *, const char *, int);
typedef void (* ttLibC_close_func)(void **);

extern ttLibC_OpusEncoder_make_func  ttLibGo_OpusEncoder_make;
extern ttLibC_codec_func             ttLibGo_OpusEncoder_encode;
extern ttLibC_opus_codecControl_func ttLibGo_OpusEncoder_codecControl;
extern ttLibC_close_func             ttLibGo_OpusEncoder_close;

extern bool ttLibGoFrameCallback(void *ptr, ttLibC_Frame *frame);
}
tOpusEncoder::tOpusEncoder(maps *mp) {
  if(ttLibGo_OpusEncoder_make != nullptr) {
    _encoder = (*ttLibGo_OpusEncoder_make)(mp->getUint32("sampleRate"), mp->getUint32("channelNum"), mp->getUint32("unitSampleNum"));
  }
}

tOpusEncoder::~tOpusEncoder() {
  if(ttLibGo_OpusEncoder_close != nullptr) {
    (*ttLibGo_OpusEncoder_close)(&_encoder);
  }
}
bool tOpusEncoder::encodeFrame(ttLibC_Frame *cFrame, ttLibGoFrame *goFrame, void *ptr) {
  bool result = false;
  if(ttLibGo_OpusEncoder_encode != nullptr) {
    if(cFrame->type != frameType_pcmS16) {
      return result;
    }
    update(cFrame, goFrame);
    result = (*ttLibGo_OpusEncoder_encode)(_encoder, cFrame, ttLibGoFrameCallback, ptr);
    reset(cFrame, goFrame);
  }
  return result;
}

int tOpusEncoder::codecControl(string control, int value) {
  if(ttLibGo_OpusEncoder_codecControl != nullptr) {
    return (*ttLibGo_OpusEncoder_codecControl)(_encoder, control.c_str(), value);
  }
  return -1;
}

extern "C" {
int OpusEncoder_codecControl(void *encoder, const char *key, int value) {
  tOpusEncoder *e = reinterpret_cast<tOpusEncoder *>(encoder);
  return e->codecControl(key, value);
}
}