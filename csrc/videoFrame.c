#include <stdio.h>
#include <string.h>
#include <ttLibC/frame/frame.h>
#include <ttLibC/frame/video/video.h>
#include <ttLibC/frame/video/bgr.h>
#include <ttLibC/frame/video/flv1.h>
#include <ttLibC/frame/video/h264.h>
#include <ttLibC/frame/video/h265.h>
#include <ttLibC/frame/video/theora.h>
#include <ttLibC/frame/video/yuv420.h>

/**
 * frameのvideoTypeの値を取得する
 * @param cFrame
 * @param buffer
 * @param buffer_size
 */
void VideoFrame_getVideoType(void *cFrame, char *buffer, size_t buffer_size) {
	if(cFrame == NULL || !ttLibC_isVideo(((ttLibC_Frame *)cFrame)->type) || buffer == NULL || buffer_size < 6) {
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
	if(frame == NULL || !ttLibC_isVideo(((ttLibC_Frame *)frame)->type)) {
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
	if(frame == NULL || !ttLibC_isVideo(((ttLibC_Frame *)frame)->type)) {
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

