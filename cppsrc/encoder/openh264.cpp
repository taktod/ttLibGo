#include "openh264.hpp"
#include "../util.hpp"
#include <iostream>
#include <regex>
#include <ttLibC/container/container.h>
#include <ttLibC/encoder/openh264Encoder.h>

using namespace std;

extern "C" {
typedef void (* ttLibC_Openh264Encoder_getDefaultSEncParamExt_func)(void *, uint32_t, uint32_t);
typedef bool (* ttLibC_Openh264Encoder_paramParse_func)(void *, const char *, const char *);
typedef bool (* ttLibC_Openh264Encoder_spatialParamParse_func)(void *, uint32_t, const char *, const char *);
typedef void *(* ttLibC_Openh264Encoder_makeWithSEncParamExt_func)(void *);
typedef bool (* ttLibC_Openh264Encoder_setRCMode_func)(void *, ttLibC_Openh264Encoder_RCType);
typedef bool (* ttLibC_Openh264Encoder_setIDRInterval_func)(void *, uint32_t);
typedef bool (* ttLibC_Openh264Encoder_forceNextKeyFrame_func)(void *);
typedef void *(* ttLibC_codec_func)(void *, void *, ttLibC_getFrameFunc, void *);
typedef void (* ttLibC_close_func)(void **);

extern ttLibC_Openh264Encoder_getDefaultSEncParamExt_func ttLibGo_Openh264Encoder_getDefaultSEncParamExt;
extern ttLibC_Openh264Encoder_paramParse_func             ttLibGo_Openh264Encoder_paramParse;
extern ttLibC_Openh264Encoder_spatialParamParse_func      ttLibGo_Openh264Encoder_spatialParamParse;
extern ttLibC_Openh264Encoder_makeWithSEncParamExt_func   ttLibGo_Openh264Encoder_makeWithSEncParamExt;
extern ttLibC_codec_func                                  ttLibGo_Openh264Encoder_encode;
extern ttLibC_close_func                                  ttLibGo_Openh264Encoder_close;
extern ttLibC_Openh264Encoder_setRCMode_func              ttLibGo_Openh264Encoder_setRCMode;
extern ttLibC_Openh264Encoder_setIDRInterval_func         ttLibGo_Openh264Encoder_setIDRInterval;
extern ttLibC_Openh264Encoder_forceNextKeyFrame_func      ttLibGo_Openh264Encoder_forceNextKeyFrame;

extern bool ttLibGoFrameCallback(void *ptr, ttLibC_Frame *frame);
}
Openh264Encoder::Openh264Encoder(maps *mp) {
  if(ttLibGo_Openh264Encoder_getDefaultSEncParamExt != nullptr
  && ttLibGo_Openh264Encoder_paramParse != nullptr
  && ttLibGo_Openh264Encoder_spatialParamParse != nullptr
  && ttLibGo_Openh264Encoder_makeWithSEncParamExt != nullptr) {
    uint32_t width = mp->getUint32("width");
    uint32_t height = mp->getUint32("height");
    list<string> params = mp->getStringList("param");
    list<string> spatialParamArrays = mp->getStringList("spatialParamArray");
    if(width == 0 || height == 0) {
      return;
    }
//    SEncParamExt paramExt;
//    printf("%d\n", sizeof(SEncParamExt)); // 916 byteらしい。
    char paramExt[4096]; // とりあえず、charで大きめに保持しておいて、誤魔化す。
    // width height params spatialParamsArrayが必要
    (*ttLibGo_Openh264Encoder_getDefaultSEncParamExt)(&paramExt, width, height);
    regex r(R"([^:]+):([^:]+)");
    for(auto iter = params.cbegin(); iter != params.cend();++ iter) {
      sregex_iterator reg_iter(iter->begin(), iter->end(), r); 
      string key = reg_iter->str();
      reg_iter ++;
      string value = reg_iter->str();
      bool result = (*ttLibGo_Openh264Encoder_paramParse)(&paramExt, key.c_str(), value.c_str());
      if(!result) {
        cout << "パラメーター設定失敗:" << key << ":" << value << endl;
      }
    }
    regex sr(R"([^:]+):([^:]+):([^:]+)");
    for(auto iter = spatialParamArrays.cbegin();iter != spatialParamArrays.cend();++ iter) {
      sregex_iterator reg_iter(iter->begin(), iter->end(), r); 
      string num = reg_iter->str();
      reg_iter ++;
      string key = reg_iter->str();
      reg_iter ++;
      string value = reg_iter->str();
      bool result = (*ttLibGo_Openh264Encoder_spatialParamParse)(&paramExt, stoi(num), key.c_str(), value.c_str());
      if(!result) {
        cout << "パラメーター設定失敗:" << key << ":" << value << endl;
      }
    }
    _encoder = (*ttLibGo_Openh264Encoder_makeWithSEncParamExt)(&paramExt);
  }
}

