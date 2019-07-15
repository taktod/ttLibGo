#include "openh264.hpp"
#include "../util.hpp"
#include <iostream>
#include <ttLibC/container/container.h>

using namespace std;

extern "C" {
typedef void *(* ttLibC_make_func)();
typedef void *(* ttLibC_codec_func)(void *, void *, ttLibC_getFrameFunc, void *);
typedef void (* ttLibC_close_func)(void **);

extern ttLibC_make_func   ttLibGo_Openh264Decoder_make;
extern ttLibC_codec_func ttLibGo_Openh264Decoder_decode;
extern ttLibC_close_func  ttLibGo_Openh264Decoder_close;

extern bool ttLibGoFrameCallback(void *ptr, ttLibC_Frame *frame);
}
Openh264Decoder::Openh264Decoder(maps *mp) {
  if(ttLibGo_Openh264Decoder_make != nullptr) {
    _decoder = (*ttLibGo_Openh264Decoder_make)();
  }
}
Openh264Decoder::~Openh264Decoder() {
  if(ttLibGo_Openh264Decoder_close != nullptr) {
    (*ttLibGo_Openh264Decoder_close)(&_decoder);
  }
}
bool Openh264Decoder::decodeFrame(ttLibC_Frame *cFrame, ttLibGoFrame *goFrame, void *ptr) {
  bool result = false;
  if(ttLibGo_Openh264Decoder_decode != nullptr) {
    update(cFrame, goFrame);
    result = (*ttLibGo_Openh264Decoder_decode)(_decoder, cFrame, ttLibGoFrameCallback, ptr);
    reset(cFrame, goFrame);
  }
  return result;
}
