#ifndef TTLIBGO_DECODER_AVCODEC_HPP
#define TTLIBGO_DECODER_AVCODEC_HPP

#include "../decoder.hpp"
#include <ttLibC/decoder/avcodecDecoder.h>

class AvcodecDecoder : public Decoder {
public:
  AvcodecDecoder(maps *mp);
  ~AvcodecDecoder();
  bool decodeFrame(ttLibC_Frame *cFrame, ttLibGoFrame *goFrame, void *ptr);
private:
  ttLibC_AvcodecDecoder *_decoder;
};

#endif
