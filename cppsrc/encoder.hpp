#ifndef TTLIBGO_ENCODER_HPP
#define TTLIBGO_ENCODER_HPP

#include "util.hpp"
#include <ttLibC/frame/frame.h>
#include "frame.hpp"

class Encoder : public FrameProcessor {
public:
  static Encoder *create(maps *m);
  Encoder();
  virtual ~Encoder();
  virtual bool encodeFrame(
    ttLibC_Frame *cFrame,
    ttLibGoFrame *goFrame,
    void *ptr) = 0;
};

#endif
