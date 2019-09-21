package test

import (
	"os"
	"testing"

	"github.com/taktod/ttLibGo"
)

/*
func TestBgrFromBinary(t *testing.T) {
	// pngをdecodeして binary化 -> 復元 -> yuvにする -> jpegにする
	// という流れでいいか・・・
	in, err := os.Open(os.Getenv("HOME") + "/tools/data/source/test.png.mkv")
	if err != nil {
		panic(err)
	}
	//	counter := 0
	reader := ttLibGo.Readers.Mkv()
	defer reader.Close()
	decoder := ttLibGo.Decoders.Png()
	defer decoder.Close()
	var resampler ttLibGo.IResampler
	defer func() {
		if resampler != nil {
			resampler.Close()
		}
	}()
	encoder := ttLibGo.Encoders.Jpeg(320, 180, 90)
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
				return decoder.DecodeFrame(frame, func(frame *ttLibGo.Frame) bool {
					bgr := ttLibGo.Bgr.Cast(frame)
					return frame.GetBinaryBuffer(func(data []byte) bool {
						newbgr := ttLibGo.Bgr.FromBinary(data, uint64(len(data)), 1, bgr.Pts, bgr.Timebase, bgr.BgrType, bgr.Width, bgr.Height, bgr.WidthStride)
						defer newbgr.Close()
						if resampler == nil {
							resampler = ttLibGo.Resamplers.Swscale(newbgr.Type, bgr.BgrType, bgr.Width, bgr.Height,
								ttLibGo.FrameTypes.Yuv420, ttLibGo.Yuv420Types.Yuv420, 320, 180, ttLibGo.SwscaleModes.Bilinear)
						}
						return resampler.ResampleFrame(frame, func(frame *ttLibGo.Frame) bool {
							return encoder.EncodeFrame(frame, func(frame *ttLibGo.Frame) bool {
								// 正しく作られていることを確認できた。
								/ *								frame.GetBinaryBuffer(func(buf []byte) bool {
																file := fmt.Sprintf("out%d.jpeg", counter)
																out, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE, 0666)
																if err != nil {
																	panic(err)
																}
																out.Write(buf)
																counter++
																return true
															})* /
								return true
							})
						})
					})
				})
			}) {
			t.Errorf("読み込み失敗")
			return
		}
		if length < 65536 {
			break
		}
	}
}
*/
func TestFlv1FromBinary(t *testing.T) {
	in, err := os.Open(os.Getenv("HOME") + "/tools/data/source/test.flv1.nellymoser.flv")
	if err != nil {
		panic(err)
	}
	reader := ttLibGo.Readers.Flv()
	defer reader.Close()
	writer := ttLibGo.Writers.Flv(ttLibGo.FrameTypes.Flv1, ttLibGo.FrameTypes.Unknown)
	defer writer.Close()
	out, err := os.OpenFile("out.flv1.flv", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
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
			if frame.Type == ttLibGo.FrameTypes.Flv1 {
				return frame.GetBinaryBuffer(func(data []byte) bool {
					newFlv1 := ttLibGo.Flv1.FromBinary(data, uint64(len(data)), frame.ID, frame.Pts, frame.Timebase)
					defer newFlv1.Close()
					return writer.WriteFrame(newFlv1, func(data []byte) bool {
						out.Write(data)
						return true
					})
				})
			}
			return true
		}) {
			t.Errorf("読み込み失敗")
			return
		}
		if length < 65536 {
			break
		}
	}
}

