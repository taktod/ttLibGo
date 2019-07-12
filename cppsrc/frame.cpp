#include <stdio.h>
#include <ttLibC/frame/frame.h>
#include <ttLibC/frame/video/bgr.h>
#include <ttLibC/frame/video/flv1.h>
#include <ttLibC/frame/video/h264.h>
#include <ttLibC/frame/video/h265.h>
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
    {
      return ttLibC_Bgr_getMinimumBinaryBuffer(
        (ttLibC_Bgr *)f,
        [](void *ptr, void *data, size_t data_size) -> bool {
          return ttLibGoDataCallback(ptr, data, data_size);
        }, (void *)ptr);
    }
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

void *BgrFrame_fromBinary(void *data, size_t data_size,
  uint32_t id, uint64_t pts, uint32_t timebase,
  const char *bgr_type, uint32_t width, uint32_t height, uint32_t width_stride) {
  // まず、emptyFrameを作って、それからデータをコピーする形にするべきかな・・・
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
  ttLibC_Bgr *bgr = ttLibC_Bgr_makeEmptyFrame(bgrType, width, height);
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
  ttLibC_Flv1 *flv1 = ttLibC_Flv1_getFrame(
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
  ttLibC_H264 *h264 = ttLibC_H264_getFrame(
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
  ttLibC_H265 *h265 = ttLibC_H265_getFrame(
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
  ttLibC_Jpeg *jpeg = ttLibC_Jpeg_getFrame(
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
  ttLibC_Png *png = ttLibC_Png_getFrame(
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
	    width = ttLibC_Theora_getWidth(nullptr, u8data, data_size);
    }
    if(height == 0) {
	    height = ttLibC_Theora_getHeight(nullptr, u8data, data_size);
    }
		if(width == 0 || height == 0) {
			return NULL;
		}
	}
	ttLibC_Theora *theora = ttLibC_Theora_make(
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
	if(data_size <= 1) {
		return NULL;
	}
	bool is_key = ttLibC_Vp6_isKey(u8data, data_size);
  if(width == 0) {
	  width = ttLibC_Vp6_getWidth(nullptr, u8data, data_size, adjustment);
  }
  if(height == 0) {
	  height = ttLibC_Vp6_getHeight(nullptr, u8data, data_size, adjustment);
  }
	if(width == 0 || height == 0) {
		return NULL;
	}
	ttLibC_Vp6 *vp6 = ttLibC_Vp6_make(
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
  uint8_t *u8data = (uint8_t *)data;
	if(data_size <= 1) {
		return NULL;
	}
	bool is_key = ttLibC_Vp8_isKey(u8data, data_size);
  if(width == 0) {
	  width = ttLibC_Vp8_getWidth(nullptr, u8data, data_size);
  }
  if(height == 0) {
	  height = ttLibC_Vp8_getHeight(nullptr, u8data, data_size);
  }
	if(width == 0 || height == 0) {
		return NULL;
	}
	ttLibC_Vp8 *vp8 = ttLibC_Vp8_make(
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
  uint8_t *u8data = (uint8_t *)data;
	if(data_size <= 1) {
		return NULL;
	}
	bool is_key = ttLibC_Vp9_isKey(u8data, data_size);
  if(width == 0) {
	  width = ttLibC_Vp9_getWidth(nullptr, u8data, data_size);
  }
  if(height == 0) {
	  height = ttLibC_Vp9_getHeight(nullptr, u8data, data_size);
  }
	if(width == 0 || height == 0) {
		return NULL;
	}
	ttLibC_Vp9 *vp9 = ttLibC_Vp9_make(
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
  ttLibC_Yuv420 *yuv = ttLibC_Yuv420_makeEmptyFrame(yuvType, width, height);
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
    ttLibC_Aac *aac = ttLibC_Aac_make(
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
  ttLibC_ByteReader *reader = ttLibC_ByteReader_make(data, data_size, ByteUtilType_default);
  if(ttLibC_ByteReader_bit(reader, 12) != 0xFFF) {
    ttLibC_ByteReader_close(&reader);
    // dsi情報を取り出し
		ttLibC_ByteReader *reader = ttLibC_ByteReader_make(data, data_size, ByteUtilType_default);
		uint64_t dsi_info;
		uint32_t bit_size = 0;
		memcpy(&dsi_info, data, data_size);
		uint32_t object_type = ttLibC_ByteReader_bit(reader, 5);
		bit_size += 5;
		if(object_type == 31) {
			object_type = ttLibC_ByteReader_bit(reader, 6);
			bit_size += 6;
		}
		uint32_t sample_rate_index = ttLibC_ByteReader_bit(reader, 4);
		bit_size += 4;
		uint32_t sample_rate = 44100;
		if(sample_rate_index == 15) {
			LOG_PRINT("sample_rate is not in index_table.");
			sample_rate = ttLibC_ByteReader_bit(reader, 24);
			bit_size += 24;
		}
		else {
			sample_rate = aac_sample_rate_table[sample_rate_index];
		}
		uint32_t channel_num = ttLibC_ByteReader_bit(reader, 4);
		bit_size += 4;
		uint32_t buffer_size = (uint32_t)((bit_size + 7) / 8);
		ttLibC_ByteReader_close(&reader);
		ttLibC_Aac *aac = ttLibC_Aac_make(
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
	ttLibC_ByteReader_bit(reader, 1);
	ttLibC_ByteReader_bit(reader, 2);
	ttLibC_ByteReader_bit(reader, 1);
	ttLibC_ByteReader_bit(reader, 2);
	uint32_t sample_rate_index = ttLibC_ByteReader_bit(reader, 4);
	uint32_t sample_rate = aac_sample_rate_table[sample_rate_index];
	ttLibC_ByteReader_bit(reader, 1);
	uint32_t channel_num = ttLibC_ByteReader_bit(reader, 3);
	ttLibC_ByteReader_bit(reader, 1);
	ttLibC_ByteReader_bit(reader, 1);
	ttLibC_ByteReader_bit(reader, 1);
	ttLibC_ByteReader_bit(reader, 1);
	uint32_t frame_size = ttLibC_ByteReader_bit(reader, 13);
	ttLibC_ByteReader_bit(reader, 11);
	ttLibC_ByteReader_bit(reader, 2);
	if(reader->error != Error_noError) {
		LOG_ERROR(reader->error);
		ttLibC_ByteReader_close(&reader);
    return nullptr;
	}
	ttLibC_ByteReader_close(&reader);
  // fffで始まってたらadtsとして処理
  ttLibC_Aac *aac = ttLibC_Aac_make(
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
  ttLibC_AdpcmImaWav *adpcm = ttLibC_AdpcmImaWav_make(
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
  ttLibC_Mp3 *mp3 = ttLibC_Mp3_getFrame(
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
  ttLibC_Nellymoser *nellymoser = ttLibC_Nellymoser_make(
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
  ttLibC_Opus *opus = ttLibC_Opus_getFrame(
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
  ttLibC_PcmAlaw *pcmAlaw = ttLibC_PcmAlaw_make(
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
  ttLibC_PcmF32 *pcmF32 = ttLibC_PcmF32_make(
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
  ttLibC_PcmMulaw *pcmMulaw = ttLibC_PcmMulaw_make(
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
  ttLibC_PcmS16 *pcmS16 = ttLibC_PcmS16_make(
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
    speex = ttLibC_Speex_make(
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
    speex = ttLibC_Speex_make(
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
    speex = ttLibC_Speex_make(
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
            vorbis = ttLibC_Vorbis_make(
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
            vorbis = ttLibC_Vorbis_make(
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
            vorbis = ttLibC_Vorbis_make(
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
    vorbis = ttLibC_Vorbis_make(
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
