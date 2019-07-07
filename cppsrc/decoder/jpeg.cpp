#include "jpeg.hpp"
#include "../util.hpp"
#include <iostream>

using namespace std;

extern "C" {
extern bool ttLibGoFrameCallback(void *ptr, ttLibC_Frame *frame);
}
JpegDecoder::JpegDecoder(maps *mp) {
#ifdef __ENABLE_JPEG__
  _decoder = ttLibC_JpegDecoder_make();
#endif
}
JpegDecoder::~JpegDecoder() {
#ifdef __ENABLE_JPEG__
  ttLibC_JpegDecoder_close(&_decoder);
#endif
}
bool JpegDecoder::decodeFrame(ttLibC_Frame *cFrame, ttLibGoFrame *goFrame, void *ptr) {
  bool result = false;
#ifdef __ENABLE_JPEG__
  update(cFrame, goFrame);
  result = ttLibC_JpegDecoder_decode(_decoder, (ttLibC_Jpeg *)cFrame, [](void *ptr, ttLibC_Yuv420 *yuv) -> bool{
    return ttLibGoFrameCallback(ptr, (ttLibC_Frame *)yuv);
  }, ptr);
  reset(cFrame, goFrame);
#endif
  return result;
}
