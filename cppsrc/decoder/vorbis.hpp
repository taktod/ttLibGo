#ifndef TTLIBGO_DECODER_VORBIS_HPP
#define TTLIBGO_DECODER_VORBIS_HPP

#include "../decoder.hpp"

class VorbisDecoder : public Decoder {
public:
  VorbisDecoder(maps *mp);
  ~VorbisDecoder();
  bool decodeFrame(ttLibC_Frame *cFrame, ttLibGoFrame *goFrame, void *ptr);
private:
  void *_decoder;
};

#endif
