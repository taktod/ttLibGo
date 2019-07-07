#ifndef TTLIBGO_FRAME_HPP
#define TTLIBGO_FRAME_HPP

// goのframeが保持しているデータをコピーしておくオブジェクト
class ttLibGoFrame {
public:
  uint64_t pts;
  uint64_t dts;
  uint32_t timebase;
  uint32_t id;
  uint32_t width;
  uint32_t height;

  uint32_t sampleRate;
  uint32_t sampleNum;
  uint32_t channelNum;

  uint64_t dataPos;
  uint32_t widthStride;

  uint64_t yDataPos;
  uint32_t yStride;
  uint64_t uDataPos;
  uint32_t uStride;
  uint64_t vDataPos;
  uint32_t vStride;

  // pcmについては、データを参照してから、binaryを作り直すことはあっても、それを利用してデータを改変することはなさそうなので
  // 改変情報を保持しなくてもいいか・・・
};

class FrameProcessor {
protected:
  void update(ttLibC_Frame *cFrame, ttLibGoFrame *goFrame);
  void reset(ttLibC_Frame *cFrame, ttLibGoFrame *goFrame);
private:
  uint64_t _pts, _dts;
  uint32_t _timebase;
  uint32_t _id;
  uint32_t _width, _height;
  uint32_t _sample_rate, _sample_num, _channel_num;
  uint8_t *_data, *_y_data, *_u_data, *_v_data;
  uint32_t _width_stride, _y_stride, _u_stride, _v_stride;
};

ttLibC_Frame_Type Frame_getFrameTypeFromString(string name);

#endif
