#ifndef TTLIBGO_ENCODER_X264_HPP
#define TTLIBGO_ENCODER_X264_HPP

#include "../encoder.hpp"

class X264Encoder : public Encoder {
public:
  X264Encoder(maps *mp);
  ~X264Encoder();
  bool encodeFrame(ttLibC_Frame *cFrame, ttLibGoFrame *goFrame, void *ptr);
  bool forceNextFrameType(string type);
private:
  void *_encoder;
};

#endif
