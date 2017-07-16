package ttLibGo

import (
	"github.com/taktod/ttLibGo/ttLibGo"
	"github.com/taktod/ttLibGo/ttLibGoFaac"
	"github.com/taktod/ttLibGo/ttLibGoFdkaac"
	"github.com/taktod/ttLibGo/ttLibGoFfmpeg"
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
}
