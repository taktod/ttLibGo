#include "jpeg.hpp"
#include "../util.hpp"
#include <iostream>
#include <ttLibC/container/container.h>

using namespace std;

extern "C" {
typedef void *(* ttLibC_make_func)();
typedef void *(* ttLibC_codec_func)(void *, void *, ttLibC_getFrameFunc, void *);
typedef void (* ttLibC_close_func)(void **);

extern ttLibC_make_func ttLibGo_JpegDecoder_make;
extern ttLibC_codec_func ttLibGo_JpegDecoder_decode;
extern ttLibC_close_func ttLibGo_JpegDecoder_close;

extern bool ttLibGoFrameCallback(void *ptr, ttLibC_Frame *frame);
}
JpegDecoder::JpegDecoder(maps *mp) {
  if(ttLibGo_JpegDecoder_make != nullptr) {
    _decoder = (*ttLibGo_JpegDecoder_make)();
  }
}
JpegDecoder::~JpegDecoder() {
  if(ttLibGo_JpegDecoder_close != nullptr) {
    (*ttLibGo_JpegDecoder_close)(&_decoder);
  }
}
bool JpegDecoder::decodeFrame(ttLibC_Frame *cFrame, ttLibGoFrame *goFrame, void *ptr) {
  bool result = false;
  if(ttLibGo_JpegDecoder_decode != nullptr) {
    update(cFrame, goFrame);
    result = (*ttLibGo_JpegDecoder_decode)(_decoder, cFrame, ttLibGoFrameCallback, ptr);
    reset(cFrame, goFrame);
  }
  return result;
}
