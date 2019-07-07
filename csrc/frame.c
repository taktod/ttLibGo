#include <stdio.h>
#include <string.h>
#include <ttLibC/frame/frame.h>

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
  ttLibC_Frame *f = (ttLibC_Frame *)frame;
  ttLibC_Frame_close(&f);
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
  ttLibC_Frame *f = (ttLibC_Frame *)frame;
  return ttLibC_Frame_clone(NULL, f);
}
