# ttLibGo

# 概要

ttLibGo
golangでttLibCのプログラムをキックできるようにしてみた。
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

基本3-Clause BSD Licenseと考えています。
利用するライブラリにしたがって、ものによってはLGPLv3やGPLv3に部分的に変えたりしたいと思います。

# 使い方

goをインストールしてもらったあとは
```
$ go get github.com/taktod/ttLibGo
```
これで必要なものが入る予定です。
