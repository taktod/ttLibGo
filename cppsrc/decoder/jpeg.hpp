#ifndef TTLIBGO_DECODER_JPEG_HPP
#define TTLIBGO_DECODER_JPEG_HPP

#include "../decoder.hpp"
#include <ttLibC/decoder/jpegDecoder.h>

class JpegDecoder : public Decoder {
public:
  JpegDecoder(maps *mp);
  ~JpegDecoder();
  bool decodeFrame(ttLibC_Frame *cFrame, ttLibGoFrame *goFrame, void *ptr);
private:
  ttLibC_JpegDecoder *_decoder;
};

#endif
