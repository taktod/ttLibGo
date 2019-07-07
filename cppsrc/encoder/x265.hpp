#ifndef TTLIBGO_ENCODER_X265_HPP
#define TTLIBGO_ENCODER_X265_HPP

#include "../encoder.hpp"
#include <ttLibC/encoder/x265Encoder.h>

class X265Encoder : public Encoder {
public:
  X265Encoder(maps *mp);
  ~X265Encoder();
  bool encodeFrame(ttLibC_Frame *cFrame, ttLibGoFrame *goFrame, void *ptr);
  bool forceNextFrameType(string type);
private:
  ttLibC_X265Encoder *_encoder;
};

#endif
