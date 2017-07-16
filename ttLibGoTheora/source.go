package ttLibGoTheora

/*
#cgo pkg-config: theora theoraenc theoradec ogg
#cgo CFLAGS: -I../ -I../ttLibC/ -I/usr/local/include -std=c99 -D__ENABLE_THEORA__=1
#cgo LDFLAGS: -lc++ -L/usr/local/lib
#cgo CXXFLAGS: -std=c++0x -I../ -I../ttLibC/ -I/usr/local/include -D__ENABLE_THEORA__=1
*/
import "C"
