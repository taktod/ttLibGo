#include "x264.hpp"
#include "../util.hpp"
#include <iostream>
#include <regex>
#include <ttLibC/container/container.h>
#include <ttLibC/encoder/x264Encoder.h>

using namespace std;

extern "C" {
typedef void (* ttLibC_X264Encoder_getDefaultX264ParamTWithPresetTune_func)(void *, uint32_t, uint32_t, const char *, const char *);
typedef int (* ttLibC_X264Encoder_paramParse_func)(void *, const char *, const char *);
typedef int (* ttLibC_X264Encoder_paramApplyProfile_func)(void *, const char *);
typedef void *(* ttLibC_X264Encoder_makeWithX264ParamT_func)(void *);
typedef bool (* ttLibC_X264Encoder_forceNextFrameType_func)(void *, ttLibC_X264Encoder_FrameType);
typedef void *(* ttLibC_codec_func)(void *, void *, ttLibC_getFrameFunc, void *);
typedef void (* ttLibC_close_func)(void **);

extern ttLibC_X264Encoder_getDefaultX264ParamTWithPresetTune_func ttLibGo_X264Encoder_getDefaultX264ParamTWithPresetTune;
extern ttLibC_X264Encoder_paramParse_func                         ttLibGo_X264Encoder_paramParse;
extern ttLibC_X264Encoder_paramApplyProfile_func                  ttLibGo_X264Encoder_paramApplyProfile;
extern ttLibC_X264Encoder_makeWithX264ParamT_func                 ttLibGo_X264Encoder_makeWithX264ParamT;
extern ttLibC_codec_func                                          ttLibGo_X264Encoder_encode;
extern ttLibC_X264Encoder_forceNextFrameType_func                 ttLibGo_X264Encoder_forceNextFrameType;
extern ttLibC_close_func                                          ttLibGo_X264Encoder_close;

extern bool ttLibGoFrameCallback(void *ptr, ttLibC_Frame *frame);
}
X264Encoder::X264Encoder(maps *mp) {
  if(ttLibGo_X264Encoder_getDefaultX264ParamTWithPresetTune != nullptr
  && ttLibGo_X264Encoder_paramParse != nullptr
  && ttLibGo_X264Encoder_paramApplyProfile != nullptr
  && ttLibGo_X264Encoder_makeWithX264ParamT != nullptr) {
    uint32_t width = mp->getUint32("width");
    uint32_t height = mp->getUint32("height");
    string preset = mp->getString("preset");
    string tune = mp->getString("tune");
    string profile = mp->getString("profile");
    list<string> params = mp->getStringList("param");
    //x264_param_t x264Param; // 936 byteらしい・・・
    char x264Param[4096];
    (*ttLibGo_X264Encoder_getDefaultX264ParamTWithPresetTune)(
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
      int result = (*ttLibGo_X264Encoder_paramParse)(&x264Param, key.c_str(), value.c_str());
      if(result < 0) {
        cout << "パラメーター設定失敗:" << key << ":" << value << endl;
      }
    }
    if((*ttLibGo_X264Encoder_paramApplyProfile)(&x264Param, profile.c_str())) {
      cout << "profile apply失敗しました。" << endl;
      _encoder = NULL;
    }
    else {
      _encoder = (*ttLibGo_X264Encoder_makeWithX264ParamT)(&x264Param);
    }
  }
}

X264Encoder::~X264Encoder() {
  if(ttLibGo_X264Encoder_close != nullptr) {
    (*ttLibGo_X264Encoder_close)(&_encoder);
  }
}
bool X264Encoder::encodeFrame(ttLibC_Frame *cFrame, ttLibGoFrame *goFrame, void *ptr) {
  bool result = false;
  if(ttLibGo_X264Encoder_encode != nullptr) {
    if(cFrame->type != frameType_yuv420) {
      return result;
    }
    update(cFrame, goFrame);
    result = (*ttLibGo_X264Encoder_encode)(_encoder, cFrame, ttLibGoFrameCallback, ptr);    
    reset(cFrame, goFrame);
  }
  return result;
}

bool X264Encoder::forceNextFrameType(string type) {
  bool result = false;
  if(ttLibGo_X264Encoder_forceNextFrameType != nullptr) {
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
    result = (* ttLibGo_X264Encoder_forceNextFrameType)(_encoder, frameType);
  }
  return result;
}

extern "C" {
bool X264Encoder_forceNextFrameType(void *encoder, const char *type) {
  X264Encoder *x = reinterpret_cast<X264Encoder *>(encoder);
  return x->forceNextFrameType(type);
}
}