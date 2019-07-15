#ifndef TTLIBGO_ENCODER_JPEG_HPP
#define TTLIBGO_ENCODER_JPEG_HPP

#include "../encoder.hpp"

class JpegEncoder : public Encoder {
public:
  JpegEncoder(maps *mp);
  ~JpegEncoder();
  bool encodeFrame(ttLibC_Frame *cFrame, ttLibGoFrame *goFrame, void *ptr);
  bool setQuality(uint32_t quality);
private:
  void *_encoder;
};

#endif
