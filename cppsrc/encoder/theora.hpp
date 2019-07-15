#ifndef TTLIBGO_ENCODER_THEORA_HPP
#define TTLIBGO_ENCODER_THEORA_HPP

#include "../encoder.hpp"

class TheoraEncoder : public Encoder {
public:
  TheoraEncoder(maps *mp);
  ~TheoraEncoder();
  bool encodeFrame(ttLibC_Frame *cFrame, ttLibGoFrame *goFrame, void *ptr);
private:
  void *_encoder;
};

#endif
