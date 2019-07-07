#include "openh264.hpp"
#include "../util.hpp"
#include <iostream>

using namespace std;

extern "C" {
extern bool ttLibGoFrameCallback(void *ptr, ttLibC_Frame *frame);
}
Openh264Decoder::Openh264Decoder(maps *mp) {
#ifdef __ENABLE_OPENH264__
  _decoder = ttLibC_Openh264Decoder_make();
#endif
}
Openh264Decoder::~Openh264Decoder() {
#ifdef __ENABLE_OPENH264__
  ttLibC_Openh264Decoder_close(&_decoder);
#endif
}
bool Openh264Decoder::decodeFrame(ttLibC_Frame *cFrame, ttLibGoFrame *goFrame, void *ptr) {
  bool result = false;
#ifdef __ENABLE_OPENH264__
  update(cFrame, goFrame);
  result = ttLibC_Openh264Decoder_decode(_decoder, (ttLibC_H264 *)cFrame, [](void *ptr, ttLibC_Yuv420 *yuv) -> bool{
    return ttLibGoFrameCallback(ptr, (ttLibC_Frame *)yuv);
  }, ptr);
  reset(cFrame, goFrame);
#endif
  return result;
}
