#include "opus.hpp"
#include "../util.hpp"
#include <iostream>

using namespace std;

extern "C" {
extern bool ttLibGoFrameCallback(void *ptr, ttLibC_Frame *frame);
}
tOpusEncoder::tOpusEncoder(maps *mp) {
#ifdef __ENABLE_OPUS__
  _encoder = ttLibC_OpusEncoder_make(mp->getUint32("sampleRate"), mp->getUint32("channelNum"), mp->getUint32("unitSampleNum"));
#endif
}

tOpusEncoder::~tOpusEncoder() {
#ifdef __ENABLE_OPUS__
  ttLibC_OpusEncoder_close(&_encoder);
#endif
}
bool tOpusEncoder::encodeFrame(ttLibC_Frame *cFrame, ttLibGoFrame *goFrame, void *ptr) {
  bool result = false;
#ifdef __ENABLE_OPUS__
  if(cFrame->type != frameType_pcmS16) {
    return result;
  }
  update(cFrame, goFrame);
  result = ttLibC_OpusEncoder_encode(_encoder, (ttLibC_PcmS16 *)cFrame, [](void *ptr, ttLibC_Opus *opus)->bool {
    return ttLibGoFrameCallback(ptr, (ttLibC_Frame *)opus);
  }, ptr);
  reset(cFrame, goFrame);
#endif
  return result;
}

int tOpusEncoder::codecControl(string control, int value) {
#ifdef __ENABLE_OPUS__
  if(_encoder == NULL) {
    puts("encoderが準備されていません。");
    return -1;
  }
  OpusEncoder *nativeEncoder = (OpusEncoder *)ttLibC_OpusEncoder_refNativeEncoder(_encoder);
#define OpusSetCtl(a, b) if(control == #a) { \
    return opus_encoder_ctl(nativeEncoder, a(b));\
  }
#define OpusCtl(a) if(control == #a) { \
    return opus_encoder_ctl(nativeEncoder, a);\
  }
#define OpusGetCtl(a) if(control == #a) { \
    int result, res; \
    res = opus_encoder_ctl(nativeEncoder, a(&result)); \
    if(res != OPUS_OK) {return res;} \
    return result; \
  }
#define OpusGetUCtl(a) if(control == #a) { \
    uint32_t result; \
    int res; \
    res = opus_encoder_ctl(nativeEncoder, a(&result)); \
    if(res != OPUS_OK) {return res;} \
    return result; \
  }
  OpusSetCtl(OPUS_SET_COMPLEXITY, value)
  OpusGetCtl(OPUS_GET_COMPLEXITY)
  OpusSetCtl(OPUS_SET_BITRATE, value)
  OpusGetCtl(OPUS_GET_BITRATE)
  OpusSetCtl(OPUS_SET_VBR, value)
  OpusGetCtl(OPUS_GET_VBR)
  OpusSetCtl(OPUS_SET_VBR_CONSTRAINT, value)
  OpusGetCtl(OPUS_GET_VBR_CONSTRAINT)
  OpusSetCtl(OPUS_SET_FORCE_CHANNELS, value)
  OpusGetCtl(OPUS_GET_FORCE_CHANNELS)
  OpusSetCtl(OPUS_SET_MAX_BANDWIDTH, value)
  OpusGetCtl(OPUS_GET_MAX_BANDWIDTH)
  OpusSetCtl(OPUS_SET_BANDWIDTH, value)
  OpusSetCtl(OPUS_SET_SIGNAL, value)
  OpusGetCtl(OPUS_GET_SIGNAL)
  OpusSetCtl(OPUS_SET_APPLICATION, value)
  OpusGetCtl(OPUS_GET_APPLICATION)
  OpusGetCtl(OPUS_GET_LOOKAHEAD)
  OpusSetCtl(OPUS_SET_INBAND_FEC, value)
  OpusGetCtl(OPUS_GET_INBAND_FEC)
  OpusSetCtl(OPUS_SET_PACKET_LOSS_PERC, value)
  OpusGetCtl(OPUS_GET_PACKET_LOSS_PERC)
  OpusSetCtl(OPUS_SET_DTX, value)
  OpusGetCtl(OPUS_GET_DTX)
  OpusSetCtl(OPUS_SET_LSB_DEPTH, value)
  OpusGetCtl(OPUS_GET_LSB_DEPTH)
  OpusSetCtl(OPUS_SET_EXPERT_FRAME_DURATION, value)
  OpusGetCtl(OPUS_GET_EXPERT_FRAME_DURATION)
  OpusSetCtl(OPUS_SET_PREDICTION_DISABLED, value)
  OpusGetCtl(OPUS_GET_PREDICTION_DISABLED)
  OpusCtl(OPUS_RESET_STATE)
  OpusGetUCtl(OPUS_GET_FINAL_RANGE)
  OpusGetCtl(OPUS_GET_BANDWIDTH)
  OpusGetCtl(OPUS_GET_SAMPLE_RATE)
  OpusSetCtl(OPUS_SET_PHASE_INVERSION_DISABLED, value)
  OpusGetCtl(OPUS_GET_PHASE_INVERSION_DISABLED)
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
int OpusEncoder_codecControl(void *encoder, const char *key, int value) {
  tOpusEncoder *e = reinterpret_cast<tOpusEncoder *>(encoder);
  return e->codecControl(key, value);
}
}