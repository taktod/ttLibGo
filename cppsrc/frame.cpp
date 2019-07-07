#include <stdio.h>
#include <ttLibC/frame/frame.h>
#include <ttLibC/frame/video/bgr.h>
#include <ttLibC/frame/video/flv1.h>
#include <ttLibC/frame/video/h264.h>
#include <ttLibC/frame/video/h265.h>
#include <ttLibC/frame/video/theora.h>
#include <ttLibC/frame/video/yuv420.h>

#include <ttLibC/frame/audio/audio.h>
#include <ttLibC/frame/audio/aac.h>
#include <ttLibC/frame/audio/mp3.h>
#include <ttLibC/frame/audio/opus.h>
#include <ttLibC/frame/audio/pcmf32.h>
#include <ttLibC/frame/audio/pcms16.h>
#include <ttLibC/frame/audio/speex.h>
#include <ttLibC/frame/audio/vorbis.h>

#include "frame.hpp"
#include <string>

using namespace std;

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
  _pts = cFrame->pts;
  _dts = cFrame->dts;
  _timebase = cFrame->timebase;
  _id = cFrame->id;
  cFrame->pts = goFrame->pts;
  cFrame->dts = goFrame->dts;
  cFrame->timebase = goFrame->timebase;
  cFrame->id = goFrame->id;
  if(ttLibC_Frame_isVideo(cFrame)) {
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
  if(ttLibC_Frame_isAudio(cFrame)) {
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
  cFrame->pts = _pts;
  cFrame->dts = _dts;
  cFrame->timebase = _timebase;
  cFrame->id = _id;
  if(ttLibC_Frame_isVideo(cFrame)) {
    ttLibC_Video *video = (ttLibC_Video *)cFrame;
    video->width = _width;
    video->height = _height;
  }
  if(ttLibC_Frame_isAudio(cFrame)) {
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
  if(frame == NULL) {
    return false;
  }
  ttLibC_Frame *f = (ttLibC_Frame *)frame;
	switch(f->type) {
	case frameType_bgr:
    break;
	case frameType_yuv420:
    {
      return ttLibC_Yuv420_getMinimumBinaryBuffer(
        (ttLibC_Yuv420 *)f,
        [](void *ptr, void *data, size_t data_size) -> bool {
          return ttLibGoDataCallback(ptr, data, data_size);
        }, (void *)ptr);
    }
		break;
	case frameType_pcmS16:
	case frameType_pcmF32:
		{
      ttLibC_Frame *cloned = ttLibC_Frame_clone(nullptr, f);
      if(cloned == nullptr) {
        return false;
      }
  		bool result = ttLibGoDataCallback((void *)ptr, cloned->data, cloned->buffer_size);
      ttLibC_Frame_close(&cloned);
      return result;
		}
		break;
	default:
		// とりあえずこれでいいはず。
		return ttLibGoDataCallback((void *)ptr, f->data, f->buffer_size);
	}
  return false;
}

}