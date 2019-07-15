#ifndef TTLIBGO_DECODER_JPEG_HPP
#define TTLIBGO_DECODER_JPEG_HPP

#include "../decoder.hpp"

class JpegDecoder : public Decoder {
public:
  JpegDecoder(maps *mp);
  ~JpegDecoder();
  bool decodeFrame(ttLibC_Frame *cFrame, ttLibGoFrame *goFrame, void *ptr);
private:
  void *_decoder;
};

#endif
