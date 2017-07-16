package ttLibGoVorbis

/*
#cgo pkg-config: vorbis vorbisenc ogg
#cgo CFLAGS: -I../ -I../ttLibC/ -I/usr/local/include -std=c99 -D__ENABLE_VORBIS_DECODE__=1 -D__ENABLE_VORBIS_ENCODE__=1
#cgo LDFLAGS: -lc++ -L/usr/local/lib
#cgo CXXFLAGS: -std=c++0x -I../ -I../ttLibC/ -I/usr/local/include -D__ENABLE_VORBIS_DECODE__=1 -D__ENABLE_VORBIS_ENCODE__=1
*/
import "C"
