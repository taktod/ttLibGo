#include <string>
#include "util.hpp"
#include <ttLibC/container/container.h>
#include <ttLibC/container/flv.h>
#include <ttLibC/container/mp4.h>
#include <ttLibC/container/mkv.h>
#include <ttLibC/container/mpegts.h>
#include <ttLibC/container/containerCommon.h>
#include "frame.hpp"

class Writer : public FrameProcessor {
public:
  Writer(maps *m) {
    string type = m->getString("type");
    if(type == "flv") {
      _writer = (ttLibC_ContainerWriter *)ttLibC_FlvWriter_make(
        Frame_getFrameTypeFromString(m->getString("videoType")),
        Frame_getFrameTypeFromString(m->getString("audioType")));
      return;
    }
    list<string> frameList = m->getStringList("frameTypes");
    ttLibC_Frame_Type *types = new ttLibC_Frame_Type[frameList.size()];
    int i = 0;
    for(auto iter = frameList.cbegin(); iter != frameList.cend(); ++ iter) {
      types[i] = Frame_getFrameTypeFromString(*iter);
      i ++;
    }
    if(type == "mp4") {
      _writer = (ttLibC_ContainerWriter *)ttLibC_Mp4Writer_make(types, frameList.size());
    }
    else if(type == "mpegts") {
      _writer = (ttLibC_ContainerWriter *)ttLibC_MpegtsWriter_make(types, frameList.size());
    }
    else if(type == "mkv") {
      _writer = (ttLibC_ContainerWriter *)ttLibC_MkvWriter_make(types, frameList.size());
    }
    else if(type == "webm") {
      _writer = (ttLibC_ContainerWriter *)ttLibC_MkvWriter_make(types, frameList.size());
		  _writer->type = containerType_webm;
    }
    delete[] types;
  }
  ~Writer() {
    ttLibC_ContainerWriter_close(&_writer);
  }
  bool writeFrame(
    ttLibC_Frame *cFrame,
    ttLibGoFrame *goFrame,
    uint32_t mode,
    uint32_t unitDuration,
    void *ptr) {
    if(_writer == nullptr) {
      return false;
    }
    bool result = false;
    update(cFrame, goFrame);
    _writer->mode = mode;
    ((ttLibC_ContainerWriter_ *)_writer)->unit_duration = unitDuration;
    switch(_writer->type) {
    case containerType_flv:
      result = ttLibC_FlvWriter_write((ttLibC_FlvWriter *)_writer, cFrame, ttLibGoDataCallback, (void *)ptr);
      break;
    case containerType_mkv:
    case containerType_webm:
      result = ttLibC_MkvWriter_write((ttLibC_MkvWriter *)_writer, cFrame, ttLibGoDataCallback, (void *)ptr);
      break;
    case containerType_mp4:
      result = ttLibC_Mp4Writer_write((ttLibC_Mp4Writer *)_writer, cFrame, ttLibGoDataCallback, (void *)ptr);
      break;
    case containerType_mpegts:
      result = ttLibC_MpegtsWriter_write((ttLibC_MpegtsWriter *)_writer, cFrame, ttLibGoDataCallback, (void *)ptr);
      break;
    default:
      break;
    }
    reset(cFrame, goFrame);
    return result;
  }
  bool writeInfo(void *ptr) {
    if(_writer == NULL) {
      return false;
    }
    if(_writer->type != containerType_mpegts) {
      return true;
    }
    return ttLibC_MpegtsWriter_writeInfo((ttLibC_MpegtsWriter *)_writer, ttLibGoDataCallback, (void *)ptr);
  }
  uint64_t getPts() {
    if(_writer == nullptr) {
      return 0;
    }
    return _writer->pts;
  }
  uint32_t getTimebase() {
    if(_writer == nullptr) {
      return 1000;
    }
    return _writer->timebase;
  }
private:
  ttLibC_ContainerWriter *_writer;
};

extern "C" {

void *ContainerWriter_make(void *mp) {
  // これでmapの情報から、初期化を実施すれば良い。
  maps *m = reinterpret_cast<maps *>(mp);
  return new Writer(m);
}
bool ContainerWriter_writeFrame(void *writer, void *frame, void *goFrame, uint32_t mode, uint32_t unitDuration, uintptr_t ptr) {
  if(writer == nullptr) {
    return false;
  }
  Writer *w = reinterpret_cast<Writer *>(writer);
  ttLibC_Frame *cF = reinterpret_cast<ttLibC_Frame *>(frame);
  ttLibGoFrame *goF = reinterpret_cast<ttLibGoFrame *>(goFrame);
  return w->writeFrame(cF, goF, mode, unitDuration, (void *)ptr);
}
bool ContainerWriter_writeInfo(void *writer, uintptr_t ptr) {
  if(writer == nullptr) {
    return false;
  }
  Writer *w = reinterpret_cast<Writer *>(writer);
  return w->writeInfo((void *)ptr);
}
void ContainerWriter_close(void *writer) {
  Writer *w = reinterpret_cast<Writer *>(writer);
  delete w;
}

uint64_t ContainerWriter_getPts(void *writer) {
  Writer *w = reinterpret_cast<Writer *>(writer);
  return w->getPts();
}

uint32_t ContainerWriter_getTimebase(void *writer) {
  Writer *w = reinterpret_cast<Writer *>(writer);
  return w->getTimebase();
}

}