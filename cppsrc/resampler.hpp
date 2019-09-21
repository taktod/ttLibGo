#ifndef TTLIBGO_RESAMPLER_HPP
#define TTLIBGO_RESAMPLER_HPP

#include "util.hpp"
#include "ttLibC/ttLibC/frame/frame.h"
#include "frame.hpp"

class Resampler : public FrameProcessor {
public:
  static Resampler *create(maps *m);
  Resampler();
  virtual ~Resampler();
  bool resampleFrame(
    ttLibC_Frame *cFrame,
    ttLibGoFrame *goFrame,
    void *ptr);
};

#endif
