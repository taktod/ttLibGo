package testEx

import (
	"fmt"
	"os"
	"testing"

	"github.com/taktod/ttLibGo/ttLibGoJpeg"

	"github.com/taktod/ttLibGo/ttLibGo"
	"github.com/taktod/ttLibGo/ttLibGoFfmpeg"
)

func TestFrameGetBinaryBuffer(t *testing.T) {
	// とりあえずjpegをdecodeしてみるか・・
	jpegIn, _ := os.Open("./target.jpeg") // このデータはyuvj444p
	st, _ := jpegIn.Stat()
	fmt.Println(st.Size())
	buf := make([]byte, st.Size())
	jpegIn.Read(buf)
	jpeg := new(ttLibGo.Frame)
	jpeg.Binary = buf
	jpeg.CodecType = "jpeg"
	jpeg.Pts = 0
	jpeg.Timebase = 1000
	jpeg.Restore()
	defer jpeg.Close()
	var cloned ttLibGo.Frame
	defer cloned.Close()
	// まぁいいや。あとはできそうだから、decodeするか・・・
	var jpegDecoder ttLibGoFfmpeg.AvcodecDecoder
	defer jpegDecoder.Close()
	jpegDecoder.InitVideo("jpeg", 0, 0)
	jpegDecoder.Decode(
		jpeg,
		func(frame *ttLibGo.Frame) bool {
			// これをencodeでもしてみるか・・・
			var jpegEncoder ttLibGoJpeg.JpegEncoder
			defer jpegEncoder.Close()
			jpegEncoder.Init(200, 200, 90)
			jpegEncoder.Encode(
				frame,
				func(frame *ttLibGo.Frame) bool {
					fmt.Println(frame)
					out, _ := os.OpenFile("gen.jpeg", os.O_WRONLY|os.O_CREATE, 0666)
					out.Write(frame.GetBinaryBuffer())
					return true
				})
			return true
		})
}
