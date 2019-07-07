package test

import (
	"fmt"
	"os"
	"testing"

	"github.com/taktod/ttLibGo"
)

func readerBody(t *testing.T, targetFile string, reader ttLibGo.IReader, type1 interface{}, type2 interface{}) {
	t.Log(fmt.Sprintf("reader:%v %v %v", reader, type1, type2))
	type1Count := 0
	type2Count := 0
	in, err := os.Open(targetFile)
	if err != nil {
		panic(err)
	}
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
				if frame.Type == type1 {
					type1Count++
				}
				if frame.Type == type2 {
					type2Count++
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
	if type1Count == 0 {
		t.Error(fmt.Sprintf("%vが取得できませんでした", type1))
	}
	if type1 != type2 && type2Count == 0 {
		t.Error(fmt.Sprintf("%vが取得できませんでした", type2))
	}
}

func TestFlvReader(t *testing.T) {
	readerBody(t, os.Getenv("HOME")+"/tools/data/source/test.flv1.nellymoser.flv", ttLibGo.Readers.Flv(), ttLibGo.FrameTypes.Flv1, ttLibGo.FrameTypes.Nellymoser)
	readerBody(t, os.Getenv("HOME")+"/tools/data/source/test.flv1.pcm_alaw.flv", ttLibGo.Readers.Flv(), ttLibGo.FrameTypes.Flv1, ttLibGo.FrameTypes.PcmAlaw)
	readerBody(t, os.Getenv("HOME")+"/tools/data/source/test.flv1.pcm_mulaw.flv", ttLibGo.Readers.Flv(), ttLibGo.FrameTypes.Flv1, ttLibGo.FrameTypes.PcmMulaw)
	readerBody(t, os.Getenv("HOME")+"/tools/data/source/test.flv1.speex.flv", ttLibGo.Readers.Flv(), ttLibGo.FrameTypes.Flv1, ttLibGo.FrameTypes.Speex)
	readerBody(t, os.Getenv("HOME")+"/tools/data/source/test.h264.aac.flv", ttLibGo.Readers.Flv(), ttLibGo.FrameTypes.H264, ttLibGo.FrameTypes.Aac)
	readerBody(t, os.Getenv("HOME")+"/tools/data/source/test.vp6.mp3.flv", ttLibGo.Readers.Flv(), ttLibGo.FrameTypes.Vp6, ttLibGo.FrameTypes.Mp3)
}

func TestMkvReader(t *testing.T) {
	readerBody(t, os.Getenv("HOME")+"/tools/data/source/test.h264.aac.mkv", ttLibGo.Readers.Mkv(), ttLibGo.FrameTypes.H264, ttLibGo.FrameTypes.Aac)
	readerBody(t, os.Getenv("HOME")+"/tools/data/source/test.h265.mp3.mkv", ttLibGo.Readers.Mkv(), ttLibGo.FrameTypes.H265, ttLibGo.FrameTypes.Mp3)
	readerBody(t, os.Getenv("HOME")+"/tools/data/source/test.mjpeg.adpcmimawav.mkv", ttLibGo.Readers.Mkv(), ttLibGo.FrameTypes.Jpeg, ttLibGo.FrameTypes.AdpcmImaWav)
	readerBody(t, os.Getenv("HOME")+"/tools/data/source/test.h264.pcms16.mkv", ttLibGo.Readers.Mkv(), ttLibGo.FrameTypes.H264, ttLibGo.FrameTypes.PcmS16)
	readerBody(t, os.Getenv("HOME")+"/tools/data/source/test.opus.mkv", ttLibGo.Readers.Mkv(), ttLibGo.FrameTypes.Opus, ttLibGo.FrameTypes.Opus)
	readerBody(t, os.Getenv("HOME")+"/tools/data/source/test.png.mkv", ttLibGo.Readers.Mkv(), ttLibGo.FrameTypes.Png, ttLibGo.FrameTypes.Png)
	readerBody(t, os.Getenv("HOME")+"/tools/data/source/test.theora.speex.mkv", ttLibGo.Readers.Mkv(), ttLibGo.FrameTypes.Theora, ttLibGo.FrameTypes.Speex)
	readerBody(t, os.Getenv("HOME")+"/tools/data/source/test.vorbis.mkv", ttLibGo.Readers.Mkv(), ttLibGo.FrameTypes.Vorbis, ttLibGo.FrameTypes.Vorbis)
}

func TestWebmReader(t *testing.T) {
	readerBody(t, os.Getenv("HOME")+"/tools/data/source/test.vp8.vorbis.webm", ttLibGo.Readers.Webm(), ttLibGo.FrameTypes.Vp8, ttLibGo.FrameTypes.Vorbis)
	readerBody(t, os.Getenv("HOME")+"/tools/data/source/test.vp9.opus.webm", ttLibGo.Readers.Webm(), ttLibGo.FrameTypes.Vp9, ttLibGo.FrameTypes.Opus)
}

func TestMp4Reader(t *testing.T) {
	readerBody(t, os.Getenv("HOME")+"/tools/data/source/test.h264.aac.frag.mp4", ttLibGo.Readers.Mp4(), ttLibGo.FrameTypes.H264, ttLibGo.FrameTypes.Aac)
	readerBody(t, os.Getenv("HOME")+"/tools/data/source/test.h264.aac.mp4", ttLibGo.Readers.Mp4(), ttLibGo.FrameTypes.H264, ttLibGo.FrameTypes.Aac)
	readerBody(t, os.Getenv("HOME")+"/tools/data/source/test.h264.aac.mp4box.mp4", ttLibGo.Readers.Mp4(), ttLibGo.FrameTypes.H264, ttLibGo.FrameTypes.Aac)
	readerBody(t, os.Getenv("HOME")+"/tools/data/source/test.h264b.aac.mp4", ttLibGo.Readers.Mp4(), ttLibGo.FrameTypes.H264, ttLibGo.FrameTypes.Aac)
	readerBody(t, os.Getenv("HOME")+"/tools/data/source/test.h264b.mp3.boxed.mp4", ttLibGo.Readers.Mp4(), ttLibGo.FrameTypes.H264, ttLibGo.FrameTypes.Mp3)
	readerBody(t, os.Getenv("HOME")+"/tools/data/source/test.h264b.mp3.mp4", ttLibGo.Readers.Mp4(), ttLibGo.FrameTypes.H264, ttLibGo.FrameTypes.Mp3)
	readerBody(t, os.Getenv("HOME")+"/tools/data/source/test.h265.mp3.frag.mp4", ttLibGo.Readers.Mp4(), ttLibGo.FrameTypes.H265, ttLibGo.FrameTypes.Mp3)
	readerBody(t, os.Getenv("HOME")+"/tools/data/source/test.h265.mp3.mp4", ttLibGo.Readers.Mp4(), ttLibGo.FrameTypes.H265, ttLibGo.FrameTypes.Mp3)
	readerBody(t, os.Getenv("HOME")+"/tools/data/source/test.h265.mp3.mp4box.mp4", ttLibGo.Readers.Mp4(), ttLibGo.FrameTypes.H265, ttLibGo.FrameTypes.Mp3)
	readerBody(t, os.Getenv("HOME")+"/tools/data/source/test.mjpeg.vorbis.frag.mp4", ttLibGo.Readers.Mp4(), ttLibGo.FrameTypes.Jpeg, ttLibGo.FrameTypes.Vorbis)
	readerBody(t, os.Getenv("HOME")+"/tools/data/source/test.mjpeg.vorbis.mp4", ttLibGo.Readers.Mp4(), ttLibGo.FrameTypes.Jpeg, ttLibGo.FrameTypes.Vorbis)
	readerBody(t, os.Getenv("HOME")+"/tools/data/source/test.mjpeg.vorbis.mp4box.mp4", ttLibGo.Readers.Mp4(), ttLibGo.FrameTypes.Jpeg, ttLibGo.FrameTypes.Vorbis)
}

func TestMpegtsReader(t *testing.T) {
	readerBody(t, os.Getenv("HOME")+"/tools/data/source/test.h264.aac.10sec.ts", ttLibGo.Readers.Mpegts(), ttLibGo.FrameTypes.H264, ttLibGo.FrameTypes.Aac)
	readerBody(t, os.Getenv("HOME")+"/tools/data/source/test.h264.aac.ts", ttLibGo.Readers.Mpegts(), ttLibGo.FrameTypes.H264, ttLibGo.FrameTypes.Aac)
	readerBody(t, os.Getenv("HOME")+"/tools/data/source/test.h264.mp3.ts", ttLibGo.Readers.Mpegts(), ttLibGo.FrameTypes.H264, ttLibGo.FrameTypes.Mp3)
	readerBody(t, os.Getenv("HOME")+"/tools/data/source/test.h264b.aac.ts", ttLibGo.Readers.Mpegts(), ttLibGo.FrameTypes.H264, ttLibGo.FrameTypes.Aac)
	readerBody(t, os.Getenv("HOME")+"/tools/data/source/test.h264b.mp3.10sec.ts", ttLibGo.Readers.Mpegts(), ttLibGo.FrameTypes.H264, ttLibGo.FrameTypes.Mp3)
	readerBody(t, os.Getenv("HOME")+"/tools/data/source/test.h264b.mp3.ts", ttLibGo.Readers.Mpegts(), ttLibGo.FrameTypes.H264, ttLibGo.FrameTypes.Mp3)
}

/*
func TestHoge(t *testing.T) {
	readerBody(t, os.Getenv("HOME")+"/Documents/node/ttLibJsGyp/js/target.flv", ttLibGo.Readers.Flv(), ttLibGo.FrameTypes.H264, ttLibGo.FrameTypes.Aac)
}
*/
