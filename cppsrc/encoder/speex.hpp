#ifndef TTLIBGO_ENCODER_SPEEX_HPP
#define TTLIBGO_ENCODER_SPEEX_HPP

#include "../encoder.hpp"

class SpeexEncoder : public Encoder {
public:
  SpeexEncoder(maps *mp);
  ~SpeexEncoder();
  bool encodeFrame(ttLibC_Frame *cFrame, ttLibGoFrame *goFrame, void *ptr);
private:
  void *_encoder;
};

#endif
