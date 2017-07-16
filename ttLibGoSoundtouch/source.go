package ttLibGoSoundtouch

/*
#cgo pkg-config: soundtouch
#cgo CFLAGS: -I../ -I../ttLibC/ -I/usr/local/include -std=c99 -D__ENABLE_SOUNDTOUCH__=1
#cgo LDFLAGS: -lc++ -L/usr/local/lib
#cgo CXXFLAGS: -std=c++0x -I../ -I../ttLibC/ -I/usr/local/include -D__ENABLE_SOUNDTOUCH__=1
*/
import "C"
