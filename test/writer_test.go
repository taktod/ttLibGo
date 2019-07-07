package test

import (
	"fmt"
	"os"
	"testing"

	"github.com/taktod/ttLibGo"
)

func writerBody(t *testing.T, targetFile string, reader ttLibGo.IReader, writer ttLibGo.IWriter) {
	t.Log(fmt.Sprintf("writer:%v", writer))
	writeCount := 0
	in, err := os.Open(targetFile)
	if err != nil {
		panic(err)
	}
	defer reader.Close()
	defer writer.Close()
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
				return writer.WriteFrame(frame, func(data []byte) bool {
					writeCount += len(data)
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
	if writeCount < 100 {
		t.Error(fmt.Sprintf("出力が小さすぎます:size:%d", writeCount))
	}
}

func TestFlvWriter(t *testing.T) {
	writerBody(t, os.Getenv("HOME")+"/tools/data/source/test.flv1.nellymoser.flv", ttLibGo.Readers.Flv(), ttLibGo.Writers.Flv(ttLibGo.FrameTypes.Flv1, ttLibGo.FrameTypes.Nellymoser))
	writerBody(t, os.Getenv("HOME")+"/tools/data/source/test.flv1.pcm_alaw.flv", ttLibGo.Readers.Flv(), ttLibGo.Writers.Flv(ttLibGo.FrameTypes.Flv1, ttLibGo.FrameTypes.PcmAlaw))
	writerBody(t, os.Getenv("HOME")+"/tools/data/source/test.flv1.pcm_mulaw.flv", ttLibGo.Readers.Flv(), ttLibGo.Writers.Flv(ttLibGo.FrameTypes.Flv1, ttLibGo.FrameTypes.PcmMulaw))
	writerBody(t, os.Getenv("HOME")+"/tools/data/source/test.flv1.speex.flv", ttLibGo.Readers.Flv(), ttLibGo.Writers.Flv(ttLibGo.FrameTypes.Flv1, ttLibGo.FrameTypes.Speex))
	writerBody(t, os.Getenv("HOME")+"/tools/data/source/test.h264.aac.flv", ttLibGo.Readers.Flv(), ttLibGo.Writers.Flv(ttLibGo.FrameTypes.H264, ttLibGo.FrameTypes.Aac))
	writerBody(t, os.Getenv("HOME")+"/tools/data/source/test.vp6.mp3.flv", ttLibGo.Readers.Flv(), ttLibGo.Writers.Flv(ttLibGo.FrameTypes.Vp6, ttLibGo.FrameTypes.Mp3))
}

func TestMkvWriter(t *testing.T) {
	writerBody(t, os.Getenv("HOME")+"/tools/data/source/test.h264.aac.mkv", ttLibGo.Readers.Mkv(), ttLibGo.Writers.Mkv(ttLibGo.FrameTypes.H264, ttLibGo.FrameTypes.Aac))
	writerBody(t, os.Getenv("HOME")+"/tools/data/source/test.h265.mp3.mkv", ttLibGo.Readers.Mkv(), func() ttLibGo.IWriter {
		writer := ttLibGo.Writers.Mkv(ttLibGo.FrameTypes.H265, ttLibGo.FrameTypes.Mp3)
		writer.Mode = ttLibGo.WriterModes.EnableDts
		return writer
	}())
	writerBody(t, os.Getenv("HOME")+"/tools/data/source/test.mjpeg.adpcmimawav.mkv", ttLibGo.Readers.Mkv(), ttLibGo.Writers.Mkv(ttLibGo.FrameTypes.Jpeg, ttLibGo.FrameTypes.AdpcmImaWav))
	writerBody(t, os.Getenv("HOME")+"/tools/data/source/test.h264.pcms16.mkv", ttLibGo.Readers.Mkv(), ttLibGo.Writers.Mkv(ttLibGo.FrameTypes.H264, ttLibGo.FrameTypes.PcmS16))
	writerBody(t, os.Getenv("HOME")+"/tools/data/source/test.opus.mkv", ttLibGo.Readers.Mkv(), ttLibGo.Writers.Mkv(ttLibGo.FrameTypes.Opus))
	writerBody(t, os.Getenv("HOME")+"/tools/data/source/test.png.mkv", ttLibGo.Readers.Mkv(), ttLibGo.Writers.Mkv(ttLibGo.FrameTypes.Png))
	writerBody(t, os.Getenv("HOME")+"/tools/data/source/test.theora.speex.mkv", ttLibGo.Readers.Mkv(), ttLibGo.Writers.Mkv(ttLibGo.FrameTypes.Theora, ttLibGo.FrameTypes.Speex))
	writerBody(t, os.Getenv("HOME")+"/tools/data/source/test.vorbis.mkv", ttLibGo.Readers.Mkv(), ttLibGo.Writers.Mkv(ttLibGo.FrameTypes.Vorbis))
}

func TestWebmWriter(t *testing.T) {
	writerBody(t, os.Getenv("HOME")+"/tools/data/source/test.vp8.vorbis.webm", ttLibGo.Readers.Webm(), ttLibGo.Writers.Webm(ttLibGo.FrameTypes.Vp8, ttLibGo.FrameTypes.Vorbis))
	writerBody(t, os.Getenv("HOME")+"/tools/data/source/test.vp9.opus.webm", ttLibGo.Readers.Webm(), ttLibGo.Writers.Webm(ttLibGo.FrameTypes.Vp9, ttLibGo.FrameTypes.Opus))
}

func TestMp4Writer(t *testing.T) {
	writerBody(t, os.Getenv("HOME")+"/tools/data/source/test.h264.aac.frag.mp4", ttLibGo.Readers.Mp4(), ttLibGo.Writers.Mp4(ttLibGo.FrameTypes.H264, ttLibGo.FrameTypes.Aac))
	writerBody(t, os.Getenv("HOME")+"/tools/data/source/test.h264.aac.mp4", ttLibGo.Readers.Mp4(), ttLibGo.Writers.Mp4(ttLibGo.FrameTypes.H264, ttLibGo.FrameTypes.Aac))
	writerBody(t, os.Getenv("HOME")+"/tools/data/source/test.h264.aac.mp4box.mp4", ttLibGo.Readers.Mp4(), ttLibGo.Writers.Mp4(ttLibGo.FrameTypes.H264, ttLibGo.FrameTypes.Aac))
	writerBody(t, os.Getenv("HOME")+"/tools/data/source/test.h264b.aac.mp4", ttLibGo.Readers.Mp4(), func() ttLibGo.IWriter {
		writer := ttLibGo.Writers.Mp4(ttLibGo.FrameTypes.H264, ttLibGo.FrameTypes.Aac)
		writer.Mode = ttLibGo.WriterModes.EnableDts
		return writer
	}())
	writerBody(t, os.Getenv("HOME")+"/tools/data/source/test.h264b.mp3.boxed.mp4", ttLibGo.Readers.Mp4(), func() ttLibGo.IWriter {
		writer := ttLibGo.Writers.Mp4(ttLibGo.FrameTypes.H264, ttLibGo.FrameTypes.Mp3)
		writer.Mode = ttLibGo.WriterModes.EnableDts
		return writer
	}())
	writerBody(t, os.Getenv("HOME")+"/tools/data/source/test.h264b.mp3.mp4", ttLibGo.Readers.Mp4(), func() ttLibGo.IWriter {
		writer := ttLibGo.Writers.Mp4(ttLibGo.FrameTypes.H264, ttLibGo.FrameTypes.Mp3)
		writer.Mode = ttLibGo.WriterModes.EnableDts
		return writer
	}())
	writerBody(t, os.Getenv("HOME")+"/tools/data/source/test.h265.mp3.frag.mp4", ttLibGo.Readers.Mp4(), func() ttLibGo.IWriter {
		writer := ttLibGo.Writers.Mp4(ttLibGo.FrameTypes.H265, ttLibGo.FrameTypes.Mp3)
		writer.Mode = ttLibGo.WriterModes.EnableDts
		return writer
	}())
	writerBody(t, os.Getenv("HOME")+"/tools/data/source/test.h265.mp3.mp4", ttLibGo.Readers.Mp4(), func() ttLibGo.IWriter {
		writer := ttLibGo.Writers.Mp4(ttLibGo.FrameTypes.H265, ttLibGo.FrameTypes.Mp3)
		writer.Mode = ttLibGo.WriterModes.EnableDts
		return writer
	}())
	writerBody(t, os.Getenv("HOME")+"/tools/data/source/test.h265.mp3.mp4box.mp4", ttLibGo.Readers.Mp4(), func() ttLibGo.IWriter {
		writer := ttLibGo.Writers.Mp4(ttLibGo.FrameTypes.H265, ttLibGo.FrameTypes.Mp3)
		writer.Mode = ttLibGo.WriterModes.EnableDts
		return writer
	}())
	writerBody(t, os.Getenv("HOME")+"/tools/data/source/test.mjpeg.vorbis.frag.mp4", ttLibGo.Readers.Mp4(), ttLibGo.Writers.Mp4(ttLibGo.FrameTypes.Jpeg, ttLibGo.FrameTypes.Vorbis))
	writerBody(t, os.Getenv("HOME")+"/tools/data/source/test.mjpeg.vorbis.mp4", ttLibGo.Readers.Mp4(), ttLibGo.Writers.Mp4(ttLibGo.FrameTypes.Jpeg, ttLibGo.FrameTypes.Vorbis))
	writerBody(t, os.Getenv("HOME")+"/tools/data/source/test.mjpeg.vorbis.mp4box.mp4", ttLibGo.Readers.Mp4(), ttLibGo.Writers.Mp4(ttLibGo.FrameTypes.Jpeg, ttLibGo.FrameTypes.Vorbis))
}

func TestMpegtsWriter(t *testing.T) {
	writerBody(t, os.Getenv("HOME")+"/tools/data/source/test.h264.aac.10sec.ts", ttLibGo.Readers.Mpegts(), ttLibGo.Writers.Mpegts(ttLibGo.FrameTypes.H264, ttLibGo.FrameTypes.Aac))
	writerBody(t, os.Getenv("HOME")+"/tools/data/source/test.h264.aac.ts", ttLibGo.Readers.Mpegts(), ttLibGo.Writers.Mpegts(ttLibGo.FrameTypes.H264, ttLibGo.FrameTypes.Aac))
	writerBody(t, os.Getenv("HOME")+"/tools/data/source/test.h264.mp3.ts", ttLibGo.Readers.Mpegts(), ttLibGo.Writers.Mpegts(ttLibGo.FrameTypes.H264, ttLibGo.FrameTypes.Mp3))
	writerBody(t, os.Getenv("HOME")+"/tools/data/source/test.h264b.aac.ts", ttLibGo.Readers.Mpegts(), func() ttLibGo.IWriter {
		writer := ttLibGo.Writers.Mpegts(ttLibGo.FrameTypes.H264, ttLibGo.FrameTypes.Aac)
		writer.Mode = ttLibGo.WriterModes.EnableDts
		return writer
	}())
	writerBody(t, os.Getenv("HOME")+"/tools/data/source/test.h264b.mp3.10sec.ts", ttLibGo.Readers.Mpegts(), func() ttLibGo.IWriter {
		writer := ttLibGo.Writers.Mpegts(ttLibGo.FrameTypes.H264, ttLibGo.FrameTypes.Mp3)
		writer.Mode = ttLibGo.WriterModes.EnableDts
		return writer
	}())
	writerBody(t, os.Getenv("HOME")+"/tools/data/source/test.h264b.mp3.ts", ttLibGo.Readers.Mpegts(), func() ttLibGo.IWriter {
		writer := ttLibGo.Writers.Mpegts(ttLibGo.FrameTypes.H264, ttLibGo.FrameTypes.Mp3)
		writer.Mode = ttLibGo.WriterModes.EnableDts
		return writer
	}())
}
