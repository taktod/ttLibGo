#include "vorbis.hpp"
#include "../util.hpp"
#include <iostream>

using namespace std;

extern "C" {
extern bool ttLibGoFrameCallback(void *ptr, ttLibC_Frame *frame);
}
VorbisDecoder::VorbisDecoder(maps *mp) {
#ifdef __ENABLE_VORBIS_DECODE__
  _decoder = ttLibC_VorbisDecoder_make();
#endif
}
VorbisDecoder::~VorbisDecoder() {
#ifdef __ENABLE_VORBIS_DECODE__
  ttLibC_VorbisDecoder_close(&_decoder);
#endif
}
bool VorbisDecoder::decodeFrame(ttLibC_Frame *cFrame, ttLibGoFrame *goFrame, void *ptr) {
  bool result = false;
#ifdef __ENABLE_VORBIS_DECODE__
  update(cFrame, goFrame);
  result = ttLibC_VorbisDecoder_decode(_decoder, (ttLibC_Vorbis *)cFrame, [](void *ptr, ttLibC_PcmF32 *pcm) -> bool{
    return ttLibGoFrameCallback(ptr, (ttLibC_Frame *)pcm);
  }, ptr);
  reset(cFrame, goFrame);
#endif
  return result;
}
