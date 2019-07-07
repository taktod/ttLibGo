package test

import (
	"os"
	"testing"

	"github.com/taktod/ttLibGo"
)

func TestYuvClip(t *testing.T) {
	t.Log("h264 / aac")
	in, err := os.Open(os.Getenv("HOME") + "/tools/data/source/test.h264.aac.flv")
	if err != nil {
		panic(err)
	}
	//	counter := 1
	decoder := ttLibGo.Decoders.AvcodecVideo(ttLibGo.FrameTypes.H264, 640, 360)
	defer decoder.Close()
	encoder := ttLibGo.Encoders.Jpeg(320, 360, 90)
	defer encoder.Close()
	reader := ttLibGo.Readers.Flv()
	defer reader.Close()
	for {
		buffer := make([]byte, 65536)
		length, err := in.Read(buffer)
		if length == 0 {
			break
		}
		if err != nil {
			panic(err)
		}
		if !reader.ReadFrame(
			buffer,
			uint64(length),
			func(frame *ttLibGo.Frame) bool {
				switch frame.Type {
				case ttLibGo.FrameTypes.H264:
					decoder.DecodeFrame(frame, func(frame *ttLibGo.Frame) bool {
						// ここでyuvのデータをclipしてみる
						if frame.Type == ttLibGo.FrameTypes.Yuv420 {
							yuv420 := ttLibGo.Yuv420.Cast(frame)
							yuv420.Width = 320
							yuv420.YData = 160
							yuv420.UData = 80
							yuv420.VData = 80
							return encoder.EncodeFrame(yuv420, func(frame *ttLibGo.Frame) bool {
								//								fmt.Println(frame)
								/*								frame.GetBinaryBuffer(func(buf []byte) bool {
																file := fmt.Sprintf("out%d.jpeg", counter)
																out, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE, 0666)
																if err != nil {
																	panic(err)
																}
																out.Write(buf)
																counter++
																return true
															})*/
								return true
							})
						}
						return true
					})
					// これを適当にdecodeしてみるか・・・
					break
				default:
					break
				}
				return true
			}) {
			t.Errorf("読み込み失敗しました")
			return
		}
		if length < 65536 {
			break
		}
	}
	//	t.Error() // とりあえずエラー完了にしとく
}
