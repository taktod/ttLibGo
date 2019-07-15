#ifndef TTLIBGO_DECODER_PNG_HPP
#define TTLIBGO_DECODER_PNG_HPP

#include "../decoder.hpp"

class PngDecoder : public Decoder {
public:
  PngDecoder(maps *mp);
  ~PngDecoder();
  bool decodeFrame(ttLibC_Frame *cFrame, ttLibGoFrame *goFrame, void *ptr);
private:
  void *_decoder;
};

#endif
