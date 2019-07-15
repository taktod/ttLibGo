#include <stdio.h>
#include <ttLibC/allocator.h>
#include <ttLibC/container/container.h>
#include <string>

using namespace std;

// go側の関数をkickできるようにしておきます。
extern "C" {

extern bool ttLibGoFrameCallback(void *ptr, ttLibC_Frame *frame);

typedef void *(* ttLibC_make_func)();
typedef bool(* ttLibC_ContainerReader_read_func)(void *, void *, size_t, bool(*)(void *, void *), void *);
typedef bool(* ttLibC_Container_GetFrame_func)(void *, ttLibC_getFrameFunc, void *);
typedef void (* ttLibC_close_func)(void **);

extern ttLibC_make_func ttLibGo_FlvReader_make;
extern ttLibC_make_func ttLibGo_MkvReader_make;
extern ttLibC_make_func ttLibGo_Mp4Reader_make;
extern ttLibC_make_func ttLibGo_MpegtsReader_make;

extern ttLibC_ContainerReader_read_func ttLibGo_FlvReader_read;
extern ttLibC_ContainerReader_read_func ttLibGo_MkvReader_read;
extern ttLibC_ContainerReader_read_func ttLibGo_Mp4Reader_read;
extern ttLibC_ContainerReader_read_func ttLibGo_MpegtsReader_read;

extern ttLibC_Container_GetFrame_func ttLibGo_Flv_getFrame;
extern ttLibC_Container_GetFrame_func ttLibGo_Mkv_getFrame;
extern ttLibC_Container_GetFrame_func ttLibGo_Mp4_getFrame;
extern ttLibC_Container_GetFrame_func ttLibGo_Mpegts_getFrame;

extern ttLibC_close_func ttLibGo_ContainerReader_close;

ttLibC_ContainerReader *ContainerReader_make(const char *format) {
  string fmt = string(format);
  if(fmt == "flv") {
		if(ttLibGo_FlvReader_make != nullptr) {
			return (ttLibC_ContainerReader *)(*ttLibGo_FlvReader_make)();
		}
  }
  if(fmt == "mkv" || fmt == "webm") {
		if(ttLibGo_MkvReader_make != nullptr) {
			return (ttLibC_ContainerReader *)(*ttLibGo_MkvReader_make)();
		}
  }
  if(fmt == "mp4") {
		if(ttLibGo_Mp4Reader_make != nullptr) {
			return (ttLibC_ContainerReader *)(*ttLibGo_Mp4Reader_make)();
		}
  }
  if(fmt == "mpegts") {
		if(ttLibGo_MpegtsReader_make != nullptr) {
			return (ttLibC_ContainerReader *)(*ttLibGo_MpegtsReader_make)();
		}
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
		if(ttLibGo_FlvReader_read != nullptr) {
			return (*ttLibGo_FlvReader_read)(reader, data, data_size, [](void *ptr, void *flv) -> bool{
				if(ttLibGo_Flv_getFrame != nullptr) {
	  	  	return (*ttLibGo_Flv_getFrame)(flv, ttLibGoFrameCallback, ptr);
				}
				return false;
	    }, (void *)ptr);
		}
		break;
	case containerType_mkv:
	case containerType_webm:
		if(ttLibGo_MkvReader_read != nullptr) {
			return (*ttLibGo_MkvReader_read)(reader, data, data_size, [](void *ptr, void *mkv) -> bool{
				if(ttLibGo_Mkv_getFrame != nullptr) {
	  	  	return (*ttLibGo_Mkv_getFrame)(mkv, ttLibGoFrameCallback, ptr);
				}
				return false;
	    }, (void *)ptr);
		}
		break;
	case containerType_mp4:
		if(ttLibGo_Mp4Reader_read != nullptr) {
			return (*ttLibGo_Mp4Reader_read)(reader, data, data_size, [](void *ptr, void *mp4) -> bool{
				if(ttLibGo_Mp4_getFrame != nullptr) {
	  	  	return (*ttLibGo_Mp4_getFrame)(mp4, ttLibGoFrameCallback, ptr);
				}
				return false;
	    }, (void *)ptr);
		}
		break;
	case containerType_mpegts:
		if(ttLibGo_MpegtsReader_read != nullptr) {
			return (*ttLibGo_MpegtsReader_read)(reader, data, data_size, [](void *ptr, void *mpegts) -> bool{
				if(ttLibGo_Mpegts_getFrame != nullptr) {
	  	  	return (*ttLibGo_Mpegts_getFrame)(mpegts, ttLibGoFrameCallback, ptr);
				}
				return false;
	    }, (void *)ptr);
		}
		break;
	default:
		break;
	}
	return false;
}

void ContainerReader_close(ttLibC_ContainerReader *reader) {
	if(ttLibGo_ContainerReader_close != nullptr) {
		(*ttLibGo_ContainerReader_close)((void **)&reader);
	}
}

}