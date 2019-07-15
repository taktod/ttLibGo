#ifndef TTLIBGO_DECODER_OPENH264_HPP
#define TTLIBGO_DECODER_OPENH264_HPP

#include "../decoder.hpp"

class Openh264Decoder : public Decoder {
public:
  Openh264Decoder(maps *mp);
  ~Openh264Decoder();
  bool decodeFrame(ttLibC_Frame *cFrame, ttLibGoFrame *goFrame, void *ptr);
private:
  void *_decoder;
};

#endif
