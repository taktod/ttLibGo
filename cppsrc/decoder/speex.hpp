#ifndef TTLIBGO_DECODER_SPEEX_HPP
#define TTLIBGO_DECODER_SPEEX_HPP

#include "../decoder.hpp"

class SpeexDecoder : public Decoder {
public:
  SpeexDecoder(maps *mp);
  ~SpeexDecoder();
  bool decodeFrame(ttLibC_Frame *cFrame, ttLibGoFrame *goFrame, void *ptr);
private:
  void *_decoder;
};

#endif
