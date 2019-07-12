package test

import (
	"fmt"
	"os"
	"testing"

	"github.com/taktod/ttLibGo"
)

func decodeBody(t *testing.T, targetFile string, reader ttLibGo.IReader, targetType interface{}, decoder ttLibGo.IDecoder) {
	t.Log(fmt.Sprintf("decoder:%v targetFrame:%v", decoder, targetType))
	count := 0
	in, err := os.Open(targetFile)
	if err != nil {
		panic(err)
	}
	defer reader.Close()
	defer decoder.Close()
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
						count++
						return true
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

func TestAvcodecVideoDecoder(t *testing.T) {
	decodeBody(t, os.Getenv("HOME")+"/tools/data/source/test.h264.aac.flv", ttLibGo.Readers.Flv(), ttLibGo.FrameTypes.H264, ttLibGo.Decoders.AvcodecVideo(ttLibGo.FrameTypes.H264, 640, 360))
	decodeBody(t, os.Getenv("HOME")+"/tools/data/source/test.flv1.speex.flv", ttLibGo.Readers.Flv(), ttLibGo.FrameTypes.Flv1, ttLibGo.Decoders.AvcodecVideo(ttLibGo.FrameTypes.Flv1, 640, 360))
	decodeBody(t, os.Getenv("HOME")+"/tools/data/source/test.vp6.mp3.flv", ttLibGo.Readers.Flv(), ttLibGo.FrameTypes.Vp6, ttLibGo.Decoders.AvcodecVideo(ttLibGo.FrameTypes.Vp6, 640, 360))
	decodeBody(t, os.Getenv("HOME")+"/tools/data/source/test.mjpeg.adpcmimawav.mkv", ttLibGo.Readers.Mkv(), ttLibGo.FrameTypes.Jpeg, ttLibGo.Decoders.AvcodecVideo(ttLibGo.FrameTypes.Jpeg, 640, 360))
	decodeBody(t, os.Getenv("HOME")+"/tools/data/source/test.png.mkv", ttLibGo.Readers.Mkv(), ttLibGo.FrameTypes.Png, ttLibGo.Decoders.AvcodecVideo(ttLibGo.FrameTypes.Png, 640, 360))
	decodeBody(t, os.Getenv("HOME")+"/tools/data/source/test.h265.mp3.mkv", ttLibGo.Readers.Mkv(), ttLibGo.FrameTypes.H265, ttLibGo.Decoders.AvcodecVideo(ttLibGo.FrameTypes.H265, 640, 360))
	decodeBody(t, os.Getenv("HOME")+"/tools/data/source/test.h264b.aac.mp4", ttLibGo.Readers.Mp4(), ttLibGo.FrameTypes.H264, ttLibGo.Decoders.AvcodecVideo(ttLibGo.FrameTypes.H264, 640, 360))
	decodeBody(t, os.Getenv("HOME")+"/tools/data/source/test.theora.speex.mkv", ttLibGo.Readers.Mkv(), ttLibGo.FrameTypes.Theora, ttLibGo.Decoders.AvcodecVideo(ttLibGo.FrameTypes.Theora, 640, 360))
	decodeBody(t, os.Getenv("HOME")+"/tools/data/source/test.vp8.vorbis.webm", ttLibGo.Readers.Webm(), ttLibGo.FrameTypes.Vp8, ttLibGo.Decoders.AvcodecVideo(ttLibGo.FrameTypes.Vp8, 640, 360))
	decodeBody(t, os.Getenv("HOME")+"/tools/data/source/test.vp9.opus.webm", ttLibGo.Readers.Webm(), ttLibGo.FrameTypes.Vp9, ttLibGo.Decoders.AvcodecVideo(ttLibGo.FrameTypes.Vp9, 640, 360))
}

func TestAvcodecAudioDecoder(t *testing.T) {
	decodeBody(t, os.Getenv("HOME")+"/tools/data/source/test.h264.aac.flv", ttLibGo.Readers.Flv(), ttLibGo.FrameTypes.Aac, ttLibGo.Decoders.AvcodecAudio(ttLibGo.FrameTypes.Aac, 44100, 2))
	decodeBody(t, os.Getenv("HOME")+"/tools/data/source/test.flv1.nellymoser.flv", ttLibGo.Readers.Flv(), ttLibGo.FrameTypes.Nellymoser, ttLibGo.Decoders.AvcodecAudio(ttLibGo.FrameTypes.Nellymoser, 44100, 1))
	decodeBody(t, os.Getenv("HOME")+"/tools/data/source/test.flv1.pcm_alaw.flv", ttLibGo.Readers.Flv(), ttLibGo.FrameTypes.PcmAlaw, ttLibGo.Decoders.AvcodecAudio(ttLibGo.FrameTypes.PcmAlaw, 8000, 1))
	decodeBody(t, os.Getenv("HOME")+"/tools/data/source/test.flv1.pcm_mulaw.flv", ttLibGo.Readers.Flv(), ttLibGo.FrameTypes.PcmMulaw, ttLibGo.Decoders.AvcodecAudio(ttLibGo.FrameTypes.PcmMulaw, 8000, 1))
	decodeBody(t, os.Getenv("HOME")+"/tools/data/source/test.flv1.speex.flv", ttLibGo.Readers.Flv(), ttLibGo.FrameTypes.Speex, ttLibGo.Decoders.AvcodecAudio(ttLibGo.FrameTypes.Speex, 16000, 1))
	decodeBody(t, os.Getenv("HOME")+"/tools/data/source/test.vp6.mp3.flv", ttLibGo.Readers.Flv(), ttLibGo.FrameTypes.Mp3, ttLibGo.Decoders.AvcodecAudio(ttLibGo.FrameTypes.Mp3, 44100, 2))
	decodeBody(t, os.Getenv("HOME")+"/tools/data/source/test.mjpeg.adpcmimawav.mkv", ttLibGo.Readers.Mkv(), ttLibGo.FrameTypes.AdpcmImaWav, ttLibGo.Decoders.AvcodecAudio(ttLibGo.FrameTypes.AdpcmImaWav, 44100, 2))
	decodeBody(t, os.Getenv("HOME")+"/tools/data/source/test.opus.mkv", ttLibGo.Readers.Mkv(), ttLibGo.FrameTypes.Opus, ttLibGo.Decoders.AvcodecAudio(ttLibGo.FrameTypes.Opus, 48000, 2))
	decodeBody(t, os.Getenv("HOME")+"/tools/data/source/test.vorbis.mkv", ttLibGo.Readers.Mkv(), ttLibGo.FrameTypes.Vorbis, ttLibGo.Decoders.AvcodecAudio(ttLibGo.FrameTypes.Vorbis, 44100, 2))
}

func TestJpegDecoder(t *testing.T) {
	decodeBody(t, os.Getenv("HOME")+"/tools/data/source/test.mjpeg.adpcmimawav.mkv", ttLibGo.Readers.Mkv(), ttLibGo.FrameTypes.Jpeg, ttLibGo.Decoders.Jpeg())
}

func TestOpenh264Decoder(t *testing.T) {
	decodeBody(t, os.Getenv("HOME")+"/tools/data/source/test.h264.aac.flv", ttLibGo.Readers.Flv(), ttLibGo.FrameTypes.H264, ttLibGo.Decoders.Openh264())
}

func TestOpusDecoder(t *testing.T) {
	decodeBody(t, os.Getenv("HOME")+"/tools/data/source/test.opus.mkv", ttLibGo.Readers.Mkv(), ttLibGo.FrameTypes.Opus, func() ttLibGo.IDecoder {
		op := ttLibGo.Decoders.Opus(48000, 2)
		op.SetCodecControl(ttLibGo.OpusDecoderControls.SetPhaseInversionDisabled, true)
		sampleRate := op.GetCodecControl(ttLibGo.OpusDecoderControls.GetSampleRate)
		if sampleRate != 48000 {
			t.Error("failed to ref sampleRate")
		}
		fmt.Println(sampleRate)
		return op
	}())
	//	decodeBody(t, os.Getenv("HOME")+"/tools/data/source/test.opus.mkv", ttLibGo.Readers.Mkv(), ttLibGo.FrameTypes.Opus, ttLibGo.Decoders.Opus(48000, 2))
}

func TestPngDecoder(t *testing.T) {
	decodeBody(t, os.Getenv("HOME")+"/tools/data/source/test.png.mkv", ttLibGo.Readers.Mkv(), ttLibGo.FrameTypes.Png, ttLibGo.Decoders.Png())
}

func TestSpeexDecoder(t *testing.T) {
	decodeBody(t, os.Getenv("HOME")+"/tools/data/source/test.flv1.speex.flv", ttLibGo.Readers.Flv(), ttLibGo.FrameTypes.Speex, ttLibGo.Decoders.Speex(16000, 1))
}

func TestTheoraDecoder(t *testing.T) {
	decodeBody(t, os.Getenv("HOME")+"/tools/data/source/test.theora.speex.mkv", ttLibGo.Readers.Mkv(), ttLibGo.FrameTypes.Theora, ttLibGo.Decoders.Theora())
}

func TestVorbisDecoder(t *testing.T) {
	decodeBody(t, os.Getenv("HOME")+"/tools/data/source/test.vorbis.mkv", ttLibGo.Readers.Mkv(), ttLibGo.FrameTypes.Vorbis, ttLibGo.Decoders.Vorbis())
}
