#include <stdio.h>
#include <string.h>
#include <ttLibC/frame/frame.h>
#include <ttLibC/frame/audio/audio.h>
#include <ttLibC/frame/audio/aac.h>
#include <ttLibC/frame/audio/mp3.h>
#include <ttLibC/frame/audio/opus.h>
#include <ttLibC/frame/audio/pcmf32.h>
#include <ttLibC/frame/audio/pcms16.h>
#include <ttLibC/frame/audio/speex.h>
#include <ttLibC/frame/audio/vorbis.h>

/**
 * frameのsampleRateの値を取得する
 * @param frame
 * @return sampleRateの値
 */
uint32_t AudioFrame_getSampleRate(void *frame) {
	if(frame == NULL || !ttLibC_isAudio(((ttLibC_Frame *)frame)->type)) {
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
	if(frame == NULL || !ttLibC_isAudio(((ttLibC_Frame *)frame)->type)) {
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
	if(frame == NULL || !ttLibC_isAudio(((ttLibC_Frame *)frame)->type)) {
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
