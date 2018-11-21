package ttLibGo

/*
#include "frame.h"
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
#include <ttLibC/frame/video/video.h>
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
#include <ttLibC/frame/video/wmv1.h>
#include <ttLibC/frame/video/wmv2.h>
#include <ttLibC/frame/video/yuv420.h>

// frameのデータは基本void *でやりとりすることにする。
uint64_t Frame_getPts(void *frame) {
	if(frame == NULL) {
		return 0;
	}
	ttLibC_Frame *f = (ttLibC_Frame *)frame;
	return f->pts;
}

uint32_t Frame_getTimebase(void *frame) {
	if(frame == NULL) {
		return 0;
	}
	ttLibC_Frame *f = (ttLibC_Frame *)frame;
	return f->timebase;
}

uint32_t Frame_getID(void *frame) {
	if(frame == NULL) {
		return 0;
	}
	ttLibC_Frame *f = (ttLibC_Frame *)frame;
	return f->id;
}

void Frame_getCodecType(void *frame, char *buffer, size_t buffer_size) {
	if(frame == NULL) {
		return;
	}
	ttLibC_Frame *f = (ttLibC_Frame *)frame;
	switch(f->type) {
	case frameType_bgr:
		strcpy(buffer, "bgr");
		break;
	case frameType_flv1:
		strcpy(buffer, "flv1");
		break;
	case frameType_h264:
		strcpy(buffer, "h264");
		break;
	case frameType_h265:
		strcpy(buffer, "h265");
		break;
	case frameType_jpeg:
		strcpy(buffer, "jpeg");
		break;
	case frameType_png:
		strcpy(buffer, "png");
		break;
	case frameType_theora:
		strcpy(buffer, "theora");
		break;
	case frameType_vp6:
		strcpy(buffer, "vp6");
		break;
	case frameType_vp8:
		strcpy(buffer, "vp8");
		break;
	case frameType_vp9:
		strcpy(buffer, "vp9");
		break;
	case frameType_wmv1:
		strcpy(buffer, "wmv1");
		break;
	case frameType_wmv2:
		strcpy(buffer, "wmv2");
		break;
	case frameType_yuv420:
		strcpy(buffer, "yuv");
		break;
	case frameType_aac:
		strcpy(buffer, "aac");
		break;
	case frameType_adpcm_ima_wav:
		strcpy(buffer, "adpcmImaWav");
		break;
	case frameType_mp3:
		strcpy(buffer, "mp3");
		break;
	case frameType_nellymoser:
		strcpy(buffer, "nellymoser");
		break;
	case frameType_opus:
		strcpy(buffer, "opus");
		break;
	case frameType_pcm_alaw:
		strcpy(buffer, "pcmAlaw");
		break;
	case frameType_pcmF32:
		strcpy(buffer, "pcmF32");
		break;
	case frameType_pcm_mulaw:
		strcpy(buffer, "pcmMulaw");
		break;
	case frameType_pcmS16:
		strcpy(buffer, "pcmS16");
		break;
	case frameType_speex:
		strcpy(buffer, "speex");
		break;
	case frameType_vorbis:
		strcpy(buffer, "vorbis");
		break;
 	default:
		break;
	}
}

void Frame_getSubType(void *frame, char *buffer, size_t buffer_size) {
	if(frame == NULL) {
		return;
	}
	ttLibC_Frame *f = (ttLibC_Frame *)frame;
	switch(f->type) {
	case frameType_aac:
		{
			ttLibC_Aac *aac = (ttLibC_Aac *)frame;
			switch(aac->type) {
			case AacType_raw:
				strcpy(buffer, "raw");
				break;
			case AacType_adts:
				strcpy(buffer, "adts");
				break;
			case AacType_dsi:
				strcpy(buffer, "dsi");
				break;
			default:
				break;
			}
		}
		break;
	case frameType_mp3:
		{
			ttLibC_Mp3 *mp3 = (ttLibC_Mp3 *)frame;
			switch(mp3->type) {
			case Mp3Type_tag:
				strcpy(buffer, "tag");
				break;
			case Mp3Type_id3:
				strcpy(buffer, "id3");
				break;
			case Mp3Type_frame:
				strcpy(buffer, "frame");
				break;
			default:
				break;
			}
		}
		break;
	case frameType_opus:
		{
			ttLibC_Opus *opus = (ttLibC_Opus *)frame;
			switch(opus->type) {
			case OpusType_header:
				strcpy(buffer, "header");
				break;
			case OpusType_comment:
				strcpy(buffer, "comment");
				break;
			case OpusType_frame:
				strcpy(buffer, "frame");
				break;
			default:
				break;
			}
		}
		break;
	case frameType_pcmF32:
		{
			ttLibC_PcmF32 *pcm = (ttLibC_PcmF32 *)frame;
			switch(pcm->type) {
			case PcmF32Type_interleave:
				strcpy(buffer, "interleave");
				break;
			case PcmF32Type_planar:
				strcpy(buffer, "planar");
				break;
			default:
				break;
			}
		}
		break;
	case frameType_pcmS16:
		{
			ttLibC_PcmS16 *pcm = (ttLibC_PcmS16 *)frame;
			switch(pcm->type) {
			case PcmS16Type_bigEndian:
				strcpy(buffer, "bigEndian");
				break;
			case PcmS16Type_bigEndian_planar:
				strcpy(buffer, "bigEndianPlanar");
				break;
			case PcmS16Type_littleEndian:
				strcpy(buffer, "littleEndian");
				break;
			case PcmS16Type_littleEndian_planar:
				strcpy(buffer, "littleEndianPlanar");
				break;
			default:
				break;
			}
		}
		break;
	case frameType_speex:
		{
			ttLibC_Speex *speex = (ttLibC_Speex *)frame;
			switch(speex->type) {
			case SpeexType_header:
				strcpy(buffer, "header");
				break;
			case SpeexType_comment:
				strcpy(buffer, "comment");
				break;
			case SpeexType_frame:
				strcpy(buffer, "frame");
				break;
			default:
				break;
			}
		}
		break;
	case frameType_vorbis:
		{
			ttLibC_Vorbis *vorbis = (ttLibC_Vorbis *)frame;
			switch(vorbis->type) {
			case VorbisType_identification:
				strcpy(buffer, "identification");
				break;
			case VorbisType_comment:
				strcpy(buffer, "comment");
				break;
			case VorbisType_setup:
				strcpy(buffer, "setup");
				break;
			case VorbisType_frame:
				strcpy(buffer, "frame");
				break;
			default:
				break;
			}
		}
		break;
	case frameType_bgr:
		{
			ttLibC_Bgr *bgr = (ttLibC_Bgr *)frame;
			switch(bgr->type) {
			case BgrType_abgr:
				strcpy(buffer, "abgr");
				break;
			case BgrType_bgr:
				strcpy(buffer, "bgr");
				break;
			case BgrType_bgra:
				strcpy(buffer, "bgra");
				break;
			case BgrType_argb:
				strcpy(buffer, "argb");
				break;
			case BgrType_rgb:
				strcpy(buffer, "rgb");
				break;
			case BgrType_rgba:
				strcpy(buffer, "rgba");
				break;
			default:
				break;
			}
		}
		break;
	case frameType_flv1:
		{
			ttLibC_Flv1 *flv1 = (ttLibC_Flv1 *)frame;
			switch(flv1->type) {
			case Flv1Type_intra:
				strcpy(buffer, "intra");
				break;
			case Flv1Type_inner:
				strcpy(buffer, "inner");
				break;
			case Flv1Type_disposableInner:
				strcpy(buffer, "disposableInner");
				break;
			default:
				break;
			}
		}
		break;
	case frameType_h264:
		{
			ttLibC_H264 *h264 = (ttLibC_H264 *)frame;
			switch(h264->type) {
			case H264Type_configData:
				strcpy(buffer, "configData");
				break;
			case H264Type_sliceIDR:
				strcpy(buffer, "sliceIDR");
				break;
			case H264Type_slice:
				strcpy(buffer, "slice");
				break;
			default:
				break;
			}
		}
		break;
	case frameType_h265:
		{
			ttLibC_H265 *h265 = (ttLibC_H265 *)frame;
			switch(h265->type) {
			case H265Type_configData:
				strcpy(buffer, "configData");
				break;
			case H265Type_sliceIDR:
				strcpy(buffer, "sliceIDR");
				break;
			case H265Type_slice:
				strcpy(buffer, "slice");
				break;
			default:
				break;
			}
		}
		break;
	case frameType_theora:
		{
			ttLibC_Theora *theora = (ttLibC_Theora *)frame;
			switch(theora->type) {
			case TheoraType_identificationHeaderDecodeFrame:
				strcpy(buffer, "identification");
				break;
			case TheoraType_commentHeaderFrame:
				strcpy(buffer, "comment");
				break;
			case TheoraType_setupHeaderFrame:
				strcpy(buffer, "setup");
				break;
			case TheoraType_intraFrame:
				strcpy(buffer, "intra");
				break;
			case TheoraType_innerFrame:
				strcpy(buffer, "inner");
				break;
			default:
				break;
			}
		}
		break;
	case frameType_yuv420:
		{
			ttLibC_Yuv420 *yuv = (ttLibC_Yuv420 *)frame;
			switch(yuv->type) {
			case Yuv420Type_planar:
				strcpy(buffer, "planar");
				break;
			case Yuv420Type_semiPlanar:
				strcpy(buffer, "semiPlanar");
				break;
			case Yvu420Type_planar:
				strcpy(buffer, "yvuPlanar");
				break;
			case Yvu420Type_semiPlanar:
				strcpy(buffer, "yvuSemiPlanar");
				break;
			default:
				break;
			}
		}
		break;
	default:
		break;
	}
}

uint32_t Frame_getWidth(void *frame) {
	if(frame == NULL) {
		return 0;
	}
	ttLibC_Frame *f = (ttLibC_Frame *)frame;
	if(ttLibC_Frame_isVideo(f)) {
		ttLibC_Video *video = (ttLibC_Video *)frame;
		return video->width;
	}
	return 0;
}

uint32_t Frame_getHeight(void *frame) {
	if(frame == NULL) {
		return 0;
	}
	ttLibC_Frame *f = (ttLibC_Frame *)frame;
	if(ttLibC_Frame_isVideo(f)) {
		ttLibC_Video *video = (ttLibC_Video *)frame;
		return video->height;
	}
	return 0;
}

void Frame_getVideoType(void *frame, char *buffer, size_t buffer_size) {
	if(frame == NULL) {
		return;
	}
	ttLibC_Frame *f = (ttLibC_Frame *)frame;
	if(ttLibC_Frame_isVideo(f)) {
		ttLibC_Video *video = (ttLibC_Video *)frame;
		switch(video->type) {
		case videoType_key:
			strcpy(buffer, "key");
			break;
		case videoType_inner:
			strcpy(buffer, "inner");
			break;
		case videoType_info:
			strcpy(buffer, "info");
			break;
		default:
			break;
		}
	}
}

void Frame_getH26xType(void *frame, char *buffer, size_t buffer_size) {
	if(frame == NULL) {
		return;
	}
	ttLibC_Frame *f = (ttLibC_Frame *)frame;
	switch(f->type) {
	case frameType_h264:
		{
			ttLibC_H264 *h264 = (ttLibC_H264 *)frame;
			switch(h264->frame_type) {
			case H264FrameType_P:
				strcpy(buffer, "p");
				break;
			case H264FrameType_B:
				strcpy(buffer, "b");
				break;
			case H264FrameType_I:
				strcpy(buffer, "i");
				break;
			case H264FrameType_SP:
				strcpy(buffer, "sp");
				break;
			case H264FrameType_SI:
				strcpy(buffer, "si");
				break;
			default:
				break;
			}
		}
		break;
	case frameType_h265:
		{
			ttLibC_H265 *h265 = (ttLibC_H265 *)frame;
			switch(h265->frame_type) {
			case H265FrameType_B:
				strcpy(buffer, "b");
				break;
			case H265FrameType_I:
				strcpy(buffer, "i");
				break;
			case H265FrameType_P:
				strcpy(buffer, "p");
				break;
			default:
				break;
			}
		}
		break;
	default:
		break;
	}
}

bool Frame_isDisposable(void *frame) {
	if(frame == NULL) {
		return false;
	}
	ttLibC_Frame *f = (ttLibC_Frame *)frame;
	switch(f->type) {
	case frameType_flv1:
		{
			ttLibC_Flv1 *flv1 = (ttLibC_Flv1 *)f;
			return flv1->type == Flv1Type_disposableInner;
		}
		break;
	case frameType_h264:
		{
			ttLibC_H264 *h264 = (ttLibC_H264 *)f;
			return h264->is_disposable;
		}
		break;
	case frameType_h265:
		{
			ttLibC_H265 *h265 = (ttLibC_H265 *)f;
			return h265->is_disposable;
		}
		break;
	default:
		break;
	}
	return false;
}

uint32_t Frame_getSampleRate(void *frame) {
	if(frame == NULL) {
		return 0;
	}
	ttLibC_Frame *f = (ttLibC_Frame *)frame;
	if(ttLibC_Frame_isAudio(f)) {
		ttLibC_Audio *audio = (ttLibC_Audio *)frame;
		return audio->sample_rate;
	}
	return 0;
}

uint32_t Frame_getSampleNum(void *frame) {
	if(frame == NULL) {
		return 0;
	}
	ttLibC_Frame *f = (ttLibC_Frame *)frame;
	if(ttLibC_Frame_isAudio(f)) {
		ttLibC_Audio *audio = (ttLibC_Audio *)frame;
		return audio->sample_num;
	}
	return 0;
}

uint32_t Frame_getChannelNum(void *frame) {
	if(frame == NULL) {
		return 0;
	}
	ttLibC_Frame *f = (ttLibC_Frame *)frame;
	if(ttLibC_Frame_isAudio(f)) {
		ttLibC_Audio *audio = (ttLibC_Audio *)frame;
		return audio->channel_num;
	}
	return 0;
}

uint32_t Frame_getStride(void *frame) {
	if(frame == NULL) {
		return 0;
	}
	ttLibC_Frame *f = (ttLibC_Frame *)frame;
	if(f->type != frameType_bgr) {
		return 0;
	}
	ttLibC_Bgr *bgr = (ttLibC_Bgr *)frame;
	return bgr->width_stride;
}

uint32_t Frame_getYStride(void *frame) {
	if(frame == NULL) {
		return 0;
	}
	ttLibC_Frame *f = (ttLibC_Frame *)frame;
	if(f->type != frameType_yuv420) {
		return 0;
	}
	ttLibC_Yuv420 *yuv = (ttLibC_Yuv420 *)frame;
	return yuv->y_stride;
}

uint32_t Frame_getUStride(void *frame) {
	if(frame == NULL) {
		return 0;
	}
	ttLibC_Frame *f = (ttLibC_Frame *)frame;
	if(f->type != frameType_yuv420) {
		return 0;
	}
	ttLibC_Yuv420 *yuv = (ttLibC_Yuv420 *)frame;
	return yuv->u_stride;
}

uint32_t Frame_getVStride(void *frame) {
	if(frame == NULL) {
		return 0;
	}
	ttLibC_Frame *f = (ttLibC_Frame *)frame;
	if(f->type != frameType_yuv420) {
		return 0;
	}
	ttLibC_Yuv420 *yuv = (ttLibC_Yuv420 *)frame;
	return yuv->v_stride;
}

uintptr_t Frame_getData1Ptr(void *frame) {
	if(frame == NULL) {
		return 0;
	}
	ttLibC_Frame *f = (ttLibC_Frame *)frame;
	switch(f->type) {
	case frameType_bgr:
		{
			ttLibC_Bgr *bgr = (ttLibC_Bgr *)frame;
			return (uintptr_t)bgr->data;
		}
		break;
	case frameType_yuv420:
		{
			ttLibC_Yuv420 *yuv = (ttLibC_Yuv420 *)frame;
			return (uintptr_t)yuv->y_data;
		}
		break;
	case frameType_pcmS16:
		{
			ttLibC_PcmS16 *pcm = (ttLibC_PcmS16 *)frame;
			return (uintptr_t)pcm->l_data;
		}
		break;
	case frameType_pcmF32:
		{
			ttLibC_PcmF32 *pcm = (ttLibC_PcmF32 *)frame;
			return (uintptr_t)pcm->l_data;
		}
		break;
	default:
		return 0;
	}
}

uintptr_t Frame_getData2Ptr(void *frame) {
	if(frame == NULL) {
		return 0;
	}
	ttLibC_Frame *f = (ttLibC_Frame *)frame;
	switch(f->type) {
	case frameType_yuv420:
		{
			ttLibC_Yuv420 *yuv = (ttLibC_Yuv420 *)frame;
			return (uintptr_t)yuv->u_data;
		}
		break;
	case frameType_pcmS16:
		{
			ttLibC_PcmS16 *pcm = (ttLibC_PcmS16 *)frame;
			return (uintptr_t)pcm->r_data;
		}
		break;
	case frameType_pcmF32:
		{
			ttLibC_PcmF32 *pcm = (ttLibC_PcmF32 *)frame;
			return (uintptr_t)pcm->r_data;
		}
		break;
	default:
		return 0;
	}
}

uintptr_t Frame_getData3Ptr(void *frame) {
	if(frame == NULL) {
		return 0;
	}
	ttLibC_Frame *f = (ttLibC_Frame *)frame;
	switch(f->type) {
	case frameType_yuv420:
		{
			ttLibC_Yuv420 *yuv = (ttLibC_Yuv420 *)frame;
			return (uintptr_t)yuv->v_data;
		}
		break;
	default:
		return 0;
	}
}




bool Frame_restore(
		void *frame,
		const char *codec_type,
		const char *sub_type,
		uint64_t pts,
		uint32_t timebase,
		uint32_t id,
		uint32_t width,
		uint32_t height,
		const char *video_type,
		uint32_t sample_rate,
		uint32_t sample_num,
		uint32_t channel_num,
		uintptr_t data,
		uintptr_t l_data,
		uintptr_t r_data,
		uintptr_t y_data,
		uintptr_t u_data,
		uintptr_t v_data,
		uint32_t stride,
		uint32_t y_stride,
		uint32_t u_stride,
		uint32_t v_stride,
		uintptr_t ptr1,
		uintptr_t ptr2,
		uintptr_t ptr3) {
	// codecTypeについては、調整しなくていいか
	if(frame == NULL) {
		return false;
	}
	ttLibC_Frame *f = (ttLibC_Frame *)frame;
	// subtypeもほっとく
	f->pts = pts;
	f->timebase = timebase;
	f->id = id;
	// width heightは普通はかえないけど、yuvとかbgrならいじることがあるわけか・・・
	// video_typeもさわらない
	switch(f->type) {
		// この4つだけ
 	case frameType_pcmF32:
	case frameType_pcmS16:
		// まぁ本当は基本さわらない
		break;
	case frameType_bgr:
		{
			ttLibC_Bgr *bgr = (ttLibC_Bgr *)frame;
			bgr->inherit_super.width = width;
			bgr->inherit_super.height = height;
			bgr->data = (uint8_t *)(ptr1 + data);
		}
		break;
	case frameType_yuv420:
		{
			ttLibC_Yuv420 *yuv = (ttLibC_Yuv420 *)frame;
			yuv->inherit_super.width = width;
			yuv->inherit_super.height = height;
			yuv->y_data = (uint8_t *)(ptr1 + y_data);
			yuv->u_data = (uint8_t *)(ptr2 + u_data);
			yuv->v_data = (uint8_t *)(ptr3 + v_data);
		}
		break;
	default:
		break;
	}
	return true;
}

// binaryのデータが設定されている場合は、これをcallする
// frameが生成された場合はそのフレームが応答される。
// エラー時には、生成されない
ttLibC_Frame *Frame_fromBinary(
		void *frame,
		const char *codec_type,
		const char *sub_type,
		uint64_t pts,
		uint32_t timebase,
		uint32_t width,
		uint32_t height,
		uint32_t sample_rate,
		uint32_t sample_num,
		uint32_t channel_num,
		void  *data, // このデータはunsafePointerとして、go側からやってくるので、参照として保持しておけばよいと思う。
		size_t data_size,
		uint32_t stride,
		uint32_t y_stride,
		uint32_t u_stride,
		uint32_t v_stride) {
	if(data == NULL || data_size == 0) {
		// 復元するフレームのbinaryデータがそもそもないので、動作しない。
		return NULL;
	}
	ttLibC_Frame *f = (ttLibC_Frame *)frame;
	ttLibC_Frame_Type frame_type = Frame_getFrameType(codec_type);
	// 前のフレームを解放したけど、データの更新な叶わなかった場合・・・こまったことになる。
	// よって複数の状態を応答したいところ // 先に解放しなければいいのか・・・
	// 違う場合は、データができたときのみ、解放するとすればいけるか？
	ttLibC_Frame *prev_frame = NULL;
	if(frame != NULL && f->type == frame_type) {
		prev_frame = frame;
	}
	ttLibC_Frame *new_frame = NULL;
	switch(frame_type) {
	case frameType_aac:
		{
			new_frame = (ttLibC_Frame *)ttLibC_Aac_getFrame(
				(ttLibC_Aac *)prev_frame,
				data,
				data_size,
				true,
				pts,
				timebase);
		}
		break;
	case frameType_adpcm_ima_wav:
		{
			if(sample_rate == 0 || channel_num == 0) {
				return NULL;
			}
			new_frame = (ttLibC_Frame *)ttLibC_AdpcmImaWav_make(
				(ttLibC_AdpcmImaWav *)prev_frame,
				sample_rate,
				sample_num,
				channel_num,
				data,
				data_size,
				true,
				pts,
				timebase);
		}
		break;
	case frameType_mp3:
		{
			new_frame = (ttLibC_Frame *)ttLibC_Mp3_getFrame(
				(ttLibC_Mp3 *)prev_frame,
				data,
				data_size,
				true,
				pts,
				timebase);
		}
		break;
	case frameType_nellymoser:
		{
			if(sample_rate == 0 || channel_num == 0) {
				return NULL;
			}
			new_frame = (ttLibC_Frame *)ttLibC_Nellymoser_make(
				(ttLibC_Nellymoser *)prev_frame,
				sample_rate,
				sample_num,
				channel_num,
				data,
				data_size,
				true,
				pts,
				timebase);
		}
		break;
	case frameType_opus:
		{
			new_frame = (ttLibC_Frame *)ttLibC_Opus_getFrame(
				(ttLibC_Opus *)prev_frame,
				data,
				data_size,
				true,
				pts,
				timebase);
		}
		break;
	case frameType_pcm_alaw:
		{
			if(sample_rate == 0) {
				return NULL;
			}
			new_frame = (ttLibC_Frame *)ttLibC_PcmAlaw_make(
				(ttLibC_PcmAlaw *)prev_frame,
				sample_rate,
				data_size,
				1,
				data,
				data_size,
				true,
				pts,
				timebase);
		}
		break;
	case frameType_pcmF32:
		{
			// subTypeを確認する必要がある
			// memory allocを実施する必要はなさそう
			if(sub_type == NULL || sample_rate == 0 || channel_num == 0) {
 				return NULL;
			}
			if(sample_num == 0) {
				sample_num = data_size / 4 / channel_num;
			}
			else {
				if(sample_num > data_size / 4 / channel_num) {
					// データサイズが小さすぎる
					return NULL;
				}
			}
			ttLibC_PcmF32_Type pcm_type = PcmF32Type_interleave;
			uint8_t *l_data = data;
			uint8_t *r_data = NULL;
			size_t l_stride = data_size;
			size_t r_stride = 0;

			if(strcmp(sub_type, "planar") == 0) {
				pcm_type = PcmF32Type_planar;
				l_stride = sample_num * 4;
				if(channel_num == 2) {
					r_data = l_data + l_stride;
					r_stride = l_stride;
				}
			}
			else if(strcmp(sub_type, "interleave") == 0) {
				pcm_type = PcmF32Type_interleave;
				l_stride = sample_num * 4 * channel_num;
			}
			else {
				return NULL;
			}
			new_frame = (ttLibC_Frame *)ttLibC_PcmF32_make(
				(ttLibC_PcmF32 *)prev_frame,
				pcm_type,
				sample_rate,
				sample_num,
				channel_num,
				data,
				data_size,
				l_data,
				l_stride,
				r_data,
				r_stride,
				true,
				pts,
				timebase);
		}
		break;
	case frameType_pcm_mulaw:
		{
			if(sample_rate == 0) {
				return NULL;
			}
			new_frame = (ttLibC_Frame *)ttLibC_PcmMulaw_make(
				(ttLibC_PcmMulaw *)prev_frame,
				sample_rate,
				data_size,
				1,
				data,
				data_size,
				true,
				pts,
				timebase);
		}
		break;
	case frameType_pcmS16:
		{
			if(sub_type == NULL || sample_rate == 0 || channel_num == 0) {
				return NULL;
			}
			if(sample_num == 0) {
				sample_num = data_size / 2 / channel_num;
			}
			else {
				if(sample_num > data_size / 2 / channel_num) {
					return NULL;
				}
			}
			ttLibC_PcmS16_Type pcm_type = PcmS16Type_littleEndian;
			uint8_t *l_data = data;
			uint8_t *r_data = NULL;
			size_t l_stride = data_size;
			size_t r_stride = 0;

			if(strcmp(sub_type, "bigEndian") == 0) {
				pcm_type = PcmS16Type_bigEndian;
				l_stride = sample_num * 2 * channel_num;
			}
			else if(strcmp(sub_type, "littleEndian") == 0) {
				pcm_type = PcmS16Type_littleEndian;
				l_stride = sample_num * 2 * channel_num;
			}
			else if(strcmp(sub_type, "bigEndianPlanar") == 0) {
				pcm_type = PcmS16Type_bigEndian_planar;
				l_stride = sample_num * 2;
				if(channel_num == 2) {
					r_data = l_data + l_stride;
					r_stride = l_stride;
				}
			}
			else if(strcmp(sub_type, "littleEndianPlanar") == 0) {
				pcm_type = PcmS16Type_littleEndian_planar;
				l_stride = sample_num * 2;
				if(channel_num == 2) {
					r_data = l_data + l_stride;
					r_stride = l_stride;
				}
			}
			else {
				return NULL;
			}
			new_frame = (ttLibC_Frame *)ttLibC_PcmS16_make(
				(ttLibC_PcmS16 *)prev_frame,
				pcm_type,
				sample_rate,
				sample_num,
				channel_num,
				data,
				data_size,
				l_data,
				l_stride,
				r_data,
				r_stride,
				true,
				pts,
				timebase);
		}
		break;
	case frameType_speex:
		{
			new_frame = (ttLibC_Frame *)ttLibC_Speex_getFrame(
				(ttLibC_Speex *)prev_frame,
				data,
				data_size,
				true,
				pts,
				timebase);
		}
		break;
	case frameType_vorbis:
		{
			new_frame = (ttLibC_Frame *)ttLibC_Vorbis_getFrame(
				(ttLibC_Vorbis *)prev_frame,
				data,
				data_size,
				true,
				pts,
				timebase);
		}
		break;

		// あとは映像系のフレーム
	case frameType_bgr:
		{
			if(width == 0 || height == 0) {
				return NULL;
			}
//			uint32_t stride = 0;
			ttLibC_Bgr_Type bgr_type = BgrType_abgr;
			if(strcmp(sub_type, "bgr") == 0) {
				bgr_type = BgrType_bgr;
				if(stride == 0) {
					stride = width * 3;
				}
			}
			else if(strcmp(sub_type, "rgb") == 0) {
				bgr_type = BgrType_rgb;
				if(stride == 0) {
					stride = width * 3;
				}
			}
			else if(strcmp(sub_type, "bgra") == 0) {
				bgr_type = BgrType_bgra;
				if(stride == 0) {
					stride = width * 4;
				}
			}
			else if(strcmp(sub_type, "abgr") == 0) {
				bgr_type = BgrType_abgr;
				if(stride == 0) {
					stride = width * 4;
				}
			}
			else if(strcmp(sub_type, "argb") == 0) {
				bgr_type = BgrType_argb;
				if(stride == 0) {
					stride = width * 4;
				}
			}
			else if(strcmp(sub_type, "rgba") == 0) {
				bgr_type = BgrType_rgba;
				if(stride == 0) {
					stride = width * 4;
				}
			}
			else {
				return NULL;
			}
			if(data_size < stride * height) {
				return NULL;
			}
			new_frame = (ttLibC_Frame *)ttLibC_Bgr_make(
				(ttLibC_Bgr *)prev_frame,
				bgr_type,
				width,
				height,
				stride,
				data,
				data_size,
				true,
				pts,
				timebase);
		}
		break;
	case frameType_flv1:
		{
			new_frame = (ttLibC_Frame *)ttLibC_Flv1_getFrame(
				(ttLibC_Flv1 *)prev_frame,
				data,
				data_size,
				true,
				pts,
				timebase);
		}
		break;
	case frameType_h264:
		{
			new_frame = (ttLibC_Frame *)ttLibC_H264_getFrame(
				(ttLibC_H264 *)prev_frame,
				data,
				data_size,
				true,
				pts,
				timebase);
		}
		break;
	case frameType_h265:
		{
			new_frame = (ttLibC_Frame *)ttLibC_H265_getFrame(
				(ttLibC_H265 *)prev_frame,
				data,
				data_size,
				true,
				pts,
				timebase);
		}
		break;
	case frameType_jpeg:
		{
			new_frame = (ttLibC_Frame *)ttLibC_Jpeg_getFrame(
				(ttLibC_Jpeg *)prev_frame,
				data,
				data_size,
				true,
				pts,
				timebase);
		}
		break;
	case frameType_png:
		{
			new_frame = (ttLibC_Frame *)ttLibC_Png_getFrame(
				(ttLibC_Png *)prev_frame,
				data,
				data_size,
				true,
				pts,
				timebase);
		}
		break;
	case frameType_theora:
		{
			new_frame = (ttLibC_Frame *)ttLibC_Theora_getFrame(
				(ttLibC_Theora *)prev_frame,
				data,
				data_size,
				true,
				pts,
				timebase);
		}
		break;
	case frameType_vp6:
		{
			uint32_t adjustment = 0x00;
			new_frame = (ttLibC_Frame *)ttLibC_Vp6_getFrame(
				(ttLibC_Vp6 *)prev_frame,
				data,
				data_size,
				true,
				pts,
				timebase,
				adjustment);
		}
		break;
	case frameType_vp8:
		{
			new_frame = (ttLibC_Frame *)ttLibC_Vp8_getFrame(
				(ttLibC_Vp8 *)prev_frame,
				data,
				data_size,
				true,
				pts,
				timebase);
		}
		break;
	case frameType_vp9:
		{
			new_frame = (ttLibC_Frame *)ttLibC_Vp9_getFrame(
				(ttLibC_Vp9 *)prev_frame,
				data,
				data_size,
				true,
				pts,
				timebase);
		}
		break;
	case frameType_yuv420:
		{
			if(width == 0 || height == 0) {
				return NULL;
			}
			if(y_stride == 0) {
				y_stride = width;
			}
			if(u_stride == 0) {
				u_stride = width >> 1;
			}
			if(v_stride == 0) {
				v_stride = width >> 1;
			}
			uint8_t *y_data;
			uint8_t *u_data;
			uint8_t *v_data;
			ttLibC_Yuv420_Type yuv_type = Yuv420Type_planar;
			if(strcmp(sub_type, "planar") == 0) {
				yuv_type = Yuv420Type_planar;
				y_data = data;
				u_data = y_data + y_stride * height;
				v_data = u_data + u_stride * (height >> 1);
			}
			else if(strcmp(sub_type, "yvuPlanar") == 0) {
				yuv_type = Yuv420Type_semiPlanar;
				y_data = data;
				v_data = y_data + y_stride * height;
				u_data = v_data + v_stride * (height >> 1);
			}
			else if(strcmp(sub_type, "semiPlanar") == 0) {
				yuv_type = Yvu420Type_planar;
				u_stride = width;
				v_stride = width;
				y_data = data;
				u_data = y_data + y_stride * height;
				v_data = u_data + 1;
			}
			else if(strcmp(sub_type, "yvuSemiPlanar") == 0) {
				yuv_type = Yvu420Type_semiPlanar;
				u_stride = width;
				v_stride = width;
				y_data = data;
				v_data = y_data + y_stride * height;
				u_data = v_data + 1;
			}
			else {
				return NULL;
			}
			new_frame = (ttLibC_Frame *)ttLibC_Yuv420_make(
				(ttLibC_Yuv420 *)prev_frame,
				yuv_type,
				width,
				height,
				data,
				data_size,
				y_data,
				y_stride,
				u_data,
				u_stride,
				v_data,
				v_stride,
				true,
				pts,
				timebase);
		}
		break;
	default:
		return NULL;
	}
	if(new_frame != NULL // new_frameが生成された場合
	&& prev_frame == NULL // reuse frameには選ばれなかった
	&& frame != NULL) { // 入力フレームはある。
		// この場合は新規フレームはできているが、前のフレームがreuseされなかったことになるため、前のフレームは解放しなければならない。
		ttLibC_Frame_close(&f);
	}
	return new_frame;
}

ttLibC_Frame *Frame_refFrame(void *frame) {
	return (ttLibC_Frame *)frame;
}

void Frame_copyData(ttLibC_Frame *frame, void *data, size_t data_size) {
	if(data == NULL) {
		// コピー先がない場合
		return;
	}
	if(frame == NULL) {
		return;
	}
	if(frame->data == NULL) {
		// 元データがない場合もしくはそのデータがない場合
		return;
	}
	if(data_size == 0 || data_size > frame->data_size) {
		// コピーするデータサイズがそもそもない場合、data_sizeが保持しているデータより大きい場合
		return;
	}
	memcpy(data, frame->data, data_size);
}

void Frame_close(void *frame) {
	ttLibC_Frame *f = (ttLibC_Frame *)frame;
	if(f == NULL) {
		return;
	}
	ttLibC_Frame_close(&f);
}

size_t Frame_refBufferSize(void *frame) {
	ttLibC_Frame *f = (ttLibC_Frame *)frame;
	if(f == NULL) {
		return 0;
	}
	return f->buffer_size;
}

*/
import "C"
import (
	"bytes"
	"unsafe"
)

