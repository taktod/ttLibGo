package ttLibGo

import (
	"github.com/taktod/ttLibGo/ttLibGo"
	"github.com/taktod/ttLibGo/ttLibGoFaac"
	"github.com/taktod/ttLibGo/ttLibGoFdkaac"
	"github.com/taktod/ttLibGo/ttLibGoFfmpeg"
	"github.com/taktod/ttLibGo/ttLibGoJpeg"
	"github.com/taktod/ttLibGo/ttLibGoLibyuv"
	"github.com/taktod/ttLibGo/ttLibGoMp3lame"
	"github.com/taktod/ttLibGo/ttLibGoOpus"
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
	var opusDecoder ttLibGoOpus.OpusDecoder
	defer opusDecoder.Close()
}
