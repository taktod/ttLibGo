package ttLibGo

/*
#include <ttLibC/frame/frame.h>
*/
import "C"
import (
	"unsafe"
)

// FrameCallback フレームが生成されたときのcallback動作
type FrameCallback func(frame *Frame) bool

// DataCallback binaryデータが生成されたときのcallback動作
type DataCallback func(data []byte) bool

// frameCall c言語側でのフレーム生成callbackの処理で利用する構造体
// これを利用してgoの世界に戻す
type frameCall struct {
	callback FrameCallback
}

//export ttLibGoFrameCallback
func ttLibGoFrameCallback(any unsafe.Pointer, cFrame *C.ttLibC_Frame) C.bool {
	// c言語側に渡してたvoid*にいれてる参照からgoで利用するcallbackを復元する
	call := (*frameCall)(any)
	// goで利用可能なフレームの形にして応答する
	frame := new(Frame)
	frame.init(cttLibCFrame(cFrame))
	return C.bool(call.callback(frame))
}

// dataCall c言語側でbinaryデータを取得したときにgoの世界に戻すための構造体
type dataCall struct {
	callback DataCallback
}

//export ttLibGoDataCallback
func ttLibGoDataCallback(any unsafe.Pointer, data unsafe.Pointer, dataSize C.size_t) C.bool {
	call := (*dataCall)(any)
	gData := C.GoBytes(data, C.int(dataSize))
	return C.bool(call.callback(gData))
}