func TestH264FromBinary(t *testing.T) {
	in, err := os.Open(os.Getenv("HOME") + "/tools/data/source/test.h264.aac.flv")
	if err != nil {
		panic(err)
	}
	reader := ttLibGo.Readers.Flv()
	defer reader.Close()
	writer := ttLibGo.Writers.Flv(ttLibGo.FrameTypes.H264, ttLibGo.FrameTypes.Unknown)
	defer writer.Close()
	out, err := os.OpenFile("out.h264.flv", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
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
			if frame.Type == ttLibGo.FrameTypes.H264 {
				return frame.GetBinaryBuffer(func(data []byte) bool {
					newH264 := ttLibGo.H264.FromBinary(data, uint64(len(data)), frame.ID, frame.Pts, frame.Timebase)
					defer newH264.Close()
					return writer.WriteFrame(newH264, func(data []byte) bool {
						out.Write(data)
						return true
					})
				})
			}
			return true
		}) {
			t.Errorf("読み込み失敗")
			return
		}
		if length < 65536 {
			break
		}
	}
}

func TestH265FromBinary(t *testing.T) {
	in, err := os.Open(os.Getenv("HOME") + "/tools/data/source/test.h265.mp3.mkv")
	if err != nil {
		panic(err)
	}
	reader := ttLibGo.Readers.Mkv()
	defer reader.Close()
	writer := ttLibGo.Writers.Mkv(ttLibGo.FrameTypes.H265)
	defer writer.Close()
	writer.Mode = ttLibGo.WriterModes.EnableDts
	out, err := os.OpenFile("out.h265.mkv", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
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
			if frame.Type == ttLibGo.FrameTypes.H265 {
				return frame.GetBinaryBuffer(func(data []byte) bool {
					newH265 := ttLibGo.H265.FromBinary(data, uint64(len(data)), frame.ID, frame.Pts, frame.Timebase)
					defer newH265.Close()
					return writer.WriteFrame(newH265, func(data []byte) bool {
						out.Write(data)
						return true
					})
				})
			}
			return true
		}) {
			t.Errorf("読み込み失敗")
			return
		}
		if length < 65536 {
			break
		}
	}
}

func TestJpegFromBinary(t *testing.T) {
	in, err := os.Open(os.Getenv("HOME") + "/tools/data/source/test.mjpeg.adpcmimawav.mkv")
	if err != nil {
		panic(err)
	}
	reader := ttLibGo.Readers.Mkv()
	defer reader.Close()
	writer := ttLibGo.Writers.Mkv(ttLibGo.FrameTypes.Jpeg)
	defer writer.Close()
	out, err := os.OpenFile("out.jpeg.mkv", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
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
			if frame.Type == ttLibGo.FrameTypes.Jpeg {
				return frame.GetBinaryBuffer(func(data []byte) bool {
					newJpeg := ttLibGo.Jpeg.FromBinary(data, uint64(len(data)), frame.ID, frame.Pts, frame.Timebase)
					defer newJpeg.Close()
					return writer.WriteFrame(newJpeg, func(data []byte) bool {
						out.Write(data)
						return true
					})
				})
			}
			return true
		}) {
			t.Errorf("読み込み失敗")
			return
		}
		if length < 65536 {
			break
		}
	}
}

func TestPngFromBinary(t *testing.T) {
	in, err := os.Open(os.Getenv("HOME") + "/tools/data/source/test.png.mkv")
	if err != nil {
		panic(err)
	}
	reader := ttLibGo.Readers.Mkv()
	defer reader.Close()
	decoder := ttLibGo.Decoders.Png()
	defer decoder.Close()
	resampler := ttLibGo.Resamplers.Image(ttLibGo.FrameTypes.Yuv420, ttLibGo.Yuv420Types.Yuv420)
	defer resampler.Close()
	encoder := ttLibGo.Encoders.X264(640, 360, "", "", "", nil)
	defer encoder.Close()
	writer := ttLibGo.Writers.Mkv(ttLibGo.FrameTypes.H264)
	defer writer.Close()
	out, err := os.OpenFile("out.png.mkv", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
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
			if frame.Type == ttLibGo.FrameTypes.Png {
				return frame.GetBinaryBuffer(func(data []byte) bool {
					newPng := ttLibGo.Png.FromBinary(data, uint64(len(data)), frame.ID, frame.Pts, frame.Timebase)
					defer newPng.Close()
					return decoder.DecodeFrame(newPng, func(frame *ttLibGo.Frame) bool {
						return resampler.ResampleFrame(frame, func(frame *ttLibGo.Frame) bool {
							return encoder.EncodeFrame(frame, func(frame *ttLibGo.Frame) bool {
								return writer.WriteFrame(frame, func(data []byte) bool {
									out.Write(data)
									return true
								})
							})
						})
					})
				})
			}
			return true
		}) {
			t.Errorf("読み込み失敗")
			return
		}
		if length < 65536 {
			break
		}
	}
}

