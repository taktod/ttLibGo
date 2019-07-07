#include "theora.hpp"
#include "../util.hpp"
#include <iostream>

using namespace std;

extern "C" {
extern bool ttLibGoFrameCallback(void *ptr, ttLibC_Frame *frame);
}
TheoraDecoder::TheoraDecoder(maps *mp) {
#ifdef __ENABLE_THEORA__
  _decoder = ttLibC_TheoraDecoder_make();
#endif
}
TheoraDecoder::~TheoraDecoder() {
#ifdef __ENABLE_THEORA__
  ttLibC_TheoraDecoder_close(&_decoder);
#endif
}
bool TheoraDecoder::decodeFrame(ttLibC_Frame *cFrame, ttLibGoFrame *goFrame, void *ptr) {
  bool result = false;
#ifdef __ENABLE_THEORA__
  update(cFrame, goFrame);
  result = ttLibC_TheoraDecoder_decode(_decoder, (ttLibC_Theora *)cFrame, [](void *ptr, ttLibC_Yuv420 *yuv) -> bool{
    return ttLibGoFrameCallback(ptr, (ttLibC_Frame *)yuv);
  }, ptr);
  reset(cFrame, goFrame);
#endif
  return result;
}
