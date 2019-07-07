#include "vorbis.hpp"
#include "../util.hpp"
#include <iostream>

using namespace std;

extern "C" {
extern bool ttLibGoFrameCallback(void *ptr, ttLibC_Frame *frame);
}
VorbisEncoder::VorbisEncoder(maps *mp) {
#ifdef __ENABLE_VORBIS_ENCODE__
  _encoder = ttLibC_VorbisEncoder_make(mp->getUint32("sampleRate"), mp->getUint32("channelNum"));
#endif
}

VorbisEncoder::~VorbisEncoder() {
#ifdef __ENABLE_VORBIS_ENCODE__
  ttLibC_VorbisEncoder_close(&_encoder);
#endif
}
bool VorbisEncoder::encodeFrame(ttLibC_Frame *cFrame, ttLibGoFrame *goFrame, void *ptr) {
  bool result = false;
#ifdef __ENABLE_VORBIS_ENCODE__
  if(cFrame->type != frameType_pcmS16
  && cFrame->type != frameType_pcmF32) {
    return result;
  }
  update(cFrame, goFrame);
  result = ttLibC_VorbisEncoder_encode(_encoder, (ttLibC_Audio *)cFrame, [](void *ptr, ttLibC_Vorbis *vorbis)->bool {
    return ttLibGoFrameCallback(ptr, (ttLibC_Frame *)vorbis);
  }, ptr);
  reset(cFrame, goFrame);
#endif
  return result;
}