func TestTheoraFromBinary(t *testing.T) {
	in, err := os.Open(os.Getenv("HOME") + "/tools/data/source/test.theora.speex.mkv")
	if err != nil {
		panic(err)
	}
	reader := ttLibGo.Readers.Mkv()
	defer reader.Close()
	writer := ttLibGo.Writers.Mkv(ttLibGo.FrameTypes.Theora)
	defer writer.Close()
	out, err := os.OpenFile("out.theora.mkv", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
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
			if frame.Type == ttLibGo.FrameTypes.Theora {
				theora := ttLibGo.Theora.Cast(frame)
				return theora.GetBinaryBuffer(func(data []byte) bool {
					newTheora := ttLibGo.Theora.FromBinary(data, uint64(len(data)), theora.ID, theora.Pts, theora.Timebase, theora.Width, theora.Height)
					defer newTheora.Close()
					return writer.WriteFrame(newTheora, func(data []byte) bool {
						out.Write(data)
						return true
					})
				})
			}
			return true
		}) {
			t.Errorf("読み込み失敗")
			return
		}
		if length < 65536 {
			break
		}
	}
}

func TestVp6FromBinary(t *testing.T) {
	in, err := os.Open(os.Getenv("HOME") + "/tools/data/source/test.vp6.mp3.flv")
	if err != nil {
		panic(err)
	}
	reader := ttLibGo.Readers.Flv()
	defer reader.Close()
	writer := ttLibGo.Writers.Flv(ttLibGo.FrameTypes.Vp6, ttLibGo.FrameTypes.Unknown)
	defer writer.Close()
	out, err := os.OpenFile("out.vp6.flv", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
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
			if frame.Type == ttLibGo.FrameTypes.Vp6 {
				vp6 := ttLibGo.Vp6.Cast(frame)
				return frame.GetBinaryBuffer(func(data []byte) bool {
					newVp6 := ttLibGo.Vp6.FromBinary(data, uint64(len(data)), vp6.ID, vp6.Pts, vp6.Timebase, vp6.Width, vp6.Height, 0)
					defer newVp6.Close()
					return writer.WriteFrame(newVp6, func(data []byte) bool {
						out.Write(data)
						return true
					})
				})
			}
			return true
		}) {
			t.Errorf("読み込み失敗")
			return
		}
		if length < 65536 {
			break
		}
	}
}

func TestVp8FromBinary(t *testing.T) {
	in, err := os.Open(os.Getenv("HOME") + "/tools/data/source/test.vp8.vorbis.webm")
	if err != nil {
		panic(err)
	}
	reader := ttLibGo.Readers.Webm()
	defer reader.Close()
	writer := ttLibGo.Writers.Webm(ttLibGo.FrameTypes.Vp8)
	defer writer.Close()
	out, err := os.OpenFile("out.vp8.webm", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
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
			if frame.Type == ttLibGo.FrameTypes.Vp8 {
				vp8 := ttLibGo.Vp8.Cast(frame)
				return frame.GetBinaryBuffer(func(data []byte) bool {
					newVp8 := ttLibGo.Vp8.FromBinary(data, uint64(len(data)), vp8.ID, vp8.Pts, vp8.Timebase, vp8.Width, vp8.Height)
					defer newVp8.Close()
					return writer.WriteFrame(newVp8, func(data []byte) bool {
						out.Write(data)
						return true
					})
				})
			}
			return true
		}) {
			t.Errorf("読み込み失敗")
			return
		}
		if length < 65536 {
			break
		}
	}
}

