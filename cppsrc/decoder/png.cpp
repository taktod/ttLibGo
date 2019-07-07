#include "png.hpp"
#include "../util.hpp"
#include <iostream>

using namespace std;

extern "C" {
extern bool ttLibGoFrameCallback(void *ptr, ttLibC_Frame *frame);
}
PngDecoder::PngDecoder(maps *mp) {
#ifdef __ENABLE_LIBPNG__
  _decoder = ttLibC_PngDecoder_make();
#endif
}
PngDecoder::~PngDecoder() {
#ifdef __ENABLE_LIBPNG__
  ttLibC_PngDecoder_close(&_decoder);
#endif
}
bool PngDecoder::decodeFrame(ttLibC_Frame *cFrame, ttLibGoFrame *goFrame, void *ptr) {
  bool result = false;
#ifdef __ENABLE_LIBPNG__
  update(cFrame, goFrame);
  result = ttLibC_PngDecoder_decode(_decoder, (ttLibC_Png *)cFrame, [](void *ptr, ttLibC_Bgr *bgr) -> bool{
    return ttLibGoFrameCallback(ptr, (ttLibC_Frame *)bgr);
  }, ptr);
  reset(cFrame, goFrame);
#endif
  return result;
}
