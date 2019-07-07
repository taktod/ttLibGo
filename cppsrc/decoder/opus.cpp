#include "opus.hpp"
#include "../util.hpp"
#include <iostream>

#ifdef __ENABLE_OPUS__
# include <opus/opus.h>
#endif

using namespace std;

extern "C" {
extern bool ttLibGoFrameCallback(void *ptr, ttLibC_Frame *frame);
}
tOpusDecoder::tOpusDecoder(maps *mp) {
#ifdef __ENABLE_OPUS__
  _decoder = ttLibC_OpusDecoder_make(mp->getUint32("sampleRate"), mp->getUint32("channelNum"));
#endif
}
tOpusDecoder::~tOpusDecoder() {
#ifdef __ENABLE_OPUS__
  ttLibC_OpusDecoder_close(&_decoder);
#endif
}
bool tOpusDecoder::decodeFrame(ttLibC_Frame *cFrame, ttLibGoFrame *goFrame, void *ptr) {
  bool result = false;
#ifdef __ENABLE_OPUS__
  update(cFrame, goFrame);
  result = ttLibC_OpusDecoder_decode(_decoder, (ttLibC_Opus *)cFrame, [](void *ptr, ttLibC_PcmS16 *pcm) -> bool{
    return ttLibGoFrameCallback(ptr, (ttLibC_Frame *)pcm);
  }, ptr);
  reset(cFrame, goFrame);
#endif
  return result;
}

int tOpusDecoder::codecControl(string control, int value) {
#ifdef __ENABLE_OPUS__
  if(_decoder == NULL) {
    puts("decoderが準備されていません。");
    return -1;
  }
  OpusDecoder *nativeDecoder = (OpusDecoder *)ttLibC_OpusDecoder_refNativeDecoder(_decoder);
#define OpusSetCtl(a, b) if(control == #a) { \
    return opus_decoder_ctl(nativeDecoder, a(b));\
  }
#define OpusCtl(a) if(control == #a) { \
    return opus_decoder_ctl(nativeDecoder, a);\
  }
#define OpusGetCtl(a) if(control == #a) { \
    int result, res; \
    res = opus_decoder_ctl(nativeDecoder, a(&result)); \
    if(res != OPUS_OK) {return res;} \
    return result; \
  }
#define OpusGetUCtl(a) if(control == #a) { \
    uint32_t result; \
    int res; \
    res = opus_decoder_ctl(nativeDecoder, a(&result)); \
    if(res != OPUS_OK) {return res;} \
    return result; \
  }
  OpusCtl(OPUS_RESET_STATE)
  OpusGetUCtl(OPUS_GET_FINAL_RANGE)
  OpusGetCtl(OPUS_GET_BANDWIDTH)
  OpusGetCtl(OPUS_GET_SAMPLE_RATE)
  OpusSetCtl(OPUS_SET_PHASE_INVERSION_DISABLED, value)
  OpusGetCtl(OPUS_GET_PHASE_INVERSION_DISABLED)
  OpusSetCtl(OPUS_SET_GAIN, value)
  OpusGetCtl(OPUS_GET_GAIN)
  OpusGetCtl(OPUS_GET_LAST_PACKET_DURATION)
  OpusGetCtl(OPUS_GET_PITCH)
#undef OpusSetCtl
#undef OpusCtl
#undef OpusGetCtl
#undef OpusGetUCtl
  return OPUS_OK;
#else
  return -1;
#endif
}

extern "C" {
int OpusDecoder_codecControl(void *decoder, const char *key, int value) {
  tOpusDecoder *d = reinterpret_cast<tOpusDecoder *>(decoder);
  return d->codecControl(key, value);
}
}