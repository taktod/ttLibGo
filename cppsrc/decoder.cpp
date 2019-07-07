#include <ttLibC/allocator.h>
#include <ttLibC/frame/frame.h>
#include <ttLibC/frame/video/video.h>
#include <ttLibC/frame/video/bgr.h>
#include <ttLibC/frame/video/flv1.h>
#include <ttLibC/frame/video/h264.h>
#include <ttLibC/frame/video/h265.h>
#include <ttLibC/frame/video/theora.h>
#include <ttLibC/frame/video/yuv420.h>
#include <ttLibC/frame/audio/audio.h>
#include <ttLibC/frame/audio/aac.h>
#include <ttLibC/frame/audio/mp3.h>
#include <ttLibC/frame/audio/opus.h>
#include <ttLibC/frame/audio/pcmf32.h>
#include <ttLibC/frame/audio/pcms16.h>
#include <ttLibC/frame/audio/speex.h>
#include <ttLibC/frame/audio/vorbis.h>
#include <string>

#include "decoder.hpp"
#include "decoder/avcodec.hpp"
#include "decoder/jpeg.hpp"
#include "decoder/openh264.hpp"
#include "decoder/opus.hpp"
#include "decoder/png.hpp"
#include "decoder/speex.hpp"
#include "decoder/theora.hpp"
#include "decoder/vorbis.hpp"

using namespace std;

Decoder *Decoder::create(maps *mp) {
  string decoderType = mp->getString("decoder");
  if(decoderType == "avcodecVideo" || decoderType == "avcodecAudio") {
#ifdef __ENABLE_AVCODEC__
    return new AvcodecDecoder(mp);
#endif
  }
  else if(decoderType == "jpeg") {
#ifdef __ENABLE_JPEG__
    return new JpegDecoder(mp);
#endif
  }
  else if(decoderType == "openh264") {
#ifdef __ENABLE_OPENH264__
    return new Openh264Decoder(mp);
#endif
  }
  else if(decoderType == "opus") {
#ifdef __ENABLE_OPUS__
    return new tOpusDecoder(mp);
#endif
  }
  else if(decoderType == "png") {
#ifdef __ENABLE_LIBPNG__
    return new PngDecoder(mp);
#endif
  }
  else if(decoderType == "speex") {
#ifdef __ENABLE_SPEEX__
    return new SpeexDecoder(mp);
#endif
  }
  else if(decoderType == "theora") {
#ifdef __ENABLE_THEORA__
    return new TheoraDecoder(mp);
#endif
  }
  else if(decoderType == "vorbis") {
#ifdef __ENABLE_VORBIS_DECODE__
    return new VorbisDecoder(mp);
#endif
  }
  return nullptr;
}
Decoder::Decoder(){
}
Decoder::~Decoder() {
}

extern "C" {

void *Decoder_make(void *mp) {
  maps *m = reinterpret_cast<maps *>(mp);
  return Decoder::create(m);
}
bool Decoder_decodeFrame(void *decoder, void *frame, void *goFrame, uintptr_t ptr) {
  if(decoder == nullptr) {
    return false;
  }
  Decoder *d = reinterpret_cast<Decoder *>(decoder);
  ttLibC_Frame *cF = reinterpret_cast<ttLibC_Frame *>(frame);
  ttLibGoFrame *goF = reinterpret_cast<ttLibGoFrame *>(goFrame);
  return d->decodeFrame(cF, goF, (void *)ptr);
}

void Decoder_close(void *decoder) {
  if(decoder == nullptr) {
    return;
  }
  Decoder *d = reinterpret_cast<Decoder *>(decoder);
  delete d;
}

}