func TestVp9FromBinary(t *testing.T) {
	in, err := os.Open(os.Getenv("HOME") + "/tools/data/source/test.vp9.opus.webm")
	if err != nil {
		panic(err)
	}
	reader := ttLibGo.Readers.Webm()
	defer reader.Close()
	writer := ttLibGo.Writers.Webm(ttLibGo.FrameTypes.Vp9)
	defer writer.Close()
	out, err := os.OpenFile("out.vp9.webm", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
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
			if frame.Type == ttLibGo.FrameTypes.Vp9 {
				vp9 := ttLibGo.Vp9.Cast(frame)
				return frame.GetBinaryBuffer(func(data []byte) bool {
					newVp9 := ttLibGo.Vp9.FromBinary(data, uint64(len(data)), vp9.ID, vp9.Pts, vp9.Timebase, vp9.Width, vp9.Height)
					defer newVp9.Close()
					return writer.WriteFrame(newVp9, func(data []byte) bool {
						out.Write(data)
						return true
					})
				})
			}
			return true
		}) {
			t.Errorf("読み込み失敗")
			return
		}
		if length < 65536 {
			break
		}
	}
}

func TestYuv420FromBinary(t *testing.T) {
	in, err := os.Open(os.Getenv("HOME") + "/tools/data/source/test.h264.aac.flv")
	if err != nil {
		panic(err)
	}
	reader := ttLibGo.Readers.Flv()
	defer reader.Close()
	decoder := ttLibGo.Decoders.AvcodecVideo(ttLibGo.FrameTypes.H264, 640, 360)
	defer decoder.Close()
	encoder := ttLibGo.Encoders.X264(640, 360, "", "", "", nil)
	defer encoder.Close()
	writer := ttLibGo.Writers.Mkv(ttLibGo.FrameTypes.H264)
	defer writer.Close()
	out, err := os.OpenFile("out.yuv.h264.mkv", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
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
			if frame.Type == ttLibGo.FrameTypes.H264 {
				return decoder.DecodeFrame(frame, func(frame *ttLibGo.Frame) bool {
					yuv := ttLibGo.Yuv420.Cast(frame)
					return yuv.GetBinaryBuffer(func(data []byte) bool {
						newYuv := ttLibGo.Yuv420.FromBinary(data, uint64(len(data)), 1, yuv.Pts, yuv.Timebase, yuv.Yuv420Type, yuv.Width, yuv.Height)
						defer newYuv.Close()
						return encoder.EncodeFrame(newYuv, func(frame *ttLibGo.Frame) bool {
							return writer.WriteFrame(frame, func(data []byte) bool {
								out.Write(data)
								return true
							})
						})
					})
				})
			}
			return true
		}) {
			t.Errorf("読み込み失敗")
			return
		}
		if length < 65536 {
			break
		}
	}
}

func TestAacRawFromBinary(t *testing.T) {
	in, err := os.Open(os.Getenv("HOME") + "/tools/data/source/test.h264.aac.flv")
	if err != nil {
		panic(err)
	}
	var dsi *ttLibGo.Frame
	defer func() {
		if dsi != nil {
			dsi.Close()
		}
	}()
	reader := ttLibGo.Readers.Flv()
	defer reader.Close()
	writer := ttLibGo.Writers.Flv(ttLibGo.FrameTypes.Unknown, ttLibGo.FrameTypes.Aac)
	defer writer.Close()
	out, err := os.OpenFile("out.aac.flv", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
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
			if frame.Type == ttLibGo.FrameTypes.Aac {
				return frame.GetBinaryBuffer(func(data []byte) bool {
					newAac := ttLibGo.Aac.FromBinary(data, uint64(len(data)), frame.ID, frame.Pts, frame.Timebase, dsi)
					if newAac != nil {
						aac := ttLibGo.Aac.Cast(newAac)
						if aac.AacType == ttLibGo.AacTypes.Dsi {
							dsi = newAac
						} else {
							defer newAac.Close()
						}
						return writer.WriteFrame(aac, func(data []byte) bool {
							out.Write(data)
							return true
						})
					}
					return true
				})
			}
			return true
		}) {
			t.Errorf("読み込み失敗")
			return
		}
		if length < 65536 {
			break
		}
	}
}

