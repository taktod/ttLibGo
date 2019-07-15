#include "vorbis.hpp"
#include "../util.hpp"
#include <iostream>
#include <ttLibC/container/container.h>

using namespace std;

extern "C" {
typedef void *(* ttLibC_make_func)();
typedef void *(* ttLibC_codec_func)(void *, void *, ttLibC_getFrameFunc, void *);
typedef void (* ttLibC_close_func)(void **);

extern ttLibC_make_func   ttLibGo_VorbisDecoder_make;
extern ttLibC_codec_func ttLibGo_VorbisDecoder_decode;
extern ttLibC_close_func  ttLibGo_VorbisDecoder_close;

extern bool ttLibGoFrameCallback(void *ptr, ttLibC_Frame *frame);
}
VorbisDecoder::VorbisDecoder(maps *mp) {
  if(ttLibGo_VorbisDecoder_make != nullptr) {
    _decoder = (*ttLibGo_VorbisDecoder_make)();
  }
}
VorbisDecoder::~VorbisDecoder() {
  if(ttLibGo_VorbisDecoder_close != nullptr) {
    (*ttLibGo_VorbisDecoder_close)(&_decoder);
  }
}
bool VorbisDecoder::decodeFrame(ttLibC_Frame *cFrame, ttLibGoFrame *goFrame, void *ptr) {
  bool result = false;
  if(ttLibGo_VorbisDecoder_decode != nullptr) {
    update(cFrame, goFrame);
    result = (*ttLibGo_VorbisDecoder_decode)(_decoder, cFrame, ttLibGoFrameCallback, ptr);
    reset(cFrame, goFrame);
  }
  return result;
}