// CttLibCFrame c言語のframeのデータに一応名前をつけておく
// 内部としては、void*扱い
type CttLibCFrame unsafe.Pointer

// Frame goで扱うframeのオブジェクトの定義
// いろいろと参照できるデータをキープしておく
// cloneしたりbinaryから復元する動作の場合は受け皿のttLibGo.Frame
// を初めに生成してから、それを使いまわす。
// 最後にcloseする必要がある
// encoder・decoder・readerで生成されたデータについては、生成したオブジェクトが管理するので
// closeする必要はない
type Frame struct {
	CFrame       CttLibCFrame
	CodecType    string
	SubType      string
	Pts          uint64
	Timebase     uint32
	ID           uint32
	Width        uint32
	Height       uint32
	VideoType    string
	H26xType     string // これ・・・あれかh264のやつか
	IsDisposable bool
	SampleRate   uint32
	SampleNum    uint32
	ChannelNum   uint32
	Data         uintptr // bgrのデータのズレ分
	LData        uintptr // pcmS16 F32のleft側音声のデータのズレ分
	RData        uintptr // right側
	YData        uintptr // yuvのy要素のデータのズレ分
	UData        uintptr // yuvのu要素のデータのズレ分
	VData        uintptr // yuvのv要素のデータのズレ分
	Stride       uint32
	YStride      uint32
	UStride      uint32
	VStride      uint32
	ptrArray     [3]uintptr
	Binary       []byte // データをbinaryから復元するときに利用する
	binary       []byte // binaryから復元したときに、元のbinaryをkeepする場所
	hasBody      bool   // このデータがあとで解放すべきであるかを指定している
}

