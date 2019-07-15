#ifndef TTLIBGO_RESAMPLER_AUDIO_HPP
#define TTLIBGO_RESAMPLER_AUDIO_HPP

#include "../resampler.hpp"

class AudioResampler : public Resampler {
public:
  AudioResampler(maps *mp);
  ~AudioResampler();
  bool resampleFrame(ttLibC_Frame *cFrame, ttLibGoFrame *goFrame, void *ptr);
private:
  uint32_t getSubType(ttLibC_Frame_Type frameType, string name);
  ttLibC_Frame_Type _targetType;
  uint32_t _subType;
  ttLibC_Audio *_pcm;
};

#endif
