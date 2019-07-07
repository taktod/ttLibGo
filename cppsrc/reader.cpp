#include <stdio.h>
#include <ttLibC/allocator.h>
#include <ttLibC/container/flv.h>
#include <ttLibC/container/mkv.h>
#include <ttLibC/container/mp4.h>
#include <ttLibC/container/mpegts.h>
#include <string>

using namespace std;

// go側の関数をkickできるようにしておきます。
extern "C" {

extern bool ttLibGoFrameCallback(void *ptr, ttLibC_Frame *frame);

ttLibC_ContainerReader *ContainerReader_make(const char *format) {
  string fmt = string(format);
  if(fmt == "flv") {
		return (ttLibC_ContainerReader *)ttLibC_FlvReader_make();
  }
  if(fmt == "mkv" || fmt == "webm") {
		return (ttLibC_ContainerReader *)ttLibC_MkvReader_make();
  }
  if(fmt == "mp4") {
		return (ttLibC_ContainerReader *)ttLibC_Mp4Reader_make();
  }
  if(fmt == "mpegts") {
		return (ttLibC_ContainerReader *)ttLibC_MpegtsReader_make();
  }
	return NULL;
}

// 読み込み実施
bool ContainerReader_read(
		ttLibC_ContainerReader *reader,
		void *data,
		size_t data_size,
		uintptr_t ptr) {
	if(reader == NULL) {
		return false;
	}
	switch(reader->type) {
	case containerType_flv:
		return ttLibC_FlvReader_read((ttLibC_FlvReader *)reader, data, data_size, [](void *ptr, ttLibC_Flv *flv) -> bool{
    	return ttLibC_Flv_getFrame(flv, ttLibGoFrameCallback, ptr);
    }, (void *)ptr);
	case containerType_mkv:
	case containerType_webm:
		return ttLibC_MkvReader_read((ttLibC_MkvReader *)reader, data, data_size, [](void *ptr, ttLibC_Mkv *mkv) -> bool{
      return ttLibC_Mkv_getFrame(mkv, ttLibGoFrameCallback, ptr);
    }, (void *)ptr);
	case containerType_mp4:
		return ttLibC_Mp4Reader_read((ttLibC_Mp4Reader *)reader, data, data_size, [](void *ptr, ttLibC_Mp4 *mp4) -> bool {
	    return ttLibC_Mp4_getFrame(mp4, ttLibGoFrameCallback, ptr);
    }, (void *)ptr);
	case containerType_mpegts:
		return ttLibC_MpegtsReader_read((ttLibC_MpegtsReader *)reader, data, data_size, [](void *ptr, ttLibC_Mpegts *mpegts) -> bool {
    	return ttLibC_Mpegts_getFrame(mpegts, ttLibGoFrameCallback, ptr);
    }, (void *)ptr);
	default:
		break;
	}
	return false;
}

}