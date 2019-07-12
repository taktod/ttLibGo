package test

import (
	"fmt"
	"os"
	"testing"

	"github.com/taktod/ttLibGo"
)

func encodeBody(t *testing.T, targetFile string, reader ttLibGo.IReader, targetType interface{}, decoder ttLibGo.IDecoder, encoder ttLibGo.IEncoder) {
	t.Log(fmt.Sprintf("encoder:%v", encoder))
	count := 0
	in, err := os.Open(targetFile)
	if err != nil {
		panic(err)
	}
	defer reader.Close()
	defer decoder.Close()
	defer encoder.Close()
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
				if frame.Type == targetType {
					decoder.DecodeFrame(frame, func(frame *ttLibGo.Frame) bool {
						return encoder.EncodeFrame(frame, func(frame *ttLibGo.Frame) bool {
							count++
							return true
						})
					})
				}
				return true
			}) {
			t.Errorf("読み込み失敗しました")
			return
		}
		if count > 5 {
			return
		}
		if length < 65536 {
			break
		}
	}
	t.Error(fmt.Sprintf("encodeデータ数が規定に達しませんでした。%d", count))
}

func TestFdkaacEncoder(t *testing.T) {
	encodeBody(t, os.Getenv("HOME")+"/tools/data/source/test.opus.mkv", ttLibGo.Readers.Mkv(), ttLibGo.FrameTypes.Opus, ttLibGo.Decoders.Opus(48000, 2), func() ttLibGo.IEncoder {
		fdk := ttLibGo.Encoders.Fdkaac(ttLibGo.FdkAacTypes.AotAacLC, 48000, 2, 48000)
		fdk.SetBitrate(96000)
		return fdk
	}())
}
func TestJpegEncoder(t *testing.T) {
	encodeBody(t, os.Getenv("HOME")+"/tools/data/source/test.h264.aac.flv", ttLibGo.Readers.Flv(), ttLibGo.FrameTypes.H264, ttLibGo.Decoders.AvcodecVideo(ttLibGo.FrameTypes.H264, 640, 360), func() ttLibGo.IEncoder {
		jpeg := ttLibGo.Encoders.Jpeg(640, 360, 90)
		jpeg.SetQuality(85)
		return jpeg
	}())
}

func TestOpenh264Encoder(t *testing.T) {
	type miia []map[interface{}]interface{}
	type mii map[interface{}]interface{}

	encodeBody(t, os.Getenv("HOME")+"/tools/data/source/test.h264.aac.flv", ttLibGo.Readers.Flv(), ttLibGo.FrameTypes.H264, ttLibGo.Decoders.AvcodecVideo(ttLibGo.FrameTypes.H264, 640, 360),
		func() ttLibGo.IEncoder {
			o := ttLibGo.Encoders.Openh264(640, 360, nil, nil)
			o.SetIDRInterval(15)
			return o
		}())
	encodeBody(t, os.Getenv("HOME")+"/tools/data/source/test.h264.aac.flv", ttLibGo.Readers.Flv(), ttLibGo.FrameTypes.H264, ttLibGo.Decoders.AvcodecVideo(ttLibGo.FrameTypes.H264, 640, 360),
		ttLibGo.Encoders.Openh264(640, 360, mii{
			ttLibGo.Openh264Params.IMinQp: 4,
			ttLibGo.Openh264Params.IMaxQp: 51,
		}, nil))
	encodeBody(t, os.Getenv("HOME")+"/tools/data/source/test.h264.aac.flv", ttLibGo.Readers.Flv(), ttLibGo.FrameTypes.H264, ttLibGo.Decoders.AvcodecVideo(ttLibGo.FrameTypes.H264, 640, 360),
		ttLibGo.Encoders.Openh264(640, 360, mii{
			ttLibGo.Openh264Params.IMinQp:         4,
			ttLibGo.Openh264Params.IMaxQp:         51,
			ttLibGo.Openh264Params.ITargetBitrate: 300000,
			ttLibGo.Openh264Params.IMaxBitrate:    310000,
		}, miia{
			mii{
				ttLibGo.Openh264SpatialParams.ISpatialBitrate:         300000,
				ttLibGo.Openh264SpatialParams.IMaxSpatialBitrate:      310000,
				ttLibGo.Openh264SpatialParams.SliceArgument.SliceMode: ttLibGo.Openh264SpatialParams.SliceArgument.SliceMode.SmSingleSLICE,
			},
		}))
}

func TestOpusEncoder(t *testing.T) {
	encodeBody(t, os.Getenv("HOME")+"/tools/data/source/test.opus.mkv", ttLibGo.Readers.Mkv(), ttLibGo.FrameTypes.Opus, ttLibGo.Decoders.Opus(48000, 2), func() ttLibGo.IEncoder {
		o := ttLibGo.Encoders.Opus(48000, 2, 480)
		first_bitrate := o.GetCodecControl(ttLibGo.OpusEncoderControls.GetBitrate)
		o.SetCodecControl(ttLibGo.OpusEncoderControls.SetBitrate, 96000)
		bitrate := o.GetCodecControl(ttLibGo.OpusEncoderControls.GetBitrate)
		if first_bitrate == bitrate || bitrate != 96000 {
			t.Error(fmt.Sprintf("first_bitrate:%d bitrate:%d", first_bitrate, bitrate))
		}
		return o
	}())
}

