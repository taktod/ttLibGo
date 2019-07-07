package ttLibGo

/*
#cgo pkg-config: libavcodec libavutil libswscale libswresample libjpeg openh264 opus libpng speex theora vorbis vorbisenc fdk-aac x264 x265 soundtouch speexdsp
#cgo CFLAGS: -I./ -I./ttLibC/ -I/usr/local/include -I./csrc/ -I./cppsrc/ -I./dummy/ -std=c99
#cgo LDFLAGS: -lc++ -L/usr/local/lib
#cgo CXXFLAGS: -std=c++0x -I./ -I./ttLibC/ -I/usr/local/include -I./csrc/ -I./cppsrc/ -I./dummy/
*/
import "C"
