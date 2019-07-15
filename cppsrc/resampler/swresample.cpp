#include "swresample.hpp"
#include "../util.hpp"
#include <iostream>

#include <ttLibC/frame/audio/pcms16.h>
#include <ttLibC/frame/audio/pcmf32.h>

using namespace std;

extern "C" {
typedef void *(* ttLibC_SwresampleResampler_make_func)(ttLibC_Frame_Type, uint32_t, uint32_t, uint32_t, ttLibC_Frame_Type, uint32_t, uint32_t, uint32_t);
typedef void *(* ttLibC_codec_func)(void *, void *, ttLibC_getFrameFunc, void *);
typedef void (* ttLibC_close_func)(void **);

extern ttLibC_SwresampleResampler_make_func ttLibGo_SwresampleResampler_make;
extern ttLibC_codec_func                    ttLibGo_SwresampleResampler_resample;
extern ttLibC_close_func                    ttLibGo_SwresampleResampler_close;

extern bool ttLibGoFrameCallback(void *ptr, ttLibC_Frame *frame);
}

uint32_t SwresampleResampler::getSubType(ttLibC_Frame_Type type, string name) {
  switch(type) {
  case frameType_pcmS16:
    if(name == "bigEndian") {return PcmS16Type_bigEndian;}
    if(name == "bigEndianPlanar") {return PcmS16Type_bigEndian_planar;}
    if(name == "littleEndian") {return PcmS16Type_littleEndian;}
    if(name == "littleEndianPlanar") {return PcmS16Type_littleEndian_planar;}
    break;
  case frameType_pcmF32:
    if(name == "interleave") {return PcmF32Type_interleave;}
    if(name == "planar") {return PcmF32Type_planar;}
    break;
  default:
    break;
  }
  return 99;
}

SwresampleResampler::SwresampleResampler(maps *mp) {
  if(ttLibGo_SwresampleResampler_make != nullptr) {
    ttLibC_Frame_Type inType = Frame_getFrameTypeFromString(mp->getString("inType"));
    uint32_t inSubType = getSubType(inType, mp->getString("inSubType"));
    uint32_t inSampleRate = mp->getUint32("inSampleRate");
    uint32_t inChannelNum = mp->getUint32("inChannelNum");
    ttLibC_Frame_Type outType = Frame_getFrameTypeFromString(mp->getString("outType"));
    uint32_t outSubType = getSubType(outType, mp->getString("outSubType"));
    uint32_t outSampleRate = mp->getUint32("outSampleRate");
    uint32_t outChannelNum = mp->getUint32("outChannelNum");
    _resampler = (*ttLibGo_SwresampleResampler_make)(
      inType, inSubType, inSampleRate, inChannelNum,
      outType, outSubType, outSampleRate, outChannelNum);
  }
}

SwresampleResampler::~SwresampleResampler() {
  if(ttLibGo_SwresampleResampler_close != nullptr) {
    (*ttLibGo_SwresampleResampler_close)(&_resampler);
  }
}

bool SwresampleResampler::resampleFrame(ttLibC_Frame *cFrame, ttLibGoFrame *goFrame, void *ptr) {
  bool result = false;
  if(ttLibGo_SwresampleResampler_resample != nullptr) {
    update(cFrame, goFrame);
    result = (*ttLibGo_SwresampleResampler_resample)(
      _resampler,
      cFrame,
      ttLibGoFrameCallback,
      ptr);
    reset(cFrame, goFrame);
  }
  return result;
}