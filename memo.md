# 概要
 
ttLibGoを作り直した。
今回は起動時にcmakeとmakeを実行
作成したshared libraryにdlopenでアクセスして利用する形にしました。

dylibかsoがあるなら、buildする必要はないんですが、ttLibGoのデータが新しくなったときにbuildしなおしてくれないので、とりあえず毎回buildにしてみた。

ひととおり動作を調整したけど、fdkaacはttLibCにはいってないので、未実装状態
これは、別でソースコードをいれておいて、対応可能にしておきたいところ。

とりあえず、ubuntu環境(docker)でテストしてみたところだけど、
stdintがみえなくて動作しないがいくつか (みえるようにした。)
c++11がはいってないので動作しないもいくつか (apt installでいれた。)
cmakeもいれないとだめ。
ttLibCのCMakeListsがこけた・・・なんだこれ・・・

add_compile_definitionをadd_definitionに変更
cmake 10.系では利用できないっぽい。
checkLibraryのheaderの確認をかきかえ。
現状のままでは、みつからないときにbuildがこける。
-> 直した

dockerのubuntuでテストしてみた。

```
$ docker run -it ubuntu:latest /bin/bash
# apt update
# apt upgrade
# apt install git golang cmake c++11 vim
# useradd taktod
# su - taktod
$ go get github.com/taktod/ttLibGo
$ vi test.go
```

```test.go
package main

import (
  "fmt"
  "github.com/taktod/ttLibGo"
  )
func main() {
  fmt.Println("hello")
  reader := ttLibGo.Readers.Flv()
  defer reader.Close()
}
```

```
$ go run test.go
/home/taktod/go/src/github.com/taktod/ttLibGo/build/libttLibC.so
hello
$ 
```

とりあえずこれで動いてくれた。
