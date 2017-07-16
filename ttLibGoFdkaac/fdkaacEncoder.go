package ttLibGoFdkaac

/*
#include <fdk-aac/aacenc_lib.h>
#include <ttLibC/allocator.h>
#include <ttLibC/frame/audio/pcms16.h>
#include <ttLibC/frame/audio/aac.h>
#include "../ttLibGo/frame.h"

extern bool ttLibGoFdkaacFrameCallback(void *ptr, ttLibC_Frame *frame);

typedef struct fdkaacEncoder_t {
	HANDLE_AACENCODER handle;
	ttLibC_Aac *aac;
	uint8_t    *pcm_buffer;
	size_t      pcm_buffer_size;
	size_t      pcm_buffer_next_pos;
	uint8_t    *data;
	size_t      data_size;
	uint64_t    pts;
	uint32_t    sample_rate;
	uint32_t    sample_num;
	bool        is_pts_initialized;
} fdkaacEncoder_t;

void FdkaacEncoder_close(fdkaacEncoder_t **encoder);

fdkaacEncoder_t *FdkaacEncoder_make(
		const char *ctype,
		uint32_t sample_rate,
		uint32_t channel_num,
		uint32_t bitrate) {
	fdkaacEncoder_t *encoder = ttLibC_malloc(sizeof(fdkaacEncoder_t));
	if(encoder == NULL) {
		return NULL;
	}
	AACENC_ERROR err;
	memset(encoder, 0, sizeof(fdkaacEncoder_t));
	if((err = aacEncOpen(&encoder->handle, 0, channel_num)) != AACENC_OK) {
		puts("aacが開けませんでした");
		ttLibC_free(encoder);
		return NULL;
	}
#define ENUM(str)	else if(!strcmp(ctype, str))
	AUDIO_OBJECT_TYPE type;
	if(0) {}
	ENUM("AOT_NONE")
		type = AOT_NONE;
	ENUM("AOT_NULL_OBJECT")
		type = AOT_NULL_OBJECT;
	ENUM("AOT_AAC_MAIN")
		type = AOT_AAC_MAIN;
	ENUM("AOT_AAC_LC")
		type = AOT_AAC_LC;
	ENUM("AOT_AAC_SSR")
		type = AOT_AAC_SSR;
	ENUM("AOT_AAC_LTP")
		type = AOT_AAC_LTP;
	ENUM("AOT_SBR")
		type = AOT_SBR;
	ENUM("AOT_AAC_SCAL")
		type = AOT_AAC_SCAL;
	ENUM("AOT_TWIN_VQ")
		type = AOT_TWIN_VQ;
	ENUM("AOT_CELP")
		type = AOT_CELP;
	ENUM("AOT_HVXC")
		type = AOT_HVXC;
	ENUM("AOT_RSVD_10")
		type = AOT_RSVD_10;
	ENUM("AOT_RSVD_11")
		type = AOT_RSVD_11;
	ENUM("AOT_TTSI")
		type = AOT_TTSI;
	ENUM("AOT_MAIN_SYNTH")
		type = AOT_MAIN_SYNTH;
	ENUM("AOT_WAV_TAB_SYNTH")
		type = AOT_WAV_TAB_SYNTH;
	ENUM("AOT_GEN_MIDI")
		type = AOT_GEN_MIDI;
	ENUM("AOT_ALG_SYNTH_AUD_FX")
		type = AOT_ALG_SYNTH_AUD_FX;
	ENUM("AOT_ER_AAC_LC")
		type = AOT_ER_AAC_LC;
	ENUM("AOT_RSVD_18")
		type = AOT_RSVD_18;
	ENUM("AOT_ER_AAC_LTP")
		type = AOT_ER_AAC_LTP;
	ENUM("AOT_ER_AAC_SCAL")
		type = AOT_ER_AAC_SCAL;
	ENUM("AOT_ER_TWIN_VQ")
		type = AOT_ER_TWIN_VQ;
	ENUM("AOT_ER_BSAC")
		type = AOT_ER_BSAC;
	ENUM("AOT_ER_AAC_LD")
		type = AOT_ER_AAC_LD;
	ENUM("AOT_ER_CELP")
		type = AOT_ER_CELP;
	ENUM("AOT_ER_HVXC")
		type = AOT_ER_HVXC;
	ENUM("AOT_ER_HILN")
		type = AOT_ER_HILN;
	ENUM("AOT_ER_PARA")
		type = AOT_ER_PARA;
	ENUM("AOT_RSVD_28")
		type = AOT_RSVD_28;
	ENUM("AOT_PS")
		type = AOT_PS;
	ENUM("AOT_MPEGS")
		type = AOT_MPEGS;
	ENUM("AOT_ESCAPE")
		type = AOT_ESCAPE;
	ENUM("AOT_MP3ONMP4_L1")
		type = AOT_MP3ONMP4_L1;
	ENUM("AOT_MP3ONMP4_L2")
		type = AOT_MP3ONMP4_L2;
	ENUM("AOT_MP3ONMP4_L3")
		type = AOT_MP3ONMP4_L3;
	ENUM("AOT_RSVD_35")
		type = AOT_RSVD_35;
	ENUM("AOT_RSVD_36")
		type = AOT_RSVD_36;
	ENUM("AOT_AAC_SLS")
		type = AOT_AAC_SLS;
	ENUM("AOT_SLS")
		type = AOT_SLS;
	ENUM("AOT_ER_AAC_ELD")
		type = AOT_ER_AAC_ELD;
	ENUM("AOT_USAC")
		type = AOT_USAC;
	ENUM("AOT_SAOC")
		type = AOT_SAOC;
	ENUM("AOT_LD_MPEGS")
		type = AOT_LD_MPEGS;
	ENUM("AOT_DRM_AAC")
		type = AOT_DRM_AAC;
	ENUM("AOT_DRM_SBR")
		type = AOT_DRM_SBR;
	ENUM("AOT_DRM_MPEG_PS")
		type = AOT_DRM_MPEG_PS;
	else {
		err = AACENC_INIT_ERROR;
	}
#undef ENUM
	if(err == AACENC_OK && (err = aacEncoder_SetParam(encoder->handle, AACENC_AOT, type)) != AACENC_OK) {
		puts("fdkaac profile設定失敗");
	}
	if(err == AACENC_OK && (err = aacEncoder_SetParam(encoder->handle, AACENC_SAMPLERATE, sample_rate)) != AACENC_OK) {
		puts("fdkaac sampleRate設定失敗");
	}
	CHANNEL_MODE mode;
	switch(channel_num) {
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
	if(err == AACENC_OK && (err = aacEncoder_SetParam(encoder->handle, AACENC_CHANNELMODE, mode)) != AACENC_OK) {
		puts("fdkaac channelNum設定失敗");
	}
	if(err == AACENC_OK && (err = aacEncoder_SetParam(encoder->handle, AACENC_CHANNELORDER, 1)) != AACENC_OK) {
		puts("fdkaac channelOrder設定失敗");
	}
	if(err == AACENC_OK && (err = aacEncoder_SetParam(encoder->handle, AACENC_BITRATE, bitrate * 1024)) != AACENC_OK) {
		puts("fdkaac bitrate設定失敗");
	}
	if(err == AACENC_OK && (err = aacEncoder_SetParam(encoder->handle, AACENC_TRANSMUX, 2)) != AACENC_OK) {
		puts("fdkaac transmux設定失敗");
	}
	if(err == AACENC_OK && (err = aacEncoder_SetParam(encoder->handle, AACENC_AFTERBURNER, 1)) != AACENC_OK) {
		puts("fdkaac afterburnar設定失敗");
	}
	// 初期化実行
	if(err == AACENC_OK && (err = aacEncEncode(encoder->handle, NULL, NULL, NULL, NULL)) != AACENC_OK) {
		puts("fdkaac 初期化失敗");
	}
	AACENC_InfoStruct info;
	if(err == AACENC_OK && (err = aacEncInfo(encoder->handle, &info)) != AACENC_OK) {
		puts("fdkaac エンコードインフォ参照失敗");
	}
	if(err != AACENC_OK) {
		if(encoder->handle != NULL) {
			aacEncClose(&encoder->handle);
			encoder->handle = NULL;
		}
		ttLibC_free(encoder);
		return NULL;
	}
	encoder->aac                 = NULL;
	encoder->data_size           = 20480;
	encoder->data                = ttLibC_malloc(encoder->data_size);
	encoder->pcm_buffer_next_pos = 0;
	encoder->pcm_buffer_size     = info.frameLength * channel_num * sizeof(int16_t);
	encoder->pcm_buffer          = ttLibC_malloc(encoder->pcm_buffer_size);
	encoder->pts                 = 0;
	encoder->sample_rate         = sample_rate;
	encoder->sample_num          = info.frameLength;
	encoder->is_pts_initialized  = false;
	if(encoder->data == NULL || encoder->pcm_buffer == NULL) {
		FdkaacEncoder_close(&encoder);
		return NULL;
	}
	return encoder;
}

bool FdkaacEncoder_encode(fdkaacEncoder_t *encoder, void *frame, uintptr_t ptr) {
	if(encoder == NULL) {
		return false;
	}
	if(frame == NULL) {
		return true;
	}
	ttLibC_Frame *f = (ttLibC_Frame *)frame;
	if(f->type != frameType_pcmS16) {
		puts("pcmS16のみ処理可能です");
		return false;
	}
	if(!encoder->is_pts_initialized) {
		encoder->is_pts_initialized = true;
		encoder->pts = f->pts * encoder->sample_rate / f->timebase;
	}
	uint8_t *data = (uint8_t *)f->data;
	size_t left_size = f->buffer_size;
	AACENC_ERROR err;
	AACENC_BufDesc in_buf   = {0}, out_buf = {0};
	AACENC_InArgs  in_args  = {0};
	AACENC_OutArgs out_args = {0};

	int in_identifier = IN_AUDIO_DATA;
	int in_size = encoder->pcm_buffer_size;
	int in_elem_size = sizeof(int16_t);
	void *in_ptr = NULL;
	void *out_ptr = encoder->data;
	int out_identifier = OUT_BITSTREAM_DATA;
	int out_size = encoder->data_size;
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

	if(encoder->pcm_buffer_next_pos != 0) {
		if(left_size < in_size - encoder->pcm_buffer_next_pos) {
			// append data is not enough for encode. just add.
			memcpy(encoder->pcm_buffer + encoder->pcm_buffer_next_pos, data, left_size);
			encoder->pcm_buffer_next_pos += left_size;
			return true;
		}
		memcpy(encoder->pcm_buffer + encoder->pcm_buffer_next_pos, data, in_size - encoder->pcm_buffer_next_pos);
		data += in_size - encoder->pcm_buffer_next_pos;
		left_size -= in_size - encoder->pcm_buffer_next_pos;
		in_ptr = encoder->pcm_buffer;
		in_buf.bufs = &in_ptr;
		out_ptr = encoder->data;
		out_buf.bufs = &out_ptr;
		// ここでエンコードを実施すればよい。
		if((err = aacEncEncode(encoder->handle, &in_buf, &out_buf, &in_args, &out_args)) != AACENC_OK) {
			puts("failed to encode data.");
			return false;;
		}
		if(out_args.numOutBytes > 0) {
			ttLibC_Aac *aac = ttLibC_Aac_getFrame(
				encoder->aac,
				encoder->data,
				out_args.numOutBytes,
				true,
				encoder->pts,
				encoder->sample_rate);
			if(aac == NULL) {
				return false;
			}
			encoder->aac = aac;
			encoder->pts += 1024;
			if(!ttLibGoFdkaacFrameCallback((void *)ptr, (ttLibC_Frame *)aac)) {
				return false;
			}
		}
		encoder->pcm_buffer_next_pos = 0;
	}
	do {
		// get the data from pcm buffer.
		if(left_size < in_size) {
			// need more data.
			if(left_size != 0) {
				memcpy(encoder->pcm_buffer, data, left_size);
				encoder->pcm_buffer_next_pos = left_size;
			}
			return true;
		}
		// こっちもここでencodeを実施すればよい。
		// data have enough size.
		in_ptr = data;
		in_buf.bufs = &in_ptr;
		out_ptr = encoder->data;
		out_buf.bufs = &out_ptr;
		// ここでエンコードを実施すればよい。
		if((err = aacEncEncode(encoder->handle, &in_buf, &out_buf, &in_args, &out_args)) != AACENC_OK) {
			puts("failed to encode data.");
			return false;
		}
		if(out_args.numOutBytes > 0) {
			ttLibC_Aac *aac = ttLibC_Aac_getFrame(
				encoder->aac,
				encoder->data,
				out_args.numOutBytes,
				true,
				encoder->pts,
				encoder->sample_rate);
			if(aac == NULL) {
				return false;
			}
			encoder->aac = aac;
			encoder->pts += 1024;
			if(!ttLibGoFdkaacFrameCallback((void *)ptr, (ttLibC_Frame *)aac)) {
				return false;
			}
		}
		data += in_size;
		left_size -= in_size;
	}while(true);
}

void FdkaacEncoder_close(fdkaacEncoder_t **encoder) {
	fdkaacEncoder_t *target = *encoder;
	if(target == NULL) {
		return;
	}
	if(target->handle != NULL) {
		aacEncClose(&target->handle);
		target->handle = NULL;
	}
	if(target->pcm_buffer) {
		ttLibC_free(target->pcm_buffer);
		target->pcm_buffer = NULL;
	}
	if(target->data) {
		ttLibC_free(target->data);
		target->data = NULL;
	}
	ttLibC_Aac_close(&target->aac);
	ttLibC_free(target);
	*encoder = NULL;
}

bool FdkaacEncoder_setBitrate(
		fdkaacEncoder_t *encoder,
		uint32_t value) {
	if(encoder == NULL) {
		return false;
	}
	return aacEncoder_SetParam(encoder->handle, AACENC_BITRATE, value) == AACENC_OK;
}
*/
import "C"
import (
	"unsafe"

	"github.com/taktod/ttLibGo/ttLibGo"
)

