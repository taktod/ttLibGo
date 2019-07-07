#ifndef TTLIBGO_DECODER_OPUS_HPP
#define TTLIBGO_DECODER_OPUS_HPP

#include "../decoder.hpp"
#include <ttLibC/decoder/opusDecoder.h>

class tOpusDecoder : public Decoder {
public:
  tOpusDecoder(maps *mp);
  ~tOpusDecoder();
  bool decodeFrame(ttLibC_Frame *cFrame, ttLibGoFrame *goFrame, void *ptr);
  int codecControl(string control, int value);
private:
  ttLibC_OpusDecoder *_decoder;
};

#endif
