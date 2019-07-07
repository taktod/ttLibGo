#ifndef TTLIBGO_ENCODER_THEORA_HPP
#define TTLIBGO_ENCODER_THEORA_HPP

#include "../encoder.hpp"
#include <ttLibC/encoder/theoraEncoder.h>

class TheoraEncoder : public Encoder {
public:
  TheoraEncoder(maps *mp);
  ~TheoraEncoder();
  bool encodeFrame(ttLibC_Frame *cFrame, ttLibGoFrame *goFrame, void *ptr);
private:
  ttLibC_TheoraEncoder *_encoder;
};

#endif
