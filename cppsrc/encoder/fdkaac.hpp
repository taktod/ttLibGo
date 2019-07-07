#ifndef TTLIBGO_ENCODER_FDKAAC_HPP
#define TTLIBGO_ENCODER_FDKAAC_HPP

#include "../encoder.hpp"
#include <ttLibC/frame/audio/aac.h>

class FdkaacEncoder : public Encoder {
public:
  FdkaacEncoder(maps *mp);
  ~FdkaacEncoder();
  bool encodeFrame(ttLibC_Frame *cFrame, ttLibGoFrame *goFrame, void *ptr);
  bool setBitrate(uint32_t value);
private:
#ifdef __ENABLE_FDKAAC_ENCODE__
  HANDLE_AACENCODER handle_;
  ttLibC_Aac *aac_;
	uint8_t    *pcm_buffer_;
	size_t      pcm_buffer_size_;
	size_t      pcm_buffer_next_pos_;
	uint8_t    *data_;
	size_t      data_size_;
	uint64_t    pts_;
	uint32_t    sample_rate_;
	uint32_t    sample_num_;
	bool        is_pts_initialized_;
#endif
};

#endif
