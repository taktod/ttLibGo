#ifndef TTLIBGO_DECODER_PNG_HPP
#define TTLIBGO_DECODER_PNG_HPP

#include "../decoder.hpp"
#include <ttLibC/decoder/pngDecoder.h>

class PngDecoder : public Decoder {
public:
  PngDecoder(maps *mp);
  ~PngDecoder();
  bool decodeFrame(ttLibC_Frame *cFrame, ttLibGoFrame *goFrame, void *ptr);
private:
  ttLibC_PngDecoder *_decoder;
};

#endif
