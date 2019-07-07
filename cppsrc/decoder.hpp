#ifndef TTLIBGO_DECODER_HPP
#define TTLIBGO_DECODER_HPP

#include "util.hpp"
#include <ttLibC/frame/frame.h>
#include "frame.hpp"

class Decoder : public FrameProcessor {
public:
  static Decoder *create(maps *m);
  Decoder();
  virtual ~Decoder();
  virtual bool decodeFrame(
    ttLibC_Frame *cFrame,
    ttLibGoFrame *goFrame,
    void *ptr) = 0;
};

#endif
