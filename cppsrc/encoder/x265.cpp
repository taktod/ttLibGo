#include "x265.hpp"
#include "../util.hpp"
#include <iostream>
#include <regex>

#ifdef __ENABLE_X265__
# include <x265.h>
#endif

using namespace std;

extern "C" {
extern bool ttLibGoFrameCallback(void *ptr, ttLibC_Frame *frame);
}
X265Encoder::X265Encoder(maps *mp) {
#ifdef __ENABLE_X265__
  uint32_t width = mp->getUint32("width");
  uint32_t height = mp->getUint32("height");
  string preset = mp->getString("preset");
  string tune = mp->getString("tune");
  string profile = mp->getString("profile");
  list<string> params = mp->getStringList("param");
  x265_api *api;
  x265_param *x265Param;
  if(!ttLibC_X265Encoder_getDefaultX265ApiAndParam((void **)&api, (void **)&x265Param, preset.c_str(), tune.c_str(), width, height)) {
    cout << "default api paramの取得に失敗しました。" << endl;
    return;
  }
  regex r(R"([^:]+):(.+)");
  for(auto iter = params.cbegin(); iter != params.cend();++ iter) {
    sregex_iterator reg_iter(iter->begin(), iter->end(), r); 
    string key = reg_iter->str();
    reg_iter ++;
    string value = reg_iter->str();
    int result = x265_param_parse(x265Param, key.c_str(), value.c_str());
    if(result < 0) {
      cout << "パラメーター設定失敗:" << key << ":" << value << endl;
    }
  }
  if(x265_param_apply_profile(x265Param, profile.c_str())) {
    cout << "profile apply失敗しました。" << endl;
    _encoder = NULL;
  }
  else {
    _encoder = ttLibC_X265Encoder_makeWithX265ApiAndParam((void *)api, (void *)x265Param);
  }
#endif
}

X265Encoder::~X265Encoder() {
#ifdef __ENABLE_X265__
  ttLibC_X265Encoder_close(&_encoder);
#endif
}
bool X265Encoder::encodeFrame(ttLibC_Frame *cFrame, ttLibGoFrame *goFrame, void *ptr) {
  bool result = false;
#ifdef __ENABLE_X265__
  if(cFrame->type != frameType_yuv420) {
    return result;
  }
  update(cFrame, goFrame);
  result = ttLibC_X265Encoder_encode(_encoder, (ttLibC_Yuv420 *)cFrame, [](void *ptr, ttLibC_H265 *h265)->bool {
    return ttLibGoFrameCallback(ptr, (ttLibC_Frame *)h265);
  }, ptr);
  reset(cFrame, goFrame);
#endif
  return result;
}

bool X265Encoder::forceNextFrameType(string type) {
  bool result = false;
#ifdef __ENABLE_X265__
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
  result = ttLibC_X265Encoder_forceNextFrameType(_encoder, frameType);
#endif
  return result;
}

extern "C" {
bool X265Encoder_forceNextFrameType(void *encoder, const char *type) {
  X265Encoder *x = reinterpret_cast<X265Encoder *>(encoder);
  return x->forceNextFrameType(type);
}
}