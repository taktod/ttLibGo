#include "theora.hpp"
#include "../util.hpp"
#include <iostream>
#include <ttLibC/container/container.h>

using namespace std;

extern "C" {
typedef void *(* ttLibC_make_func)();
typedef void *(* ttLibC_codec_func)(void *, void *, ttLibC_getFrameFunc, void *);
typedef void (* ttLibC_close_func)(void **);

extern ttLibC_make_func   ttLibGo_TheoraDecoder_make;
extern ttLibC_codec_func ttLibGo_TheoraDecoder_decode;
extern ttLibC_close_func  ttLibGo_TheoraDecoder_close;

extern bool ttLibGoFrameCallback(void *ptr, ttLibC_Frame *frame);
}
TheoraDecoder::TheoraDecoder(maps *mp) {
  if(ttLibGo_TheoraDecoder_make != nullptr) {
    _decoder = (*ttLibGo_TheoraDecoder_make)();
  }
}
TheoraDecoder::~TheoraDecoder() {
  if(ttLibGo_TheoraDecoder_close != nullptr) {
    (*ttLibGo_TheoraDecoder_close)(&_decoder);
  }
}
bool TheoraDecoder::decodeFrame(ttLibC_Frame *cFrame, ttLibGoFrame *goFrame, void *ptr) {
  bool result = false;
  if(ttLibGo_TheoraDecoder_decode != nullptr) {
    update(cFrame, goFrame);
    result = (*ttLibGo_TheoraDecoder_decode)(_decoder, cFrame, ttLibGoFrameCallback, ptr);
    reset(cFrame, goFrame);
  }
  return result;
}
