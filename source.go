package ttLibGo

/*
#cgo CFLAGS: -I./ -I./ttLibC/ -I./cppsrc/ -std=c99 -Wno-multichar
#cgo LDFLAGS: -L/usr/local/lib -ldl
#cgo CXXFLAGS: -std=c++0x -I./ -I./ttLibC/ -I./cppsrc/ -Wno-multichar

#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <dlfcn.h>

#include "ttLibC/ttLibC/container/container.h"
#include "ttLibC/ttLibC/encoder/openh264Encoder.h"
#include "ttLibC/ttLibC/encoder/x264Encoder.h"
#include "ttLibC/ttLibC/encoder/x265Encoder.h"
#include "ttLibC/ttLibC/frame/video/bgr.h"
#include "ttLibC/ttLibC/frame/video/theora.h"
#include "ttLibC/ttLibC/frame/video/video.h"
#include "ttLibC/ttLibC/frame/video/yuv420.h"
#include "ttLibC/ttLibC/frame/audio/aac.h"
#include "ttLibC/ttLibC/frame/audio/pcmf32.h"
#include "ttLibC/ttLibC/frame/audio/pcms16.h"
#include "ttLibC/ttLibC/frame/audio/speex.h"
#include "ttLibC/ttLibC/frame/audio/vorbis.h"
#include "ttLibC/ttLibC/resampler/swscaleResampler.h"
#include "ttLibC/ttLibC/resampler/libyuvResampler.h"
#include "ttLibC/ttLibC/util/byteUtil.h"

typedef void *(* ttLibC_make_func)();
typedef bool (* ttLibC_ContainerReader_read_func)(void *, void *, size_t, bool(*)(void *, void *), void *);
typedef bool (* ttLibC_Container_GetFrame_func)(void *, ttLibC_getFrameFunc, void *);
typedef void (* ttLibC_close_func)(void **);

ttLibC_make_func ttLibGo_FlvReader_make = NULL;
ttLibC_make_func ttLibGo_MkvReader_make = NULL;
ttLibC_make_func ttLibGo_Mp4Reader_make = NULL;
ttLibC_make_func ttLibGo_MpegtsReader_make = NULL;

ttLibC_ContainerReader_read_func ttLibGo_FlvReader_read = NULL;
ttLibC_ContainerReader_read_func ttLibGo_MkvReader_read = NULL;
ttLibC_ContainerReader_read_func ttLibGo_Mp4Reader_read = NULL;
ttLibC_ContainerReader_read_func ttLibGo_MpegtsReader_read = NULL;

ttLibC_Container_GetFrame_func ttLibGo_Flv_getFrame = NULL;
ttLibC_Container_GetFrame_func ttLibGo_Mkv_getFrame = NULL;
ttLibC_Container_GetFrame_func ttLibGo_Mp4_getFrame = NULL;
ttLibC_Container_GetFrame_func ttLibGo_Mpegts_getFrame = NULL;

ttLibC_close_func ttLibGo_ContainerReader_close = NULL;

typedef void *(* ttLibC_AvcodecDecoder_make_func)(ttLibC_Frame_Type, uint32_t, uint32_t);
typedef void *(* ttLibC_codec_func)(void *, void *, ttLibC_getFrameFunc, void *);

ttLibC_AvcodecDecoder_make_func ttLibGo_AvcodecVideoDecoder_make = NULL;
ttLibC_AvcodecDecoder_make_func ttLibGo_AvcodecAudioDecoder_make = NULL;
ttLibC_codec_func               ttLibGo_AvcodecDecoder_decode = NULL;
ttLibC_close_func               ttLibGo_AvcodecDecoder_close = NULL;

ttLibC_make_func  ttLibGo_JpegDecoder_make = NULL;
ttLibC_codec_func ttLibGo_JpegDecoder_decode = NULL;
ttLibC_close_func ttLibGo_JpegDecoder_close = NULL;

ttLibC_make_func  ttLibGo_Openh264Decoder_make = NULL;
ttLibC_codec_func ttLibGo_Openh264Decoder_decode = NULL;
ttLibC_close_func ttLibGo_Openh264Decoder_close = NULL;

typedef void *(* ttLibC_OpusDecoder_make_func)(uint32_t, uint32_t);
typedef int (* ttLibC_opus_codecControl_func)(void *, const char *, int);

ttLibC_OpusDecoder_make_func  ttLibGo_OpusDecoder_make = NULL;
ttLibC_codec_func             ttLibGo_OpusDecoder_decode = NULL;
ttLibC_close_func             ttLibGo_OpusDecoder_close = NULL;
ttLibC_opus_codecControl_func ttLibGo_OpusDecoder_codecControl = NULL;

ttLibC_make_func  ttLibGo_PngDecoder_make = NULL;
ttLibC_codec_func ttLibGo_PngDecoder_decode = NULL;
ttLibC_close_func ttLibGo_PngDecoder_close = NULL;

typedef void *(* ttLibC_SpeexDecoder_make_func)(uint32_t, uint32_t);

ttLibC_SpeexDecoder_make_func ttLibGo_SpeexDecoder_make = NULL;
ttLibC_codec_func             ttLibGo_SpeexDecoder_decode = NULL;
ttLibC_close_func             ttLibGo_SpeexDecoder_close = NULL;

ttLibC_make_func  ttLibGo_TheoraDecoder_make = NULL;
ttLibC_codec_func ttLibGo_TheoraDecoder_decode = NULL;
ttLibC_close_func ttLibGo_TheoraDecoder_close = NULL;

ttLibC_make_func  ttLibGo_VorbisDecoder_make = NULL;
ttLibC_codec_func ttLibGo_VorbisDecoder_decode = NULL;
ttLibC_close_func ttLibGo_VorbisDecoder_close = NULL;

typedef void *(* ttLibC_FdkaacEncoder_make_func)(const char *,uint32_t,uint32_t,uint32_t);
typedef bool (* ttLibC_FdkaacEncoder_setBitrate_func)(void *,uint32_t);

ttLibC_FdkaacEncoder_make_func       ttLibGo_FdkaacEncoder_make = NULL;
ttLibC_codec_func                    ttLibGo_FdkaacEncoder_encode = NULL;
ttLibC_close_func                    ttLibGo_FdkaacEncoder_close = NULL;
ttLibC_FdkaacEncoder_setBitrate_func ttLibGo_FdkaacEncoder_setBitrate = NULL;

typedef void *(* ttLibC_JpegEncoder_make_func)(uint32_t, uint32_t, uint32_t);
typedef void *(* ttLibC_JpegEncoder_setQuality_func)(void *, uint32_t);

ttLibC_JpegEncoder_make_func       ttLibGo_JpegEncoder_make = NULL;
ttLibC_codec_func                  ttLibGo_JpegEncoder_encode = NULL;
ttLibC_JpegEncoder_setQuality_func ttLibGo_JpegEncoder_setQuality = NULL;
ttLibC_close_func                  ttLibGo_JpegEncoder_close = NULL;

typedef void (* ttLibC_Openh264Encoder_getDefaultSEncParamExt_func)(void *, uint32_t, uint32_t);
typedef bool (* ttLibC_Openh264Encoder_paramParse_func)(void *, const char *, const char *);
typedef bool (* ttLibC_Openh264Encoder_spatialParamParse_func)(void *, uint32_t, const char *, const char *);
typedef void *(* ttLibC_Openh264Encoder_makeWithSEncParamExt_func)(void *);
typedef bool (* ttLibC_Openh264Encoder_setRCMode_func)(void *, ttLibC_Openh264Encoder_RCType);
typedef bool (* ttLibC_Openh264Encoder_setIDRInterval_func)(void *, uint32_t);
typedef bool (* ttLibC_Openh264Encoder_forceNextKeyFrame_func)(void *);

ttLibC_Openh264Encoder_getDefaultSEncParamExt_func ttLibGo_Openh264Encoder_getDefaultSEncParamExt = NULL;
ttLibC_Openh264Encoder_paramParse_func             ttLibGo_Openh264Encoder_paramParse = NULL;
ttLibC_Openh264Encoder_spatialParamParse_func      ttLibGo_Openh264Encoder_spatialParamParse = NULL;
ttLibC_Openh264Encoder_makeWithSEncParamExt_func   ttLibGo_Openh264Encoder_makeWithSEncParamExt = NULL;
ttLibC_codec_func                                  ttLibGo_Openh264Encoder_encode = NULL;
ttLibC_close_func                                  ttLibGo_Openh264Encoder_close = NULL;
ttLibC_Openh264Encoder_setRCMode_func              ttLibGo_Openh264Encoder_setRCMode = NULL;
ttLibC_Openh264Encoder_setIDRInterval_func         ttLibGo_Openh264Encoder_setIDRInterval = NULL;
ttLibC_Openh264Encoder_forceNextKeyFrame_func      ttLibGo_Openh264Encoder_forceNextKeyFrame = NULL;

typedef void *(* ttLibC_OpusEncoder_make_func)(uint32_t, uint32_t, uint32_t);

ttLibC_OpusEncoder_make_func  ttLibGo_OpusEncoder_make = NULL;
ttLibC_codec_func             ttLibGo_OpusEncoder_encode = NULL;
ttLibC_opus_codecControl_func ttLibGo_OpusEncoder_codecControl = NULL;
ttLibC_close_func             ttLibGo_OpusEncoder_close = NULL;

typedef void *(* ttLibC_SpeexEncoder_make_func)(uint32_t, uint32_t, uint32_t);

ttLibC_SpeexEncoder_make_func ttLibGo_SpeexEncoder_make = NULL;
ttLibC_codec_func             ttLibGo_SpeexEncoder_encode = NULL;
ttLibC_close_func             ttLibGo_SpeexEncoder_close = NULL;

typedef void *(* ttLibC_TheoraEncoder_make_ex_func)(uint32_t, uint32_t, uint32_t, uint32_t, uint32_t);

ttLibC_TheoraEncoder_make_ex_func ttLibGo_TheoraEncoder_make_ex = NULL;
ttLibC_codec_func                 ttLibGo_TheoraEncoder_encode = NULL;
ttLibC_close_func                 ttLibGo_TheoraEncoder_close = NULL;

typedef void *(* ttLibC_VorbisEncoder_make_func)(uint32_t, uint32_t);

ttLibC_VorbisEncoder_make_func ttLibGo_VorbisEncoder_make = NULL;
ttLibC_codec_func ttLibGo_VorbisEncoder_encode = NULL;
ttLibC_close_func ttLibGo_VorbisEncoder_close = NULL;

typedef void (* ttLibC_X264Encoder_getDefaultX264ParamTWithPresetTune_func)(void *, uint32_t, uint32_t, const char *, const char *);
typedef int (* ttLibC_X264Encoder_paramParse_func)(void *, const char *, const char *);
typedef int (* ttLibC_X264Encoder_paramApplyProfile_func)(void *, const char *);
typedef void *(* ttLibC_X264Encoder_makeWithX264ParamT_func)(void *);
typedef bool (* ttLibC_X264Encoder_forceNextFrameType_func)(void *, ttLibC_X264Encoder_FrameType);

ttLibC_X264Encoder_getDefaultX264ParamTWithPresetTune_func ttLibGo_X264Encoder_getDefaultX264ParamTWithPresetTune = NULL;
ttLibC_X264Encoder_paramParse_func                         ttLibGo_X264Encoder_paramParse = NULL;
ttLibC_X264Encoder_paramApplyProfile_func                  ttLibGo_X264Encoder_paramApplyProfile = NULL;
ttLibC_X264Encoder_makeWithX264ParamT_func                 ttLibGo_X264Encoder_makeWithX264ParamT = NULL;
ttLibC_codec_func                                          ttLibGo_X264Encoder_encode = NULL;
ttLibC_X264Encoder_forceNextFrameType_func                 ttLibGo_X264Encoder_forceNextFrameType = NULL;
ttLibC_close_func                                          ttLibGo_X264Encoder_close = NULL;

typedef bool (* ttLibC_X265Encoder_getDefaultX265ApiAndParam_func)(void **, void **, const char *, const char *, uint32_t, uint32_t);
typedef int (* ttLibC_X265Encoder_paramParse_func)(void *param, const char *key, const char *value);
typedef int (* ttLibC_X265Encoder_paramApplyProfile_func)(void *param, const char *profile);
typedef void *(* ttLibC_X265Encoder_makeWithX265ApiAndParam_func)(void *, void *);
typedef bool (* ttLibC_X265Encoder_forceNextFrameType_func)(void *, ttLibC_X265Encoder_FrameType);

ttLibC_X265Encoder_getDefaultX265ApiAndParam_func ttLibGo_X265Encoder_getDefaultX265ApiAndParam = NULL;
ttLibC_X265Encoder_paramParse_func                ttLibGo_X265Encoder_paramParse = NULL;
ttLibC_X265Encoder_paramApplyProfile_func         ttLibGo_X265Encoder_paramApplyProfile = NULL;
ttLibC_X265Encoder_makeWithX265ApiAndParam_func   ttLibGo_X265Encoder_makeWithX265ApiAndParam = NULL;
ttLibC_codec_func                                 ttLibGo_X265Encoder_encode = NULL;
ttLibC_X265Encoder_forceNextFrameType_func        ttLibGo_X265Encoder_forceNextFrameType = NULL;
ttLibC_close_func                                 ttLibGo_X265Encoder_close = NULL;

typedef bool (* ttLibC_Resampler_convert_func)(void *, void *);

typedef void *(* ttLibC_AudioResampler_convertFormat_func)(void *, ttLibC_Frame_Type, uint32_t, uint32_t, void *);

ttLibC_AudioResampler_convertFormat_func ttLibGo_AudioResampler_convertFormat = NULL;

ttLibC_Resampler_convert_func ttLibGo_ImageResampler_ToBgr    = NULL;
ttLibC_Resampler_convert_func ttLibGo_ImageResampler_ToYuv420 = NULL;

typedef void *(* ttLibC_ImageResizer_resize_func)(void *, void *, bool);

ttLibC_ImageResizer_resize_func ttLibGo_ImageResizer_resize = NULL;

typedef bool (* ttLibC_LibyuvResampler_resize_func)(void *, void *, ttLibC_LibyuvFilter_Mode, ttLibC_LibyuvFilter_Mode);
typedef bool (* ttLibC_LibyuvResampler_rotate_func)(void *, void *, ttLibC_LibyuvRotate_Mode);

ttLibC_LibyuvResampler_resize_func ttLibGo_LibyuvResampler_resize   = NULL;
ttLibC_LibyuvResampler_rotate_func ttLibGo_LibyuvResampler_rotate   = NULL;
ttLibC_Resampler_convert_func      ttLibGo_LibyuvResampler_ToBgr    = NULL;
ttLibC_Resampler_convert_func      ttLibGo_LibyuvResampler_ToYuv420 = NULL;

typedef void *(* ttLibC_Soundtouch_make_func)(uint32_t, uint32_t);
typedef void (* ttLibC_soundtouch_set_func)(void *, double);

ttLibC_Soundtouch_make_func ttLibGo_Soundtouch_make = NULL;
ttLibC_codec_func           ttLibGo_Soundtouch_resample = NULL;
ttLibC_close_func           ttLibGo_Soundtouch_close = NULL;
ttLibC_soundtouch_set_func  ttLibGo_Soundtouch_setRate = NULL;
ttLibC_soundtouch_set_func  ttLibGo_Soundtouch_setTempo = NULL;
ttLibC_soundtouch_set_func  ttLibGo_Soundtouch_setRateChange = NULL;
ttLibC_soundtouch_set_func  ttLibGo_Soundtouch_setTempoChange = NULL;
ttLibC_soundtouch_set_func  ttLibGo_Soundtouch_setPitch = NULL;
ttLibC_soundtouch_set_func  ttLibGo_Soundtouch_setPitchOctaves = NULL;
ttLibC_soundtouch_set_func  ttLibGo_Soundtouch_setPitchSemiTones = NULL;

typedef void *(* ttLibC_SpeexdspResampler_make_func)(uint32_t,uint32_t,uint32_t,uint32_t);
typedef void *(* ttLibC_SpeexdspResampler_resample_func)(void *, void *, void *);

ttLibC_SpeexdspResampler_make_func     ttLibGo_SpeexdspResampler_make = NULL;
ttLibC_close_func                      ttLibGo_SpeexdspResampler_close = NULL;
ttLibC_SpeexdspResampler_resample_func ttLibGo_SpeexdspResampler_resample = NULL;

typedef void *(* ttLibC_SwresampleResampler_make_func)(ttLibC_Frame_Type, uint32_t, uint32_t, uint32_t, ttLibC_Frame_Type, uint32_t, uint32_t, uint32_t);

ttLibC_SwresampleResampler_make_func ttLibGo_SwresampleResampler_make = NULL;
ttLibC_codec_func                    ttLibGo_SwresampleResampler_resample = NULL;
ttLibC_close_func                    ttLibGo_SwresampleResampler_close = NULL;

typedef void *(* ttLibC_SwscaleResampler_make_func)(ttLibC_Frame_Type, uint32_t, uint32_t, uint32_t, ttLibC_Frame_Type, uint32_t, uint32_t, uint32_t, ttLibC_SwscaleResampler_Mode);
typedef void *(* ttLibC_SwscaleResampler_resample_func)(void *, void *, void *);

ttLibC_SwscaleResampler_make_func     ttLibGo_SwscaleResampler_make = NULL;
ttLibC_SwscaleResampler_resample_func ttLibGo_SwscaleResampler_resample = NULL;
ttLibC_close_func                     ttLibGo_SwscaleResampler_close = NULL;

typedef void *(* ttLibC_FlvWriter_make_func)(ttLibC_Frame_Type, ttLibC_Frame_Type);
typedef void *(* ttLibC_containerWriter_make_func)(ttLibC_Frame_Type *, uint32_t);
typedef bool (* ttLibC_containerWriter_write_func)(void *, void *, bool(*)(void *, void *, size_t), void *);
typedef bool (* ttLibC_mpegtsWriter_writeInfo_func)(void *, bool(*)(void *, void *,size_t), void *);

ttLibC_FlvWriter_make_func       ttLibGo_FlvWriter_make = NULL;
ttLibC_containerWriter_make_func ttLibGo_Mp4Writer_make = NULL;
ttLibC_containerWriter_make_func ttLibGo_MpegtsWriter_make = NULL;
ttLibC_containerWriter_make_func ttLibGo_MkvWriter_make = NULL;
ttLibC_containerWriter_write_func ttLibGo_FlvWriter_write = NULL;
ttLibC_containerWriter_write_func ttLibGo_MkvWriter_write = NULL;
ttLibC_containerWriter_write_func ttLibGo_Mp4Writer_write = NULL;
ttLibC_containerWriter_write_func ttLibGo_MpegtsWriter_write = NULL;
ttLibC_mpegtsWriter_writeInfo_func ttLibGo_MpegtsWriter_writeInfo = NULL;
ttLibC_close_func ttLibGo_ContainerWriter_close = NULL;

typedef bool (* ttLibC_isType_func)(ttLibC_Frame_Type);
typedef void *(* ttLibC_Frame_clone_func)(void *, void *);
typedef bool (* ttLibC_isFrameType_func)(void *);

ttLibC_isType_func ttLibGo_isAudio = NULL;
ttLibC_isType_func ttLibGo_isVideo = NULL;
ttLibC_Frame_clone_func ttLibGo_Frame_clone = NULL;
ttLibC_close_func       ttLibGo_Frame_close = NULL;
ttLibC_isFrameType_func ttLibGo_Frame_isVideo = NULL;
ttLibC_isFrameType_func ttLibGo_Frame_isAudio = NULL;

typedef bool (* ttLibC_video_getMinimumBinaryBuffer_func)(void *, bool(*)(void *, void *, size_t), void *);

ttLibC_video_getMinimumBinaryBuffer_func ttLibGo_Bgr_getMinimumBinaryBuffer = NULL;
ttLibC_video_getMinimumBinaryBuffer_func ttLibGo_Yuv420_getMinimumBinaryBuffer = NULL;

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

ttLibC_Bgr_makeEmptyFrame_func ttLibGo_Bgr_makeEmptyFrame = NULL;
ttLibC_frame_getFrame_func ttLibGo_Flv1_getFrame = NULL;
ttLibC_frame_getFrame_func ttLibGo_H264_getFrame = NULL;
ttLibC_frame_getFrame_func ttLibGo_H265_getFrame = NULL;
ttLibC_frame_getFrame_func ttLibGo_Jpeg_getFrame = NULL;
ttLibC_frame_getFrame_func ttLibGo_Png_getFrame = NULL;
ttLibC_video_getWidth_func  ttLibGo_Theora_getWidth = NULL;
ttLibC_video_getHeight_func ttLibGo_Theora_getHeight = NULL;
ttLibC_Theora_make_func     ttLibGo_Theora_make = NULL;
ttLibC_video_isKey_func   ttLibGo_Vp6_isKey = NULL;
ttLibC_Vp6_getWidth_func  ttLibGo_Vp6_getWidth = NULL;
ttLibC_Vp6_getHeight_func ttLibGo_Vp6_getHeight = NULL;
ttLibC_video_make_func    ttLibGo_Vp6_make = NULL;
ttLibC_video_isKey_func     ttLibGo_Vp8_isKey = NULL;
ttLibC_video_getWidth_func  ttLibGo_Vp8_getWidth = NULL;
ttLibC_video_getHeight_func ttLibGo_Vp8_getHeight = NULL;
ttLibC_video_make_func      ttLibGo_Vp8_make = NULL;
ttLibC_video_isKey_func     ttLibGo_Vp9_isKey = NULL;
ttLibC_video_getWidth_func  ttLibGo_Vp9_getWidth = NULL;
ttLibC_video_getHeight_func ttLibGo_Vp9_getHeight = NULL;
ttLibC_video_make_func      ttLibGo_Vp9_make = NULL;
ttLibC_Yuv420_makeEmptyFrame_func ttLibGo_Yuv420_makeEmptyFrame = NULL;

typedef void *(* ttLibC_Aac_make_func)(void *,ttLibC_Aac_Type,uint32_t,uint32_t,uint32_t,void*,size_t,bool,uint64_t,uint32_t,uint64_t);
typedef void *(* ttLibC_audio_make_func)(void *,uint32_t,uint32_t,uint32_t,void *,size_t,bool,uint64_t,uint32_t);
typedef void *(* ttLibC_PcmF32_make_func)(void *,ttLibC_PcmF32_Type,uint32_t,uint32_t,uint32_t,void*,size_t,void *,uint32_t,void*,uint32_t,bool,uint64_t,uint32_t);
typedef void *(* ttLibC_PcmS16_make_func)(void *,ttLibC_PcmS16_Type,uint32_t,uint32_t,uint32_t,void*,size_t,void *,uint32_t,void*,uint32_t,bool,uint64_t,uint32_t);
typedef void *(* ttLibC_Speex_make_func)(void *,ttLibC_Speex_Type,uint32_t,uint32_t,uint32_t,void *,size_t,bool,uint64_t,uint32_t);
typedef void *(* ttLibC_Vorbis_make_func)(void *,ttLibC_Vorbis_Type,uint32_t,uint32_t,uint32_t,void *,size_t,bool,uint64_t,uint32_t);

ttLibC_Aac_make_func ttLibGo_Aac_make = NULL;
ttLibC_audio_make_func ttLibGo_AdpcmImaWav_make = NULL;
ttLibC_frame_getFrame_func ttLibGo_Mp3_getFrame = NULL;
ttLibC_audio_make_func ttLibGo_Nellymoser_make = NULL;
ttLibC_frame_getFrame_func ttLibGo_Opus_getFrame = NULL;
ttLibC_audio_make_func ttLibGo_PcmAlaw_make = NULL;
ttLibC_PcmF32_make_func ttLibGo_PcmF32_make = NULL;
ttLibC_audio_make_func ttLibGo_PcmMulaw_make = NULL;
ttLibC_PcmS16_make_func ttLibGo_PcmS16_make = NULL;
ttLibC_Speex_make_func ttLibGo_Speex_make = NULL;
ttLibC_Vorbis_make_func ttLibGo_Vorbis_make = NULL;

typedef void *(* ttLibC_ByteReader_make_func)(void *,size_t,ttLibC_ByteUtil_Type);
typedef uint64_t (* ttLibC_ByteReader_bit_func)(void *,uint32_t);

ttLibC_ByteReader_make_func ttLibGo_ByteReader_make = NULL;
ttLibC_ByteReader_bit_func  ttLibGo_ByteReader_bit = NULL;
ttLibC_close_func           ttLibGo_ByteReader_close = NULL;

static void *lib_handle;

void terminate() {
	if(lib_handle != NULL) {
		ttLibGo_FlvReader_make    = NULL;
		ttLibGo_MkvReader_make    = NULL;
		ttLibGo_Mp4Reader_make    = NULL;
		ttLibGo_MpegtsReader_make = NULL;

		ttLibGo_FlvReader_read    = NULL;
		ttLibGo_MkvReader_read    = NULL;
		ttLibGo_Mp4Reader_read    = NULL;
		ttLibGo_MpegtsReader_read = NULL;

		ttLibGo_Flv_getFrame    = NULL;
		ttLibGo_Mkv_getFrame    = NULL;
		ttLibGo_Mp4_getFrame    = NULL;
		ttLibGo_Mpegts_getFrame = NULL;

		ttLibGo_ContainerReader_close = NULL;

		dlclose(lib_handle);
		lib_handle = NULL;
	}
}

bool setupLibrary(const char *lib_path) {
	atexit(terminate);
	puts(lib_path);
	lib_handle =dlopen(lib_path, RTLD_LAZY);
	if(lib_handle == NULL) {
		return false;
	}
	// containerReader
	ttLibGo_FlvReader_make    = (ttLibC_make_func)dlsym(lib_handle, "ttLibC_FlvReader_make");
	ttLibGo_MkvReader_make    = (ttLibC_make_func)dlsym(lib_handle, "ttLibC_MkvReader_make");
	ttLibGo_Mp4Reader_make    = (ttLibC_make_func)dlsym(lib_handle, "ttLibC_Mp4Reader_make");
	ttLibGo_MpegtsReader_make = (ttLibC_make_func)dlsym(lib_handle, "ttLibC_MpegtsReader_make");
	ttLibGo_FlvReader_read    = (ttLibC_ContainerReader_read_func)dlsym(lib_handle, "ttLibC_FlvReader_read");
	ttLibGo_MkvReader_read    = (ttLibC_ContainerReader_read_func)dlsym(lib_handle, "ttLibC_MkvReader_read");
	ttLibGo_Mp4Reader_read    = (ttLibC_ContainerReader_read_func)dlsym(lib_handle, "ttLibC_Mp4Reader_read");
	ttLibGo_MpegtsReader_read = (ttLibC_ContainerReader_read_func)dlsym(lib_handle, "ttLibC_MpegtsReader_read");
	ttLibGo_Flv_getFrame    = (ttLibC_Container_GetFrame_func)dlsym(lib_handle, "ttLibC_Flv_getFrame");
	ttLibGo_Mkv_getFrame    = (ttLibC_Container_GetFrame_func)dlsym(lib_handle, "ttLibC_Mkv_getFrame");
	ttLibGo_Mp4_getFrame    = (ttLibC_Container_GetFrame_func)dlsym(lib_handle, "ttLibC_Mp4_getFrame");
	ttLibGo_Mpegts_getFrame = (ttLibC_Container_GetFrame_func)dlsym(lib_handle, "ttLibC_Mpegts_getFrame");
	ttLibGo_ContainerReader_close = (ttLibC_close_func)dlsym(lib_handle, "ttLibC_ContainerReader_close");

	ttLibGo_AvcodecVideoDecoder_make = (ttLibC_AvcodecDecoder_make_func)dlsym(lib_handle, "ttLibC_AvcodecVideoDecoder_make");
	ttLibGo_AvcodecAudioDecoder_make = (ttLibC_AvcodecDecoder_make_func)dlsym(lib_handle, "ttLibC_AvcodecAudioDecoder_make");
	ttLibGo_AvcodecDecoder_decode    = (ttLibC_codec_func)dlsym(lib_handle,               "ttLibC_AvcodecDecoder_decode");
	ttLibGo_AvcodecDecoder_close     = (ttLibC_close_func)dlsym(lib_handle,               "ttLibC_AvcodecDecoder_close");

	ttLibGo_JpegDecoder_make   = (ttLibC_make_func)dlsym(lib_handle,  "ttLibC_JpegDecoder_make");
	ttLibGo_JpegDecoder_decode = (ttLibC_codec_func)dlsym(lib_handle, "ttLibC_JpegDecoder_decode");
	ttLibGo_JpegDecoder_close  = (ttLibC_close_func)dlsym(lib_handle, "ttLibC_JpegDecoder_close");

	ttLibGo_Openh264Decoder_make   = (ttLibC_make_func)dlsym(lib_handle,  "ttLibC_Openh264Decoder_make");
	ttLibGo_Openh264Decoder_decode = (ttLibC_codec_func)dlsym(lib_handle, "ttLibC_Openh264Decoder_decode");
	ttLibGo_Openh264Decoder_close  = (ttLibC_close_func)dlsym(lib_handle, "ttLibC_Openh264Decoder_close");

	ttLibGo_OpusDecoder_make         = (ttLibC_OpusDecoder_make_func)dlsym(lib_handle,         "ttLibC_OpusDecoder_make");
	ttLibGo_OpusDecoder_decode       = (ttLibC_codec_func)dlsym(lib_handle,                    "ttLibC_OpusDecoder_decode");
	ttLibGo_OpusDecoder_close        = (ttLibC_close_func)dlsym(lib_handle,                    "ttLibC_OpusDecoder_close");
	ttLibGo_OpusDecoder_codecControl = (ttLibC_opus_codecControl_func)dlsym(lib_handle, "ttLibC_OpusDecoder_codecControl");

	ttLibGo_PngDecoder_make   = (ttLibC_make_func)dlsym(lib_handle,  "ttLibC_PngDecoder_make");
	ttLibGo_PngDecoder_decode = (ttLibC_codec_func)dlsym(lib_handle, "ttLibC_PngDecoder_decode");
	ttLibGo_PngDecoder_close  = (ttLibC_close_func)dlsym(lib_handle, "ttLibC_PngDecoder_close");

	ttLibGo_SpeexDecoder_make         = (ttLibC_SpeexDecoder_make_func)dlsym(lib_handle, "ttLibC_SpeexDecoder_make");
	ttLibGo_SpeexDecoder_decode       = (ttLibC_codec_func)dlsym(lib_handle,             "ttLibC_SpeexDecoder_decode");
	ttLibGo_SpeexDecoder_close        = (ttLibC_close_func)dlsym(lib_handle,             "ttLibC_SpeexDecoder_close");

	ttLibGo_TheoraDecoder_make   = (ttLibC_make_func)dlsym(lib_handle,  "ttLibC_TheoraDecoder_make");
	ttLibGo_TheoraDecoder_decode = (ttLibC_codec_func)dlsym(lib_handle, "ttLibC_TheoraDecoder_decode");
	ttLibGo_TheoraDecoder_close  = (ttLibC_close_func)dlsym(lib_handle, "ttLibC_TheoraDecoder_close");

	ttLibGo_VorbisDecoder_make   = (ttLibC_make_func)dlsym(lib_handle,  "ttLibC_VorbisDecoder_make");
	ttLibGo_VorbisDecoder_decode = (ttLibC_codec_func)dlsym(lib_handle, "ttLibC_VorbisDecoder_decode");
	ttLibGo_VorbisDecoder_close  = (ttLibC_close_func)dlsym(lib_handle, "ttLibC_VorbisDecoder_close");

	ttLibGo_FdkaacEncoder_make        = (ttLibC_FdkaacEncoder_make_func)dlsym(lib_handle,       "ttLibC_FdkaacEncoder_make");
	ttLibGo_FdkaacEncoder_encode      = (ttLibC_codec_func)dlsym(lib_handle,                    "ttLibC_FdkaacEncoder_encode");
	ttLibGo_FdkaacEncoder_close       = (ttLibC_close_func)dlsym(lib_handle,                    "ttLibC_FdkaacEncoder_close");
	ttLibGo_FdkaacEncoder_setBitrate  = (ttLibC_FdkaacEncoder_setBitrate_func)dlsym(lib_handle, "ttLibC_FdkaacEncoder_setBitrate");

	ttLibGo_JpegEncoder_make       = (ttLibC_JpegEncoder_make_func)dlsym(lib_handle,       "ttLibC_JpegEncoder_make");
	ttLibGo_JpegEncoder_encode     = (ttLibC_codec_func)dlsym(lib_handle,                  "ttLibC_JpegEncoder_encode");
	ttLibGo_JpegEncoder_setQuality = (ttLibC_JpegEncoder_setQuality_func)dlsym(lib_handle, "ttLibC_JpegEncoder_setQuality");
	ttLibGo_JpegEncoder_close      = (ttLibC_close_func)dlsym(lib_handle,                  "ttLibC_JpegEncoder_close");

	ttLibGo_Openh264Encoder_getDefaultSEncParamExt = (ttLibC_Openh264Encoder_getDefaultSEncParamExt_func)dlsym(lib_handle, "ttLibC_Openh264Encoder_getDefaultSEncParamExt");
	ttLibGo_Openh264Encoder_paramParse             = (ttLibC_Openh264Encoder_paramParse_func)dlsym(lib_handle,             "ttLibC_Openh264Encoder_paramParse");
	ttLibGo_Openh264Encoder_spatialParamParse      = (ttLibC_Openh264Encoder_spatialParamParse_func)dlsym(lib_handle,      "ttLibC_Openh264Encoder_spatialParamParse");
	ttLibGo_Openh264Encoder_makeWithSEncParamExt   = (ttLibC_Openh264Encoder_makeWithSEncParamExt_func)dlsym(lib_handle,   "ttLibC_Openh264Encoder_makeWithSEncParamExt");
	ttLibGo_Openh264Encoder_encode                 = (ttLibC_codec_func)dlsym(lib_handle,                                  "ttLibC_Openh264Encoder_encode");
	ttLibGo_Openh264Encoder_close                  = (ttLibC_close_func)dlsym(lib_handle,                                  "ttLibC_Openh264Encoder_close");
	ttLibGo_Openh264Encoder_setRCMode              = (ttLibC_Openh264Encoder_setRCMode_func)dlsym(lib_handle,              "ttLibC_Openh264Encoder_setRCMode");
	ttLibGo_Openh264Encoder_setIDRInterval         = (ttLibC_Openh264Encoder_setIDRInterval_func)dlsym(lib_handle,         "ttLibC_Openh264Encoder_setIDRInterval");
	ttLibGo_Openh264Encoder_forceNextKeyFrame      = (ttLibC_Openh264Encoder_forceNextKeyFrame_func)dlsym(lib_handle,      "ttLibC_Openh264Encoder_forceNextKeyFrame");

	ttLibGo_OpusEncoder_make         = (ttLibC_OpusEncoder_make_func)dlsym(lib_handle,  "ttLibC_OpusEncoder_make");
	ttLibGo_OpusEncoder_encode       = (ttLibC_codec_func)dlsym(lib_handle,             "ttLibC_OpusEncoder_encode");
	ttLibGo_OpusEncoder_codecControl = (ttLibC_opus_codecControl_func)dlsym(lib_handle, "ttLibC_OpusEncoder_codecControl");
	ttLibGo_OpusEncoder_close        = (ttLibC_close_func)dlsym(lib_handle,             "ttLibC_OpusEncoder_close");

	ttLibGo_SpeexEncoder_make   = (ttLibC_SpeexEncoder_make_func)dlsym(lib_handle, "ttLibC_SpeexEncoder_make");
	ttLibGo_SpeexEncoder_encode = (ttLibC_codec_func)dlsym(lib_handle,             "ttLibC_SpeexEncoder_encode");
	ttLibGo_SpeexEncoder_close  = (ttLibC_close_func)dlsym(lib_handle,             "ttLibC_SpeexEncoder_close");

	ttLibGo_TheoraEncoder_make_ex = (ttLibC_TheoraEncoder_make_ex_func)dlsym(lib_handle, "ttLibC_TheoraEncoder_make_ex");
	ttLibGo_TheoraEncoder_encode  = (ttLibC_codec_func)dlsym(lib_handle,                 "ttLibC_TheoraEncoder_encode");
	ttLibGo_TheoraEncoder_close   = (ttLibC_close_func)dlsym(lib_handle,                 "ttLibC_TheoraEncoder_close");

	ttLibGo_VorbisEncoder_make   = (ttLibC_VorbisEncoder_make_func)dlsym(lib_handle, "ttLibC_VorbisEncoder_make");
	ttLibGo_VorbisEncoder_encode = (ttLibC_codec_func)dlsym(lib_handle,              "ttLibC_VorbisEncoder_encode");
	ttLibGo_VorbisEncoder_close  = (ttLibC_close_func)dlsym(lib_handle,              "ttLibC_VorbisEncoder_close");

	ttLibGo_X264Encoder_getDefaultX264ParamTWithPresetTune = (ttLibC_X264Encoder_getDefaultX264ParamTWithPresetTune_func)dlsym(lib_handle, "ttLibC_X264Encoder_getDefaultX264ParamTWithPresetTune");
	ttLibGo_X264Encoder_paramParse                         = (ttLibC_X264Encoder_paramParse_func)dlsym(lib_handle,                         "ttLibC_X264Encoder_paramParse");
	ttLibGo_X264Encoder_paramApplyProfile                  = (ttLibC_X264Encoder_paramApplyProfile_func)dlsym(lib_handle,                  "ttLibC_X264Encoder_paramApplyProfile");
	ttLibGo_X264Encoder_makeWithX264ParamT                 = (ttLibC_X264Encoder_makeWithX264ParamT_func)dlsym(lib_handle,                 "ttLibC_X264Encoder_makeWithX264ParamT");
	ttLibGo_X264Encoder_encode                             = (ttLibC_codec_func)dlsym(lib_handle,                                          "ttLibC_X264Encoder_encode");
	ttLibGo_X264Encoder_forceNextFrameType                 = (ttLibC_X264Encoder_forceNextFrameType_func)dlsym(lib_handle,                 "ttLibC_X264Encoder_forceNextFrameType");
	ttLibGo_X264Encoder_close                              = (ttLibC_close_func)dlsym(lib_handle,                                          "ttLibC_X264Encoder_close");

	ttLibGo_X265Encoder_getDefaultX265ApiAndParam = (ttLibC_X265Encoder_getDefaultX265ApiAndParam_func)dlsym(lib_handle, "ttLibC_X265Encoder_getDefaultX265ApiAndParam");
	ttLibGo_X265Encoder_paramParse                = (ttLibC_X265Encoder_paramParse_func)dlsym(lib_handle,                "ttLibC_X265Encoder_paramParse");
	ttLibGo_X265Encoder_paramApplyProfile         = (ttLibC_X265Encoder_paramApplyProfile_func)dlsym(lib_handle,         "ttLibC_X265Encoder_paramApplyProfile");
	ttLibGo_X265Encoder_makeWithX265ApiAndParam   = (ttLibC_X265Encoder_makeWithX265ApiAndParam_func)dlsym(lib_handle,   "ttLibC_X265Encoder_makeWithX265ApiAndParam");
	ttLibGo_X265Encoder_encode                    = (ttLibC_codec_func)dlsym(lib_handle,                                 "ttLibC_X265Encoder_encode");
	ttLibGo_X265Encoder_forceNextFrameType        = (ttLibC_X265Encoder_forceNextFrameType_func)dlsym(lib_handle,        "ttLibC_X265Encoder_forceNextFrameType_func");
	ttLibGo_X265Encoder_close                     = (ttLibC_close_func)dlsym(lib_handle,                                 "ttLibC_X265Encoder_close");

	ttLibGo_AudioResampler_convertFormat = (ttLibC_AudioResampler_convertFormat_func)dlsym(lib_handle, "ttLibC_AudioResampler_convertFormat");

	ttLibGo_ImageResampler_ToBgr    = (ttLibC_Resampler_convert_func)dlsym(lib_handle, "ttLibC_ImageResampler_ToBgr");
	ttLibGo_ImageResampler_ToYuv420 = (ttLibC_Resampler_convert_func)dlsym(lib_handle, "ttLibC_ImageResampler_ToYuv420");

	ttLibGo_ImageResizer_resize    = (ttLibC_ImageResizer_resize_func)dlsym(lib_handle,    "ttLibC_ImageResizer_resize");

	ttLibGo_LibyuvResampler_resize   = (ttLibC_LibyuvResampler_resize_func)dlsym(lib_handle, "ttLibC_LibyuvResampler_resize");
	ttLibGo_LibyuvResampler_rotate   = (ttLibC_LibyuvResampler_rotate_func)dlsym(lib_handle, "ttLibC_LibyuvResampler_rotate");
	ttLibGo_LibyuvResampler_ToBgr    = (ttLibC_Resampler_convert_func)dlsym(lib_handle,      "ttLibC_LibyuvResampler_ToBgr");
	ttLibGo_LibyuvResampler_ToYuv420 = (ttLibC_Resampler_convert_func)dlsym(lib_handle,      "ttLibC_LibyuvResampler_ToYuv420");

	ttLibGo_Soundtouch_make              = (ttLibC_Soundtouch_make_func)dlsym(lib_handle, "ttLibC_Soundtouch_make");
	ttLibGo_Soundtouch_resample          = (ttLibC_codec_func)dlsym(lib_handle,           "ttLibC_Soundtouch_resample");
	ttLibGo_Soundtouch_close             = (ttLibC_close_func)dlsym(lib_handle,           "ttLibC_Soundtouch_close");
	ttLibGo_Soundtouch_setRate           = (ttLibC_soundtouch_set_func)dlsym(lib_handle,  "ttLibC_Soundtouch_setRate");
	ttLibGo_Soundtouch_setTempo          = (ttLibC_soundtouch_set_func)dlsym(lib_handle,  "ttLibC_Soundtouch_setTempo");
	ttLibGo_Soundtouch_setRateChange     = (ttLibC_soundtouch_set_func)dlsym(lib_handle,  "ttLibC_Soundtouch_setRateChange");
	ttLibGo_Soundtouch_setTempoChange    = (ttLibC_soundtouch_set_func)dlsym(lib_handle,  "ttLibC_Soundtouch_setTempoChange");
	ttLibGo_Soundtouch_setPitch          = (ttLibC_soundtouch_set_func)dlsym(lib_handle,  "ttLibC_Soundtouch_setPitch");
	ttLibGo_Soundtouch_setPitchOctaves   = (ttLibC_soundtouch_set_func)dlsym(lib_handle,  "ttLibC_Soundtouch_setPitchOctaves");
	ttLibGo_Soundtouch_setPitchSemiTones = (ttLibC_soundtouch_set_func)dlsym(lib_handle,  "ttLibC_Soundtouch_setPitchSemiTones");

	ttLibGo_SpeexdspResampler_make     = (ttLibC_SpeexdspResampler_make_func)dlsym(lib_handle,     "ttLibC_SpeexdspResampler_make");
	ttLibGo_SpeexdspResampler_close    = (ttLibC_close_func)dlsym(lib_handle,                      "ttLibC_SpeexdspResampler_close");
	ttLibGo_SpeexdspResampler_resample = (ttLibC_SpeexdspResampler_resample_func)dlsym(lib_handle, "ttLibC_SpeexdspResampler_resample");

	ttLibGo_SwresampleResampler_make     = (ttLibC_SwresampleResampler_make_func)dlsym(lib_handle, "ttLibC_SwresampleResampler_make");
	ttLibGo_SwresampleResampler_resample = (ttLibC_codec_func)dlsym(lib_handle,                    "ttLibC_SwresampleResampler_resample");
	ttLibGo_SwresampleResampler_close    = (ttLibC_close_func)dlsym(lib_handle,                    "ttLibC_SwresampleResampler_close");

	ttLibGo_SwscaleResampler_make     = (ttLibC_SwscaleResampler_make_func)dlsym(lib_handle,     "ttLibC_SwscaleResampler_make");
	ttLibGo_SwscaleResampler_resample = (ttLibC_SwscaleResampler_resample_func)dlsym(lib_handle, "ttLibC_SwscaleResampler_resample");
	ttLibGo_SwscaleResampler_close    = (ttLibC_close_func)dlsym(lib_handle,                     "ttLibC_SwscaleResampler_close");

	ttLibGo_FlvWriter_make    = (ttLibC_FlvWriter_make_func)dlsym(lib_handle,       "ttLibC_FlvWriter_make");
	ttLibGo_Mp4Writer_make    = (ttLibC_containerWriter_make_func)dlsym(lib_handle, "ttLibC_Mp4Writer_make");
	ttLibGo_MpegtsWriter_make = (ttLibC_containerWriter_make_func)dlsym(lib_handle, "ttLibC_MpegtsWriter_make");
	ttLibGo_MkvWriter_make    = (ttLibC_containerWriter_make_func)dlsym(lib_handle, "ttLibC_MkvWriter_make");
	ttLibGo_FlvWriter_write    = (ttLibC_containerWriter_write_func)dlsym(lib_handle, "ttLibC_FlvWriter_write");
	ttLibGo_MkvWriter_write    = (ttLibC_containerWriter_write_func)dlsym(lib_handle, "ttLibC_MkvWriter_write");
	ttLibGo_Mp4Writer_write    = (ttLibC_containerWriter_write_func)dlsym(lib_handle, "ttLibC_Mp4Writer_write");
	ttLibGo_MpegtsWriter_write = (ttLibC_containerWriter_write_func)dlsym(lib_handle, "ttLibC_MpegtsWriter_write");
	ttLibGo_MpegtsWriter_writeInfo = (ttLibC_mpegtsWriter_writeInfo_func)dlsym(lib_handle, "ttLibC_MpegtsWriter_writeInfo");
	ttLibGo_ContainerWriter_close = (ttLibC_close_func)dlsym(lib_handle, "ttLibC_ContainerWriter_close");

	ttLibGo_isAudio = (ttLibC_isType_func)dlsym(lib_handle, "ttLibC_isAudio");
	ttLibGo_isVideo = (ttLibC_isType_func)dlsym(lib_handle, "ttLibC_isVideo");
	ttLibGo_Frame_clone = (ttLibC_Frame_clone_func)dlsym(lib_handle, "ttLibC_Frame_clone");
	ttLibGo_Frame_close = (ttLibC_close_func)dlsym(lib_handle,       "ttLibC_Frame_close");
	ttLibGo_Frame_isVideo = (ttLibC_isFrameType_func)dlsym(lib_handle, "ttLibC_Frame_isVideo");
	ttLibGo_Frame_isAudio = (ttLibC_isFrameType_func)dlsym(lib_handle, "ttLibC_Frame_isAudio");
	ttLibGo_Bgr_getMinimumBinaryBuffer    = (ttLibC_video_getMinimumBinaryBuffer_func)dlsym(lib_handle, "ttLibC_Bgr_getMinimumBinaryBuffer");
	ttLibGo_Yuv420_getMinimumBinaryBuffer = (ttLibC_video_getMinimumBinaryBuffer_func)dlsym(lib_handle, "ttLibC_Yuv420_getMinimumBinaryBuffer");

	ttLibGo_Bgr_makeEmptyFrame = (ttLibC_Bgr_makeEmptyFrame_func)dlsym(lib_handle, "ttLibC_Bgr_makeEmptyFrame");
	ttLibGo_Flv1_getFrame = (ttLibC_frame_getFrame_func)dlsym(lib_handle, "ttLibC_Flv1_getFrame");
	ttLibGo_H264_getFrame = (ttLibC_frame_getFrame_func)dlsym(lib_handle, "ttLibC_H264_getFrame");
	ttLibGo_H265_getFrame = (ttLibC_frame_getFrame_func)dlsym(lib_handle, "ttLibC_H265_getFrame");
	ttLibGo_Jpeg_getFrame = (ttLibC_frame_getFrame_func)dlsym(lib_handle, "ttLibC_Jpeg_getFrame");
	ttLibGo_Png_getFrame  = (ttLibC_frame_getFrame_func)dlsym(lib_handle, "ttLibC_Png_getFrame");
	ttLibGo_Theora_getWidth  = (ttLibC_video_getWidth_func)dlsym(lib_handle,  "ttLibC_Theora_getWidth");
	ttLibGo_Theora_getHeight = (ttLibC_video_getHeight_func)dlsym(lib_handle, "ttLibC_Theora_getHeight");
	ttLibGo_Theora_make      = (ttLibC_Theora_make_func)dlsym(lib_handle,    "ttLibC_Theora_make");
	ttLibGo_Vp6_isKey     = (ttLibC_video_isKey_func)dlsym(lib_handle,   "ttLibC_Vp6_isKey");
	ttLibGo_Vp6_getWidth  = (ttLibC_Vp6_getWidth_func)dlsym(lib_handle,  "ttLibC_Vp6_getWidth");
	ttLibGo_Vp6_getHeight = (ttLibC_Vp6_getHeight_func)dlsym(lib_handle, "ttLibC_Vp6_getHeight");
	ttLibGo_Vp6_make      = (ttLibC_video_make_func)dlsym(lib_handle,    "ttLibC_Vp6_make");
	ttLibGo_Vp8_isKey     = (ttLibC_video_isKey_func)dlsym(lib_handle,     "ttLibC_Vp8_isKey");
	ttLibGo_Vp8_getWidth  = (ttLibC_video_getWidth_func)dlsym(lib_handle,  "ttLibC_Vp8_getWidth");
	ttLibGo_Vp8_getHeight = (ttLibC_video_getHeight_func)dlsym(lib_handle, "ttLibC_Vp8_getHeight");
	ttLibGo_Vp8_make      = (ttLibC_video_make_func)dlsym(lib_handle,      "ttLibC_Vp8_make");
	ttLibGo_Vp9_isKey     = (ttLibC_video_isKey_func)dlsym(lib_handle,     "ttLibC_Vp9_isKey");
	ttLibGo_Vp9_getWidth  = (ttLibC_video_getWidth_func)dlsym(lib_handle,  "ttLibC_Vp9_getWidth");
	ttLibGo_Vp9_getHeight = (ttLibC_video_getHeight_func)dlsym(lib_handle, "ttLibC_Vp9_getHeight");
	ttLibGo_Vp9_make      = (ttLibC_video_make_func)dlsym(lib_handle,      "ttLibC_Vp9_make");
	ttLibGo_Yuv420_makeEmptyFrame = (ttLibC_Yuv420_makeEmptyFrame_func)dlsym(lib_handle, "ttLibC_Yuv420_makeEmptyFrame");

	ttLibGo_Aac_make = (ttLibC_Aac_make_func)dlsym(lib_handle, "ttLibC_Aac_make");
	ttLibGo_AdpcmImaWav_make = (ttLibC_audio_make_func)dlsym(lib_handle, "ttLibC_AdpcmImaWav_make");
	ttLibGo_Mp3_getFrame = (ttLibC_frame_getFrame_func)dlsym(lib_handle, "ttLibC_Mp3_getFrame");
	ttLibGo_Nellymoser_make = (ttLibC_audio_make_func)dlsym(lib_handle, "ttLibC_Nellymoser_make");
	ttLibGo_Opus_getFrame = (ttLibC_frame_getFrame_func)dlsym(lib_handle, "ttLibC_Opus_getFrame");
	ttLibGo_PcmAlaw_make = (ttLibC_audio_make_func)dlsym(lib_handle, "ttLibC_PcmAlaw_make");
	ttLibGo_PcmF32_make = (ttLibC_PcmF32_make_func)dlsym(lib_handle, "ttLibC_PcmF32_make");
	ttLibGo_PcmMulaw_make = (ttLibC_audio_make_func)dlsym(lib_handle, "ttLibC_PcmMulaw_make");
	ttLibGo_PcmS16_make = (ttLibC_PcmS16_make_func)dlsym(lib_handle, "ttLibC_PcmS16_make");
	ttLibGo_Speex_make = (ttLibC_Speex_make_func)dlsym(lib_handle, "ttLibC_Speex_make");
	ttLibGo_Vorbis_make = (ttLibC_Vorbis_make_func)dlsym(lib_handle, "ttLibC_Vorbis_make");

	ttLibGo_ByteReader_make  = (ttLibC_ByteReader_make_func)dlsym(lib_handle, "ttLibC_ByteReader_make");
	ttLibGo_ByteReader_bit   = (ttLibC_ByteReader_bit_func)dlsym(lib_handle,  "ttLibC_ByteReader_bit");
	ttLibGo_ByteReader_close = (ttLibC_close_func)dlsym(lib_handle,           "ttLibC_ByteReader_close");

	return true;
}

*/
import "C"
import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"unsafe"
)