func TestAacAdtsFromBinary(t *testing.T) {
	in, err := os.Open(os.Getenv("HOME") + "/tools/data/source/test.h264.aac.ts")
	if err != nil {
		panic(err)
	}
	var dsi *ttLibGo.Frame
	defer func() {
		if dsi != nil {
			dsi.Close()
		}
	}()
	reader := ttLibGo.Readers.Mpegts()
	defer reader.Close()
	writer := ttLibGo.Writers.Mpegts(ttLibGo.FrameTypes.Aac)
	defer writer.Close()
	out, err := os.OpenFile("out.aac.ts", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
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
			if frame.Type == ttLibGo.FrameTypes.Aac {
				return frame.GetBinaryBuffer(func(data []byte) bool {
					newAac := ttLibGo.Aac.FromBinary(data, uint64(len(data)), frame.ID-1, frame.Pts, frame.Timebase, dsi)
					if newAac != nil {
						aac := ttLibGo.Aac.Cast(newAac)
						if aac.AacType == ttLibGo.AacTypes.Dsi {
							dsi = newAac
						} else {
							defer newAac.Close()
						}
						return writer.WriteFrame(aac, func(data []byte) bool {
							out.Write(data)
							return true
						})
					}
					return true
				})
			}
			return true
		}) {
			t.Errorf("読み込み失敗")
			return
		}
		if length < 65536 {
			break
		}
	}
}

func TestAdpcmImaWavFromBinary(t *testing.T) {
	in, err := os.Open(os.Getenv("HOME") + "/tools/data/source/test.mjpeg.adpcmimawav.mkv")
	if err != nil {
		panic(err)
	}
	reader := ttLibGo.Readers.Mkv()
	defer reader.Close()
	writer := ttLibGo.Writers.Mkv(ttLibGo.FrameTypes.AdpcmImaWav)
	defer writer.Close()
	out, err := os.OpenFile("out.adpcmimawav.mkv", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
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
			if frame.Type == ttLibGo.FrameTypes.AdpcmImaWav {
				adpcmImaWav := ttLibGo.AdpcmImaWav.Cast(frame)
				return frame.GetBinaryBuffer(func(data []byte) bool {
					newAdpcmImaWav := ttLibGo.AdpcmImaWav.FromBinary(data, uint64(len(data)), 1, adpcmImaWav.Pts, adpcmImaWav.Timebase, adpcmImaWav.SampleRate, adpcmImaWav.SampleNum, adpcmImaWav.ChannelNum)
					defer newAdpcmImaWav.Close()
					return writer.WriteFrame(newAdpcmImaWav, func(data []byte) bool {
						out.Write(data)
						return true
					})
				})
			}
			return true
		}) {
			t.Errorf("読み込み失敗")
			return
		}
		if length < 65536 {
			break
		}
	}
}

func TestMp3FromBinary(t *testing.T) {
	in, err := os.Open(os.Getenv("HOME") + "/tools/data/source/test.vp6.mp3.flv")
	if err != nil {
		panic(err)
	}
	reader := ttLibGo.Readers.Flv()
	defer reader.Close()
	writer := ttLibGo.Writers.Mkv(ttLibGo.FrameTypes.Mp3)
	defer writer.Close()
	out, err := os.OpenFile("out.mp3.mkv", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
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
			if frame.Type == ttLibGo.FrameTypes.Mp3 {
				return frame.GetBinaryBuffer(func(data []byte) bool {
					newMp3 := ttLibGo.Mp3.FromBinary(data, uint64(len(data)), 1, frame.Pts, frame.Timebase)
					defer newMp3.Close()
					return writer.WriteFrame(newMp3, func(data []byte) bool {
						out.Write(data)
						return true
					})
				})
			}
			return true
		}) {
			t.Errorf("読み込み失敗")
			return
		}
		if length < 65536 {
			break
		}
	}
}

