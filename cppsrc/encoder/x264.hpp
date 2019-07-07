#ifndef TTLIBGO_ENCODER_X264_HPP
#define TTLIBGO_ENCODER_X264_HPP

#include "../encoder.hpp"
#include <ttLibC/encoder/x264Encoder.h>

class X264Encoder : public Encoder {
public:
  X264Encoder(maps *mp);
  ~X264Encoder();
  bool encodeFrame(ttLibC_Frame *cFrame, ttLibGoFrame *goFrame, void *ptr);
  bool forceNextFrameType(string type);
private:
  ttLibC_X264Encoder *_encoder;
};

#endif
