# ttLibGO

# 概要

go言語でttLibCの変換ライブラリを使う動作

これを使うとflvやmkvといったコンテナデータの読み込み
内部に定義されているframeデータを取り出すことができます。

取り出したframeデータはffmpegのavcodecでdecodeしたり
fdkaacやx264でencodeしたりできる。

もっと簡単にいろんな人が映像変換を楽しめるようになったらいいなと思います。

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

使い方は、testディレクトリの中身をみてもらえればと思います。

# のこりやること

とりあえず、frameのbinaryからframeを復元する動作をのこっているので、やらないと・・・