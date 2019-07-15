#include "avcodec.hpp"
#include "../util.hpp"
#include <iostream>
#include <ttLibC/container/container.h>

using namespace std;

extern "C" {
typedef void *(* ttLibC_AvcodecDecoder_make_func)(ttLibC_Frame_Type, uint32_t, uint32_t);
typedef void *(* ttLibC_codec_func)(void *, void *, ttLibC_getFrameFunc, void *);
typedef void (* ttLibC_close_func)(void **);
typedef bool (* ttLibC_isType_func)(ttLibC_Frame_Type);

extern ttLibC_AvcodecDecoder_make_func ttLibGo_AvcodecVideoDecoder_make;
extern ttLibC_AvcodecDecoder_make_func ttLibGo_AvcodecAudioDecoder_make;
extern ttLibC_codec_func ttLibGo_AvcodecDecoder_decode;
extern ttLibC_close_func ttLibGo_AvcodecDecoder_close;
extern ttLibC_isType_func ttLibGo_isAudio;
extern ttLibC_isType_func ttLibGo_isVideo;

extern bool ttLibGoFrameCallback(void *ptr, ttLibC_Frame *frame);
}
AvcodecDecoder::AvcodecDecoder(maps *mp) {
  if(ttLibGo_AvcodecVideoDecoder_make != nullptr && ttLibGo_AvcodecAudioDecoder_make != nullptr
  && ttLibGo_isAudio != nullptr && ttLibGo_isVideo != nullptr) {
    ttLibC_Frame_Type type = Frame_getFrameTypeFromString(mp->getString("frameType"));
    if((*ttLibGo_isAudio)(type)) {
      _decoder = (*ttLibGo_AvcodecAudioDecoder_make)(
        type,
        mp->getUint32("sampleRate"),
        mp->getUint32("channelNum"));
    }
    else if((*ttLibGo_isVideo)(type)) {
      _decoder = (*ttLibGo_AvcodecVideoDecoder_make)(
        type,
        mp->getUint32("width"),
        mp->getUint32("height"));
    }
  }
}

AvcodecDecoder::~AvcodecDecoder() {
  if(ttLibGo_AvcodecDecoder_close != nullptr) {
    (*ttLibGo_AvcodecDecoder_close)(&_decoder);
  }
}
bool AvcodecDecoder::decodeFrame(ttLibC_Frame *cFrame, ttLibGoFrame *goFrame, void *ptr) {
  bool result = false;
  if(ttLibGo_AvcodecDecoder_decode != nullptr) {
    update(cFrame, goFrame);
    result = (*ttLibGo_AvcodecDecoder_decode)(_decoder, cFrame, ttLibGoFrameCallback, ptr);
    reset(cFrame, goFrame);
  }
  return result;
}