func TestNellymoserFromBinary(t *testing.T) {
	in, err := os.Open(os.Getenv("HOME") + "/tools/data/source/test.flv1.nellymoser.flv")
	if err != nil {
		panic(err)
	}
	reader := ttLibGo.Readers.Flv()
	defer reader.Close()
	writer := ttLibGo.Writers.Flv(ttLibGo.FrameTypes.Unknown, ttLibGo.FrameTypes.Nellymoser)
	defer writer.Close()
	out, err := os.OpenFile("out.nellymoser.flv", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
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
			if frame.Type == ttLibGo.FrameTypes.Nellymoser {
				nellymoser := ttLibGo.Nellymoser.Cast(frame)
				return frame.GetBinaryBuffer(func(data []byte) bool {
					newNellymoser := ttLibGo.Nellymoser.FromBinary(data, uint64(len(data)), nellymoser.ID, nellymoser.Pts, nellymoser.Timebase, nellymoser.SampleRate, nellymoser.SampleNum, nellymoser.ChannelNum)
					defer newNellymoser.Close()
					return writer.WriteFrame(newNellymoser, func(data []byte) bool {
						out.Write(data)
						return true
					})
				})
			}
			return true
		}) {
			t.Errorf("読み込み失敗")
			return
		}
		if length < 65536 {
			break
		}
	}
}

func TestOpusFromBinary(t *testing.T) {
	in, err := os.Open(os.Getenv("HOME") + "/tools/data/source/test.vp9.opus.webm")
	if err != nil {
		panic(err)
	}
	reader := ttLibGo.Readers.Webm()
	defer reader.Close()
	writer := ttLibGo.Writers.Webm(ttLibGo.FrameTypes.Opus)
	defer writer.Close()
	out, err := os.OpenFile("out.opus.webm", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
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
			if frame.Type == ttLibGo.FrameTypes.Opus {
				return frame.GetBinaryBuffer(func(data []byte) bool {
					newOpus := ttLibGo.Opus.FromBinary(data, uint64(len(data)), 1, frame.Pts, frame.Timebase)
					defer newOpus.Close()
					return writer.WriteFrame(newOpus, func(data []byte) bool {
						out.Write(data)
						return true
					})
				})
			}
			return true
		}) {
			t.Errorf("読み込み失敗")
			return
		}
		if length < 65536 {
			break
		}
	}
}

func TestPcmAlawFromBinary(t *testing.T) {
	in, err := os.Open(os.Getenv("HOME") + "/tools/data/source/test.flv1.pcm_alaw.flv")
	if err != nil {
		panic(err)
	}
	reader := ttLibGo.Readers.Flv()
	defer reader.Close()
	writer := ttLibGo.Writers.Flv(ttLibGo.FrameTypes.Unknown, ttLibGo.FrameTypes.PcmAlaw)
	defer writer.Close()
	out, err := os.OpenFile("out.pcm_alaw.flv", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
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
			if frame.Type == ttLibGo.FrameTypes.PcmAlaw {
				pcmAlaw := ttLibGo.PcmAlaw.Cast(frame)
				return frame.GetBinaryBuffer(func(data []byte) bool {
					newPcmAlaw := ttLibGo.PcmAlaw.FromBinary(data, uint64(len(data)), pcmAlaw.ID, pcmAlaw.Pts, pcmAlaw.Timebase, pcmAlaw.SampleRate, pcmAlaw.SampleNum, pcmAlaw.ChannelNum)
					defer newPcmAlaw.Close()
					return writer.WriteFrame(newPcmAlaw, func(data []byte) bool {
						out.Write(data)
						return true
					})
				})
			}
			return true
		}) {
			t.Errorf("読み込み失敗")
			return
		}
		if length < 65536 {
			break
		}
	}
}

