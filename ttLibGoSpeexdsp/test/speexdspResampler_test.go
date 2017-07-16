package test

import (
	"fmt"
	"os"
	"testing"

	"github.com/taktod/ttLibGo/ttLibGo"
	"github.com/taktod/ttLibGo/ttLibGoSpeex"
	"github.com/taktod/ttLibGo/ttLibGoSpeexdsp"
)

func TestSpeexdspResampler(t *testing.T) {
	{
		in, err := os.Open(os.Getenv("HOME") + "/tools/data/source/test.flv1.speex.flv")
		if err != nil {
			panic(err)
		}
		var count uint32
		var reader ttLibGo.Reader
		if !reader.Init("flv") {
			t.Errorf("reader初期化失敗")
			return
		}
		defer reader.Close()

		var decoder ttLibGoSpeex.SpeexDecoder
		defer decoder.Close()
		var resampler ttLibGoSpeexdsp.SpeexdspResampler
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
					if frame.CodecType == "speex" {
						decoder.Init(frame.SampleRate, frame.ChannelNum)
						return decoder.Decode(
							frame,
							func(frame *ttLibGo.Frame) bool {
								resampler.Init(frame.ChannelNum, frame.SampleRate, 48000, 8)
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
