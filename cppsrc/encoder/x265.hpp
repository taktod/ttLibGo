#ifndef TTLIBGO_ENCODER_X265_HPP
#define TTLIBGO_ENCODER_X265_HPP

#include "../encoder.hpp"

class X265Encoder : public Encoder {
public:
  X265Encoder(maps *mp);
  ~X265Encoder();
  bool encodeFrame(ttLibC_Frame *cFrame, ttLibGoFrame *goFrame, void *ptr);
  bool forceNextFrameType(string type);
private:
  void *_encoder;
};

#endif
