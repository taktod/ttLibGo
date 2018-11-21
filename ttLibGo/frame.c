#include <ttLibC/frame/frame.h>
#include "log.h"
#include <string.h>

ttLibC_Frame_Type __attribute__ ((visibility("hidden")))Frame_getFrameType(const char *type) {
	if(type == NULL) {
		return frameType_unknown;
	}
	if(strcmp(type, "aac") == 0) {
		return frameType_aac;
	}
	else if(strcmp(type, "bgr") == 0) {
		return frameType_bgr;
	}
	else if(strcmp(type, "mp3") == 0) {
		return frameType_mp3;
	}
	else if(strcmp(type, "png") == 0) {
		return frameType_png;
	}
	else if(strcmp(type, "vp6") == 0) {
		return frameType_vp6;
	}
	else if(strcmp(type, "vp8") == 0) {
		return frameType_vp8;
	}
	else if(strcmp(type, "vp9") == 0) {
		return frameType_vp9;
	}
	else if(strcmp(type, "yuv") == 0) {
		return frameType_yuv420;
	}
	else if(strcmp(type, "flv1") == 0) {
		return frameType_flv1;
	}
	else if(strcmp(type, "h264") == 0) {
		return frameType_h264;
	}
	else if(strcmp(type, "h265") == 0) {
		return frameType_h265;
	}
	else if(strcmp(type, "jpeg") == 0) {
		return frameType_jpeg;
	}
	else if(strcmp(type, "opus") == 0) {
		return frameType_opus;
	}
	else if(strcmp(type, "wmv1") == 0) {
		return frameType_wmv1;
	}
	else if(strcmp(type, "wmv2") == 0) {
		return frameType_wmv2;
	}
	else if(strcmp(type, "speex") == 0) {
		return frameType_speex;
	}
	else if(strcmp(type, "pcmF32") == 0) {
		return frameType_pcmF32;
	}
	else if(strcmp(type, "pcmS16") == 0) {
		return frameType_pcmS16;
	}
	else if(strcmp(type, "theora") == 0) {
		return frameType_theora;
	}
	else if(strcmp(type, "vorbis") == 0) {
		return frameType_vorbis;
	}
	else if(strcmp(type, "pcmAlaw") == 0) {
		return frameType_pcm_alaw;
	}
	else if(strcmp(type, "pcmMulaw") == 0) {
		return frameType_pcm_mulaw;
	}
	else if(strcmp(type, "nellymoser") == 0) {
		return frameType_nellymoser;
	}
	else if(strcmp(type, "adpcmImaWav") == 0) {
		return frameType_adpcm_ima_wav;
	}
	ERR_PRINT("フレームのtypeが不正でした");
	return frameType_unknown;
}