// FdkaacEncoder fdkaacを利用したencoder
type FdkaacEncoder struct {
	call     frameCall
	cEncoder *C.fdkaacEncoder_t
}

// Init 初期化
func (fdkaac *FdkaacEncoder) Init(
	aacType string,
	sampleRate uint32,
	channelNum uint32,
	bitrate uint32) bool {
	if fdkaac.cEncoder == nil {
		cType := C.CString(aacType)
		defer C.free(unsafe.Pointer(cType))
		fdkaac.cEncoder = C.FdkaacEncoder_make(
			cType,
			C.uint32_t(sampleRate),
			C.uint32_t(channelNum),
			C.uint32_t(bitrate))
	}
	return fdkaac.cEncoder != nil
}

// Encode pcmS16のフレームをエンコードします
func (fdkaac *FdkaacEncoder) Encode(
	frame *ttLibGo.Frame,
	callback ttLibGo.FrameCallback) bool {
	if fdkaac.cEncoder == nil {
		return false
	}
	if frame == nil {
		return true
	}
	frame.Restore()
	fdkaac.call.callback = callback
	if !C.FdkaacEncoder_encode(
		fdkaac.cEncoder,
		unsafe.Pointer(frame.CFrame),
		C.uintptr_t(uintptr(unsafe.Pointer(fdkaac)))) {
		return false
	}
	return true
}

// Close 解放します
func (fdkaac *FdkaacEncoder) Close() {
	C.FdkaacEncoder_close(&fdkaac.cEncoder)
}

// SetBitrate bitrateを任意のタイミングで変更します
func (fdkaac *FdkaacEncoder) SetBitrate(value uint32) bool {
	if fdkaac.cEncoder == nil {
		return false
	}
	return bool(C.FdkaacEncoder_setBitrate(fdkaac.cEncoder, C.uint32_t(value)))
}
