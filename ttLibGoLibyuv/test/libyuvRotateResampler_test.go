package test

import (
	"fmt"
	"os"
	"testing"

	"github.com/taktod/ttLibGo/ttLibGo"
	"github.com/taktod/ttLibGo/ttLibGoFfmpeg"
	"github.com/taktod/ttLibGo/ttLibGoLibyuv"
)

func TestLibyuvRotateResampler(t *testing.T) {
	{
		in, err := os.Open(os.Getenv("HOME") + "/tools/data/source/test.h264.aac.mkv")
		if err != nil {
			panic(err)
		}
		var count uint32
		var reader ttLibGo.Reader
		if !reader.Init("mkv") {
			t.Errorf("reader初期化失敗")
		}
		defer reader.Close()
		var decoder ttLibGoFfmpeg.AvcodecDecoder
		defer decoder.Close()
		var resampler ttLibGoLibyuv.LibyuvRotateResampler
		defer resampler.Close()
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
						decoder.InitVideo("h264", frame.Width, frame.Height)
						return decoder.Decode(
							frame,
							func(frame *ttLibGo.Frame) bool {
								resampler.Init(90)
								return resampler.Resample(
									frame,
									func(frame *ttLibGo.Frame) bool {
										fmt.Println(frame)
										count++
										return true
									})
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
