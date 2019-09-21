#include "fdkaac.hpp"
#include "../util.hpp"
#include <iostream>
#include "ttLibC/ttLibC/container/container.h"

using namespace std;

extern "C" {
typedef void *(* ttLibC_FdkaacEncoder_make_func)(const char *,uint32_t,uint32_t,uint32_t);
typedef bool (* ttLibC_FdkaacEncoder_setBitrate_func)(void *,uint32_t);
typedef void (* ttLibC_close_func)(void **);
typedef void *(* ttLibC_codec_func)(void *, void *, ttLibC_getFrameFunc, void *);

extern ttLibC_FdkaacEncoder_make_func       ttLibGo_FdkaacEncoder_make;
extern ttLibC_codec_func                    ttLibGo_FdkaacEncoder_encode;
extern ttLibC_close_func                    ttLibGo_FdkaacEncoder_close;
extern ttLibC_FdkaacEncoder_setBitrate_func ttLibGo_FdkaacEncoder_setBitrate;

extern bool ttLibGoFrameCallback(void *ptr, ttLibC_Frame *frame);
}
FdkaacEncoder::FdkaacEncoder(maps *mp) {
	if(ttLibGo_FdkaacEncoder_make != nullptr) {
		string type = mp->getString("aacType");
		uint32_t sampleRate = mp->getUint32("sampleRate");
		uint32_t channelNum = mp->getUint32("channelNum");
		uint32_t bitrate = mp->getUint32("bitrate");
		_encoder = (*ttLibGo_FdkaacEncoder_make)(type.c_str(), sampleRate, channelNum, bitrate);
	}
}

FdkaacEncoder::~FdkaacEncoder() {
	if(ttLibGo_FdkaacEncoder_close != nullptr) {
		(*ttLibGo_FdkaacEncoder_close)(&_encoder);
	}
}
bool FdkaacEncoder::encodeFrame(ttLibC_Frame *cFrame, ttLibGoFrame *goFrame, void *ptr) {
	bool result = false;
	if(ttLibGo_FdkaacEncoder_encode != nullptr) {
		if(cFrame->type != frameType_pcmS16) {
			return result;
		}
    update(cFrame, goFrame);
    result = (*ttLibGo_FdkaacEncoder_encode)(_encoder, cFrame, ttLibGoFrameCallback, ptr);
    reset(cFrame, goFrame);
	}
	return result;
}

bool FdkaacEncoder::setBitrate(uint32_t value) {
	if(ttLibGo_FdkaacEncoder_setBitrate != nullptr) {
		return (* ttLibGo_FdkaacEncoder_setBitrate)(_encoder, value);
	}
	return false;
}
extern "C" {
bool FdkaacEncoder_setBitrate(void *encoder, uint32_t value) {
  FdkaacEncoder *fdk = reinterpret_cast<FdkaacEncoder *>(encoder);
	return fdk->setBitrate(value);
}
}
