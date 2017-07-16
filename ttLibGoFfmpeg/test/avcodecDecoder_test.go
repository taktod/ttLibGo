package test

import (
	"os"
	"testing"

	"github.com/taktod/ttLibGo/ttLibGo"
	"github.com/taktod/ttLibGo/ttLibGoFfmpeg"
)

func TestAvcodecVideoDecoder(t *testing.T) {
	{
		targetCodec := "h264"
		t.Log(targetCodec)
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
						decoder.InitVideo(targetCodec, frame.Width, frame.Height)
						return decoder.Decode(
							frame,
							func(frame *ttLibGo.Frame) bool {
								count++
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
			t.Errorf("デコードされませんでした")
			return
		}
		t.Logf("デコードされた数:%d", count)
	}
}

func TestAvcodecAudioDecoder(t *testing.T) {
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
								count++
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
			t.Errorf("デコードされませんでした")
			return
		}
		t.Logf("デコードされた数:%d", count)
	}
}
