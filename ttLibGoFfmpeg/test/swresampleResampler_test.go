package test

import (
	"fmt"
	"os"
	"testing"

	"github.com/taktod/ttLibGo/ttLibGo"
	"github.com/taktod/ttLibGo/ttLibGoFfmpeg"
)

func TestSwresampleResampler(t *testing.T) {
	{
		t.Log("swresample")
		in, err := os.Open(os.Getenv("HOME") + "/tools/data/source/test.h264.aac.mkv")
		if err != nil {
			panic(err)
		}
		var count uint32
		var reader ttLibGo.Reader
		if !reader.Init("mkv") {
			t.Errorf("reader初期化失敗")
			return
		}
		defer reader.Close()

		var decoder ttLibGoFfmpeg.AvcodecDecoder
		defer decoder.Close()

		var resampler ttLibGoFfmpeg.SwresampleResampler
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
					if frame.CodecType == "aac" {
						decoder.InitAudio(frame.CodecType, frame.SampleRate, frame.ChannelNum)
						return decoder.Decode(
							frame,
							func(frame *ttLibGo.Frame) bool {
								fmt.Println(frame)
								resampler.Init(frame.CodecType, frame.SubType, frame.SampleRate, frame.ChannelNum,
									"pcmS16", "littleEndian", frame.SampleRate, frame.ChannelNum)
								return resampler.Resample(
									frame,
									func(frame *ttLibGo.Frame) bool {
										fmt.Println(frame) // このフレームptsが更新されないんですね、ちょっと怖い
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
