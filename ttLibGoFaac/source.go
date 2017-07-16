package ttLibGoFaac

/*
#cgo CFLAGS: -I../ -I../ttLibC/ -I/usr/local/include -std=c99 -D__ENABLE_FAAC_ENCODE__=1
#cgo LDFLAGS: -lc++ -L/usr/local/lib -lfaac
#cgo CXXFLAGS: -std=c++0x -I../ -I../ttLibC/ -I/usr/local/include -D__ENABLE_FAAC_ENCODE__=1
*/
import "C"

// faacについては/usr/local/include、/usr/local/libにインストールされていることを期待しておきます