Openh264Encoder::~Openh264Encoder() {
  if(ttLibGo_Openh264Encoder_close != nullptr) {
    (*ttLibGo_Openh264Encoder_close)(&_encoder);
  }
}
bool Openh264Encoder::encodeFrame(ttLibC_Frame *cFrame, ttLibGoFrame *goFrame, void *ptr) {
  bool result = false;
  if(ttLibGo_Openh264Encoder_encode != nullptr) {
    if(cFrame->type != frameType_yuv420) {
      return result;
    }
    update(cFrame, goFrame);
    result = (*ttLibGo_Openh264Encoder_encode)(_encoder, cFrame, ttLibGoFrameCallback, ptr);
    reset(cFrame, goFrame);
  }
  return result;
}

bool Openh264Encoder::setRCMode(string mode) {
  if(ttLibGo_Openh264Encoder_setRCMode != nullptr) {
    ttLibC_Openh264Encoder_RCType type;
    if(mode == "QualityMode") {
      type = Openh264EncoderRCType_QualityMode;
    }
    else if(mode == "BitrateMode") {
      type = Openh264EncoderRCType_BitrateMode;
    }
    else if(mode == "BufferbasedMode") {
      type = Openh264EncoderRCType_BufferbasedMode;
    }
    else if(mode == "TimestampMode") {
      type = Openh264EncoderRCType_TimestampMode;
    }
    else if(mode == "BitrateModePostSkip") {
      type = Openh264EncoderRCType_BitrateModePostSkip;
    }
    else if(mode == "OffMode") {
      type = Openh264EncoderRCType_OffMode;
    }
    else {
      return false;
    }
    return (*ttLibGo_Openh264Encoder_setRCMode)(_encoder, type);
  }
  return false;
}
bool Openh264Encoder::setIDRInterval(uint32_t value) {
  bool result = false;
  if(ttLibGo_Openh264Encoder_setIDRInterval != nullptr) {
    result = (*ttLibGo_Openh264Encoder_setIDRInterval)(_encoder, value);
  }
  return result;
}
bool Openh264Encoder::forceNextKeyFrame() {
  bool result = false;
  if(ttLibGo_Openh264Encoder_forceNextKeyFrame != nullptr) {
    (*ttLibGo_Openh264Encoder_forceNextKeyFrame)(_encoder);
  }
  return result;
}

extern "C" {
bool Openh264Encoder_setRCMode(void *encoder, const char *mode) {
  Openh264Encoder *o = reinterpret_cast<Openh264Encoder *>(encoder);
  return o->setRCMode(mode);
}
bool Openh264Encoder_setIDRInterval(void *encoder, uint32_t interval) {
  Openh264Encoder *o = reinterpret_cast<Openh264Encoder *>(encoder);
  return o->setIDRInterval(interval);
}
bool Openh264Encoder_forceNextKeyFrame(void *encoder) {
  Openh264Encoder *o = reinterpret_cast<Openh264Encoder *>(encoder);
  return o->forceNextKeyFrame();
}
}
