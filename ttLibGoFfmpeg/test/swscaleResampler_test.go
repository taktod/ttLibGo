package test

import (
	"os"
	"testing"

	"fmt"

	"github.com/taktod/ttLibGo/ttLibGo"
	"github.com/taktod/ttLibGo/ttLibGoFfmpeg"
)

func TestSwscaleResampler(t *testing.T) {
	{
		t.Log("swscale")
		in, err := os.Open(os.Getenv("HOME") + "/tools/data/source/test.h264.aac.flv")
		if err != nil {
			panic(err)
		}
		var count uint32
		var reader ttLibGo.Reader
		if !reader.Init("flv") {
			t.Errorf("reader初期化失敗")
		}
		defer reader.Close()

		var decoder ttLibGoFfmpeg.AvcodecDecoder
		defer decoder.Close()

		var resampler ttLibGoFfmpeg.SwscaleResampler
		defer resampler.Close()

		cloned := new(ttLibGo.Frame)
		defer cloned.Close()

		for {
			buffer := make([]byte, 65536)
			length, err := in.Read(buffer)
			if err != nil {
				panic(err)
			}
			if !reader.ReadFrame(
				buffer,
				uint64(length),
				func(frame *ttLibGo.Frame) bool {
					if frame.CodecType == "h264" {
						decoder.InitVideo(frame.CodecType, frame.Width, frame.Height)
						return decoder.Decode(
							frame,
							func(frame *ttLibGo.Frame) bool {
								resampler.Init(frame.CodecType, frame.SubType, frame.Width, frame.Height,
									"yuv", "planar", 360, 240, "Bilinear")
								resampler.Resample(
									frame,
									func(frame *ttLibGo.Frame) bool {
										fmt.Println(frame)
										count++
										return true
									})
								return true
							})
					}
					return true
				}) {
				t.Errorf("コンテナ読み込み失敗")
				return
			}
			if length < 65536 {
				break
			}
		}
		if count == 0 {
			t.Errorf("リサンプルされませんでした")
			return
		}
		t.Logf("リサンプルされた数:%d", count)
	}
}
