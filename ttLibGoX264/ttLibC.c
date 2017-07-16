#define TT_VISIBILITY_HIDDEN __attribute__ ((visibility("hidden")))
#define TT_VISIBILITY_DEFAULT __attribute__ ((visibility("hidden")))

#include "../ttLibC/ttLibC/encoder/x264Encoder.c"

#include "../ttLibGo/frame.c"
#include "../ttLibC/ttLibC/frame/audio/aac.c"
#include "../ttLibC/ttLibC/frame/audio/adpcmImaWav.c"
#include "../ttLibC/ttLibC/frame/audio/audio.c"
#include "../ttLibC/ttLibC/frame/audio/mp3.c"
#include "../ttLibC/ttLibC/frame/audio/nellymoser.c"
#include "../ttLibC/ttLibC/frame/audio/opus.c"
#include "../ttLibC/ttLibC/frame/audio/pcmAlaw.c"
#include "../ttLibC/ttLibC/frame/audio/pcmf32.c"
#include "../ttLibC/ttLibC/frame/audio/pcmMulaw.c"
#include "../ttLibC/ttLibC/frame/audio/pcms16.c"
#include "../ttLibC/ttLibC/frame/audio/speex.c"
#include "../ttLibC/ttLibC/frame/audio/vorbis.c"
#include "../ttLibC/ttLibC/frame/video/bgr.c"
#include "../ttLibC/ttLibC/frame/video/flv1.c"
#include "../ttLibC/ttLibC/frame/video/h264.c"
#include "../ttLibC/ttLibC/frame/video/h265.c"
#include "../ttLibC/ttLibC/frame/video/jpeg.c"
#include "../ttLibC/ttLibC/frame/video/theora.c"
#include "../ttLibC/ttLibC/frame/video/video.c"
#include "../ttLibC/ttLibC/frame/video/vp6.c"
#include "../ttLibC/ttLibC/frame/video/vp8.c"
#include "../ttLibC/ttLibC/frame/video/vp9.c"
#include "../ttLibC/ttLibC/frame/video/wmv1.c"
#include "../ttLibC/ttLibC/frame/video/wmv2.c"
#include "../ttLibC/ttLibC/frame/video/yuv420.c"
#include "../ttLibC/ttLibC/frame/frame.c"
#include "../ttLibC/ttLibC/util/byteUtil.c"
#include "../ttLibC/ttLibC/util/crc32Util.c"
#include "../ttLibC/ttLibC/util/dynamicBufferUtil.c"
#include "../ttLibC/ttLibC/util/hexUtil.c"
#include "../ttLibC/ttLibC/util/ioUtil.c"
#include "../ttLibC/ttLibC/allocator.c"
#include "../ttLibC/ttLibC/ttLibC.c"