func TestSpeexEncoder(t *testing.T) {
	encodeBody(t, os.Getenv("HOME")+"/tools/data/source/test.flv1.speex.flv", ttLibGo.Readers.Flv(), ttLibGo.FrameTypes.Speex, ttLibGo.Decoders.Speex(16000, 1), ttLibGo.Encoders.Speex(16000, 1, 8))
}

func TestTheoraEncoder(t *testing.T) {
	encodeBody(t, os.Getenv("HOME")+"/tools/data/source/test.h264.aac.flv", ttLibGo.Readers.Flv(), ttLibGo.FrameTypes.H264, ttLibGo.Decoders.AvcodecVideo(ttLibGo.FrameTypes.H264, 640, 360), ttLibGo.Encoders.Theora(640, 360, 25, 320000, 15))
}

func TestVorbisEncoder(t *testing.T) {
	encodeBody(t, os.Getenv("HOME")+"/tools/data/source/test.opus.mkv", ttLibGo.Readers.Mkv(), ttLibGo.FrameTypes.Opus, ttLibGo.Decoders.Opus(48000, 2), ttLibGo.Encoders.Vorbis(48000, 2))
	encodeBody(t, os.Getenv("HOME")+"/tools/data/source/test.h264.aac.flv", ttLibGo.Readers.Flv(), ttLibGo.FrameTypes.Aac, ttLibGo.Decoders.AvcodecAudio(ttLibGo.FrameTypes.Aac, 44100, 2), ttLibGo.Encoders.Vorbis(44100, 2))
}

func TestX264Encoder(t *testing.T) {
	encodeBody(t, os.Getenv("HOME")+"/tools/data/source/test.h264.aac.flv", ttLibGo.Readers.Flv(), ttLibGo.FrameTypes.H264, ttLibGo.Decoders.AvcodecVideo(ttLibGo.FrameTypes.H264, 640, 360),
		ttLibGo.Encoders.X264(640, 360, "", "", "", nil))
	encodeBody(t, os.Getenv("HOME")+"/tools/data/source/test.h264.aac.flv", ttLibGo.Readers.Flv(), ttLibGo.FrameTypes.H264, ttLibGo.Decoders.AvcodecVideo(ttLibGo.FrameTypes.H264, 640, 360),
		ttLibGo.Encoders.X264(640, 360, "veryfast", "zerolatency", "main", nil))
	encodeBody(t, os.Getenv("HOME")+"/tools/data/source/test.h264.aac.flv", ttLibGo.Readers.Flv(), ttLibGo.FrameTypes.H264, ttLibGo.Decoders.AvcodecVideo(ttLibGo.FrameTypes.H264, 640, 360),
		ttLibGo.Encoders.X264(640, 360, "veryfast", "zerolatency", "main", map[string]string{
			"open-gop":    "1",
			"threads":     "1",
			"merange":     "16",
			"qcomp":       "0.6",
			"ip-factor":   "0.71",
			"bitrate":     "300",
			"qp":          "21",
			"crf":         "23",
			"crf-max":     "23",
			"fps":         "30/1",
			"keyint":      "150",
			"keyint-min":  "150",
			"bframes":     "3",
			"vbv-maxrate": "0",
			"vbv-bufsize": "1024",
			"qp-max":      "40",
			"qp-min":      "21",
			"qp-step":     "4",
		}))
}

func TestX265Encoder(t *testing.T) {
	encodeBody(t, os.Getenv("HOME")+"/tools/data/source/test.h264.aac.flv", ttLibGo.Readers.Flv(), ttLibGo.FrameTypes.H264, ttLibGo.Decoders.AvcodecVideo(ttLibGo.FrameTypes.H264, 640, 360),
		ttLibGo.Encoders.X265(640, 360, "", "", "", nil))
	encodeBody(t, os.Getenv("HOME")+"/tools/data/source/test.h264.aac.flv", ttLibGo.Readers.Flv(), ttLibGo.FrameTypes.H264, ttLibGo.Decoders.AvcodecVideo(ttLibGo.FrameTypes.H264, 640, 360),
		ttLibGo.Encoders.X265(640, 360, "veryfast", "zerolatency", "main", nil))
	encodeBody(t, os.Getenv("HOME")+"/tools/data/source/test.h264.aac.flv", ttLibGo.Readers.Flv(), ttLibGo.FrameTypes.H264, ttLibGo.Decoders.AvcodecVideo(ttLibGo.FrameTypes.H264, 640, 360),
		ttLibGo.Encoders.X265(640, 360, "veryfast", "zerolatency", "main", map[string]string{
			"bitrate": "300",
			"qpmax":   "40",
			"qpmin":   "21",
		}))
}
