#include <stdio.h>
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

#include "encoder.hpp"
#include "encoder/fdkaac.hpp"
#include "encoder/jpeg.hpp"
#include "encoder/openh264.hpp"
#include "encoder/opus.hpp"
#include "encoder/speex.hpp"
#include "encoder/theora.hpp"
#include "encoder/vorbis.hpp"
#include "encoder/x264.hpp"
#include "encoder/x265.hpp"

#include <iostream>

using namespace std;

Encoder *Encoder::create(maps *mp) {
  string encoderType = mp->getString("encoder");
  if(encoderType == "fdkaac") {
    return new FdkaacEncoder(mp);
  }
  else if(encoderType == "jpeg") {
    return new JpegEncoder(mp);
  }
  else if(encoderType == "openh264") {
    return new Openh264Encoder(mp);
  }
  else if(encoderType == "opus") {
    return new tOpusEncoder(mp);
  }
  else if(encoderType == "speex") {
    return new SpeexEncoder(mp);
  }
  else if(encoderType == "theora") {
    return new TheoraEncoder(mp);
  }
  else if(encoderType == "vorbis") {
    return new VorbisEncoder(mp);
  }
  else if(encoderType == "x264") {
    return new X264Encoder(mp);
  }
  else if(encoderType == "x265") {
    return new X265Encoder(mp);
  }
  return nullptr;
}
Encoder::Encoder() {
}
Encoder::~Encoder() {
}

extern "C" {
void *Encoder_make(void *mp) {
  maps *m = reinterpret_cast<maps *>(mp);
  return Encoder::create(m);
}
bool Encoder_encodeFrame(void *encoder, void *frame, void *goFrame, uintptr_t ptr) {
  if(encoder == nullptr) {
    return false;
  }
  Encoder *e = reinterpret_cast<Encoder *>(encoder);
  ttLibC_Frame *cF = reinterpret_cast<ttLibC_Frame *>(frame);
  ttLibGoFrame *goF = reinterpret_cast<ttLibGoFrame *>(goFrame);
  return e->encodeFrame(cF, goF, (void *)ptr);
}
void Encoder_close(void *encoder) {
  if(encoder == nullptr) {
    return;
  }
  Encoder *e = reinterpret_cast<Encoder *>(encoder);
  delete e;
}

}