package test

import (
	"fmt"
	"os"
	"testing"

	"github.com/taktod/ttLibGo"
)

func TestX265EncodeDataCheck(t *testing.T) {
	in, err := os.Open(os.Getenv("HOME") + "/tools/data/source/test.h264.aac.flv")
	if err != nil {
		panic(err)
	}
	counter := 0
	frameTypes := ttLibGo.FrameTypes
	reader := ttLibGo.Readers.Flv()
	defer reader.Close()
	decoder := ttLibGo.Decoders.AvcodecVideo(frameTypes.H264, 640, 360)
	defer decoder.Close()
	encoder := ttLibGo.Encoders.X265(640, 360, "", "", "", nil)
	defer encoder.Close()
	decoder2 := ttLibGo.Decoders.AvcodecVideo(frameTypes.H265, 640, 360)
	defer decoder2.Close()
	encoder2 := ttLibGo.Encoders.Jpeg(640, 360, 90)
	defer encoder2.Close()
	for {
		buffer := make([]byte, 65536)
		length, err := in.Read(buffer)
		if length == 0 {
			break
		}
		if err != nil {
			panic(err)
		}
		if !reader.ReadFrame(buffer, uint64(length), func(frame *ttLibGo.Frame) bool {
			if frame.Type != frameTypes.H264 {
				return true
			}
			return decoder.DecodeFrame(frame, func(frame *ttLibGo.Frame) bool {
				if counter%5 == 0 {
					encoder.ForceNextFrameType(ttLibGo.X265EncoderFrameTypes.IDR)
				} else {
					encoder.ForceNextFrameType(ttLibGo.X265EncoderFrameTypes.P)
				}
				return encoder.EncodeFrame(frame, func(frame *ttLibGo.Frame) bool {
					fmt.Println(ttLibGo.H265.Cast(frame))
					return decoder2.DecodeFrame(frame, func(frame *ttLibGo.Frame) bool {
						return encoder2.EncodeFrame(frame, func(frame *ttLibGo.Frame) bool {
							counter++
							/*							return frame.GetBinaryBuffer(func(buf []byte) bool {
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
					})
				})
			})
		}) {
			t.Error("読み込み失敗")
			return
		}
		if length < 65536 {
			break
		}
	}
}
