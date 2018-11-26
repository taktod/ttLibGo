package ttLibGo

import (
	"github.com/taktod/ttLibGo/ttLibGo"
	"github.com/taktod/ttLibGo/ttLibGoFaac"
	"github.com/taktod/ttLibGo/ttLibGoFdkaac"
	"github.com/taktod/ttLibGo/ttLibGoFfmpeg"
	"github.com/taktod/ttLibGo/ttLibGoJpeg"
	"github.com/taktod/ttLibGo/ttLibGoLibyuv"
	"github.com/taktod/ttLibGo/ttLibGoMp3lame"
	"github.com/taktod/ttLibGo/ttLibGoOpenh264"
	"github.com/taktod/ttLibGo/ttLibGoOpus"
	"github.com/taktod/ttLibGo/ttLibGoPng"
	"github.com/taktod/ttLibGo/ttLibGoSoundtouch"
	"github.com/taktod/ttLibGo/ttLibGoSpeex"
	"github.com/taktod/ttLibGo/ttLibGoSpeexdsp"
	"github.com/taktod/ttLibGo/ttLibGoTheora"
	"github.com/taktod/ttLibGo/ttLibGoVorbis"
	"github.com/taktod/ttLibGo/ttLibGoX264"
	"github.com/taktod/ttLibGo/ttLibGoX265"
)

func install() {
	var reader ttLibGo.Reader
	defer reader.Close()
	var faacEncoder ttLibGoFaac.FaacEncoder
	defer faacEncoder.Close()
	var fdkaacEncoder ttLibGoFdkaac.FdkaacEncoder
	defer fdkaacEncoder.Close()
	var avcodecDecoder ttLibGoFfmpeg.AvcodecDecoder
	defer avcodecDecoder.Close()
	var jpegEncoder ttLibGoJpeg.JpegEncoder
	defer jpegEncoder.Close()
	var libyuvRotateResampler ttLibGoLibyuv.LibyuvRotateResampler
	defer libyuvRotateResampler.Close()
	var mp3lameDecoder ttLibGoMp3lame.Mp3lameDecoder
	defer mp3lameDecoder.Close()
	var openh264Decoder ttLibGoOpenh264.Openh264Decoder
	defer openh264Decoder.Close()
	var opusDecoder ttLibGoOpus.OpusDecoder
	defer opusDecoder.Close()
	var pngDecoder ttLibGoPng.PngDecoder
	defer pngDecoder.Close()
	var soundtouchResampler ttLibGoSoundtouch.SoundtouchResampler
	defer soundtouchResampler.Close()
	var speexDecoder ttLibGoSpeex.SpeexDecoder
	defer speexDecoder.Close()
	var speexdspResampler ttLibGoSpeexdsp.SpeexdspResampler
	defer speexdspResampler.Close()
	var theoraDecoder ttLibGoTheora.TheoraDecoder
	defer theoraDecoder.Close()
	var vorbisDecoder ttLibGoVorbis.VorbisDecoder
	defer vorbisDecoder.Close()
	var x264Encoder ttLibGoX264.X264Encoder
	defer x264Encoder.Close()
	var x265Encoder ttLibGoX265.X265Encoder
	defer x265Encoder.Close()
}
