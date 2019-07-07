#ifndef TTLIBGO_ENCODER_VORBIS_HPP
#define TTLIBGO_ENCODER_VORBIS_HPP

#include "../encoder.hpp"
#include <ttLibC/encoder/vorbisEncoder.h>

class VorbisEncoder : public Encoder {
public:
  VorbisEncoder(maps *mp);
  ~VorbisEncoder();
  bool encodeFrame(ttLibC_Frame *cFrame, ttLibGoFrame *goFrame, void *ptr);
private:
  ttLibC_VorbisEncoder *_encoder;
};

#endif
