#include <stdio.h>
#include <ttLibC/frame/frame.h>
#include <ttLibC/frame/video/bgr.h>
#include <ttLibC/frame/video/flv1.h>
#include <ttLibC/frame/video/h264.h>
#include <ttLibC/frame/video/h265.h>
#include <ttLibC/frame/video/jpeg.h>
#include <ttLibC/frame/video/png.h>
#include <ttLibC/frame/video/theora.h>
#include <ttLibC/frame/video/vp6.h>
#include <ttLibC/frame/video/vp8.h>
#include <ttLibC/frame/video/vp9.h>
#include <ttLibC/frame/video/yuv420.h>

#include <ttLibC/frame/audio/audio.h>
#include <ttLibC/frame/audio/aac.h>
#include <ttLibC/frame/audio/adpcmImaWav.h>
#include <ttLibC/frame/audio/mp3.h>
#include <ttLibC/frame/audio/nellymoser.h>
#include <ttLibC/frame/audio/opus.h>
#include <ttLibC/frame/audio/pcmAlaw.h>
#include <ttLibC/frame/audio/pcmf32.h>
#include <ttLibC/frame/audio/pcmMulaw.h>
#include <ttLibC/frame/audio/pcms16.h>
#include <ttLibC/frame/audio/speex.h>
#include <ttLibC/frame/audio/vorbis.h>

#include <ttLibC/util/byteUtil.h>
#include <ttLibC/util/ioUtil.h>

#include "frame.hpp"
#include <string>

using namespace std;

extern "C" {
typedef bool (* ttLibC_isFrameType_func)(void *);
typedef void *(* ttLibC_Frame_clone_func)(void *, void *);
typedef void (* ttLibC_close_func)(void **);
typedef bool (* ttLibC_isType_func)(ttLibC_Frame_Type);

extern ttLibC_Frame_clone_func ttLibGo_Frame_clone;
extern ttLibC_close_func       ttLibGo_Frame_close;
extern ttLibC_isFrameType_func ttLibGo_Frame_isVideo;
extern ttLibC_isFrameType_func ttLibGo_Frame_isAudio;
extern ttLibC_isType_func ttLibGo_isAudio;
extern ttLibC_isType_func ttLibGo_isVideo;

typedef bool (* ttLibC_video_getMinimumBinaryBuffer_func)(void *, bool(*)(void *, void *, size_t), void *);

extern ttLibC_video_getMinimumBinaryBuffer_func ttLibGo_Bgr_getMinimumBinaryBuffer;
extern ttLibC_video_getMinimumBinaryBuffer_func ttLibGo_Yuv420_getMinimumBinaryBuffer;

typedef void *(* ttLibC_Bgr_makeEmptyFrame_func)(ttLibC_Bgr_Type, uint32_t, uint32_t);
typedef void *(* ttLibC_frame_getFrame_func)(void *, void *, size_t, bool, uint64_t, uint32_t);
typedef uint32_t (* ttLibC_video_getWidth_func)(void *, void *, size_t);
typedef uint32_t (* ttLibC_video_getHeight_func)(void *, void *, size_t);
typedef void *(* ttLibC_Theora_make_func)(void *,ttLibC_Theora_Type,uint32_t,uint32_t,void *,size_t,bool,uint64_t,uint32_t);
typedef bool (* ttLibC_video_isKey_func)(void *, size_t);
typedef uint32_t (* ttLibC_Vp6_getWidth_func)(void *, void *, size_t, uint8_t);
typedef uint32_t (* ttLibC_Vp6_getHeight_func)(void *, void *, size_t, uint8_t);
typedef void *(* ttLibC_video_make_func)(void *,ttLibC_Video_Type,uint32_t,uint32_t,void *,size_t,bool,uint64_t,uint32_t);
typedef void *(* ttLibC_Yuv420_makeEmptyFrame_func)(ttLibC_Yuv420_Type, uint32_t, uint32_t);

extern ttLibC_Bgr_makeEmptyFrame_func ttLibGo_Bgr_makeEmptyFrame;
extern ttLibC_frame_getFrame_func ttLibGo_Flv1_getFrame;
extern ttLibC_frame_getFrame_func ttLibGo_H264_getFrame;
extern ttLibC_frame_getFrame_func ttLibGo_H265_getFrame;
extern ttLibC_frame_getFrame_func ttLibGo_Jpeg_getFrame;
extern ttLibC_frame_getFrame_func ttLibGo_Png_getFrame;
extern ttLibC_video_getWidth_func  ttLibGo_Theora_getWidth;
extern ttLibC_video_getHeight_func ttLibGo_Theora_getHeight;
extern ttLibC_Theora_make_func     ttLibGo_Theora_make;
extern ttLibC_video_isKey_func   ttLibGo_Vp6_isKey;
extern ttLibC_Vp6_getWidth_func  ttLibGo_Vp6_getWidth;
extern ttLibC_Vp6_getHeight_func ttLibGo_Vp6_getHeight;
extern ttLibC_video_make_func    ttLibGo_Vp6_make;
extern ttLibC_video_isKey_func     ttLibGo_Vp8_isKey;
extern ttLibC_video_getWidth_func  ttLibGo_Vp8_getWidth;
extern ttLibC_video_getHeight_func ttLibGo_Vp8_getHeight;
extern ttLibC_video_make_func      ttLibGo_Vp8_make;
extern ttLibC_video_isKey_func     ttLibGo_Vp9_isKey;
extern ttLibC_video_getWidth_func  ttLibGo_Vp9_getWidth;
extern ttLibC_video_getHeight_func ttLibGo_Vp9_getHeight;
extern ttLibC_video_make_func      ttLibGo_Vp9_make;
extern ttLibC_Yuv420_makeEmptyFrame_func ttLibGo_Yuv420_makeEmptyFrame;

typedef void *(* ttLibC_Aac_make_func)(void *,ttLibC_Aac_Type,uint32_t,uint32_t,uint32_t,void*,size_t,bool,uint64_t,uint32_t,uint64_t);
typedef void *(* ttLibC_audio_make_func)(void *,uint32_t,uint32_t,uint32_t,void *,size_t,bool,uint64_t,uint32_t);
typedef void *(* ttLibC_PcmF32_make_func)(void *,ttLibC_PcmF32_Type,uint32_t,uint32_t,uint32_t,void*,size_t,void *,uint32_t,void*,uint32_t,bool,uint64_t,uint32_t);
typedef void *(* ttLibC_PcmS16_make_func)(void *,ttLibC_PcmS16_Type,uint32_t,uint32_t,uint32_t,void*,size_t,void *,uint32_t,void*,uint32_t,bool,uint64_t,uint32_t);
typedef void *(* ttLibC_Speex_make_func)(void *,ttLibC_Speex_Type,uint32_t,uint32_t,uint32_t,void *,size_t,bool,uint64_t,uint32_t);
typedef void *(* ttLibC_Vorbis_make_func)(void *,ttLibC_Vorbis_Type,uint32_t,uint32_t,uint32_t,void *,size_t,bool,uint64_t,uint32_t);

extern ttLibC_Aac_make_func ttLibGo_Aac_make;
extern ttLibC_audio_make_func ttLibGo_AdpcmImaWav_make;
extern ttLibC_frame_getFrame_func ttLibGo_Mp3_getFrame;
extern ttLibC_audio_make_func ttLibGo_Nellymoser_make;
extern ttLibC_frame_getFrame_func ttLibGo_Opus_getFrame;
extern ttLibC_audio_make_func ttLibGo_PcmAlaw_make;
extern ttLibC_PcmF32_make_func ttLibGo_PcmF32_make;
extern ttLibC_audio_make_func ttLibGo_PcmMulaw_make;
extern ttLibC_PcmS16_make_func ttLibGo_PcmS16_make;
extern ttLibC_Speex_make_func ttLibGo_Speex_make;
extern ttLibC_Vorbis_make_func ttLibGo_Vorbis_make;

typedef void *(* ttLibC_ByteReader_make_func)(void *,size_t,ttLibC_ByteUtil_Type);
typedef uint64_t (* ttLibC_ByteReader_bit_func)(void *,uint32_t);

extern ttLibC_ByteReader_make_func ttLibGo_ByteReader_make;
extern ttLibC_ByteReader_bit_func  ttLibGo_ByteReader_bit;
extern ttLibC_close_func           ttLibGo_ByteReader_close;

}