func TestPcmF32FromBinary(t *testing.T) {
	// なぜかクオリティが悪い (pcmF32のfromBinaryが原因ではなさそう)
	in, err := os.Open(os.Getenv("HOME") + "/tools/data/source/test.vp8.vorbis.webm")
	if err != nil {
		panic(err)
	}
	reader := ttLibGo.Readers.Webm()
	defer reader.Close()
	decoder := ttLibGo.Decoders.Vorbis()
	defer decoder.Close()
	resampler := ttLibGo.Resamplers.Swresample(ttLibGo.FrameTypes.PcmF32, ttLibGo.PcmF32Types.Interleave, 44100, 2,
		ttLibGo.FrameTypes.PcmS16, ttLibGo.PcmS16Types.LittleEndian, 44100, 2)
	defer resampler.Close()
	encoder := ttLibGo.Encoders.Fdkaac(ttLibGo.FdkAacTypes.AotAacLC, 44100, 2, 96000)
	defer encoder.Close()
	writer := ttLibGo.Writers.Flv(ttLibGo.FrameTypes.Unknown, ttLibGo.FrameTypes.Aac)
	defer writer.Close()
	out, err := os.OpenFile("out.pcmf32_aac.flv", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
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
			if frame.Type == ttLibGo.FrameTypes.Vorbis {
				return decoder.DecodeFrame(frame, func(frame *ttLibGo.Frame) bool {
					pcmF32 := ttLibGo.PcmF32.Cast(frame)
					return pcmF32.GetBinaryBuffer(func(data []byte) bool {
						length := len(data)
						newPcmF32 := ttLibGo.PcmF32.FromBinary(data, uint64(length), pcmF32.ID, pcmF32.Pts, pcmF32.Timebase, pcmF32.PcmF32Type, pcmF32.SampleRate, pcmF32.SampleNum, pcmF32.ChannelNum, 0, 0)
						defer newPcmF32.Close()
						return resampler.ResampleFrame(newPcmF32, func(frame *ttLibGo.Frame) bool {
							return encoder.EncodeFrame(frame, func(frame *ttLibGo.Frame) bool {
								return writer.WriteFrame(frame, func(data []byte) bool {
									out.Write(data)
									return true
								})
							})
						})
					})
				})
			}
			return true
		}) {
			t.Errorf("読み込み失敗")
			return
		}
		if length < 65536 {
			break
		}
	}
}

func TestPcmMulawFromBinary(t *testing.T) {
	in, err := os.Open(os.Getenv("HOME") + "/tools/data/source/test.flv1.pcm_mulaw.flv")
	if err != nil {
		panic(err)
	}
	reader := ttLibGo.Readers.Flv()
	defer reader.Close()
	writer := ttLibGo.Writers.Flv(ttLibGo.FrameTypes.Unknown, ttLibGo.FrameTypes.PcmMulaw)
	defer writer.Close()
	out, err := os.OpenFile("out.pcm_mulaw.flv", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
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
			if frame.Type == ttLibGo.FrameTypes.PcmMulaw {
				pcmMulaw := ttLibGo.PcmMulaw.Cast(frame)
				return frame.GetBinaryBuffer(func(data []byte) bool {
					newPcmMulaw := ttLibGo.PcmMulaw.FromBinary(data, uint64(len(data)), pcmMulaw.ID, pcmMulaw.Pts, pcmMulaw.Timebase, pcmMulaw.SampleRate, pcmMulaw.SampleNum, pcmMulaw.ChannelNum)
					defer newPcmMulaw.Close()
					return writer.WriteFrame(newPcmMulaw, func(data []byte) bool {
						out.Write(data)
						return true
					})
				})
			}
			return true
		}) {
			t.Errorf("読み込み失敗")
			return
		}
		if length < 65536 {
			break
		}
	}
}

