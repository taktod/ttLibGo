package test

import (
	"fmt"
	"os"
	"testing"

	"github.com/taktod/ttLibGo/ttLibGo"
)

func TestFlvReader(t *testing.T) {
	// 僕の場合は~/tools/data/sourceの部分にtest.vvv.aaa.mkvといったテスト用のファイルを大量につくってあり
	// cやnode、php、goで動作テストするときに利用しています。
	// このテストでもそういったデータをつかって動作テストしています。
	// ffmpegで適当に動画を変換してテストデータを作ってもらえればと思います。
	{
		t.Log("h264 / aac")
		in, err := os.Open(os.Getenv("HOME") + "/tools/data/source/test.h264.aac.flv")
		if err != nil {
			panic(err)
		}
		var reader ttLibGo.Reader
		if !reader.Init("flv") {
			t.Errorf("flv読み込みオブジェクト作成失敗")
			return
		}
		defer reader.Close()

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
					fmt.Println(frame)
					return true
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
