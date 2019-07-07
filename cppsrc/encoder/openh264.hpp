#ifndef TTLIBGO_ENCODER_OPENH264_HPP
#define TTLIBGO_ENCODER_OPENH264_HPP

#include "../encoder.hpp"
#include <ttLibC/encoder/openh264Encoder.h>

class Openh264Encoder : public Encoder {
public:
  Openh264Encoder(maps *mp);
  ~Openh264Encoder();
  bool encodeFrame(ttLibC_Frame *cFrame, ttLibGoFrame *goFrame, void *ptr);
  bool setRCMode(string mode);
  bool setIDRInterval(uint32_t value);
  bool forceNextKeyFrame();
private:
  ttLibC_Openh264Encoder *_encoder;
};

#endif
