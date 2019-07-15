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
    return new AvcodecDecoder(mp);
  }
  else if(decoderType == "jpeg") {
    return new JpegDecoder(mp);
  }
  else if(decoderType == "openh264") {
    return new Openh264Decoder(mp);
  }
  else if(decoderType == "opus") {
    return new tOpusDecoder(mp);
  }
  else if(decoderType == "png") {
    return new PngDecoder(mp);
  }
  else if(decoderType == "speex") {
    return new SpeexDecoder(mp);
  }
  else if(decoderType == "theora") {
    return new TheoraDecoder(mp);
  }
  else if(decoderType == "vorbis") {
    return new VorbisDecoder(mp);
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
