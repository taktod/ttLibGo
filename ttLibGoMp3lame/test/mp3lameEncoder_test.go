package test

import (
	"fmt"
	"os"
	"testing"

	"github.com/taktod/ttLibGo/ttLibGo"
	"github.com/taktod/ttLibGo/ttLibGoMp3lame"
)

func TestMp3lameEncoder(t *testing.T) {
	{
		in, err := os.Open(os.Getenv("HOME") + "/tools/data/source/test.h265.mp3.mkv")
		if err != nil {
			panic(err)
		}
		var count uint32
		var reader ttLibGo.Reader
		if !reader.Init("mkv") {
			t.Errorf("reader初期化失敗")
		}
		defer reader.Close()
		var decoder ttLibGoMp3lame.Mp3lameDecoder
		defer decoder.Close()
		var encoder ttLibGoMp3lame.Mp3lameEncoder
		defer encoder.Close()
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
					if frame.CodecType == "mp3" {
						decoder.Init()
						return decoder.Decode(
							frame,
							func(frame *ttLibGo.Frame) bool {
								encoder.Init(frame.SampleRate, frame.ChannelNum, 5)
								return encoder.Encode(
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
			t.Errorf("エンコードされませんでした")
			return
		}
		t.Logf("エンコードされた数:%d", count)
	}
}
