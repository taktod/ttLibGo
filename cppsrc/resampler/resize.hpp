#ifndef TTLIBGO_RESAMPLER_RESIZE_HPP
#define TTLIBGO_RESAMPLER_RESIZE_HPP

#include "../resampler.hpp"
#include <ttLibC/resampler/imageResizer.h>

class ResizeResampler : public Resampler {
public:
  ResizeResampler(maps *mp);
  ~ResizeResampler();
  bool resampleFrame(ttLibC_Frame *cFrame, ttLibGoFrame *goFrame, void *ptr);
private:
  ttLibC_Video *_image;
  uint32_t _width;
  uint32_t _height;
  bool _isQuick;
};

#endif
