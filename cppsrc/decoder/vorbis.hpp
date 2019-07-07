#ifndef TTLIBGO_DECODER_VORBIS_HPP
#define TTLIBGO_DECODER_VORBIS_HPP

#include "../decoder.hpp"
#include <ttLibC/decoder/vorbisDecoder.h>

class VorbisDecoder : public Decoder {
public:
  VorbisDecoder(maps *mp);
  ~VorbisDecoder();
  bool decodeFrame(ttLibC_Frame *cFrame, ttLibGoFrame *goFrame, void *ptr);
private:
  ttLibC_VorbisDecoder *_decoder;
};

#endif
