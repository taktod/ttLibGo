package ttLibGoMp3lame

/*
#cgo CFLAGS: -I../ -I../ttLibC/ -I/usr/local/include -std=c99 -D__ENABLE_MP3LAME_ENCODE__=1 -D__ENABLE_MP3LAME_DECODE__=1
#cgo LDFLAGS: -lc++ -L/usr/local/lib -lmp3lame
#cgo CXXFLAGS: -std=c++0x -I../ -I../ttLibC/ -I/usr/local/include -D__ENABLE_MP3LAME_ENCODE__=1 -D__ENABLE_MP3LAME_DECODE__=1
*/
import "C"
