#ifndef TTLIBGO_RESAMPLER_SOUNDTOUCH_HPP
#define TTLIBGO_RESAMPLER_SOUNDTOUCH_HPP

#include "../resampler.hpp"

class SoundtouchResampler : public Resampler {
public:
  SoundtouchResampler(maps *mp);
  ~SoundtouchResampler();
  bool resampleFrame(ttLibC_Frame *cFrame, ttLibGoFrame *goFrame, void *ptr);
  // ここに追加でいろんな処理を入れておく必要があるわけか・・・
  void setRate(double newRate);
  void setTempo(double newTempo);
  void setRateChange(double newRate);
  void setTempoChange(double newTempo);
  void setPitch(double newPitch);
  void setPitchOctaves(double newPitch);
  void setPitchSemiTones(double newPitch);
private:
  void *_resampler;
};

#endif
