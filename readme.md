# ttLibGO

# 概要

go言語でttLibCの変換ライブラリを使う動作

これを使うとflvやmkvといったコンテナデータの読み込み
内部に定義されているframeデータを取り出すことができます。

取り出したframeデータはffmpegのavcodecでdecodeしたり
fdkaacやx264でencodeしたりできる。

もっと簡単にいろんな人が映像変換を楽しめるようになったらいいなと思います。

なお、このライブラリは、latencyを極力減らしてリアルタイム処理をする関係上
意図的に出力コンテナの情報をカットしたりしたりしてます。

# 作者

taktod
twitter: https://twitter.com/taktod
email: poepoemix@hotmail.com

# ライセンス

GPLv3

# 使い方

linuxもしくはmacで

```
$ go get github.com/taktod/ttLibGo
```

これで使えるようになると思います。

# 具体的な書き方

## まず適当にmp4でも読み込むことにします。

```test.go
package main

import (
	"os"
)

func main() {
	in, err := os.Open("test.h264.aac.mp4")
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
	}
}
```

## とりあえず、mp4を読み込むので、ttLibGo -> Readers -> Mp4を選択
コンテナReaderはttLibGo.Readersにまとめてあります。

```
	in, err := os.Open("test.h264.aac.mp4")
	if err != nil {
		panic(err)
	}
	reader := ttLibGo.Readers.Mp4()
	defer reader.Close()
	for {
		buffer := make([]byte, 65536)
```

処理がおわったらcloseしたいので、deferで指定しとく

## bufferに読み取ったデータをreaderに渡してframeを取り出す。

```
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
			fmt.Println(frame)
			return true
		}) {
			panic("error on read frame.")
		}
	}
```

これでフレームがざざーっと表示される。

```
...
&{{h264} 2102368 0 16000 1 0x6502160 false}
&{{h264} 2102896 0 16000 1 0x6502160 false}
&{{h264} 2103440 0 16000 1 0x6502160 false}
&{{h264} 2103968 0 16000 1 0x6502160 false}
&{{h264} 2104512 0 16000 1 0x6502160 false}
&{{h264} 2105040 0 16000 1 0x6502160 false}
&{{h264} 2105568 0 16000 1 0x6502160 false}
&{{h264} 2106112 0 16000 1 0x6502160 false}
&{{aac} 0 0 44100 2 0x6400730 false}
&{{aac} 0 0 44100 2 0x6400730 false}
&{{aac} 1024 0 44100 2 0x6400730 false}
&{{aac} 2048 0 44100 2 0x6400730 false}
&{{aac} 3072 0 44100 2 0x6400730 false}
&{{aac} 4096 0 44100 2 0x6400730 false}
&{{aac} 5120 0 44100 2 0x6400730 false}
&{{aac} 6144 0 44100 2 0x6400730 false}
&{{aac} 7168 0 44100 2 0x6400730 false}
...
```


## typeでどんなフレームか確認

プログラム的にフレームがなんであるか知りたい場合は、Frame.Typeを利用すればよい。
サポートしているタイプはttLibGo.FrameTypesにまとめてあります。

```
switch(frame.Type) {
case ttLibGo.FrameTypes.H264:
  // h264の場合の処理
case ttLibGo.FrameTypes.Aac:
  // aacの場合の処理
}
```

## より詳細な情報を取得

より詳細な情報を取得する場合は、こんな具合
ttLibGo.{Frame}.Castを利用すると詳細なデータにできます。

```
switch(frame.Type) {
case ttLibGo.FrameTypes.H264:
  fmt.Println(ttLibGo.H264.Cast(frame))
}
```

結果
```
&{{h264} 0 0 16000 1 0xad00270 false}
&{{h264} 0 0 16000 1 {info} 640 360 {configData} 4294967295 false 0xad00270}
&{{h264} 0 0 16000 1 0xad00270 false}
&{{h264} 0 0 16000 1 {key} 640 360 {sliceIDR} 2 false 0xad00270}
&{{h264} 528 0 16000 1 0xad00270 false}
&{{h264} 528 0 16000 1 {inner} 640 360 {slice} 0 false 0xad00270}
```

## avcodecDecoderでh264をdecodeしてみる。

avcodecDecoder (ffmpegのやつ)でdecodeしてみる。
readerの後ろのところで、decoderを準備してやる。

```
	reader := ttLibGo.Readers.Mp4()
	defer reader.Close()
	var decoder ttLibGo.IDecoder
	defer func() {
		if decoder != nil {
			decoder.Close()
		}
	}()
	for {
		buffer := make([]byte, 65536)
```

とりあえず、interfaceだけ準備しておいて値が設定されてたらcloseする形にする。

## 実際にdecodeするコードを書く

typeを確認したところで、decoderが準備されてなかったら、準備するようにしとく
decoderは例によってttLibGo.Decodersにすべてまとめてある。
今回は映像の変換なので、AvcodecVideoを選択

```
			switch frame.Type {
			case ttLibGo.FrameTypes.H264:
				h264 := ttLibGo.H264.Cast(frame)
				if decoder == nil {
					decoder = ttLibGo.Decoders.AvcodecVideo(h264.Type, h264.Width, h264.Height)
				}
				return decoder.DecodeFrame(h264, func(frame *ttLibGo.Frame) bool {
					fmt.Println(frame)
					return true
				})
			case ttLibGo.FrameTypes.Aac:
```

結果

```
&{{yuv420} 0 0 16000 1 0x7507f50 false}
&{{yuv420} 528 0 16000 1 0x7507f50 false}
&{{yuv420} 1072 0 16000 1 0x7507f50 false}
&{{yuv420} 1600 0 16000 1 0x7507f50 false}
&{{yuv420} 2128 0 16000 1 0x7507f50 false}
&{{yuv420} 2672 0 16000 1 0x7507f50 false}
```

h264のデータがyuv420に変換されてることがわかる。