ttLibC_Frame_Type Frame_getFrameTypeFromString(string name) {
  if(name == "bgr") {return frameType_bgr;}
  if(name == "flv1") {return frameType_flv1;}
  if(name == "h264") {return frameType_h264;}
  if(name == "h265") {return frameType_h265;}
  if(name == "jpeg") {return frameType_jpeg;}
  if(name == "png") {return frameType_png;}
  if(name == "theora") {return frameType_theora;}
  if(name == "vp6") {return frameType_vp6;}
  if(name == "vp8") {return frameType_vp8;}
  if(name == "vp9") {return frameType_vp9;}
  if(name == "wmv1") {return frameType_wmv1;}
  if(name == "wmv2") {return frameType_wmv2;}
  if(name == "yuv420") {return frameType_yuv420;}
  if(name == "aac") {return frameType_aac;}
  if(name == "adpcmImaWav") {return frameType_adpcm_ima_wav;}
  if(name == "mp3") {return frameType_mp3;}
  if(name == "nellymoser") {return frameType_nellymoser;}
  if(name == "opus") {return frameType_opus;}
  if(name == "pcmAlaw") {return frameType_pcm_alaw;}
  if(name == "pcmF32") {return frameType_pcmF32;}
  if(name == "pcmMulaw") {return frameType_pcm_mulaw;}
  if(name == "pcmS16") {return frameType_pcmS16;}
  if(name == "speex") {return frameType_speex;}
  if(name == "vorbis") {return frameType_vorbis;}
  return frameType_unknown;
}

void FrameProcessor::update(ttLibC_Frame *cFrame, ttLibGoFrame *goFrame) {
  if(ttLibGo_Frame_isVideo == nullptr
  || ttLibGo_Frame_isAudio == nullptr) {
    return;
  }
  _pts = cFrame->pts;
  _dts = cFrame->dts;
  _timebase = cFrame->timebase;
  _id = cFrame->id;
  cFrame->pts = goFrame->pts;
  cFrame->dts = goFrame->dts;
  cFrame->timebase = goFrame->timebase;
  cFrame->id = goFrame->id;
  if((*ttLibGo_Frame_isVideo)(cFrame)) {
    ttLibC_Video *video = (ttLibC_Video *)cFrame;
    _width = video->width;
    _height = video->height;
    if(goFrame->width != 0) {
      video->width = goFrame->width;
    }
    if(goFrame->height != 0) {
      video->height = goFrame->height;
    }
  }
  if((*ttLibGo_Frame_isAudio)(cFrame)) {
    ttLibC_Audio *audio = (ttLibC_Audio *)cFrame;
    _sample_rate = audio->sample_rate;
    _sample_num = audio->sample_num;
    _channel_num = audio->channel_num;
    if(goFrame->sampleRate != 0) {
      audio->sample_rate = goFrame->sampleRate;
    }
    if(goFrame->sampleNum != 0) {
      audio->sample_num = goFrame->sampleNum;
    }
    if(goFrame->channelNum != 0) {
      audio->channel_num = goFrame->channelNum;
    }
  }
  switch(cFrame->type) {
	case frameType_bgr:
    {
      ttLibC_Bgr *bgr = (ttLibC_Bgr *)cFrame;
      _data = bgr->data;
      _width_stride = bgr->width_stride;
      if(goFrame->dataPos != 0) {
        bgr->data = bgr->data + goFrame->dataPos;
      }
      if(goFrame->widthStride != 0) {
        bgr->width_stride = goFrame->widthStride;
      }
    }
    break;
	case frameType_yuv420:
    {
      ttLibC_Yuv420 *yuv420 = (ttLibC_Yuv420 *)cFrame;
      _y_data = yuv420->y_data;
      _y_stride = yuv420->y_stride;
      _u_data = yuv420->u_data;
      _u_stride = yuv420->u_stride;
      _v_data = yuv420->v_data;
      _v_stride = yuv420->v_stride;
      if(goFrame->yDataPos != 0) {
        yuv420->y_data = yuv420->y_data + goFrame->yDataPos;
      }
      if(goFrame->yStride != 0) {
        yuv420->y_stride = goFrame->yStride;
      }
      if(goFrame->uDataPos != 0) {
        yuv420->u_data = yuv420->u_data + goFrame->uDataPos;
      }
      if(goFrame->uStride != 0) {
        yuv420->u_stride = goFrame->uStride;
      }
      if(goFrame->vDataPos != 0) {
        yuv420->v_data = yuv420->v_data + goFrame->vDataPos;
      }
      if(goFrame->vStride != 0) {
        yuv420->v_stride = goFrame->vStride;
      }
    }
    break;
  default:
    break;
  }
}
void FrameProcessor::reset(ttLibC_Frame *cFrame, ttLibGoFrame *goFrame) {
  if(ttLibGo_Frame_isVideo == nullptr
  || ttLibGo_Frame_isAudio == nullptr) {
    return;
  }
  cFrame->pts = _pts;
  cFrame->dts = _dts;
  cFrame->timebase = _timebase;
  cFrame->id = _id;
  if((*ttLibGo_Frame_isVideo)(cFrame)) {
    ttLibC_Video *video = (ttLibC_Video *)cFrame;
    video->width = _width;
    video->height = _height;
  }
  if((*ttLibGo_Frame_isAudio)(cFrame)) {
    ttLibC_Audio *audio = (ttLibC_Audio *)cFrame;
    audio->sample_rate = _sample_rate;
    audio->sample_num = _sample_num;
    audio->channel_num = _channel_num;
  }
  switch(cFrame->type) {
	case frameType_bgr:
    {
      ttLibC_Bgr *bgr = (ttLibC_Bgr *)cFrame;
      bgr->data = _data;
      bgr->width_stride = _width_stride;
    }
    break;
	case frameType_yuv420:
    {
      ttLibC_Yuv420 *yuv420 = (ttLibC_Yuv420 *)cFrame;
      yuv420->y_data = _y_data;
      yuv420->y_stride = _y_stride;
      yuv420->u_data = _u_data;
      yuv420->u_stride = _u_stride;
      yuv420->v_data = _v_data;
      yuv420->v_stride = _v_stride;
    }
    break;
  default:
    break;
  }
}

