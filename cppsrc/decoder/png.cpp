#include "png.hpp"
#include "../util.hpp"
#include <iostream>
#include "ttLibC/ttLibC/container/container.h"

using namespace std;

extern "C" {
typedef void *(* ttLibC_make_func)();
typedef void *(* ttLibC_codec_func)(void *, void *, ttLibC_getFrameFunc, void *);
typedef void (* ttLibC_close_func)(void **);

extern ttLibC_make_func   ttLibGo_PngDecoder_make;
extern ttLibC_codec_func ttLibGo_PngDecoder_decode;
extern ttLibC_close_func  ttLibGo_PngDecoder_close;

extern bool ttLibGoFrameCallback(void *ptr, ttLibC_Frame *frame);
}
PngDecoder::PngDecoder(maps *mp) {
  if(ttLibGo_PngDecoder_make != nullptr) {
    _decoder = (*ttLibGo_PngDecoder_make)();
  }
}
PngDecoder::~PngDecoder() {
  if(ttLibGo_PngDecoder_close != nullptr) {
    (*ttLibGo_PngDecoder_close)(&_decoder);
  }
}
bool PngDecoder::decodeFrame(ttLibC_Frame *cFrame, ttLibGoFrame *goFrame, void *ptr) {
  bool result = false;
  if(ttLibGo_PngDecoder_decode != nullptr) {
    update(cFrame, goFrame);
    result = (*ttLibGo_PngDecoder_decode)(_decoder, cFrame, ttLibGoFrameCallback, ptr);
    reset(cFrame, goFrame);
  }
  return result;
}
