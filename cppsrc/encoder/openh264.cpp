#include "openh264.hpp"
#include "../util.hpp"
#include <iostream>
#include <regex>

#ifdef __ENABLE_OPENH264__
# include <wels/codec_api.h>
#endif

using namespace std;

extern "C" {
extern bool ttLibGoFrameCallback(void *ptr, ttLibC_Frame *frame);
}
Openh264Encoder::Openh264Encoder(maps *mp) {
#ifdef __ENABLE_OPENH264__
  uint32_t width = mp->getUint32("width");
  uint32_t height = mp->getUint32("height");
  list<string> params = mp->getStringList("param");
  list<string> spatialParamArrays = mp->getStringList("spatialParamArray");
  if(width == 0 || height == 0) {
    return;
  }
  SEncParamExt paramExt;
  // width height params spatialParamsArrayが必要
  ttLibC_Openh264Encoder_getDefaultSEncParamExt(&paramExt, width, height);
  regex r(R"([^:]+):([^:]+)");
  for(auto iter = params.cbegin(); iter != params.cend();++ iter) {
    sregex_iterator reg_iter(iter->begin(), iter->end(), r); 
    string key = reg_iter->str();
    reg_iter ++;
    string value = reg_iter->str();
    bool result = ttLibC_Openh264Encoder_paramParse(&paramExt, key.c_str(), value.c_str());
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
    bool result = ttLibC_Openh264Encoder_spatialParamParse(&paramExt, stoi(num), key.c_str(), value.c_str());
    if(!result) {
      cout << "パラメーター設定失敗:" << key << ":" << value << endl;
    }
  }
  _encoder = ttLibC_Openh264Encoder_makeWithSEncParamExt(&paramExt);
#endif
}

Openh264Encoder::~Openh264Encoder() {
#ifdef __ENABLE_OPENH264__
  ttLibC_Openh264Encoder_close(&_encoder);
#endif
}
bool Openh264Encoder::encodeFrame(ttLibC_Frame *cFrame, ttLibGoFrame *goFrame, void *ptr) {
  bool result = false;
#ifdef __ENABLE_OPENH264__
  if(cFrame->type != frameType_yuv420) {
    return result;
  }
  update(cFrame, goFrame);
  result = ttLibC_Openh264Encoder_encode(_encoder, (ttLibC_Yuv420 *)cFrame, [](void *ptr, ttLibC_H264 *h264)->bool {
    return ttLibGoFrameCallback(ptr, (ttLibC_Frame *)h264);
  }, ptr);
  reset(cFrame, goFrame);
#endif
  return result;
}

bool Openh264Encoder::setRCMode(string mode) {
  bool result = false;
#ifdef __ENABLE_OPENH264__
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
	return ttLibC_Openh264Encoder_setRCMode(_encoder, type);
#endif
  return result;
}
bool Openh264Encoder::setIDRInterval(uint32_t value) {
  bool result = false;
#ifdef __ENABLE_OPENH264__
  result = ttLibC_Openh264Encoder_setIDRInterval(_encoder, value);
#endif
  return result;
}
bool Openh264Encoder::forceNextKeyFrame() {
  bool result = false;
#ifdef __ENABLE_OPENH264__
  result = ttLibC_Openh264Encoder_forceNextKeyFrame(_encoder);
#endif
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
