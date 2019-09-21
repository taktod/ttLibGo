#include <string>
#include "util.hpp"
#include "ttLibC/ttLibC/container/container.h"
#include "ttLibC/ttLibC/container/flv.h"
#include "ttLibC/ttLibC/container/mp4.h"
#include "ttLibC/ttLibC/container/mkv.h"
#include "ttLibC/ttLibC/container/mpegts.h"
#include "ttLibC/ttLibC/container/containerCommon.h"
#include "frame.hpp"

extern "C" {
typedef void *(* ttLibC_FlvWriter_make_func)(ttLibC_Frame_Type, ttLibC_Frame_Type);
typedef void *(* ttLibC_containerWriter_make_func)(ttLibC_Frame_Type *, uint32_t);
typedef bool (* ttLibC_containerWriter_write_func)(void *, void *, bool(*)(void *, void *, size_t), void *);
typedef bool (* ttLibC_mpegtsWriter_writeInfo_func)(void *, bool(*)(void *, void *,size_t), void *);
typedef void (* ttLibC_close_func)(void **);

extern ttLibC_FlvWriter_make_func       ttLibGo_FlvWriter_make;
extern ttLibC_containerWriter_make_func ttLibGo_Mp4Writer_make;
extern ttLibC_containerWriter_make_func ttLibGo_MpegtsWriter_make;
extern ttLibC_containerWriter_make_func ttLibGo_MkvWriter_make;
extern ttLibC_containerWriter_write_func ttLibGo_FlvWriter_write;
extern ttLibC_containerWriter_write_func ttLibGo_MkvWriter_write;
extern ttLibC_containerWriter_write_func ttLibGo_Mp4Writer_write;
extern ttLibC_containerWriter_write_func ttLibGo_MpegtsWriter_write;
extern ttLibC_mpegtsWriter_writeInfo_func ttLibGo_MpegtsWriter_writeInfo;
extern ttLibC_close_func ttLibGo_ContainerWriter_close;

extern bool ttLibGoDataCallback(void *ptr, void *data, size_t data_size);
}

class Writer : public FrameProcessor {
public:
  Writer(maps *m) {
    if(ttLibGo_FlvWriter_make != nullptr
    && ttLibGo_Mp4Writer_make != nullptr
    && ttLibGo_MpegtsWriter_make != nullptr
    && ttLibGo_MkvWriter_make != nullptr) {
      string type = m->getString("type");
      if(type == "flv") {
        _writer = (ttLibC_ContainerWriter *)(*ttLibGo_FlvWriter_make)(
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
        _writer = (ttLibC_ContainerWriter *)(*ttLibGo_Mp4Writer_make)(types, frameList.size());
      }
      else if(type == "mpegts") {
        _writer = (ttLibC_ContainerWriter *)(*ttLibGo_MpegtsWriter_make)(types, frameList.size());
      }
      else if(type == "mkv") {
        _writer = (ttLibC_ContainerWriter *)(*ttLibGo_MkvWriter_make)(types, frameList.size());
      }
      else if(type == "webm") {
        _writer = (ttLibC_ContainerWriter *)(*ttLibGo_MkvWriter_make)(types, frameList.size());
        _writer->type = containerType_webm;
      }
      delete[] types;
    }
  }
  ~Writer() {
    if(ttLibGo_ContainerWriter_close != nullptr) {
      (*ttLibGo_ContainerWriter_close)((void **)&_writer);
    }
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
    if(ttLibGo_FlvWriter_write != nullptr
    && ttLibGo_MkvWriter_write != nullptr
    && ttLibGo_MpegtsWriter_write != nullptr
    && ttLibGo_Mp4Writer_write != nullptr) {
      update(cFrame, goFrame);
      _writer->mode = mode;
      switch(_writer->type) {
      case containerType_flv:
        result = (*ttLibGo_FlvWriter_write)((ttLibC_FlvWriter *)_writer, cFrame, ttLibGoDataCallback, (void *)ptr);
        break;
      case containerType_mkv:
      case containerType_webm:
        ((ttLibC_ContainerWriter_ *)_writer)->unit_duration = unitDuration;
        result = (*ttLibGo_MkvWriter_write)((ttLibC_MkvWriter *)_writer, cFrame, ttLibGoDataCallback, (void *)ptr);
        break;
      case containerType_mp4:
        ((ttLibC_ContainerWriter_ *)_writer)->unit_duration = unitDuration;
        result = (*ttLibGo_Mp4Writer_write)((ttLibC_Mp4Writer *)_writer, cFrame, ttLibGoDataCallback, (void *)ptr);
        break;
      case containerType_mpegts:
        ((ttLibC_ContainerWriter_ *)_writer)->unit_duration = unitDuration;
        result = (*ttLibGo_MpegtsWriter_write)((ttLibC_MpegtsWriter *)_writer, cFrame, ttLibGoDataCallback, (void *)ptr);
        break;
      default:
        break;
      }
      reset(cFrame, goFrame);
    }
    return result;
  }
  bool writeInfo(void *ptr) {
    if(_writer == NULL) {
      return false;
    }
    if(_writer->type != containerType_mpegts) {
      return true;
    }
    if(ttLibGo_MpegtsWriter_writeInfo != nullptr) {
      return (*ttLibGo_MpegtsWriter_writeInfo)((ttLibC_MpegtsWriter *)_writer, ttLibGoDataCallback, (void *)ptr);
    }
    return false;
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