// Init データによって、フレームを初期化します。
// 他のttLibGo拡張でも利用するので、公開状態にしておきます
func (frame *Frame) Init(cFrame CttLibCFrame) {
	// フレームを初期化する
	frame.CFrame = cFrame
	buf := make([]byte, 256)
	ptr := unsafe.Pointer(cFrame)
	C.Frame_getCodecType(ptr, (*C.char)(unsafe.Pointer(&buf[0])), C.size_t(256))
	frame.CodecType = string(buf[:bytes.IndexByte(buf, 0)])
	buf = make([]byte, 256)
	C.Frame_getSubType(ptr, (*C.char)(unsafe.Pointer(&buf[0])), C.size_t(256))
	frame.SubType = string(buf[:bytes.IndexByte(buf, 0)])
	frame.Pts = uint64(C.Frame_getPts(ptr))
	frame.Timebase = uint32(C.Frame_getTimebase(ptr))
	frame.ID = uint32(C.Frame_getID(ptr))
	frame.Width = uint32(C.Frame_getWidth(ptr))
	frame.Height = uint32(C.Frame_getHeight(ptr))
	buf = make([]byte, 256)
	C.Frame_getVideoType(ptr, (*C.char)(unsafe.Pointer(&buf[0])), C.size_t(256))
	frame.VideoType = string(buf[:bytes.IndexByte(buf, 0)])
	buf = make([]byte, 256)
	C.Frame_getH26xType(ptr, (*C.char)(unsafe.Pointer(&buf[0])), C.size_t(256))
	frame.H26xType = string(buf[:bytes.IndexByte(buf, 0)])
	frame.SampleRate = uint32(C.Frame_getSampleRate(ptr))
	frame.SampleNum = uint32(C.Frame_getSampleNum(ptr))
	frame.ChannelNum = uint32(C.Frame_getChannelNum(ptr))

	frame.Data = 0
	frame.LData = 0
	frame.RData = 0
	frame.YData = 0
	frame.UData = 0
	frame.VData = 0

	frame.Stride = uint32(C.Frame_getStride(ptr))
	frame.YStride = uint32(C.Frame_getYStride(ptr))
	frame.UStride = uint32(C.Frame_getUStride(ptr))
	frame.VStride = uint32(C.Frame_getVStride(ptr))

	frame.ptrArray[0] = uintptr(C.Frame_getData1Ptr(ptr))
	frame.ptrArray[1] = uintptr(C.Frame_getData2Ptr(ptr))
	frame.ptrArray[2] = uintptr(C.Frame_getData3Ptr(ptr))
	// あと他の情報も更新しなければならないな。
	frame.hasBody = false
}

