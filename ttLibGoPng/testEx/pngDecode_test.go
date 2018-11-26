package testEx

import (
	"fmt"
	"os"
	"testing"

	"github.com/taktod/ttLibGo/ttLibGoPng"

	"github.com/taktod/ttLibGo/ttLibGoFfmpeg"
	"github.com/taktod/ttLibGo/ttLibGoJpeg"

	"github.com/taktod/ttLibGo/ttLibGo"
)

func TestPngDecode(t *testing.T) {
	pngIn, _ := os.Open("./target.png")
	st, _ := pngIn.Stat()
	fmt.Println(st.Size())
	buf := make([]byte, st.Size())
	pngIn.Read(buf)
	png := new(ttLibGo.Frame)
	png.Binary = buf
	png.CodecType = "png"
	png.Pts = 0
	png.Timebase = 1000
	png.Restore()
	defer png.Close()
	var pngDecoder ttLibGoPng.PngDecoder
	defer pngDecoder.Close()
	pngDecoder.Init()
	pngDecoder.Decode(
		png,
		func(frame *ttLibGo.Frame) bool {
			var resampler ttLibGoFfmpeg.SwscaleResampler
			defer resampler.Close()
			resampler.Init(
				frame.CodecType,
				frame.SubType,
				frame.Width,
				frame.Height,
				"yuv", "planar", frame.Width, frame.Height, "Bilinear")
			resampler.Resample(
				frame,
				func(frame *ttLibGo.Frame) bool {
					fmt.Println(frame)
					var jpegEncoder ttLibGoJpeg.JpegEncoder
					defer jpegEncoder.Close()
					jpegEncoder.Init(frame.Width, frame.Height, 90)
					jpegEncoder.Encode(
						frame,
						func(frame *ttLibGo.Frame) bool {
							// あれ・・・なんかただしい映像になってない・・・
							fmt.Println(frame)
							out, _ := os.OpenFile("gen.jpeg", os.O_WRONLY|os.O_CREATE, 0666)
							out.Write(frame.GetBinaryBuffer())
							return true
						})
					return true
				})
			return true
		})
}