func TestPcmS16FromBinary(t *testing.T) {
	in, err := os.Open(os.Getenv("HOME") + "/tools/data/source/test.vp9.opus.webm")
	if err != nil {
		panic(err)
	}
	reader := ttLibGo.Readers.Webm()
	defer reader.Close()
	decoder := ttLibGo.Decoders.Opus(48000, 2)
	defer decoder.Close()
	encoder := ttLibGo.Encoders.Fdkaac(ttLibGo.FdkAacTypes.AotAacLC, 48000, 2, 96000)
	defer encoder.Close()
	writer := ttLibGo.Writers.Flv(ttLibGo.FrameTypes.Unknown, ttLibGo.FrameTypes.Aac)
	defer writer.Close()
	out, err := os.OpenFile("out.pcms16_aac.flv", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
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
			if frame.Type == ttLibGo.FrameTypes.Opus {
				return decoder.DecodeFrame(frame, func(frame *ttLibGo.Frame) bool {
					pcmS16 := ttLibGo.PcmS16.Cast(frame)
					return pcmS16.GetBinaryBuffer(func(data []byte) bool {
						length := len(data)
						newPcmS16 := ttLibGo.PcmS16.FromBinary(data, uint64(length), pcmS16.ID, pcmS16.Pts, pcmS16.Timebase, pcmS16.PcmS16Type, pcmS16.SampleRate, pcmS16.SampleNum, pcmS16.ChannelNum, 0, 0)
						defer newPcmS16.Close()
						return encoder.EncodeFrame(newPcmS16, func(frame *ttLibGo.Frame) bool {
							return writer.WriteFrame(frame, func(data []byte) bool {
								out.Write(data)
								return true
							})
						})
					})
				})
			}
			return true
		}) {
			t.Errorf("読み込み失敗")
			return
		}
		if length < 65536 {
			break
		}
	}
}

func TestSpeexFromBinary(t *testing.T) {
	in, err := os.Open(os.Getenv("HOME") + "/tools/data/source/test.flv1.speex.flv")
	if err != nil {
		panic(err)
	}
	reader := ttLibGo.Readers.Flv()
	defer reader.Close()
	writer := ttLibGo.Writers.Flv(ttLibGo.FrameTypes.Unknown, ttLibGo.FrameTypes.Speex)
	defer writer.Close()
	out, err := os.OpenFile("out.speex.flv", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
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
			if frame.Type == ttLibGo.FrameTypes.Speex {
				speex := ttLibGo.Speex.Cast(frame)
				return speex.GetBinaryBuffer(func(data []byte) bool {
					newSpeex := ttLibGo.Speex.FromBinary(data, uint64(len(data)), speex.ID, speex.Pts, speex.Timebase, speex.SampleRate, speex.SampleNum, speex.ChannelNum)
					return writer.WriteFrame(newSpeex, func(data []byte) bool {
						out.Write(data)
						return true
					})
				})
			}
			return true
		}) {
			t.Errorf("読み込み失敗")
			return
		}
		if length < 65536 {
			break
		}
	}
}

func TestVorbisFromBinary(t *testing.T) {
	in, err := os.Open(os.Getenv("HOME") + "/tools/data/source/test.vp8.vorbis.webm")
	if err != nil {
		panic(err)
	}
	reader := ttLibGo.Readers.Mkv()
	defer reader.Close()
	writer := ttLibGo.Writers.Mkv(ttLibGo.FrameTypes.Vorbis)
	defer writer.Close()
	out, err := os.OpenFile("out.vorbis.mkv", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
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
			if frame.Type == ttLibGo.FrameTypes.Vorbis {
				vorbis := ttLibGo.Vorbis.Cast(frame)
				return vorbis.GetBinaryBuffer(func(data []byte) bool {
					newVorbis := ttLibGo.Vorbis.FromBinary(data, uint64(len(data)), 1, vorbis.Pts, vorbis.Timebase, vorbis.SampleRate, vorbis.SampleNum, vorbis.ChannelNum)
					return writer.WriteFrame(newVorbis, func(data []byte) bool {
						out.Write(data)
						return true
					})
				})
			}
			return true
		}) {
			t.Errorf("読み込み失敗")
			return
		}
		if length < 65536 {
			break
		}
	}
}