extern "C" {

/**
 * frameTypeを参照する
 * @param cFrame 対象frameオブジェクトポインタ
 * @param buffer 文字列を格納するメモリー
 * @param buffer_size メモリーのサイズ
 */
void Frame_getFrameType(void *cFrame, char *buffer, size_t buffer_size) {
	if(cFrame == NULL || buffer == NULL || buffer_size < 256) {
		return;
	}
	ttLibC_Frame *f = (ttLibC_Frame *)cFrame;
	switch(f->type) {
	case frameType_bgr:           strcpy(buffer, "bgr");         break;
	case frameType_flv1:          strcpy(buffer, "flv1");        break;
	case frameType_h264:          strcpy(buffer, "h264");        break;
	case frameType_h265:          strcpy(buffer, "h265");        break;
	case frameType_jpeg:          strcpy(buffer, "jpeg");        break;
	case frameType_png:           strcpy(buffer, "png");         break;
	case frameType_theora:        strcpy(buffer, "theora");      break;
	case frameType_vp6:           strcpy(buffer, "vp6");         break;
	case frameType_vp8:           strcpy(buffer, "vp8");         break;
	case frameType_vp9:           strcpy(buffer, "vp9");         break;
	case frameType_wmv1:          strcpy(buffer, "wmv1");        break;
	case frameType_wmv2:          strcpy(buffer, "wmv2");        break;
	case frameType_yuv420:        strcpy(buffer, "yuv420");      break;
	case frameType_aac:           strcpy(buffer, "aac");         break;
	case frameType_adpcm_ima_wav: strcpy(buffer, "adpcmImaWav"); break;
	case frameType_mp3:           strcpy(buffer, "mp3");         break;
	case frameType_nellymoser:    strcpy(buffer, "nellymoser");  break;
	case frameType_opus:          strcpy(buffer, "opus");        break;
	case frameType_pcm_alaw:      strcpy(buffer, "pcmAlaw");     break;
	case frameType_pcmF32:        strcpy(buffer, "pcmF32");      break;
	case frameType_pcm_mulaw:     strcpy(buffer, "pcmMulaw");    break;
	case frameType_pcmS16:        strcpy(buffer, "pcmS16");      break;
	case frameType_speex:         strcpy(buffer, "speex");       break;
	case frameType_vorbis:        strcpy(buffer, "vorbis");      break;
	default: strcpy(buffer, ""); break;
	}
}
/**
 * frameのptsの値を取得する
 * @param frame
 * @return ptsの値
 */
uint64_t Frame_getPts(void *frame) {
	if(frame == NULL) {
		return 0;
	}
	ttLibC_Frame *f = (ttLibC_Frame *)frame;
	return f->pts;
}

/**
 * frameのdtsの値を取得する
 * @param frame
 * @return ptsの値
 */
uint64_t Frame_getDts(void *frame) {
	if(frame == NULL) {
		return 0;
	}
	ttLibC_Frame *f = (ttLibC_Frame *)frame;
	return f->dts;
}

/**
 * frameのtimebaseの値を取得する
 * @param frame
 * @return timebaseの値
 */
uint32_t Frame_getTimebase(void *frame) {
	if(frame == NULL) {
		return 0;
	}
	ttLibC_Frame *f = (ttLibC_Frame *)frame;
	return f->timebase;
}

/**
 * frameのIDの値を取得する
 * @param frame
 * @return IDの値
 */
uint32_t Frame_getID(void *frame) {
	if(frame == NULL) {
		return 0;
	}
	ttLibC_Frame *f = (ttLibC_Frame *)frame;
	return f->id;
}

/**
 * frameを解放する動作
 * @param frame
 */
void Frame_close(void *frame) {
  if(frame == NULL) {
    return;
  }
	if(ttLibGo_Frame_close != NULL) {
		(*ttLibGo_Frame_close)(&frame);
	}
}

/**
 * frameのコピーを作成する動作
 * @param frame
 * @return 作成したフレーム
 */
ttLibC_Frame *Frame_clone(void *frame) {
  if(frame == NULL) {
    return NULL;
  }
	if(ttLibGo_Frame_clone != NULL) {
		return (ttLibC_Frame *)(*ttLibGo_Frame_clone)(NULL, frame);
	}
	return NULL;
}

/**
 * frameのsampleRateの値を取得する
 * @param frame
 * @return sampleRateの値
 */
uint32_t AudioFrame_getSampleRate(void *frame) {
	if(frame == NULL || ttLibGo_isAudio == NULL || !(*ttLibGo_isAudio)(((ttLibC_Frame *)frame)->type)) {
		return 0;
	}
	ttLibC_Audio *audio = (ttLibC_Audio *)frame;
	return audio->sample_rate;
}
/**
 * frameのsampleNumの値を取得する
 * @param frame
 * @return sampleNumの値
 */
uint32_t AudioFrame_getSampleNum(void *frame) {
	if(frame == NULL || ttLibGo_isAudio == NULL || !(*ttLibGo_isAudio)(((ttLibC_Frame *)frame)->type)) {
		return 0;
	}
	ttLibC_Audio *audio = (ttLibC_Audio *)frame;
	return audio->sample_num;
}
/**
 * frameのchannelNumの値を取得する
 * @param frame
 * @return channelNumの値
 */
uint32_t AudioFrame_getChannelNum(void *frame) {
	if(frame == NULL || ttLibGo_isAudio == NULL || !(*ttLibGo_isAudio)(((ttLibC_Frame *)frame)->type)) {
		return 0;
	}
	ttLibC_Audio *audio = (ttLibC_Audio *)frame;
	return audio->channel_num;
}

void AacFrame_getAacType(void *cFrame, char *buffer, size_t buffer_size) {
  if(cFrame == NULL || ((ttLibC_Frame *)cFrame)->type != frameType_aac || buffer == NULL || buffer_size < 5) {
    return;
  }
  ttLibC_Aac *aac = (ttLibC_Aac *)cFrame;
	switch(aac->type) {
	case AacType_adts: strcpy(buffer, "adts"); break;
	case AacType_dsi:  strcpy(buffer, "dsi");  break;
	case AacType_raw:  strcpy(buffer, "raw");  break;
	default:
		break;
	}
}

void Mp3Frame_getMp3Type(void *cFrame, char *buffer, size_t buffer_size) {
  if(cFrame == NULL || ((ttLibC_Frame *)cFrame)->type != frameType_mp3 || buffer == NULL || buffer_size < 6) {
    return;
  }
  ttLibC_Mp3 *mp3 = (ttLibC_Mp3 *)cFrame;
	switch(mp3->type) {
	case Mp3Type_frame: strcpy(buffer, "frame"); break;
	case Mp3Type_id3:   strcpy(buffer, "id3");   break;
	case Mp3Type_tag:   strcpy(buffer, "tag");   break;
	default:
		break;
	}
}

void OpusFrame_getOpusType(void *cFrame, char *buffer, size_t buffer_size) {
  if(cFrame == NULL || ((ttLibC_Frame *)cFrame)->type != frameType_opus || buffer == NULL || buffer_size < 6) {
    return;
  }
  ttLibC_Opus *opus = (ttLibC_Opus *)cFrame;
	switch(opus->type) {
	case OpusType_header:  strcpy(buffer, "header");  break;
	case OpusType_comment: strcpy(buffer, "comment"); break;
	case OpusType_frame:   strcpy(buffer, "frame");   break;
	default:
		break;
	}
}

void PcmF32Frame_getPcmF32Type(void *cFrame, char *buffer, size_t buffer_size) {
  if(cFrame == NULL || ((ttLibC_Frame *)cFrame)->type != frameType_pcmF32 || buffer == NULL || buffer_size < 11) {
    return;
  }
  ttLibC_PcmF32 *pcm = (ttLibC_PcmF32 *)cFrame;
	switch(pcm->type) {
	case PcmF32Type_interleave: strcpy(buffer, "interleave");  break;
	case PcmF32Type_planar:     strcpy(buffer, "planar"); break;
	default:
		break;
	}
}

uint32_t PcmF32Frame_getLStride(void *cFrame) {
	if(cFrame == NULL || ((ttLibC_Frame *)cFrame)->type != frameType_pcmF32) {
		return 0;
	}
	ttLibC_PcmF32 *pcm = (ttLibC_PcmF32 *)cFrame;
	return pcm->l_stride;
}

uint32_t PcmF32Frame_getLStep(void *cFrame) {
	if(cFrame == NULL || ((ttLibC_Frame *)cFrame)->type != frameType_pcmF32) {
		return 0;
	}
	ttLibC_PcmF32 *pcm = (ttLibC_PcmF32 *)cFrame;
	return pcm->l_step;
}

uint32_t PcmF32Frame_getRStride(void *cFrame) {
	if(cFrame == NULL || ((ttLibC_Frame *)cFrame)->type != frameType_pcmF32) {
		return 0;
	}
	ttLibC_PcmF32 *pcm = (ttLibC_PcmF32 *)cFrame;
	return pcm->r_stride;
}

uint32_t PcmF32Frame_getRStep(void *cFrame) {
	if(cFrame == NULL || ((ttLibC_Frame *)cFrame)->type != frameType_pcmF32) {
		return 0;
	}
	ttLibC_PcmF32 *pcm = (ttLibC_PcmF32 *)cFrame;
	return pcm->r_step;
}

void PcmS16Frame_getPcmS16Type(void *cFrame, char *buffer, size_t buffer_size) {
  if(cFrame == NULL || ((ttLibC_Frame *)cFrame)->type != frameType_pcmS16 || buffer == NULL || buffer_size < 19) {
    return;
  }
  ttLibC_PcmS16 *pcm = (ttLibC_PcmS16 *)cFrame;
	switch(pcm->type) {
	case PcmS16Type_bigEndian:           strcpy(buffer, "bigEndian");  break;
	case PcmS16Type_bigEndian_planar:    strcpy(buffer, "bigEndianPlanar"); break;
	case PcmS16Type_littleEndian:        strcpy(buffer, "littleEndian");  break;
	case PcmS16Type_littleEndian_planar: strcpy(buffer, "littleEndianPlanar"); break;
	default:
		break;
	}
}

uint32_t PcmS16Frame_getLStride(void *cFrame) {
	if(cFrame == NULL || ((ttLibC_Frame *)cFrame)->type != frameType_pcmS16) {
		return 0;
	}
	ttLibC_PcmS16 *pcm = (ttLibC_PcmS16 *)cFrame;
	return pcm->l_stride;
}

uint32_t PcmS16Frame_getLStep(void *cFrame) {
	if(cFrame == NULL || ((ttLibC_Frame *)cFrame)->type != frameType_pcmS16) {
		return 0;
	}
	ttLibC_PcmS16 *pcm = (ttLibC_PcmS16 *)cFrame;
	return pcm->l_step;
}

uint32_t PcmS16Frame_getRStride(void *cFrame) {
	if(cFrame == NULL || ((ttLibC_Frame *)cFrame)->type != frameType_pcmS16) {
		return 0;
	}
	ttLibC_PcmS16 *pcm = (ttLibC_PcmS16 *)cFrame;
	return pcm->r_stride;
}

uint32_t PcmS16Frame_getRStep(void *cFrame) {
	if(cFrame == NULL || ((ttLibC_Frame *)cFrame)->type != frameType_pcmS16) {
		return 0;
	}
	ttLibC_PcmS16 *pcm = (ttLibC_PcmS16 *)cFrame;
	return pcm->r_step;
}

void SpeexFrame_getSpeexType(void *cFrame, char *buffer, size_t buffer_size) {
  if(cFrame == NULL || ((ttLibC_Frame *)cFrame)->type != frameType_speex || buffer == NULL || buffer_size < 19) {
    return;
  }
  ttLibC_Speex *speex = (ttLibC_Speex *)cFrame;
	switch(speex->type) {
	case SpeexType_comment: strcpy(buffer, "comment"); break;
	case SpeexType_frame:   strcpy(buffer, "frame");   break;
	case SpeexType_header:  strcpy(buffer, "header");  break;
	default:
		break;
	}
}

void VorbisFrame_getVorbisType(void *cFrame, char *buffer, size_t buffer_size) {
  if(cFrame == NULL || ((ttLibC_Frame *)cFrame)->type != frameType_vorbis || buffer == NULL || buffer_size < 19) {
    return;
  }
  ttLibC_Vorbis *vorbis = (ttLibC_Vorbis *)cFrame;
	switch(vorbis->type) {
	case VorbisType_comment:        strcpy(buffer, "comment");        break;
	case VorbisType_frame:          strcpy(buffer, "frame");          break;
	case VorbisType_identification: strcpy(buffer, "identification"); break;
	case VorbisType_setup:          strcpy(buffer, "setup");          break;
	default:
		break;
	}
}

/**
 * frameのvideoTypeの値を取得する
 * @param cFrame
 * @param buffer
 * @param buffer_size
 */
void VideoFrame_getVideoType(void *cFrame, char *buffer, size_t buffer_size) {
	if(cFrame == NULL || ttLibGo_isVideo == NULL || !(*ttLibGo_isVideo)(((ttLibC_Frame *)cFrame)->type) || buffer == NULL || buffer_size < 6) {
		return;
	}
	ttLibC_Video *video = (ttLibC_Video *)cFrame;
	switch(video->type) {
	case videoType_key:   strcpy(buffer, "key");   break;
	case videoType_inner: strcpy(buffer, "inner"); break;
	case videoType_info:  strcpy(buffer, "info");  break;
	default:
		break;
	}
}
/**
 * frameのwidthの値を取得する
 * @param frame
 * @return widthの値
 */
uint32_t VideoFrame_getWidth(void *frame) {
	if(frame == NULL || ttLibGo_isVideo == NULL || !(*ttLibGo_isVideo)(((ttLibC_Frame *)frame)->type)) {
		return 0;
	}
	ttLibC_Video *video = (ttLibC_Video *)frame;
	return video->width;
}
/**
 * frameのheightの値を取得する
 * @param frame
 * @return heightの値
 */
uint32_t VideoFrame_getHeight(void *frame) {
	if(frame == NULL || ttLibGo_isVideo == NULL || !(*ttLibGo_isVideo)(((ttLibC_Frame *)frame)->type)) {
		return 0;
	}
	ttLibC_Video *video = (ttLibC_Video *)frame;
	return video->height;
}

void BgrFrame_getBgrType(void *cFrame, char *buffer, size_t buffer_size) {
  if(cFrame == NULL || ((ttLibC_Frame *)cFrame)->type != frameType_bgr || buffer == NULL || buffer_size < 5) {
    return;
  }
  ttLibC_Bgr *bgr = (ttLibC_Bgr *)cFrame;
  switch(bgr->type) {
  case BgrType_abgr: strcpy(buffer, "abgr"); break;
  case BgrType_argb: strcpy(buffer, "argb"); break;
  case BgrType_bgr:  strcpy(buffer, "bgr");  break;
  case BgrType_bgra: strcpy(buffer, "bgra"); break;
  case BgrType_rgb:  strcpy(buffer, "rgb");  break;
  case BgrType_rgba: strcpy(buffer, "rgba"); break;
  default:
    break;
  }
}

uint32_t BgrFrame_getWidthStride(void *frame) {
	if(frame == NULL || ((ttLibC_Frame *)frame)->type != frameType_bgr) {
		return 0;
	}
  ttLibC_Bgr *bgr = (ttLibC_Bgr *)frame;
  return bgr->width_stride;
}

void Flv1Frame_getFlv1Type(void *cFrame, char *buffer, size_t buffer_size) {
  if(cFrame == NULL || ((ttLibC_Frame *)cFrame)->type != frameType_bgr || buffer == NULL || buffer_size < 16) {
    return;
  }
  ttLibC_Flv1 *flv1 = (ttLibC_Flv1 *)cFrame;
  switch(flv1->type) {
  case Flv1Type_disposableInner: strcpy(buffer, "disposableInner"); break;
  case Flv1Type_inner:           strcpy(buffer, "inner");           break;
  case Flv1Type_intra:           strcpy(buffer, "intra");           break;
  default:
    break;
  }
}

void H264Frame_getH264Type(void *cFrame, char *buffer, size_t buffer_size) {
  if(cFrame == NULL || ((ttLibC_Frame *)cFrame)->type != frameType_h264 || buffer == NULL || buffer_size < 11) {
    return;
  }
  ttLibC_H264 *h264 = (ttLibC_H264 *)cFrame;
  switch(h264->type) {
  case H264Type_configData: strcpy(buffer, "configData"); break;
  case H264Type_slice:      strcpy(buffer, "slice");      break;
  case H264Type_sliceIDR:   strcpy(buffer, "sliceIDR");   break;
  case H264Type_unknown:    strcpy(buffer, "unknown");    break;
  default:
    break;
  }
}

uint32_t H264Frame_getH264FrameType(void *cFrame) {
  if(cFrame == NULL || ((ttLibC_Frame *)cFrame)->type != frameType_h264) {
    return -1;
  }
  ttLibC_H264 *h264 = (ttLibC_H264 *)cFrame;
  return h264->frame_type;
}

bool H264Frame_isDisposable(void *cFrame) {
  if(cFrame == NULL || ((ttLibC_Frame *)cFrame)->type != frameType_h264) {
    return false;
  }
  ttLibC_H264 *h264 = (ttLibC_H264 *)cFrame;
  return h264->is_disposable;
}

void H265Frame_getH265Type(void *cFrame, char *buffer, size_t buffer_size) {
  if(cFrame == NULL || ((ttLibC_Frame *)cFrame)->type != frameType_h265 || buffer == NULL || buffer_size < 11) {
    return;
  }
  ttLibC_H265 *h265 = (ttLibC_H265 *)cFrame;
  switch(h265->type) {
  case H265Type_configData: strcpy(buffer, "configData"); break;
  case H265Type_slice:      strcpy(buffer, "slice");      break;
  case H265Type_sliceIDR:   strcpy(buffer, "sliceIDR");   break;
  case H265Type_unknown:    strcpy(buffer, "unknown");    break;
  default:
    break;
  }
}

uint32_t H265Frame_getH265FrameType(void *cFrame) {
  if(cFrame == NULL || ((ttLibC_Frame *)cFrame)->type != frameType_h265) {
    return -1;
  }
  ttLibC_H265 *h265 = (ttLibC_H265 *)cFrame;
  return h265->frame_type;
}

bool H265Frame_isDisposable(void *cFrame) {
  if(cFrame == NULL || ((ttLibC_Frame *)cFrame)->type != frameType_h265) {
    return false;
  }
  ttLibC_H265 *h265 = (ttLibC_H265 *)cFrame;
  return h265->is_disposable;
}

void TheoraFrame_getTheoraType(void *cFrame, char *buffer, size_t buffer_size) {
  if(cFrame == NULL || ((ttLibC_Frame *)cFrame)->type != frameType_theora || buffer == NULL || buffer_size < 22) {
    return;
  }
  ttLibC_Theora *theora = (ttLibC_Theora *)cFrame;
  switch(theora->type) {
  case TheoraType_identificationHeaderDecodeFrame: strcpy(buffer, "identificationHeaders"); break;
  case TheoraType_commentHeaderFrame:              strcpy(buffer, "commentHeader");         break;
  case TheoraType_setupHeaderFrame:                strcpy(buffer, "setupHeader");           break;
  case TheoraType_intraFrame:                      strcpy(buffer, "intra");                 break;
  case TheoraType_innerFrame:                      strcpy(buffer, "inner");                 break;
  default:
    break;
  }
}

uint64_t TheoraFrame_getGranulePos(void *cFrame) {
  if(cFrame == NULL || ((ttLibC_Frame *)cFrame)->type != frameType_theora) {
    return false;
  }
  ttLibC_Theora *theora = (ttLibC_Theora *)cFrame;
  return theora->granule_pos;
}

void Yuv420Frame_getYuv420Type(void *cFrame, char *buffer, size_t buffer_size) {
  if(cFrame == NULL || ((ttLibC_Frame *)cFrame)->type != frameType_yuv420 || buffer == NULL || buffer_size < 17) {
    return;
  }
  ttLibC_Yuv420 *yuv420 = (ttLibC_Yuv420 *)cFrame;
  switch(yuv420->type) {
  case Yuv420Type_planar:     strcpy(buffer, "yuv420"); break;
  case Yuv420Type_semiPlanar: strcpy(buffer, "yuv420SemiPlanar"); break;
  case Yvu420Type_planar:     strcpy(buffer, "yvu420"); break;
  case Yvu420Type_semiPlanar: strcpy(buffer, "yvu420SemiPlanar"); break;
  default:
    break;
  }
}

uint32_t Yuv420Frame_getYStride(void *cFrame) {
  if(cFrame == NULL || ((ttLibC_Frame *)cFrame)->type != frameType_yuv420) {
    return false;
  }
  ttLibC_Yuv420 *yuv420 = (ttLibC_Yuv420 *)cFrame;
  return yuv420->y_stride;
}

uint32_t Yuv420Frame_getUStride(void *cFrame) {
  if(cFrame == NULL || ((ttLibC_Frame *)cFrame)->type != frameType_yuv420) {
    return false;
  }
  ttLibC_Yuv420 *yuv420 = (ttLibC_Yuv420 *)cFrame;
  return yuv420->u_stride;
}

uint32_t Yuv420Frame_getVStride(void *cFrame) {
  if(cFrame == NULL || ((ttLibC_Frame *)cFrame)->type != frameType_yuv420) {
    return false;
  }
  ttLibC_Yuv420 *yuv420 = (ttLibC_Yuv420 *)cFrame;
  return yuv420->v_stride;
}

extern bool ttLibGoDataCallback(void *ptr, void *data, size_t data_size);

void *newGoFrame(uint64_t pts, uint64_t dts, uint32_t timebase, uint32_t id,
    uint32_t width, uint32_t height,
    uint32_t sampleRate, uint32_t sampleNum, uint32_t channelNum,
    uint64_t dataPos, uint32_t widthStride,
    uint64_t yDataPos, uint32_t yStride,
    uint64_t uDataPos, uint32_t uStride,
    uint64_t vDataPos, uint32_t vStride) {
  ttLibGoFrame *frame = new ttLibGoFrame();
  frame->pts         = pts;
  frame->dts         = dts;
  frame->timebase    = timebase;
  frame->id          = id;
  frame->width       = width;
  frame->height      = height;
  frame->sampleRate  = sampleRate;
  frame->sampleNum   = sampleNum;
  frame->channelNum  = channelNum;
  frame->dataPos     = dataPos;
  frame->widthStride = widthStride;
  frame->yDataPos    = yDataPos;
  frame->yStride     = yStride;
  frame->uDataPos    = uDataPos;
  frame->uStride     = uStride;
  frame->vDataPos    = vDataPos;
  frame->vStride     = vStride;
  return frame;
}
void deleteGoFrame(void *frame) {
  ttLibGoFrame *goFrame = reinterpret_cast<ttLibGoFrame*>(frame);
  delete goFrame;
}

bool Frame_getBinaryBuffer(uintptr_t ptr, void *frame) {
  if(ttLibGo_Frame_close == nullptr || ttLibGo_Frame_clone == nullptr
  || ttLibGo_Bgr_getMinimumBinaryBuffer == nullptr
  || ttLibGo_Yuv420_getMinimumBinaryBuffer == nullptr) {
    return false;
  }
  if(frame == NULL) {
    return false;
  }
  ttLibC_Frame *f = (ttLibC_Frame *)frame;
	switch(f->type) {
	case frameType_bgr:
    {
      return (*ttLibGo_Bgr_getMinimumBinaryBuffer)(f, ttLibGoDataCallback, (void *)ptr);
    }
    break;
	case frameType_yuv420:
    {
      return (*ttLibGo_Yuv420_getMinimumBinaryBuffer)(f, ttLibGoDataCallback, (void *)ptr);
    }
		break;
	case frameType_pcmS16:
	case frameType_pcmF32:
		{
      ttLibC_Frame *cloned = (ttLibC_Frame *)(*ttLibGo_Frame_clone)(nullptr, f);
      if(cloned == nullptr) {
        return false;
      }
  		bool result = ttLibGoDataCallback((void *)ptr, cloned->data, cloned->buffer_size);
      (*ttLibGo_Frame_close)((void **)&cloned);
      return result;
		}
		break;
	default:
		// とりあえずこれでいいはず。
		return ttLibGoDataCallback((void *)ptr, f->data, f->buffer_size);
	}
  return false;
}

void *BgrFrame_fromBinary(void *data, size_t data_size,
  uint32_t id, uint64_t pts, uint32_t timebase,
  const char *bgr_type, uint32_t width, uint32_t height, uint32_t width_stride) {
  // まず、emptyFrameを作って、それからデータをコピーする形にするべきかな・・・
  if(ttLibGo_Bgr_makeEmptyFrame == nullptr) {
    return nullptr;
  }
  string bgrTypeName(bgr_type);
  ttLibC_Bgr_Type bgrType = BgrType_bgr;
  if(bgrTypeName == "bgr") {
    bgrType = BgrType_bgr;
  }
  else if(bgrTypeName == "abgr") {
    bgrType = BgrType_abgr;
  }
  else if(bgrTypeName == "bgra") {
    bgrType = BgrType_bgra;
  }
  else if(bgrTypeName == "rgb") {
    bgrType = BgrType_rgb;
  }
  else if(bgrTypeName == "argb") {
    bgrType = BgrType_argb;
  }
  else if(bgrTypeName == "rgba") {
    bgrType = BgrType_rgba;
  }
  ttLibC_Bgr *bgr = (ttLibC_Bgr *)(*ttLibGo_Bgr_makeEmptyFrame)(bgrType, width, height);
  if(bgr != nullptr) {
    bgr->inherit_super.inherit_super.id       = id;
    bgr->inherit_super.inherit_super.pts      = pts;
    bgr->inherit_super.inherit_super.timebase = timebase;
    uint8_t *dst_data = bgr->data;
    uint8_t *src_data = (uint8_t *)data;
    for(int i = 0;i < height;++ i) {
      memcpy(dst_data, src_data, width);
      dst_data += bgr->width_stride;
      src_data += width_stride;
    }
  }
  return bgr;
}

void *Flv1Frame_fromBinary(void *data, size_t data_size,
  uint32_t id, uint64_t pts, uint32_t timebase) {
  if(ttLibGo_Flv1_getFrame == nullptr) {
    return nullptr;
  }
  ttLibC_Flv1 *flv1 = (ttLibC_Flv1 *)(*ttLibGo_Flv1_getFrame)(
    nullptr,
    (uint8_t *)data,
    data_size,
    false,
    pts,
    timebase);
  if(flv1 != nullptr) {
    flv1->inherit_super.inherit_super.id = id;
  }
  return flv1;
}

void *H264Frame_fromBinary(void *data, size_t data_size,
  uint32_t id, uint64_t pts, uint32_t timebase) {
  if(ttLibGo_H264_getFrame == nullptr) {
    return nullptr;
  }
  ttLibC_H264 *h264 = (ttLibC_H264 *)(*ttLibGo_H264_getFrame)(
    nullptr,
    (uint8_t *)data,
    data_size,
    false,
    pts,
    timebase);
  if(h264 != nullptr) {
    h264->inherit_super.inherit_super.id = id;
  }
  return h264;
}

void *H265Frame_fromBinary(void *data, size_t data_size,
  uint32_t id, uint64_t pts, uint32_t timebase) {
  if(ttLibGo_H265_getFrame == nullptr) {
    return nullptr;
  }
  ttLibC_H265 *h265 = (ttLibC_H265 *)(*ttLibGo_H265_getFrame)(
    nullptr,
    (uint8_t *)data,
    data_size,
    false,
    pts,
    timebase);
  if(h265 != nullptr) {
    h265->inherit_super.inherit_super.id = id;
  }
  return h265;
}

void *JpegFrame_fromBinary(void *data, size_t data_size,
  uint32_t id, uint64_t pts, uint32_t timebase) {
  if(ttLibGo_Jpeg_getFrame == nullptr) {
    return nullptr;
  }
  ttLibC_Jpeg *jpeg = (ttLibC_Jpeg *)(*ttLibGo_Jpeg_getFrame)(
    nullptr,
    (uint8_t *)data,
    data_size,
    false,
    pts,
    timebase);
  if(jpeg != nullptr) {
    jpeg->inherit_super.inherit_super.id = id;
  }
  return jpeg;
}

void *PngFrame_fromBinary(void *data, size_t data_size,
  uint32_t id, uint64_t pts, uint32_t timebase) {
  if(ttLibGo_Png_getFrame == nullptr) {
    return nullptr;
  }
  ttLibC_Png *png = (ttLibC_Png *)(*ttLibGo_Png_getFrame)(
    nullptr,
    (uint8_t *)data,
    data_size,
    false,
    pts,
    timebase);
  if(png != nullptr) {
    png->inherit_super.inherit_super.id = id;
  }
  return png;
}

void *TheoraFrame_fromBinary(void *data, size_t data_size,
  uint32_t id, uint64_t pts, uint32_t timebase, uint32_t width, uint32_t height) {
  if(ttLibGo_Theora_getWidth == nullptr
  || ttLibGo_Theora_getHeight == nullptr
  || ttLibGo_Theora_make == nullptr) {
    return nullptr;
  }
  uint8_t *u8data = (uint8_t *)data;
	uint8_t first_byte = *u8data;
	ttLibC_Theora_Type type = TheoraType_innerFrame;
	if((first_byte & 0x80) != 0) {
		switch(first_byte) {
		case 0x80: // info
			type = TheoraType_identificationHeaderDecodeFrame;
			break;
		case 0x81: // comment
			type = TheoraType_commentHeaderFrame;
			break;
		case 0x82: // setup
			type = TheoraType_setupHeaderFrame;
			break;
		default:
			ERR_PRINT("unknown theora header frame.");
			return NULL;
		}
	}
	else {
		if((first_byte & 0x40) == 0) {
			type = TheoraType_intraFrame;
		}
    if(width == 0) {
	    width = (*ttLibGo_Theora_getWidth)(nullptr, u8data, data_size);
    }
    if(height == 0) {
	    height = (*ttLibGo_Theora_getHeight)(nullptr, u8data, data_size);
    }
		if(width == 0 || height == 0) {
			return NULL;
		}
	}
	ttLibC_Theora *theora = (ttLibC_Theora *)(*ttLibGo_Theora_make)(
			nullptr,
			type,
			width,
			height,
			data,
			data_size,
			false,
			pts,
			timebase);
  if(theora != nullptr) {
    theora->inherit_super.inherit_super.id = id;
  }
  return theora;
}

void *Vp6Frame_fromBinary(void *data, size_t data_size,
  uint32_t id, uint64_t pts, uint32_t timebase, uint32_t width, uint32_t height, uint8_t adjustment) {
  uint8_t *u8data = (uint8_t *)data;
  if(ttLibGo_Vp6_isKey == nullptr
  || ttLibGo_Vp6_getWidth == nullptr
  || ttLibGo_Vp6_getHeight == nullptr
  || ttLibGo_Vp6_make == nullptr) {
    return nullptr;
  }
	if(data_size <= 1) {
		return NULL;
	}
	bool is_key = (*ttLibGo_Vp6_isKey)(u8data, data_size);
  if(width == 0) {
	  width = (*ttLibGo_Vp6_getWidth)(nullptr, u8data, data_size, adjustment);
  }
  if(height == 0) {
	  height = (*ttLibGo_Vp6_getHeight)(nullptr, u8data, data_size, adjustment);
  }
	if(width == 0 || height == 0) {
		return NULL;
	}
	ttLibC_Vp6 *vp6 = (ttLibC_Vp6 *)(*ttLibGo_Vp6_make)(
			nullptr,
			is_key ? videoType_key : videoType_inner,
			width,
			height,
			u8data, data_size,
			false,
			pts,
			timebase);
  if(vp6 != nullptr) {
    vp6->inherit_super.inherit_super.id = id;
  }
  return vp6;
}

void *Vp8Frame_fromBinary(void *data, size_t data_size,
  uint32_t id, uint64_t pts, uint32_t timebase, uint32_t width, uint32_t height) {
  if(ttLibGo_Vp8_isKey == nullptr
  || ttLibGo_Vp8_getWidth == nullptr
  || ttLibGo_Vp8_getHeight == nullptr
  || ttLibGo_Vp8_make == nullptr) {
    return nullptr;
  }
  uint8_t *u8data = (uint8_t *)data;
	if(data_size <= 1) {
		return NULL;
	}
	bool is_key = (*ttLibGo_Vp8_isKey)(u8data, data_size);
  if(width == 0) {
	  width = (*ttLibGo_Vp8_getWidth)(nullptr, u8data, data_size);
  }
  if(height == 0) {
	  height = (*ttLibGo_Vp8_getHeight)(nullptr, u8data, data_size);
  }
	if(width == 0 || height == 0) {
		return NULL;
	}
	ttLibC_Vp8 *vp8 = (ttLibC_Vp8 *)(*ttLibGo_Vp8_make)(
			nullptr,
			is_key ? videoType_key : videoType_inner,
			width,
			height,
			u8data, data_size,
			false,
			pts,
			timebase);
  if(vp8 != nullptr) {
    vp8->inherit_super.inherit_super.id = id;
  }
  return vp8;
}

void *Vp9Frame_fromBinary(void *data, size_t data_size,
  uint32_t id, uint64_t pts, uint32_t timebase, uint32_t width, uint32_t height) {
  if(ttLibGo_Vp9_isKey == nullptr
  || ttLibGo_Vp9_getWidth == nullptr
  || ttLibGo_Vp9_getHeight == nullptr
  || ttLibGo_Vp9_make == nullptr) {
    return nullptr;
  }
  uint8_t *u8data = (uint8_t *)data;
	if(data_size <= 1) {
		return NULL;
	}
	bool is_key = (*ttLibGo_Vp9_isKey)(u8data, data_size);
  if(width == 0) {
	  width = (*ttLibGo_Vp9_getWidth)(nullptr, u8data, data_size);
  }
  if(height == 0) {
	  height = (*ttLibGo_Vp9_getHeight)(nullptr, u8data, data_size);
  }
	if(width == 0 || height == 0) {
		return NULL;
	}
	ttLibC_Vp9 *vp9 = (ttLibC_Vp9 *)(*ttLibGo_Vp9_make)(
			nullptr,
			is_key ? videoType_key : videoType_inner,
			width,
			height,
			u8data, data_size,
			false,
			pts,
			timebase);
  if(vp9 != nullptr) {
    vp9->inherit_super.inherit_super.id = id;
  }
  return vp9;
}

void *Yuv420Frame_fromPlaneBinaries(uint32_t id, uint64_t pts, uint32_t timebase,
  const char *yuv_type, uint32_t width, uint32_t height,
  void *y_data, uint32_t y_stride,
  void *u_data, uint32_t u_stride,
  void *v_data, uint32_t v_stride) {
  if(ttLibGo_Yuv420_makeEmptyFrame == nullptr) {
    return nullptr;
  }
  string yuvTypeName(yuv_type);
  ttLibC_Yuv420_Type yuvType = Yuv420Type_planar;
  uint32_t y_step = 1;
  uint32_t u_step = 1;
  uint32_t v_step = 1;
  uint32_t half_width  = ((width  + 1) >> 1);
  uint32_t half_height = ((height + 1) >> 1);
  if(yuvTypeName == "yuv420") {
    yuvType = Yuv420Type_planar;
  }
  else if(yuvTypeName == "yuv420SemiPlanar") {
    yuvType = Yuv420Type_semiPlanar;
    u_step = 2;
    v_step = 2;
  }
  else if(yuvTypeName == "yvu420") {
    yuvType = Yvu420Type_planar;
  }
  else if(yuvTypeName == "yvu420SemiPlanar") {
    yuvType = Yvu420Type_semiPlanar;
    u_step = 2;
    v_step = 2;
  }
  ttLibC_Yuv420 *yuv = (ttLibC_Yuv420 *)(*ttLibGo_Yuv420_makeEmptyFrame)(yuvType, width, height);
  if(yuv != nullptr) {
    yuv->inherit_super.inherit_super.id       = id;
    yuv->inherit_super.inherit_super.pts      = pts;
    yuv->inherit_super.inherit_super.timebase = timebase;
    uint8_t *dst_y_data = yuv->y_data;
    uint8_t *dst_u_data = yuv->u_data;
    uint8_t *dst_v_data = yuv->v_data;
    uint8_t *src_y_data = (uint8_t *)y_data;
    uint8_t *src_u_data = (uint8_t *)u_data;
    uint8_t *src_v_data = (uint8_t *)v_data;
    for(int i = 0;i < height;++ i) {
      memcpy(dst_y_data, src_y_data, width);
      dst_y_data += yuv->y_stride;
      src_y_data += y_stride;
    }
    for(int i = 0;i < half_height;++ i) {
      uint8_t *dst_u = dst_u_data;
      uint8_t *src_u = src_u_data;
      uint8_t *dst_v = dst_v_data;
      uint8_t *src_v = src_v_data;
      for(int j = 0;j < half_width;++ j) {
        *dst_u = *src_u;
        dst_u += u_step;
        src_u += u_step;
        *dst_v = *src_v;
        dst_v += v_step;
        src_v += v_step;
      }
      dst_u_data += yuv->u_stride;
      src_u_data += u_stride;
      dst_v_data += yuv->v_stride;
      src_v_data += v_stride;
    }
  }
  return yuv;
}

typedef struct {
	ttLibC_Aac inherit_super;
	uint64_t dsi_info; // for raw, need to have dsi_information.
} ttLibC_Frame_Audio_Aac_;

typedef ttLibC_Frame_Audio_Aac_ ttLibC_Aac_;

static uint32_t aac_sample_rate_table[] = {
		96000, 88200, 64000, 48000, 44100, 32000, 24000, 22050, 16000, 12000, 11025, 8000
};

void *AacFrame_fromBinary(void *data, size_t data_size,
  uint32_t id, uint64_t pts, uint32_t timebase,
  void *dsiFrame) {
  if(ttLibGo_Aac_make == nullptr
  || ttLibGo_ByteReader_make == nullptr
  || ttLibGo_ByteReader_bit == nullptr
  || ttLibGo_ByteReader_close == nullptr) {
    return nullptr;
  }
  // データを確認する。
  if(dsiFrame != nullptr) {
    ttLibC_Aac_ *dsi = reinterpret_cast<ttLibC_Aac_ *>(dsiFrame);
    // dsiをベースにデータを復元して応答する
    uint8_t *u8data = (uint8_t *)data;
    if(data_size < 2) {
      return nullptr;
    }
    if(u8data[0] == 0xFF
    && (u8data[1] & 0xF0) == 0xF0) {
      // rawデータなのに、なぜか入力 binaryがadtsっぽいので、nullptr応答しておく
      return nullptr;
    }
    ttLibC_Aac *aac = (ttLibC_Aac *)(*ttLibGo_Aac_make)(
      nullptr,
      AacType_raw,
      dsi->inherit_super.inherit_super.sample_rate,
      1024,
      dsi->inherit_super.inherit_super.channel_num,
      data,
      data_size,
      false,
      pts,
      timebase,
      dsi->dsi_info);
    if(aac != nullptr) {
      aac->inherit_super.inherit_super.id = id;
    }
    return aac;
  }
  // あとはこっち。
  ttLibC_ByteReader *reader = (ttLibC_ByteReader *)(*ttLibGo_ByteReader_make)(data, data_size, ByteUtilType_default);
  if((*ttLibGo_ByteReader_bit)(reader, 12) != 0xFFF) {
    (*ttLibGo_ByteReader_close)((void **)&reader);
    // dsi情報を取り出し
		ttLibC_ByteReader *reader = (ttLibC_ByteReader *)(*ttLibGo_ByteReader_make)(data, data_size, ByteUtilType_default);
		uint64_t dsi_info;
		uint32_t bit_size = 0;
		memcpy(&dsi_info, data, data_size);
		uint32_t object_type = (*ttLibGo_ByteReader_bit)(reader, 5);
		bit_size += 5;
		if(object_type == 31) {
			object_type = (*ttLibGo_ByteReader_bit)(reader, 6);
			bit_size += 6;
		}
		uint32_t sample_rate_index = (*ttLibGo_ByteReader_bit)(reader, 4);
		bit_size += 4;
		uint32_t sample_rate = 44100;
		if(sample_rate_index == 15) {
			LOG_PRINT("sample_rate is not in index_table.");
			sample_rate = (*ttLibGo_ByteReader_bit)(reader, 24);
			bit_size += 24;
		}
		else {
			sample_rate = aac_sample_rate_table[sample_rate_index];
		}
		uint32_t channel_num = (*ttLibGo_ByteReader_bit)(reader, 4);
		bit_size += 4;
		uint32_t buffer_size = (uint32_t)((bit_size + 7) / 8);
    (*ttLibGo_ByteReader_close)((void **)&reader);
		ttLibC_Aac *aac = (ttLibC_Aac *)(*ttLibGo_Aac_make)(
				nullptr,
				AacType_dsi,
				sample_rate,
				0,
				channel_num,
				data,
				data_size,
				false,
				pts,
				timebase,
				dsi_info);
		if(aac != NULL) {
			aac->inherit_super.inherit_super.buffer_size = buffer_size;
      aac->inherit_super.inherit_super.id = id;
		}
		return aac;
  }
  // adtsの取り出し
	(*ttLibGo_ByteReader_bit)(reader, 1);
	(*ttLibGo_ByteReader_bit)(reader, 2);
	(*ttLibGo_ByteReader_bit)(reader, 1);
	(*ttLibGo_ByteReader_bit)(reader, 2);
	uint32_t sample_rate_index = (*ttLibGo_ByteReader_bit)(reader, 4);
	uint32_t sample_rate = aac_sample_rate_table[sample_rate_index];
	(*ttLibGo_ByteReader_bit)(reader, 1);
	uint32_t channel_num = (*ttLibGo_ByteReader_bit)(reader, 3);
	(*ttLibGo_ByteReader_bit)(reader, 1);
	(*ttLibGo_ByteReader_bit)(reader, 1);
	(*ttLibGo_ByteReader_bit)(reader, 1);
	(*ttLibGo_ByteReader_bit)(reader, 1);
	uint32_t frame_size = (*ttLibGo_ByteReader_bit)(reader, 13);
	(*ttLibGo_ByteReader_bit)(reader, 11);
	(*ttLibGo_ByteReader_bit)(reader, 2);
	if(reader->error != Error_noError) {
		//LOG_ERROR(reader->error); // この動作、ttLibCの内部関数をkickしていたため、buildが失敗してた。
    LOG_PRINT("error:%d", reader->error);
    (*ttLibGo_ByteReader_close)((void **)&reader);
    return nullptr;
	}
    (*ttLibGo_ByteReader_close)((void **)&reader);
  // fffで始まってたらadtsとして処理
  ttLibC_Aac *aac = (ttLibC_Aac *)(*ttLibGo_Aac_make)(
			nullptr,
			AacType_adts,
			sample_rate,
			1024,
			channel_num,
			data,
			frame_size,
			false,
			pts,
			timebase,
			0);
  if(aac != nullptr) {
    aac->inherit_super.inherit_super.id = id;
  }
  return aac;
}

void *AdpcmImaWavFrame_fromBinary(void *data, size_t data_size, 
  uint32_t id, uint64_t pts, uint32_t timebase,
  uint32_t sample_rate, uint32_t sample_num, uint32_t channel_num) {
  if(ttLibGo_AdpcmImaWav_make == nullptr) {
    return nullptr;
  }
  ttLibC_AdpcmImaWav *adpcm = (ttLibC_AdpcmImaWav *)(*ttLibGo_AdpcmImaWav_make)(
    nullptr,
    sample_rate,
    sample_num,
    channel_num,
    data,
    data_size,
    false,
    pts,
    timebase);
  if(adpcm != nullptr) {
    adpcm->inherit_super.inherit_super.id = id;
  }
  return adpcm;
}

void *Mp3Frame_fromBinary(void *data, size_t data_size,
  uint32_t id, uint64_t pts, uint32_t timebase) {
  if(ttLibGo_Mp3_getFrame == nullptr) {
    return nullptr;
  }
  ttLibC_Mp3 *mp3 = (ttLibC_Mp3 *)(*ttLibGo_Mp3_getFrame)(
    nullptr,
    data,
    data_size,
    false,
    pts,
    timebase);
  if(mp3 != nullptr) {
    mp3->inherit_super.inherit_super.id = id;
  }
  return mp3;
}

void *NellymoserFrame_fromBinary(void *data, size_t data_size, 
  uint32_t id, uint64_t pts, uint32_t timebase,
  uint32_t sample_rate, uint32_t sample_num, uint32_t channel_num) {
  if(ttLibGo_Nellymoser_make == nullptr) {
    return nullptr;
  }
  ttLibC_Nellymoser *nellymoser = (ttLibC_Nellymoser *)(*ttLibGo_Nellymoser_make)(
    nullptr,
    sample_rate,
    sample_num,
    channel_num,
    data,
    data_size,
    false,
    pts,
    timebase);
  if(nellymoser != nullptr) {
    nellymoser->inherit_super.inherit_super.id = id;
  }
  return nellymoser;
}

void *OpusFrame_fromBinary(void *data, size_t data_size,
  uint32_t id, uint64_t pts, uint32_t timebase) {
  if(ttLibGo_Opus_getFrame == nullptr) {
    return nullptr;
  }
  ttLibC_Opus *opus = (ttLibC_Opus *)(*ttLibGo_Opus_getFrame)(
    nullptr,
    data,
    data_size,
    false,
    pts,
    timebase);
  if(opus != nullptr) {
    opus->inherit_super.inherit_super.id = id;
  }
  return opus;
}

void *PcmAlawsFrame_fromBinary(void *data, size_t data_size, 
  uint32_t id, uint64_t pts, uint32_t timebase,
  uint32_t sample_rate, uint32_t sample_num, uint32_t channel_num) {
  if(ttLibGo_PcmAlaw_make == nullptr) {
    return nullptr;
  }
  ttLibC_PcmAlaw *pcmAlaw = (ttLibC_PcmAlaw *)(*ttLibGo_PcmAlaw_make)(
    nullptr,
    sample_rate,
    sample_num,
    channel_num,
    data,
    data_size,
    false,
    pts,
    timebase);
  if(pcmAlaw != nullptr) {
    pcmAlaw->inherit_super.inherit_super.id = id;
  }
  return pcmAlaw;
}

void *PcmF32Frame_fromBinary(void *data, size_t data_size, 
  uint32_t id, uint64_t pts, uint32_t timebase,
  const char *pcmF32_type, uint32_t sample_rate, uint32_t sample_num, uint32_t channel_num,
  uint32_t l_index, uint32_t r_index) {
  if(ttLibGo_PcmF32_make == nullptr) {
    return nullptr;
  }
  string pcmF32Name(pcmF32_type);
  ttLibC_PcmF32_Type pcmF32Type = PcmF32Type_interleave;
  uint32_t l_stride = 0;
  uint32_t r_stride = 0;
  if(pcmF32Name == "interleave") {
    pcmF32Type = PcmF32Type_interleave;
    l_stride = data_size;
    r_index = 0;
    r_stride = 0;
  }
  else if(pcmF32Name == "planar") {
    pcmF32Type = PcmF32Type_planar;
    l_stride = r_index;
    r_stride = data_size - r_index;
  }
  uint8_t *u8data = (uint8_t *)data;
  ttLibC_PcmF32 *pcmF32 = (ttLibC_PcmF32 *)(*ttLibGo_PcmF32_make)(
    nullptr,
    pcmF32Type,
    sample_rate,
    sample_num,
    channel_num,
    data,
    data_size,
    u8data + l_index,
    l_stride,
    r_index == 0 ? nullptr : u8data + r_index,
    r_stride,
    false,
    pts,
    timebase);
  if(pcmF32 != nullptr) {
    pcmF32->inherit_super.inherit_super.id = id;
  }
  return pcmF32;
}

void *PcmMulawsFrame_fromBinary(void *data, size_t data_size, 
  uint32_t id, uint64_t pts, uint32_t timebase,
  uint32_t sample_rate, uint32_t sample_num, uint32_t channel_num) {
  if(ttLibGo_PcmMulaw_make == nullptr) {
    return nullptr;
  }
  ttLibC_PcmMulaw *pcmMulaw = (ttLibC_PcmMulaw *)(*ttLibGo_PcmMulaw_make)(
    nullptr,
    sample_rate,
    sample_num,
    channel_num,
    data,
    data_size,
    false,
    pts,
    timebase);
  if(pcmMulaw != nullptr) {
    pcmMulaw->inherit_super.inherit_super.id = id;
  }
  return pcmMulaw;
}

void *PcmS16Frame_fromBinary(void *data, size_t data_size,
  uint32_t id, uint64_t pts, uint32_t timebase,
  const char *pcmS16_type, uint32_t sample_rate, uint32_t sample_num, uint32_t channel_num,
  uint32_t l_index, uint32_t r_index) {
  if(ttLibGo_PcmS16_make == nullptr) {
    return nullptr;
  }
  string pcmS16Name(pcmS16_type);
  ttLibC_PcmS16_Type pcmS16Type = PcmS16Type_littleEndian;
  uint32_t l_stride = 0;
  uint32_t r_stride = 0;
  if(pcmS16Name == "bigEndian") {
    pcmS16Type = PcmS16Type_bigEndian;
    l_stride = data_size;
    r_index = 0;
    r_stride = 0;
  }
  else if(pcmS16Name == "bigEndianPlanar") {
    pcmS16Type = PcmS16Type_bigEndian_planar;
    l_stride = r_index;
    r_stride = data_size - r_index;
  }
  else if(pcmS16Name == "littleEndian") {
    pcmS16Type = PcmS16Type_littleEndian;
    l_stride = data_size;
    r_index = 0;
    r_stride = 0;
  }
  else if(pcmS16Name == "littleEndianPlanar") {
    pcmS16Type = PcmS16Type_littleEndian_planar;
    l_stride = r_index;
    r_stride = data_size - r_index;
  }
  uint8_t *u8data = (uint8_t *)data;
  ttLibC_PcmS16 *pcmS16 = (ttLibC_PcmS16 *)(*ttLibGo_PcmS16_make)(
    nullptr,
    pcmS16Type,
    sample_rate,
    sample_num,
    channel_num,
    data,
    data_size,
    u8data + l_index,
    l_stride,
    r_index == 0 ? nullptr : u8data + r_index,
    r_stride,
    false,
    pts,
    timebase);
  if(pcmS16 != nullptr) {
    pcmS16->inherit_super.inherit_super.id = id;
  }
  return pcmS16;
}

void *SpeexFrame_fromBinary(void *data, size_t data_size, 
  uint32_t id, uint64_t pts, uint32_t timebase,
  uint32_t sample_rate, uint32_t sample_num, uint32_t channel_num) {
  if(ttLibGo_Speex_make == nullptr) {
    return nullptr;
  }
  if(data_size < 8) {
    return nullptr;
  }
  uint8_t *buf = (uint8_t *)data;
  ttLibC_Speex *speex = nullptr;
  if(buf[0] == 'S'
  && buf[1] == 'p'
  && buf[2] == 'e'
  && buf[3] == 'e'
  && buf[4] == 'x'
  && buf[5] == ' '
  && buf[6] == ' '
  && buf[7] == ' ') {
    if(data_size < 0x44) {
      return nullptr;
    }
    uint32_t *buf_int = (uint32_t *)buf;
    uint32_t sample_rate = le_uint32_t(*buf_int);
    buf_int += 3;
    uint32_t channel_num = le_uint32_t(*buf_int);
    buf_int += 2;
    uint32_t sample_num = le_uint32_t(*buf_int);
    buf_int += 2;
    sample_num *= le_uint32_t(*buf_int);
    speex = (ttLibC_Speex *)(ttLibGo_Speex_make)(
        nullptr,
        SpeexType_header,
        sample_rate,
        sample_num,
        channel_num,
        data,
        data_size,
        false,
        pts,
        timebase);
  }
  else if(buf[2] == 0x00 && buf[3] == 0x00) {
    // commentであるかの判定が欲しい。
    // とりあえず3byte 4byteが00 00になってたら、コメントということで・・・
    // 強制コメント扱いにしておく。
    // speexのcommentは先頭が4byte little endianのコメントデータ長になってる。非常に長いコメントが存在しない
    // 通常のフレームだと何かしらの値が入ることから、これで判断してお茶を濁しておこう。
    speex = (ttLibC_Speex *)(ttLibGo_Speex_make)(
        nullptr,
        SpeexType_comment,
        sample_rate,
        0,
        channel_num,
        data,
        data_size,
        false,
        pts,
        timebase);
  }
  else {
    // どれにも当てはまらなかったらframeということで・・・
    speex = (ttLibC_Speex *)(ttLibGo_Speex_make)(
        nullptr,
        SpeexType_frame,
        sample_rate,
        sample_num,
        channel_num,
        data,
        data_size,
        false,
        pts,
        timebase);
  }
  if(speex != nullptr) {
    speex->inherit_super.inherit_super.id = id;
  }
  return speex;
}

void *VorbisFrame_fromBinary(void *data, size_t data_size, 
  uint32_t id, uint64_t pts, uint32_t timebase,
  uint32_t sample_rate, uint32_t sample_num, uint32_t channel_num) {
  if(ttLibGo_Vorbis_make == nullptr) {
    return nullptr;
  }
  uint8_t *buf = (uint8_t *)data;
  ttLibC_Vorbis *vorbis = nullptr;
  if(data_size > 7) {
    if(buf[1] == 'v'
		&& buf[2] == 'o'
		&& buf[3] == 'r'
		&& buf[4] == 'b'
		&& buf[5] == 'i'
		&& buf[6] == 's') {
      switch(buf[0]) {
        case 0x01: // identification
          {
            if(data_size < 29) {
              // too small data size.
              return NULL;
            }
            uint32_t channel_num = buf[11];
            uint32_t *buf32 = (uint32_t *)(buf + 12);
            uint32_t sample_rate = le_uint32_t(*buf32);
            uint32_t block0 = buf[28] & 0x0F;
            uint32_t block1 = (buf[28] >> 4) & 0x0F;
            uint32_t block0_value = (1 << block0);
            uint32_t block1_value = (1 << block1);
            vorbis = (ttLibC_Vorbis *)(*ttLibGo_Vorbis_make)(
                nullptr,
                VorbisType_identification,
                sample_rate,
                0,
                channel_num,
                data,
                data_size,
                false,
                pts,
                timebase);
            if(vorbis == NULL) {
              ERR_PRINT("failed to make vorbisFrame.");
              return NULL;
            }
            vorbis->block0 = block0_value;
            vorbis->block1 = block1_value;
            vorbis->block_type = 0;
          }
          break;
        case 0x03: // comment
          {
            vorbis = (ttLibC_Vorbis *)(*ttLibGo_Vorbis_make)(
                nullptr,
                VorbisType_comment,
                sample_rate,
                0,
                channel_num,
                data,
                data_size,
                false,
                pts,
                timebase);
            if(vorbis == NULL) {
              ERR_PRINT("failed to make vorbisFrame.");
              return NULL;
            }
          }
          break;
        case 0x05: // setup
          {
            vorbis = (ttLibC_Vorbis *)(*ttLibGo_Vorbis_make)(
                nullptr,
                VorbisType_setup,
                sample_rate,
                0,
                channel_num,
                data,
                data_size,
                false,
                pts,
                timebase);
            if(vorbis == NULL) {
              ERR_PRINT("failed to make vorbisFrame.");
              return NULL;
            }
          }
          break;
        default:
          return nullptr;
      }
    }
  }
  if(vorbis == nullptr) {
    // 強制フレーム扱い
    vorbis = (ttLibC_Vorbis *)(*ttLibGo_Vorbis_make)(
        nullptr,
        VorbisType_frame,
        sample_rate,
        sample_num,
        channel_num,
        data,
        data_size,
        false,
        pts,
        timebase);
  }
  if(vorbis != nullptr) {
    vorbis->inherit_super.inherit_super.id = id;
  }
  return vorbis;
}

}
