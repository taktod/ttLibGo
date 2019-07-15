#include "vorbis.hpp"
#include "../util.hpp"
#include <iostream>

using namespace std;

extern "C" {
typedef void *(* ttLibC_VorbisEncoder_make_func)(uint32_t, uint32_t);
typedef void *(* ttLibC_codec_func)(void *, void *, ttLibC_getFrameFunc, void *);
typedef void (* ttLibC_close_func)(void **);

extern ttLibC_VorbisEncoder_make_func ttLibGo_VorbisEncoder_make;
extern ttLibC_codec_func ttLibGo_VorbisEncoder_encode;
extern ttLibC_close_func ttLibGo_VorbisEncoder_close;

extern bool ttLibGoFrameCallback(void *ptr, ttLibC_Frame *frame);
}
VorbisEncoder::VorbisEncoder(maps *mp) {
  if(ttLibGo_VorbisEncoder_make != nullptr) {
    _encoder = (*ttLibGo_VorbisEncoder_make)(mp->getUint32("sampleRate"), mp->getUint32("channelNum"));
  }
}

VorbisEncoder::~VorbisEncoder() {
  if(ttLibGo_VorbisEncoder_close != nullptr) {
    (*ttLibGo_VorbisEncoder_close)(&_encoder);
  }
}
bool VorbisEncoder::encodeFrame(ttLibC_Frame *cFrame, ttLibGoFrame *goFrame, void *ptr) {
  bool result = false;
  if(ttLibGo_VorbisEncoder_encode != nullptr) {
    if(cFrame->type != frameType_pcmS16
    && cFrame->type != frameType_pcmF32) {
      return result;
    }
    update(cFrame, goFrame);
    result = (*ttLibGo_VorbisEncoder_encode)(_encoder, cFrame, ttLibGoFrameCallback, ptr);
    reset(cFrame, goFrame);
  }
  return result;
}
