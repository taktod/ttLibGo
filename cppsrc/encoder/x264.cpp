#include "x264.hpp"
#include "../util.hpp"
#include <iostream>
#include <regex>

#ifdef __ENABLE_X264__
# include <x264.h>
#endif

using namespace std;

extern "C" {
extern bool ttLibGoFrameCallback(void *ptr, ttLibC_Frame *frame);
}
X264Encoder::X264Encoder(maps *mp) {
#ifdef __ENABLE_X264__
  uint32_t width = mp->getUint32("width");
  uint32_t height = mp->getUint32("height");
  string preset = mp->getString("preset");
  string tune = mp->getString("tune");
  string profile = mp->getString("profile");
  list<string> params = mp->getStringList("param");
//  _encoder = ttLibC_X264Encoder_make(mp->getUint32("width"), mp->getUint32("height"));
  x264_param_t x264Param;
  ttLibC_X264Encoder_getDefaultX264ParamTWithPresetTune(
    &x264Param,
    width,
    height,
    preset.c_str(),
    tune.c_str());
  regex r(R"([^:]+):(.+)");
  for(auto iter = params.cbegin(); iter != params.cend();++ iter) {
    sregex_iterator reg_iter(iter->begin(), iter->end(), r); 
    string key = reg_iter->str();
    reg_iter ++;
    string value = reg_iter->str();
    int result = x264_param_parse(&x264Param, key.c_str(), value.c_str());
    if(result < 0) {
      cout << "パラメーター設定失敗:" << key << ":" << value << endl;
    }
  }
  if(x264_param_apply_profile(&x264Param, profile.c_str())) {
    cout << "profile apply失敗しました。" << endl;
    _encoder = NULL;
  }
  else {
    _encoder = ttLibC_X264Encoder_makeWithX264ParamT(&x264Param);
  }
#endif
}

X264Encoder::~X264Encoder() {
#ifdef __ENABLE_X264__
  ttLibC_X264Encoder_close(&_encoder);
#endif
}
bool X264Encoder::encodeFrame(ttLibC_Frame *cFrame, ttLibGoFrame *goFrame, void *ptr) {
  bool result = false;
#ifdef __ENABLE_X264__
  if(cFrame->type != frameType_yuv420) {
    return result;
  }
  update(cFrame, goFrame);
  result = ttLibC_X264Encoder_encode(_encoder, (ttLibC_Yuv420 *)cFrame, [](void *ptr, ttLibC_H264 *h264)->bool {
    return ttLibGoFrameCallback(ptr, (ttLibC_Frame *)h264);
  }, ptr);
  reset(cFrame, goFrame);
#endif
  return result;
}

bool X264Encoder::forceNextFrameType(string type) {
  bool result = false;
#ifdef __ENABLE_X264__
  ttLibC_X264Encoder_FrameType frameType = X264FrameType_Auto;
  if(type == "I") {
    frameType = X264FrameType_I;
  }
  else if(type == "P") {
    frameType = X264FrameType_P;
  }
  else if(type == "B") {
    frameType = X264FrameType_B;
  }
  else if(type == "IDR") {
    frameType = X264FrameType_IDR;
  }
  else if(type == "Auto") {
    frameType = X264FrameType_Auto;
  }
  else if(type == "Bref") {
    frameType = X264FrameType_Bref;
  }
  else if(type == "KeyFrame") {
    frameType = X264FrameType_KeyFrame;
  }
  else {
    return false;
  }
  result = ttLibC_X264Encoder_forceNextFrameType(_encoder, frameType);
#endif
  return result;
}

extern "C" {
bool X264Encoder_forceNextFrameType(void *encoder, const char *type) {
  X264Encoder *x = reinterpret_cast<X264Encoder *>(encoder);
  return x->forceNextFrameType(type);
}
}