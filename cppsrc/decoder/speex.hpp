#ifndef TTLIBGO_DECODER_SPEEX_HPP
#define TTLIBGO_DECODER_SPEEX_HPP

#include "../decoder.hpp"
#include <ttLibC/decoder/speexDecoder.h>

class SpeexDecoder : public Decoder {
public:
  SpeexDecoder(maps *mp);
  ~SpeexDecoder();
  bool decodeFrame(ttLibC_Frame *cFrame, ttLibGoFrame *goFrame, void *ptr);
private:
  ttLibC_SpeexDecoder *_decoder;
};

#endif