func init() {
	// 拡張子
	var ext string
	gopath := os.Getenv("GOPATH")
	switch runtime.GOOS {
	case "darwin":
		ext = "dylib"
		if gopath == "" {
			gopath = fmt.Sprintf("%s/go", os.Getenv("HOME"))
		}
	case "linux":
		ext = "so"
		if gopath == "" {
			gopath = fmt.Sprintf("%s/go", os.Getenv("HOME"))
		}
	default:
		panic("run only in linux or macOS")
	}
	// 現在のパスをなんとかする
	// 毎回ライブラリを使うときに、これを走らせておく。.
	_, err := exec.LookPath("cmake")
	if err != nil {
		panic(err) // cmakeが見つからない場合は動作しない。
	}
	pwd, _ := filepath.Abs(".") // 現在の位置を保存後で、戻す
	buildPath := fmt.Sprintf("%s/src/github.com/taktod/ttLibGo/build", gopath)
	// 移動
	err = os.Chdir(buildPath)
	if err != nil {
		panic(err)
	}
	// cmake
	cmd := exec.Command("cmake", "../ttLibC")
	cmd.Start()
	if err := cmd.Wait(); err != nil {
		panic(err)
	}
	// make
	cmd = exec.Command("make")
	cmd.Start()
	if err := cmd.Wait(); err != nil {
		panic(err)
	}
	os.Chdir(pwd)

	// これでライブラリbuildはいけた。
	// 次はこのライブラリを使ったプログラムに変えていく必要があるわけだが・・・
	targetLibrary := fmt.Sprintf("%s/libttLibC.%s", buildPath, ext)

	// ライブラリをc言語的に読み込んでおく
	//	fmt.Println(targetLibrary)
	cTargetLibrary := C.CString(targetLibrary)
	defer C.free(unsafe.Pointer(cTargetLibrary))
	result := bool(C.setupLibrary(cTargetLibrary))
	if !result {
		panic("failed to load library.")
	}
}

// Close ttLibGoのライブラリをunloadする
func Close() {
	C.terminate()
}
