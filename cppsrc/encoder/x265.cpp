#include "x265.hpp"
#include "../util.hpp"
#include <iostream>
#include <regex>
#include <ttLibC/container/container.h>
#include <ttLibC/encoder/x265Encoder.h>

using namespace std;

extern "C" {
typedef bool (* ttLibC_X265Encoder_getDefaultX265ApiAndParam_func)(void **, void **, const char *, const char *, uint32_t, uint32_t);
typedef int (* ttLibC_X265Encoder_paramParse_func)(void *param, const char *key, const char *value);
typedef int (* ttLibC_X265Encoder_paramApplyProfile_func)(void *param, const char *profile);
typedef void *(* ttLibC_X265Encoder_makeWithX265ApiAndParam_func)(void *, void *);
typedef bool (* ttLibC_X265Encoder_forceNextFrameType_func)(void *, ttLibC_X265Encoder_FrameType);
typedef void *(* ttLibC_codec_func)(void *, void *, ttLibC_getFrameFunc, void *);
typedef void (* ttLibC_close_func)(void **);

extern ttLibC_X265Encoder_getDefaultX265ApiAndParam_func ttLibGo_X265Encoder_getDefaultX265ApiAndParam;
extern ttLibC_X265Encoder_paramParse_func                ttLibGo_X265Encoder_paramParse;
extern ttLibC_X265Encoder_paramApplyProfile_func         ttLibGo_X265Encoder_paramApplyProfile;
extern ttLibC_X265Encoder_makeWithX265ApiAndParam_func   ttLibGo_X265Encoder_makeWithX265ApiAndParam;
extern ttLibC_codec_func                                 ttLibGo_X265Encoder_encode;
extern ttLibC_X265Encoder_forceNextFrameType_func        ttLibGo_X265Encoder_forceNextFrameType;
extern ttLibC_close_func                                 ttLibGo_X265Encoder_close;

extern bool ttLibGoFrameCallback(void *ptr, ttLibC_Frame *frame);
}
X265Encoder::X265Encoder(maps *mp) {
  if(ttLibGo_X265Encoder_getDefaultX265ApiAndParam != nullptr
  && ttLibGo_X265Encoder_paramParse != nullptr
  && ttLibGo_X265Encoder_paramApplyProfile != nullptr
  && ttLibGo_X265Encoder_makeWithX265ApiAndParam != nullptr) {
    uint32_t width = mp->getUint32("width");
    uint32_t height = mp->getUint32("height");
    string preset = mp->getString("preset");
    string tune = mp->getString("tune");
    string profile = mp->getString("profile");
    list<string> params = mp->getStringList("param");
    void *api;
    void *x265Param;
    if(!(*ttLibGo_X265Encoder_getDefaultX265ApiAndParam)(&api, &x265Param, preset.c_str(), tune.c_str(), width, height)) {
      cout << "default api paramの取得に失敗しました。" << endl;
      return;
    }
    regex r(R"([^:]+):(.+)");
    for(auto iter = params.cbegin(); iter != params.cend();++ iter) {
      sregex_iterator reg_iter(iter->begin(), iter->end(), r); 
      string key = reg_iter->str();
      reg_iter ++;
      string value = reg_iter->str();
      int result = (*ttLibGo_X265Encoder_paramParse)(x265Param, key.c_str(), value.c_str());
      if(result < 0) {
        cout << "パラメーター設定失敗:" << key << ":" << value << endl;
      }
    }
    if((*ttLibGo_X265Encoder_paramApplyProfile)(x265Param, profile.c_str())) {
      cout << "profile apply失敗しました。" << endl;
      _encoder = NULL;
    }
    else {
      _encoder = (*ttLibGo_X265Encoder_makeWithX265ApiAndParam)((void *)api, (void *)x265Param);
    }
  }
}

X265Encoder::~X265Encoder() {
  if(ttLibGo_X265Encoder_close != nullptr) {
    (*ttLibGo_X265Encoder_close)(&_encoder);
  }
}
bool X265Encoder::encodeFrame(ttLibC_Frame *cFrame, ttLibGoFrame *goFrame, void *ptr) {
  bool result = false;
  if(ttLibGo_X265Encoder_encode != nullptr) {
    if(cFrame->type != frameType_yuv420) {
      return result;
    }
    update(cFrame, goFrame);
    result = (*ttLibGo_X265Encoder_encode)(_encoder, cFrame, ttLibGoFrameCallback, ptr);
    reset(cFrame, goFrame);
  }
  return result;
}

bool X265Encoder::forceNextFrameType(string type) {
  bool result = false;
  if(ttLibGo_X265Encoder_forceNextFrameType != nullptr) {
    ttLibC_X265Encoder_FrameType frameType = X265FrameType_Auto;
    if(type == "I") {
      frameType = X265FrameType_I;
    }
    else if(type == "P") {
      frameType = X265FrameType_P;
    }
    else if(type == "B") {
      frameType = X265FrameType_B;
    }
    else if(type == "IDR") {
      frameType = X265FrameType_IDR;
    }
    else if(type == "Auto") {
      frameType = X265FrameType_Auto;
    }
    else if(type == "Bref") {
      frameType = X265FrameType_Bref;
    }
    else {
      return false;
    }
    result = (*ttLibGo_X265Encoder_forceNextFrameType)(_encoder, frameType);
  }
  return result;
}

extern "C" {
bool X265Encoder_forceNextFrameType(void *encoder, const char *type) {
  X265Encoder *x = reinterpret_cast<X265Encoder *>(encoder);
  return x->forceNextFrameType(type);
}
}