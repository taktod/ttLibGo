#include "audio.hpp"
#include "../util.hpp"
#include "../frame.hpp"
#include <iostream>

using namespace std;

extern "C" {
extern bool ttLibGoFrameCallback(void *ptr, ttLibC_Frame *frame);
}

uint32_t AudioResampler::getSubType(ttLibC_Frame_Type frameType, string name) {
  switch(frameType) {
  case frameType_pcmS16:
    if(name == "bigEndian") {
      return PcmS16Type_bigEndian;
    }
    else if(name == "bigEndianPlanar"){
      return PcmS16Type_bigEndian_planar;
    }
    else if(name == "littleEndian"){
      return PcmS16Type_littleEndian;
    }
    else if(name == "littleEndianPlanar"){
      return PcmS16Type_littleEndian_planar;
    }
    break;
  case frameType_pcmF32:
    if(name == "interleave") {
      return PcmF32Type_interleave;
    }
    else if(name == "planar") {
      return PcmF32Type_planar;
    }
    break;
  default:
    break;
  }
  return 99;
}

AudioResampler::AudioResampler(maps *mp) {
  _targetType = Frame_getFrameTypeFromString(mp->getString("frameType"));
  _subType = getSubType(_targetType, mp->getString("subType"));
  _pcm = nullptr;
}

AudioResampler::~AudioResampler() {
  ttLibC_Audio_close(&_pcm);
}
bool AudioResampler::resampleFrame(ttLibC_Frame *cFrame, ttLibGoFrame *goFrame, void *ptr) {
  switch(cFrame->type) {
  case frameType_pcmF32:
  case frameType_pcmS16:
    break;
  default:
    return false;
  }
  bool result = false;
  update(cFrame, goFrame);
  ttLibC_Audio *pcm = (ttLibC_Audio *)cFrame;
  ttLibC_Audio *resampled = ttLibC_AudioResampler_convertFormat(_pcm, _targetType, _subType, pcm->channel_num, pcm);
  if(resampled != nullptr) {
    _pcm = resampled;
    result = ttLibGoFrameCallback(ptr, (ttLibC_Frame *)_pcm);
  }
  reset(cFrame, goFrame);
  return result;
}
