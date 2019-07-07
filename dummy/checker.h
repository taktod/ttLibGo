#ifdef __cplusplus
extern "C" {
#endif

// ここでは、dummyのheaderをincludeできた場合は、該当ライブラリが存在しないということにしていく。

#include <libavcodec/avcodec.h>
#ifndef NO_AVCODEC
# define __ENABLE_AVCODEC__
#endif
#include <libswscale/swscale.h>
#ifndef NO_SWSCALE
# define __ENABLE_SWSCALE__
#endif
#include <libswresample/swresample.h>
#ifndef NO_SWRESAMPLE
# define __ENABLE_SWRESAMPLE__
#endif
#include <jpeglib.h>
#ifndef NO_JPEG
# define __ENABLE_JPEG__
#endif
#include <wels/codec_api.h>
#ifndef NO_OPENH264
# define __ENABLE_OPENH264__
#endif
#include <opus/opus.h>
#ifndef NO_OPUS
# define __ENABLE_OPUS__
#endif
#include <png.h>
#ifndef NO_LIBPNG
# define __ENABLE_LIBPNG__
#endif
#include <speex/speex.h>
#ifndef NO_SPEEX
# define __ENABLE_SPEEX__
#endif
#include <theora/theoradec.h>
#ifndef NO_THEORA
# define __ENABLE_THEORA__
#endif
#include <vorbis/codec.h>
#ifndef NO_VORBIS
# define __ENABLE_VORBIS_DECODE__
#endif
#include <fdk-aac/aacenc_lib.h>
#ifndef NO_FDKAAC
# define __ENABLE_FDKAAC_ENCODE__
#endif
#include <vorbis/vorbisenc.h>
#ifndef NO_VORBISENC
# define __ENABLE_VORBIS_ENCODE__
#endif
#ifdef __cplusplus
}
# include <soundtouch/SoundTouch.h>
# ifndef NO_SOUNDTOUCH
#  define __ENABLE_SOUNDTOUCH__
# endif
extern "C" {
#endif
#include <speex/speex_resampler.h>
#ifndef NO_SPEEXDSP
# define __ENABLE_SPEEXDSP__
#endif

#ifdef __cplusplus
}
#endif

