#ifndef TTLIBGO_DECODER_THEORA_HPP
#define TTLIBGO_DECODER_THEORA_HPP

#include "../decoder.hpp"
#include <ttLibC/decoder/theoraDecoder.h>

class TheoraDecoder : public Decoder {
public:
  TheoraDecoder(maps *mp);
  ~TheoraDecoder();
  bool decodeFrame(ttLibC_Frame *cFrame, ttLibGoFrame *goFrame, void *ptr);
private:
  ttLibC_TheoraDecoder *_decoder;
};

#endif
