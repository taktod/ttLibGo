#ifndef TTLIBGO_RESAMPLER_SWRESAMPLE_HPP
#define TTLIBGO_RESAMPLER_SWRESAMPLE_HPP

#include "../resampler.hpp"
#include <ttLibC/resampler/swresampleResampler.h>

class SwresampleResampler : public Resampler {
public:
  SwresampleResampler(maps *mp);
  ~SwresampleResampler();
  bool resampleFrame(ttLibC_Frame *cFrame, ttLibGoFrame *goFrame, void *ptr);
private:
  uint32_t getSubType(ttLibC_Frame_Type type, string name);
  ttLibC_SwresampleResampler *_resampler;
};

#endif
