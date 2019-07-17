#ifndef TTLIBGO_ENCODER_FDKAAC_HPP
#define TTLIBGO_ENCODER_FDKAAC_HPP

#include "../encoder.hpp"
#include <ttLibC/frame/audio/aac.h>

class FdkaacEncoder : public Encoder {
public:
  FdkaacEncoder(maps *mp);
  ~FdkaacEncoder();
  bool encodeFrame(ttLibC_Frame *cFrame, ttLibGoFrame *goFrame, void *ptr);
  bool setBitrate(uint32_t value);
private:
	void *_encoder;
};

#endif
