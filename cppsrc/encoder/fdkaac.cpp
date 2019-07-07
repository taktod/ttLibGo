#include "fdkaac.hpp"
#include "../util.hpp"
#include <iostream>
#include <ttLibC/allocator.h>

using namespace std;

extern "C" {
extern bool ttLibGoFrameCallback(void *ptr, ttLibC_Frame *frame);
}
FdkaacEncoder::FdkaacEncoder(maps *mp) {
#ifdef __ENABLE_FDKAAC_ENCODE__
  string type = mp->getString("aacType");
  uint32_t sampleRate = mp->getUint32("sampleRate");
  uint32_t channelNum = mp->getUint32("channelNum");
  uint32_t bitrate = mp->getUint32("bitrate");
  handle_ = nullptr;
  data_ = nullptr;
  pcm_buffer_ = nullptr;
  aac_ = nullptr;
  AACENC_ERROR err;
	if((err = aacEncOpen(&handle_, 0, channelNum)) != AACENC_OK) {
		puts("aacが開けませんでした");
		return;
	}
#define ENUM(str)	else if(type == str)
	AUDIO_OBJECT_TYPE aactype;
	if(0) {}
	ENUM("AOT_NONE")
		aactype = AOT_NONE;
	ENUM("AOT_NULL_OBJECT")
		aactype = AOT_NULL_OBJECT;
	ENUM("AOT_AAC_MAIN")
		aactype = AOT_AAC_MAIN;
	ENUM("AOT_AAC_LC")
		aactype = AOT_AAC_LC;
	ENUM("AOT_AAC_SSR")
		aactype = AOT_AAC_SSR;
	ENUM("AOT_AAC_LTP")
		aactype = AOT_AAC_LTP;
	ENUM("AOT_SBR")
		aactype = AOT_SBR;
	ENUM("AOT_AAC_SCAL")
		aactype = AOT_AAC_SCAL;
	ENUM("AOT_TWIN_VQ")
		aactype = AOT_TWIN_VQ;
	ENUM("AOT_CELP")
		aactype = AOT_CELP;
	ENUM("AOT_HVXC")
		aactype = AOT_HVXC;
	ENUM("AOT_RSVD_10")
		aactype = AOT_RSVD_10;
	ENUM("AOT_RSVD_11")
		aactype = AOT_RSVD_11;
	ENUM("AOT_TTSI")
		aactype = AOT_TTSI;
	ENUM("AOT_MAIN_SYNTH")
		aactype = AOT_MAIN_SYNTH;
	ENUM("AOT_WAV_TAB_SYNTH")
		aactype = AOT_WAV_TAB_SYNTH;
	ENUM("AOT_GEN_MIDI")
		aactype = AOT_GEN_MIDI;
	ENUM("AOT_ALG_SYNTH_AUD_FX")
		aactype = AOT_ALG_SYNTH_AUD_FX;
	ENUM("AOT_ER_AAC_LC")
		aactype = AOT_ER_AAC_LC;
	ENUM("AOT_RSVD_18")
		aactype = AOT_RSVD_18;
	ENUM("AOT_ER_AAC_LTP")
		aactype = AOT_ER_AAC_LTP;
	ENUM("AOT_ER_AAC_SCAL")
		aactype = AOT_ER_AAC_SCAL;
	ENUM("AOT_ER_TWIN_VQ")
		aactype = AOT_ER_TWIN_VQ;
	ENUM("AOT_ER_BSAC")
		aactype = AOT_ER_BSAC;
	ENUM("AOT_ER_AAC_LD")
		aactype = AOT_ER_AAC_LD;
	ENUM("AOT_ER_CELP")
		aactype = AOT_ER_CELP;
	ENUM("AOT_ER_HVXC")
		aactype = AOT_ER_HVXC;
	ENUM("AOT_ER_HILN")
		aactype = AOT_ER_HILN;
	ENUM("AOT_ER_PARA")
		aactype = AOT_ER_PARA;
	ENUM("AOT_RSVD_28")
		aactype = AOT_RSVD_28;
	ENUM("AOT_PS")
		aactype = AOT_PS;
	ENUM("AOT_MPEGS")
		aactype = AOT_MPEGS;
	ENUM("AOT_ESCAPE")
		aactype = AOT_ESCAPE;
	ENUM("AOT_MP3ONMP4_L1")
		aactype = AOT_MP3ONMP4_L1;
	ENUM("AOT_MP3ONMP4_L2")
		aactype = AOT_MP3ONMP4_L2;
	ENUM("AOT_MP3ONMP4_L3")
		aactype = AOT_MP3ONMP4_L3;
	ENUM("AOT_RSVD_35")
		aactype = AOT_RSVD_35;
	ENUM("AOT_RSVD_36")
		aactype = AOT_RSVD_36;
	ENUM("AOT_AAC_SLS")
		aactype = AOT_AAC_SLS;
	ENUM("AOT_SLS")
		aactype = AOT_SLS;
	ENUM("AOT_ER_AAC_ELD")
		aactype = AOT_ER_AAC_ELD;
	ENUM("AOT_USAC")
		aactype = AOT_USAC;
	ENUM("AOT_SAOC")
		aactype = AOT_SAOC;
	ENUM("AOT_LD_MPEGS")
		aactype = AOT_LD_MPEGS;
	ENUM("AOT_DRM_AAC")
		aactype = AOT_DRM_AAC;
	ENUM("AOT_DRM_SBR")
		aactype = AOT_DRM_SBR;
	ENUM("AOT_DRM_MPEG_PS")
		aactype = AOT_DRM_MPEG_PS;
	else {
		err = AACENC_INIT_ERROR;
	}
#undef ENUM
	if(err == AACENC_OK && (err = aacEncoder_SetParam(handle_, AACENC_AOT, aactype)) != AACENC_OK) {
		puts("fdkaac profile設定失敗");
	}
	if(err == AACENC_OK && (err = aacEncoder_SetParam(handle_, AACENC_SAMPLERATE, sampleRate)) != AACENC_OK) {
		puts("fdkaac sampleRate設定失敗");
	}
	CHANNEL_MODE mode;
	switch(channelNum) {
	case 1:
		mode = MODE_1;
		break;
	case 2:
		mode = MODE_2;
		break;
	default:
		err = AACENC_INIT_ERROR;
		break;
	}
	if(err == AACENC_OK && (err = aacEncoder_SetParam(handle_, AACENC_CHANNELMODE, mode)) != AACENC_OK) {
		puts("fdkaac channelNum設定失敗");
	}
	if(err == AACENC_OK && (err = aacEncoder_SetParam(handle_, AACENC_CHANNELORDER, 1)) != AACENC_OK) {
		puts("fdkaac channelOrder設定失敗");
	}
	if(err == AACENC_OK && (err = aacEncoder_SetParam(handle_, AACENC_BITRATE, bitrate)) != AACENC_OK) {
		puts("fdkaac bitrate設定失敗");
	}
	if(err == AACENC_OK && (err = aacEncoder_SetParam(handle_, AACENC_TRANSMUX, 2)) != AACENC_OK) {
		puts("fdkaac transmux設定失敗");
	}
	if(err == AACENC_OK && (err = aacEncoder_SetParam(handle_, AACENC_AFTERBURNER, 1)) != AACENC_OK) {
		puts("fdkaac afterburnar設定失敗");
	}
	// 初期化実行
	if(err == AACENC_OK && (err = aacEncEncode(handle_, NULL, NULL, NULL, NULL)) != AACENC_OK) {
		puts("fdkaac 初期化失敗");
	}
	AACENC_InfoStruct info;
	if(err == AACENC_OK && (err = aacEncInfo(handle_, &info)) != AACENC_OK) {
		puts("fdkaac エンコードインフォ参照失敗");
	}
	if(err != AACENC_OK) {
		if(handle_ != NULL) {
			aacEncClose(&handle_);
			handle_ = NULL;
		}
    return;
	}
  aac_ = NULL;
  data_size_ = 2048;
  data_ = (uint8_t *)ttLibC_malloc(data_size_);
  if(data_ == NULL) {
    aacEncClose(&handle_);
    handle_ = NULL;
    return;
  }
  pcm_buffer_next_pos_ = 0;
  pcm_buffer_size_ = info.frameLength * channelNum * sizeof(int16_t);
  pcm_buffer_ = (uint8_t *)ttLibC_malloc(pcm_buffer_size_);
  if(pcm_buffer_ == NULL) {
    aacEncClose(&handle_);
    handle_ = NULL;
    ttLibC_free(data_);
    data_ = NULL;
  }
  pts_ = 0;
  sample_rate_ = sampleRate;
  sample_num_ = info.frameLength;
  is_pts_initialized_ = false;
#endif
}

