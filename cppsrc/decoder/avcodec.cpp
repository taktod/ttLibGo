#include "avcodec.hpp"
#include "../util.hpp"
#include <iostream>

using namespace std;

extern "C" {
extern bool ttLibGoFrameCallback(void *ptr, ttLibC_Frame *frame);
}
AvcodecDecoder::AvcodecDecoder(maps *mp) {
#ifdef __ENABLE_AVCODEC__
  ttLibC_Frame_Type type = Frame_getFrameTypeFromString(mp->getString("frameType"));
  if(ttLibC_isAudio(type)) {
    _decoder = ttLibC_AvcodecAudioDecoder_make(
      type,
      mp->getUint32("sampleRate"),
      mp->getUint32("channelNum"));
  }
  else if(ttLibC_isVideo(type)) {
    _decoder = ttLibC_AvcodecVideoDecoder_make(
      type,
      mp->getUint32("width"),
      mp->getUint32("height"));
  }
#endif
}

AvcodecDecoder::~AvcodecDecoder() {
#ifdef __ENABLE_AVCODEC__
  ttLibC_AvcodecDecoder_close(&_decoder);
#endif
}
bool AvcodecDecoder::decodeFrame(ttLibC_Frame *cFrame, ttLibGoFrame *goFrame, void *ptr) {
  bool result = false;
#ifdef __ENABLE_AVCODEC__
  update(cFrame, goFrame);
  result = ttLibC_AvcodecDecoder_decode(_decoder, cFrame, ttLibGoFrameCallback, ptr);
  reset(cFrame, goFrame);
#endif
  return result;
}
