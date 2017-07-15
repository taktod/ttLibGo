package test

import (
	"os"
	"testing"

	"github.com/taktod/ttLibGo/ttLibGo"
)

func TestFlvWriter(t *testing.T) {
	{
		t.Log("h264 / aac")
		in, err := os.Open(os.Getenv("HOME") + "/tools/data/source/test.h264.aac.flv")
		if err != nil {
			panic(err)
		}
		out, err := os.OpenFile("test.h264.aac.flv", os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			panic(err)
		}
		var reader ttLibGo.Reader
		if !reader.Init("flv") {
			t.Errorf("flv読み込みオブジェクト作成失敗")
			return
		}
		defer reader.Close()

		var writer ttLibGo.Writer
		if !writer.Init("flv", 1000, "h264", "aac") {
			t.Errorf("flv書き出しオブジェクト作成失敗")
			return
		}
		defer writer.Close()

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
					return writer.WriteFrame(
						frame,
						func(data []byte) bool {
							out.Write(data)
							return true
						})
				}) {
				t.Errorf("読み込み失敗しました")
				return
			}
			if length < 65536 {
				break
			}
		}
	}
}
