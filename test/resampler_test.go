package test

import (
	"fmt"
	"os"
	"testing"

	"github.com/taktod/ttLibGo"
)

func resampleBody(t *testing.T, targetFile string, reader ttLibGo.IReader, targetType interface{}, decoder ttLibGo.IDecoder, resampler ttLibGo.IResampler) {
	t.Log(fmt.Sprintf("resampler:%v", resampler))
	count := 0
	in, err := os.Open(targetFile)
	if err != nil {
		panic(err)
	}
	defer reader.Close()
	defer decoder.Close()
	defer resampler.Close()
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
						return resampler.ResampleFrame(frame, func(frame *ttLibGo.Frame) bool {
							//							fmt.Println(frame)
							//							fmt.Println(ttLibGo.Video.Cast(frame))
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
	t.Error(fmt.Sprintf("decodeデータ数が規定に達しませんでした。%d", count))
}

func TestAudioResampler(t *testing.T) {
	resampleBody(t, os.Getenv("HOME")+"/tools/data/source/test.opus.mkv", ttLibGo.Readers.Mkv(), ttLibGo.FrameTypes.Opus, ttLibGo.Decoders.AvcodecAudio(ttLibGo.FrameTypes.Opus, 48000, 2),
		ttLibGo.Resamplers.Audio(ttLibGo.FrameTypes.PcmS16, ttLibGo.PcmS16Types.LittleEndian))
	resampleBody(t, os.Getenv("HOME")+"/tools/data/source/test.opus.mkv", ttLibGo.Readers.Mkv(), ttLibGo.FrameTypes.Opus, ttLibGo.Decoders.AvcodecAudio(ttLibGo.FrameTypes.Opus, 48000, 2),
		ttLibGo.Resamplers.Audio(ttLibGo.FrameTypes.PcmS16, ttLibGo.PcmS16Types.LittleEndianPlanar))
	resampleBody(t, os.Getenv("HOME")+"/tools/data/source/test.opus.mkv", ttLibGo.Readers.Mkv(), ttLibGo.FrameTypes.Opus, ttLibGo.Decoders.AvcodecAudio(ttLibGo.FrameTypes.Opus, 48000, 2),
		ttLibGo.Resamplers.Audio(ttLibGo.FrameTypes.PcmF32, ttLibGo.PcmF32Types.Interleave))
	resampleBody(t, os.Getenv("HOME")+"/tools/data/source/test.opus.mkv", ttLibGo.Readers.Mkv(), ttLibGo.FrameTypes.Opus, ttLibGo.Decoders.AvcodecAudio(ttLibGo.FrameTypes.Opus, 48000, 2),
		ttLibGo.Resamplers.Audio(ttLibGo.FrameTypes.PcmF32, ttLibGo.PcmF32Types.Planar))

	resampleBody(t, os.Getenv("HOME")+"/tools/data/source/test.vorbis.mkv", ttLibGo.Readers.Mkv(), ttLibGo.FrameTypes.Vorbis, ttLibGo.Decoders.AvcodecAudio(ttLibGo.FrameTypes.Vorbis, 44100, 2),
		ttLibGo.Resamplers.Audio(ttLibGo.FrameTypes.PcmS16, ttLibGo.PcmS16Types.LittleEndian))
	resampleBody(t, os.Getenv("HOME")+"/tools/data/source/test.vorbis.mkv", ttLibGo.Readers.Mkv(), ttLibGo.FrameTypes.Vorbis, ttLibGo.Decoders.AvcodecAudio(ttLibGo.FrameTypes.Vorbis, 44100, 2),
		ttLibGo.Resamplers.Audio(ttLibGo.FrameTypes.PcmS16, ttLibGo.PcmS16Types.LittleEndianPlanar))
	resampleBody(t, os.Getenv("HOME")+"/tools/data/source/test.vorbis.mkv", ttLibGo.Readers.Mkv(), ttLibGo.FrameTypes.Vorbis, ttLibGo.Decoders.AvcodecAudio(ttLibGo.FrameTypes.Vorbis, 44100, 2),
		ttLibGo.Resamplers.Audio(ttLibGo.FrameTypes.PcmF32, ttLibGo.PcmF32Types.Interleave))
	resampleBody(t, os.Getenv("HOME")+"/tools/data/source/test.vorbis.mkv", ttLibGo.Readers.Mkv(), ttLibGo.FrameTypes.Vorbis, ttLibGo.Decoders.AvcodecAudio(ttLibGo.FrameTypes.Vorbis, 44100, 2),
		ttLibGo.Resamplers.Audio(ttLibGo.FrameTypes.PcmF32, ttLibGo.PcmF32Types.Planar))
}

func TestImageResampler(t *testing.T) {
	resampleBody(t, os.Getenv("HOME")+"/tools/data/source/test.h264.aac.flv", ttLibGo.Readers.Flv(), ttLibGo.FrameTypes.H264, ttLibGo.Decoders.AvcodecVideo(ttLibGo.FrameTypes.H264, 640, 360),
		ttLibGo.Resamplers.Image(ttLibGo.FrameTypes.Bgr, ttLibGo.BgrTypes.Bgra))
	resampleBody(t, os.Getenv("HOME")+"/tools/data/source/test.h264.aac.flv", ttLibGo.Readers.Flv(), ttLibGo.FrameTypes.H264, ttLibGo.Decoders.AvcodecVideo(ttLibGo.FrameTypes.H264, 640, 360),
		ttLibGo.Resamplers.Image(ttLibGo.FrameTypes.Bgr, ttLibGo.BgrTypes.Bgr))
	resampleBody(t, os.Getenv("HOME")+"/tools/data/source/test.h264.aac.flv", ttLibGo.Readers.Flv(), ttLibGo.FrameTypes.H264, ttLibGo.Decoders.AvcodecVideo(ttLibGo.FrameTypes.H264, 640, 360),
		ttLibGo.Resamplers.Image(ttLibGo.FrameTypes.Bgr, ttLibGo.BgrTypes.Abgr))
	resampleBody(t, os.Getenv("HOME")+"/tools/data/source/test.h264.aac.flv", ttLibGo.Readers.Flv(), ttLibGo.FrameTypes.H264, ttLibGo.Decoders.AvcodecVideo(ttLibGo.FrameTypes.H264, 640, 360),
		ttLibGo.Resamplers.Image(ttLibGo.FrameTypes.Bgr, ttLibGo.BgrTypes.Rgba))
	resampleBody(t, os.Getenv("HOME")+"/tools/data/source/test.h264.aac.flv", ttLibGo.Readers.Flv(), ttLibGo.FrameTypes.H264, ttLibGo.Decoders.AvcodecVideo(ttLibGo.FrameTypes.H264, 640, 360),
		ttLibGo.Resamplers.Image(ttLibGo.FrameTypes.Bgr, ttLibGo.BgrTypes.Rgb))
	resampleBody(t, os.Getenv("HOME")+"/tools/data/source/test.h264.aac.flv", ttLibGo.Readers.Flv(), ttLibGo.FrameTypes.H264, ttLibGo.Decoders.AvcodecVideo(ttLibGo.FrameTypes.H264, 640, 360),
		ttLibGo.Resamplers.Image(ttLibGo.FrameTypes.Bgr, ttLibGo.BgrTypes.Argb))
	resampleBody(t, os.Getenv("HOME")+"/tools/data/source/test.png.mkv", ttLibGo.Readers.Mkv(), ttLibGo.FrameTypes.Png, ttLibGo.Decoders.Png(),
		ttLibGo.Resamplers.Image(ttLibGo.FrameTypes.Yuv420, ttLibGo.Yuv420Types.Yuv420))
	resampleBody(t, os.Getenv("HOME")+"/tools/data/source/test.png.mkv", ttLibGo.Readers.Mkv(), ttLibGo.FrameTypes.Png, ttLibGo.Decoders.Png(),
		ttLibGo.Resamplers.Image(ttLibGo.FrameTypes.Yuv420, ttLibGo.Yuv420Types.Yuv420SemiPlanar))
	resampleBody(t, os.Getenv("HOME")+"/tools/data/source/test.png.mkv", ttLibGo.Readers.Mkv(), ttLibGo.FrameTypes.Png, ttLibGo.Decoders.Png(),
		ttLibGo.Resamplers.Image(ttLibGo.FrameTypes.Yuv420, ttLibGo.Yuv420Types.Yvu420))
	resampleBody(t, os.Getenv("HOME")+"/tools/data/source/test.png.mkv", ttLibGo.Readers.Mkv(), ttLibGo.FrameTypes.Png, ttLibGo.Decoders.Png(),
		ttLibGo.Resamplers.Image(ttLibGo.FrameTypes.Yuv420, ttLibGo.Yuv420Types.Yvu420SemiPlanar))
}

func TestResizeResampler(t *testing.T) {
	resampleBody(t, os.Getenv("HOME")+"/tools/data/source/test.h264.aac.flv", ttLibGo.Readers.Flv(), ttLibGo.FrameTypes.H264, ttLibGo.Decoders.AvcodecVideo(ttLibGo.FrameTypes.H264, 640, 360),
		ttLibGo.Resamplers.Resize(320, 180, false))
	resampleBody(t, os.Getenv("HOME")+"/tools/data/source/test.png.mkv", ttLibGo.Readers.Mkv(), ttLibGo.FrameTypes.Png, ttLibGo.Decoders.Png(),
		ttLibGo.Resamplers.Resize(480, 320, false))
}

func TestSoundtouchResampler(t *testing.T) {
	resampleBody(t, os.Getenv("HOME")+"/tools/data/source/test.opus.mkv", ttLibGo.Readers.Mkv(), ttLibGo.FrameTypes.Opus, ttLibGo.Decoders.Opus(48000, 2),
		ttLibGo.Resamplers.Soundtouch(48000, 2))

	resampleBody(t, os.Getenv("HOME")+"/tools/data/source/test.vorbis.mkv", ttLibGo.Readers.Mkv(), ttLibGo.FrameTypes.Vorbis, ttLibGo.Decoders.AvcodecAudio(ttLibGo.FrameTypes.Vorbis, 44100, 2),
		ttLibGo.Resamplers.Soundtouch(44100, 2))
}

func TestSpeexdspResampler(t *testing.T) {
	resampleBody(t, os.Getenv("HOME")+"/tools/data/source/test.opus.mkv", ttLibGo.Readers.Mkv(), ttLibGo.FrameTypes.Opus, ttLibGo.Decoders.Opus(48000, 2),
		ttLibGo.Resamplers.Speexdsp(48000, 44100, 2, 8))
}

func TestSwresampleResampler(t *testing.T) {
	resampleBody(t, os.Getenv("HOME")+"/tools/data/source/test.opus.mkv", ttLibGo.Readers.Mkv(), ttLibGo.FrameTypes.Opus, ttLibGo.Decoders.AvcodecAudio(ttLibGo.FrameTypes.Opus, 48000, 2),
		ttLibGo.Resamplers.Swresample(
			ttLibGo.FrameTypes.PcmF32, ttLibGo.PcmF32Types.Interleave, 48000, 2,
			ttLibGo.FrameTypes.PcmS16, ttLibGo.PcmS16Types.LittleEndian, 44100, 2))
}

func TestSwscaleResampler(t *testing.T) {
	resampleBody(t, os.Getenv("HOME")+"/tools/data/source/test.h264.aac.flv", ttLibGo.Readers.Flv(), ttLibGo.FrameTypes.H264, ttLibGo.Decoders.AvcodecVideo(ttLibGo.FrameTypes.H264, 640, 360),
		ttLibGo.Resamplers.Swscale(ttLibGo.FrameTypes.Yuv420, ttLibGo.Yuv420Types.Yuv420, 640, 360,
			ttLibGo.FrameTypes.Yuv420, ttLibGo.Yuv420Types.Yvu420, 320, 180, ttLibGo.SwscaleModes.Bicubic))
	resampleBody(t, os.Getenv("HOME")+"/tools/data/source/test.h264.aac.flv", ttLibGo.Readers.Flv(), ttLibGo.FrameTypes.H264, ttLibGo.Decoders.AvcodecVideo(ttLibGo.FrameTypes.H264, 640, 360),
		ttLibGo.Resamplers.Swscale(ttLibGo.FrameTypes.Yuv420, ttLibGo.Yuv420Types.Yuv420SemiPlanar, 640, 360,
			ttLibGo.FrameTypes.Yuv420, ttLibGo.Yuv420Types.Yvu420, 320, 180, ttLibGo.SwscaleModes.Bilinear))
	resampleBody(t, os.Getenv("HOME")+"/tools/data/source/test.h264.aac.flv", ttLibGo.Readers.Flv(), ttLibGo.FrameTypes.H264, ttLibGo.Decoders.AvcodecVideo(ttLibGo.FrameTypes.H264, 640, 360),
		ttLibGo.Resamplers.Swscale(ttLibGo.FrameTypes.Yuv420, ttLibGo.Yuv420Types.Yuv420, 640, 360,
			ttLibGo.FrameTypes.Yuv420, ttLibGo.Yuv420Types.Yvu420, 320, 180, ttLibGo.SwscaleModes.X))
	resampleBody(t, os.Getenv("HOME")+"/tools/data/source/test.h264.aac.flv", ttLibGo.Readers.Flv(), ttLibGo.FrameTypes.H264, ttLibGo.Decoders.AvcodecVideo(ttLibGo.FrameTypes.H264, 640, 360),
		ttLibGo.Resamplers.Swscale(ttLibGo.FrameTypes.Yuv420, ttLibGo.Yuv420Types.Yuv420SemiPlanar, 640, 360,
			ttLibGo.FrameTypes.Yuv420, ttLibGo.Yuv420Types.Yvu420, 320, 180, ttLibGo.SwscaleModes.Lanczos))
}
