package ttLibGoFfmpeg

/*
#cgo pkg-config: libavcodec libavutil libswscale libswresample
#cgo CFLAGS: -I../ -I../ttLibC/ -std=c99 -D__ENABLE_AVCODEC__=1 -D__ENABLE_SWRESAMPLE__=1 -D__ENABLE_SWSCALE__=1
#cgo LDFLAGS: -lc++
#cgo CXXFLAGS: -std=c++0x -I../ -I../ttLibC/ -D__ENABLE_AVCODEC__=1 -D__ENABLE_SWRESAMPLE__=1 -D__ENABLE_SWSCALE__=1
*/
import "C"

// ここでttLibCとffmpegのライブラリとのリンクを実施しておきます。
// pkg-configで実施してるので、windowsだとうまく使えないかも
