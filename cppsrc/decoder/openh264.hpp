#ifndef TTLIBGO_DECODER_OPENH264_HPP
#define TTLIBGO_DECODER_OPENH264_HPP

#include "../decoder.hpp"
#include <ttLibC/decoder/openh264Decoder.h>

class Openh264Decoder : public Decoder {
public:
  Openh264Decoder(maps *mp);
  ~Openh264Decoder();
  bool decodeFrame(ttLibC_Frame *cFrame, ttLibGoFrame *goFrame, void *ptr);
private:
  ttLibC_Openh264Decoder *_decoder;
};

#endif
