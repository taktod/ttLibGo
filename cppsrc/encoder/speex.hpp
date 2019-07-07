#ifndef TTLIBGO_ENCODER_SPEEX_HPP
#define TTLIBGO_ENCODER_SPEEX_HPP

#include "../encoder.hpp"
#include <ttLibC/encoder/speexEncoder.h>

class SpeexEncoder : public Encoder {
public:
  SpeexEncoder(maps *mp);
  ~SpeexEncoder();
  bool encodeFrame(ttLibC_Frame *cFrame, ttLibGoFrame *goFrame, void *ptr);
private:
  ttLibC_SpeexEncoder *_encoder;
};

#endif
