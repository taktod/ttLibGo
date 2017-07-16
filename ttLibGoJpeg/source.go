package ttLibGoJpeg

/*
#cgo CFLAGS: -I../ -I../ttLibC/ -I/usr/local/include -std=c99 -D__ENABLE_JPEG__=1
#cgo LDFLAGS: -lc++ -L/usr/local/lib -ljpeg
#cgo CXXFLAGS: -std=c++0x -I../ -I../ttLibC/ -I/usr/local/include -D__ENABLE_JPEG__=1
*/
import "C"

// jpegが/usr/local/lib /usr/local/includeにインストールされてることを期待しています。
// たぶん、centOSとかでyunつかってると/usr/libや/usr/includeにはいってる気がします