// Restore frameの内容を構造体の値で更新します
func (frame *Frame) Restore() {
	// 設定はするけど、更新しない・・・という状況が必要になるのか微妙なところなので
	// 今回はこの処理はださないでおいておこうかな
	cCodecType := C.CString(frame.CodecType)
	defer C.free(unsafe.Pointer(cCodecType))
	cSubType := C.CString(frame.SubType)
	defer C.free(unsafe.Pointer(cSubType))
	cVideoType := C.CString(frame.VideoType)
	defer C.free(unsafe.Pointer(cVideoType))
	var binaryPointer unsafe.Pointer
	var binarySize int
	// binaryからすでにデータを復元している場合は・・・という判定が必要だったのね・・・
	if frame.Binary != nil {
		// ちゃんと動作した
		frame.binary = frame.Binary // こっちに保持させておいて、データが蒸発しないようにしておく
		binaryPointer = unsafe.Pointer(&frame.binary[0])
		binarySize = len(frame.binary)
		// binaryデータがある場合は、binaryから復元しなければならない
		cFrame := C.Frame_fromBinary(
			unsafe.Pointer(frame.CFrame), // 前まで利用していたフレーム
			cCodecType,
			cSubType,
			C.uint64_t(frame.Pts),
			C.uint32_t(frame.Timebase),
			C.uint32_t(frame.Width),
			C.uint32_t(frame.Height),
			C.uint32_t(frame.SampleRate),
			C.uint32_t(frame.SampleNum),
			C.uint32_t(frame.ChannelNum),
			binaryPointer,
			C.size_t(binarySize),
			C.uint32_t(frame.Stride),
			C.uint32_t(frame.YStride),
			C.uint32_t(frame.UStride),
			C.uint32_t(frame.VStride))

		if cFrame != nil {
			frame.CFrame = CttLibCFrame(cFrame)
			// 特定のフレームの場合はポインタの位置を復元しなければならない
			frame.ptrArray[0] = uintptr(C.Frame_getData1Ptr(unsafe.Pointer(cFrame)))
			frame.ptrArray[1] = uintptr(C.Frame_getData2Ptr(unsafe.Pointer(cFrame)))
			frame.ptrArray[2] = uintptr(C.Frame_getData3Ptr(unsafe.Pointer(cFrame)))
		}
		frame.Binary = nil
		frame.hasBody = true
	}
	// 保持している情報でframeを復元します
	C.Frame_restore(
		unsafe.Pointer(frame.CFrame),
		cCodecType,
		cSubType,
		C.uint64_t(frame.Pts),
		C.uint32_t(frame.Timebase),
		C.uint32_t(frame.ID),
		C.uint32_t(frame.Width),
		C.uint32_t(frame.Height),
		cVideoType,
		C.uint32_t(frame.SampleRate),
		C.uint32_t(frame.SampleNum),
		C.uint32_t(frame.ChannelNum),
		C.uintptr_t(frame.Data),
		C.uintptr_t(frame.LData),
		C.uintptr_t(frame.RData),
		C.uintptr_t(frame.YData),
		C.uintptr_t(frame.UData),
		C.uintptr_t(frame.VData),
		C.uint32_t(frame.Stride),
		C.uint32_t(frame.YStride),
		C.uint32_t(frame.UStride),
		C.uint32_t(frame.VStride),
		C.uintptr_t(frame.ptrArray[0]),
		C.uintptr_t(frame.ptrArray[1]),
		C.uintptr_t(frame.ptrArray[2]))
}

