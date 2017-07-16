package test

import (
	"fmt"
	"os"
	"testing"

	"github.com/taktod/ttLibGo/ttLibGo"
	"github.com/taktod/ttLibGo/ttLibGoFaac"
	"github.com/taktod/ttLibGo/ttLibGoFfmpeg"
)

func TestFaacEncoder(t *testing.T) {
	{
		targetCodec := "aac"
		t.Log(targetCodec)
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

		var resampler ttLibGoFfmpeg.SwresampleResampler
		defer resampler.Close()

		var encoder ttLibGoFaac.FaacEncoder
		defer encoder.Close()

		var pts uint64
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
					if frame.CodecType == targetCodec {
						decoder.InitAudio(targetCodec, frame.SampleRate, frame.ChannelNum)
						return decoder.Decode(
							frame,
							func(frame *ttLibGo.Frame) bool {
								resampler.Init(frame.CodecType, frame.SubType, frame.SampleRate, frame.ChannelNum,
									"pcmS16", "littleEndian", frame.SampleRate, frame.ChannelNum)
								return resampler.Resample(
									frame,
									func(frame *ttLibGo.Frame) bool {
										frame.Pts = pts
										pts += uint64(frame.SampleNum)
										frame.Timebase = frame.SampleRate
										encoder.Init("Low", frame.SampleRate, frame.ChannelNum, 96000)
										return encoder.Encode(
											frame,
											func(frame *ttLibGo.Frame) bool {
												fmt.Println(frame)
												count++
												return true
											})
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
			t.Errorf("エンコードされませんでした")
			return
		}
		t.Logf("エンコードされた数:%d", count)
	}
}
