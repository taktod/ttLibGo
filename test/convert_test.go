package test

import (
	"encoding/hex"
	"math"
	"testing"

	"github.com/taktod/ttLibGo"
)

func checkEq(t *testing.T, targetHex string, frame ttLibGo.IFrame) {
	targetArray := []byte(targetHex)
	frame.GetBinaryBuffer(func(data []byte) bool {
		l := len(data)
		l2 := len(targetArray)
		if l > l2 {
			l = l2
		}
		err := false
		for i := 0; i < l; i++ {
			if math.Abs(float64(data[i])-float64(targetArray[i])) > 10 {
				err = true
			}
		}
		if err {
			t.Error(hex.Dump(targetArray))
			t.Error(hex.Dump(data))
		}
		return true
	})
}

func scaleYuvPlanar(t *testing.T, targetCall func(dst ttLibGo.IFrame, src ttLibGo.IFrame) bool) {
	t.Log("scaleYuvPlanar")
	src := ttLibGo.Yuv420.FromBinary([]byte(""+
		"\x00\x00\x80\x80"+
		"\x00\x00\x40\x20"+
		"\xFF\xFF"+
		"\x80\xF0"),
		12, 1, 0, 1000, ttLibGo.Yuv420Types.Yuv420, 4, 2)
	defer src.Close()
	dst := ttLibGo.Yuv420.FromBinary([]byte(""+
		"\x00\x00\x00\x00\x00\x00\x00\x00"+
		"\x00\x00\x00\x00\x00\x00\x00\x00"+
		"\x00\x00\x00\x00"+
		"\x00\x00\x00\x00"),
		24, 1, 0, 1000, ttLibGo.Yuv420Types.Yuv420, 8, 2)
	defer dst.Close()
	targetCall(dst, src)
	checkEq(t, ""+
		"\x00\x00\x00\x20\x60\x80\x80\x80"+
		"\x00\x00\x00\x10\x30\x38\x28\x20"+
		"\xFF\xFF\xFF\xFF"+
		"\x80\x80\xF0\xF0",
		dst)
}

func scaleYuvPlanarClip(t *testing.T, targetCall func(dst ttLibGo.IFrame, src ttLibGo.IFrame) bool) {
	t.Log("scaleYuvPlanarClip")
	src := ttLibGo.Yuv420.FromBinary([]byte(""+
		"\x00\x00\x80\x80"+
		"\x00\x00\x40\x20"+
		"\xFF\xFF"+
		"\x80\xF0"),
		12, 1, 0, 1000, ttLibGo.Yuv420Types.Yuv420, 4, 2)
	defer src.Close()
	dst := ttLibGo.Yuv420.FromBinary([]byte(""+
		"\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF"+
		"\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF"+
		"\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF"+
		"\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF"+
		"\xFF\xFF\xFF\xFF\xFF"+
		"\xFF\xFF\xFF\xFF\xFF"+
		"\xFF\xFF\xFF\xFF\xFF"+
		"\xFF\xFF\xFF\xFF\xFF"),
		60, 1, 0, 1000, ttLibGo.Yuv420Types.Yuv420, 10, 4)
	defer dst.Close()
	ydst := ttLibGo.Yuv420.Cast(dst)
	ydst.Width = 8
	ydst.Height = 2
	yDataPos := ydst.YData
	ydst.YData = uintptr(uint64(ydst.YData) + uint64(ydst.YStride) + 1)
	targetCall(ydst, src)
	ydst.Width = 10
	ydst.Height = 4
	ydst.YData = yDataPos
	checkEq(t, ""+
		"\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF"+
		"\xFF\x00\x00\x00\x20\x60\x80\x80\x80\xFF"+
		"\xFF\x00\x00\x00\x10\x30\x38\x28\x20\xFF"+
		"\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF"+
		"\xFF\xFF\xFF\xFF\xFF"+
		"\xFF\xFF\xFF\xFF\xFF"+
		"\x80\x80\xF0\xF0\xFF"+
		"\xFF\xFF\xFF\xFF\xFF",
		ydst)
}

func TestSwscale(t *testing.T) {
	{
		swscale := ttLibGo.Resamplers.Swscale(ttLibGo.FrameTypes.Yuv420, ttLibGo.Yuv420Types.Yuv420, 4, 2,
			ttLibGo.FrameTypes.Yuv420, ttLibGo.Yuv420Types.Yuv420, 8, 2, ttLibGo.SwscaleModes.FastBilinear)
		defer swscale.Close()
		scaleYuvPlanar(t, swscale.Resample)
	}
	{
		swscale := ttLibGo.Resamplers.Swscale(ttLibGo.FrameTypes.Yuv420, ttLibGo.Yuv420Types.Yuv420, 4, 2,
			ttLibGo.FrameTypes.Yuv420, ttLibGo.Yuv420Types.Yuv420, 8, 2, ttLibGo.SwscaleModes.FastBilinear)
		defer swscale.Close()
		scaleYuvPlanarClip(t, swscale.Resample)
	}
}

func TestLibyuv(t *testing.T) {
	libyuv := ttLibGo.Resamplers.Libyuv()
	defer libyuv.Close()
	{
		scaleYuvPlanar(t, func(dst ttLibGo.IFrame, src ttLibGo.IFrame) bool {
			return libyuv.Resize(dst, src, ttLibGo.LibyuvModes.Linear, ttLibGo.LibyuvModes.None)
		})
	}
	{
		scaleYuvPlanarClip(t, func(dst ttLibGo.IFrame, src ttLibGo.IFrame) bool {
			return libyuv.Resize(dst, src, ttLibGo.LibyuvModes.Linear, ttLibGo.LibyuvModes.None)
		})
	}
}