// Copy オリジナルのフレームの内容を自身にコピーする
func (frame *Frame) Copy(original *Frame) bool {
	// もともとcloneという動作でやってたけど、copyにしようと思います。
	// 内容としては originalのフレームの内容を自分のフレームにコピーするというもの
	// やってることはcloneと同じだけど、取り回しを考えるとこの方がいいかなと思った
	clonedCFrame := C.ttLibC_Frame_clone(C.Frame_refFrame(unsafe.Pointer(frame.CFrame)), C.Frame_refFrame(unsafe.Pointer(original.CFrame)))
	if clonedCFrame == nil {
		return false
	}
	frame.Init(CttLibCFrame(clonedCFrame))
	frame.hasBody = true
	return true
}

// Close 内部で保持しているデータを解放する
// 内部のttLibCがgoで生成されたものではなければ解放する必要はない
func (frame *Frame) Close() {
	// cloneしたもの、bufferから復元してつくったものはCFrameを解放しなければならない
	if frame.hasBody {
		C.Frame_close(unsafe.Pointer(frame.CFrame))
	}
	frame.CFrame = nil
}

// GetBinaryBuffer フレームが保持しているbinaryデータを参照する
func (frame *Frame) GetBinaryBuffer() []byte {
	if frame.CFrame == nil {
		// フレームが存在しない場合は処理しない
		return nil
	}
	if frame.binary != nil {
		// binaryから復元したデータの場合は、そのbinaryがあるので応答してやる
		return frame.binary
	}
	// 保持しているframeをttLibC_Frameとしてcastしておく
	cFrame := C.Frame_refFrame(unsafe.Pointer(frame.CFrame))
	switch frame.CodecType {
	case "bgr":
		fallthrough
	case "yuv":
		fallthrough
	case "pcmF32":
		fallthrough
	case "pcmS16":
		cloned := C.ttLibC_Frame_clone(nil, cFrame)
		bufferSize := cloned.buffer_size
		if bufferSize == 0 {
			// c言語でcloneしても保持しているbuffer_sizeが0だったら応答できるものがないので、0を応答
			return nil
		}
		// bufferを作る
		buf := make([]byte, bufferSize)
		// データをコピーしてgoの世界にもってくる
		C.Frame_copyData(cloned, unsafe.Pointer(&buf[0]), bufferSize)
		// cloneしたフレームはいらないので、解放する
		C.ttLibC_Frame_close(&cloned)
		// bufferを応答しておわり。
		return buf
	default:
		if cFrame.data == nil {
			return nil
		} else {
			bufferSize := cFrame.buffer_size
			if bufferSize == 0 {
				return nil
			}
			buf := make([]byte, bufferSize)
			C.Frame_copyData(cFrame, unsafe.Pointer(&buf[0]), bufferSize)
			return buf
		}
	}
}
