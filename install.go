package ttLibGo

import (
	"github.com/taktod/ttLibGo/ttLibGo"
	"github.com/taktod/ttLibGo/ttLibGoFfmpeg"
)

func install() {
	var reader ttLibGo.Reader
	defer reader.Close()
	var avcodecDecoder ttLibGoFfmpeg.AvcodecDecoder
	defer avcodecDecoder.Close()
}
