#ifndef TTLIBGO_ENCODER_OPUS_HPP
#define TTLIBGO_ENCODER_OPUS_HPP

#include "../encoder.hpp"

class tOpusEncoder : public Encoder {
public:
  tOpusEncoder(maps *mp);
  ~tOpusEncoder();
  bool encodeFrame(ttLibC_Frame *cFrame, ttLibGoFrame *goFrame, void *ptr);
  int codecControl(string control, int value);
private:
  void *_encoder;
};

#endif