FdkaacEncoder::~FdkaacEncoder() {
#ifdef __ENABLE_FDKAAC_ENCODE__
  if(handle_ != NULL) {
    aacEncClose(&handle_);
    handle_ = NULL;
  }
  if(data_ != NULL) {
    ttLibC_free(data_);
    data_ = NULL;
  }
  if(pcm_buffer_ != NULL) {
    ttLibC_free(pcm_buffer_);
    pcm_buffer_ = NULL;
  }
  ttLibC_Aac_close(&aac_);
#endif
}
bool FdkaacEncoder::encodeFrame(ttLibC_Frame *cFrame, ttLibGoFrame *goFrame, void *ptr) {
#ifdef __ENABLE_FDKAAC_ENCODE__
  // 中の準備ができてなかったら動作させない。
  if(handle_ == NULL) {
    puts("初期化されていません。");
    return false;
  }
  if(cFrame == NULL) {
    return true;
  }
  if(cFrame->type != frameType_pcmS16) {
    puts("pcmS16のみ処理可能です。");
    return false;
  }
  update(cFrame, goFrame);
  // あとはaacをつくって、うまいことcallbackしなければならない。
  if(!is_pts_initialized_) {
    is_pts_initialized_ = true;
    pts_ = cFrame->pts * sample_rate_ / cFrame->timebase;
  }
  uint8_t *data = (uint8_t *)cFrame->data;
  size_t left_size = cFrame->buffer_size;
	AACENC_ERROR err;
	AACENC_BufDesc in_buf   = {}, out_buf = {};
	AACENC_InArgs  in_args  = {};
	AACENC_OutArgs out_args = {};
	int in_identifier = IN_AUDIO_DATA;
	int in_size = pcm_buffer_size_;
	int in_elem_size = sizeof(int16_t);
	void *in_ptr = NULL;
	void *out_ptr = data_;
	int out_identifier = OUT_BITSTREAM_DATA;
	int out_size = data_size_;
	int out_elem_size = 1;
	in_args.numInSamples = in_size >> 1;

	in_buf.numBufs = 1;
	in_buf.bufferIdentifiers = &in_identifier;
//	in_buf.bufs = &in_ptr;
	in_buf.bufSizes = &in_size;
	in_buf.bufElSizes = &in_elem_size;

	out_buf.numBufs = 1;
	out_buf.bufferIdentifiers = &out_identifier;
//	out_buf.bufs = &out_ptr;
	out_buf.bufSizes = &out_size;
	out_buf.bufElSizes = &out_elem_size;
	if(pcm_buffer_next_pos_ != 0) {
		if(left_size < in_size - pcm_buffer_next_pos_) {
			// append data is not enough for encode. just add.
			memcpy(pcm_buffer_ + pcm_buffer_next_pos_, data, left_size);
			pcm_buffer_next_pos_ += left_size;
  reset(cFrame, goFrame);
			return true;
		}
		memcpy(pcm_buffer_ + pcm_buffer_next_pos_, data, in_size - pcm_buffer_next_pos_);
		data += in_size - pcm_buffer_next_pos_;
		left_size -= in_size - pcm_buffer_next_pos_;
		in_ptr = pcm_buffer_;
		in_buf.bufs = &in_ptr;
		out_ptr = data_;
		out_buf.bufs = &out_ptr;
		// ここでエンコードを実施すればよい。
		if((err = aacEncEncode(handle_, &in_buf, &out_buf, &in_args, &out_args)) != AACENC_OK) {
			puts("failed to encode data.");
  reset(cFrame, goFrame);
			return false;
		}
		if(out_args.numOutBytes > 0) {
			ttLibC_Aac *aac = ttLibC_Aac_getFrame(
				aac_,
				data_,
				out_args.numOutBytes,
				true,
				pts_,
				sample_rate_);
			if(aac == NULL) {
  reset(cFrame, goFrame);
				return false;
			}
			aac_ = aac;
			pts_ += 1024;
      if(!ttLibGoFrameCallback(ptr, (ttLibC_Frame *)aac_)) {
  reset(cFrame, goFrame);
        return false;
      }
		}
		pcm_buffer_next_pos_ = 0;
	}
	do {
		// get the data from pcm buffer.
		if((int)left_size < in_size) {
			// need more data.
			if(left_size != 0) {
				memcpy(pcm_buffer_, data, left_size);
				pcm_buffer_next_pos_ = left_size;
			}
  reset(cFrame, goFrame);
			return true;
		}
		// こっちもここでencodeを実施すればよい。
		// data have enough size.
		in_ptr = data;
		in_buf.bufs = &in_ptr;
		out_ptr = data_;
		out_buf.bufs = &out_ptr;
		// ここでエンコードを実施すればよい。
		if((err = aacEncEncode(handle_, &in_buf, &out_buf, &in_args, &out_args)) != AACENC_OK) {
			puts("failed to encode data.");
  reset(cFrame, goFrame);
			return false;
		}
		if(out_args.numOutBytes > 0) {
			ttLibC_Aac *aac = ttLibC_Aac_getFrame(
				aac_,
				data_,
				out_args.numOutBytes,
				true,
				pts_,
				sample_rate_);
			if(aac == NULL) {
  reset(cFrame, goFrame);
				return false;
			}
			aac_ = aac;
			pts_ += 1024;
      if(!ttLibGoFrameCallback(ptr, (ttLibC_Frame *)aac_)) {
  reset(cFrame, goFrame);
        return false;
      }
		}
		data += in_size;
		left_size -= in_size;
	}while(true);
  reset(cFrame, goFrame);
  return true;
#else
  return false;
#endif
}

bool FdkaacEncoder::setBitrate(uint32_t value) {
  if(handle_ == NULL) {
    cout << "handleが準備されていません。" << endl;
    return false;
  }
	return aacEncoder_SetParam(handle_, AACENC_BITRATE, value);
}
extern "C" {
bool FdkaacEncoder_setBitrate(void *encoder, uint32_t value) {
  FdkaacEncoder *fdk = reinterpret_cast<FdkaacEncoder *>(encoder);
	return fdk->setBitrate(value);
}
}
