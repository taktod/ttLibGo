# 概要
 
ttLibGoを作り直した。
今回は起動時にcmakeとmakeを実行
作成したshared libraryにdlopenでアクセスして利用する形にしました。

dylibかsoがあるなら、buildする必要はないんですが、ttLibGoのデータが新しくなったときにbuildしなおしてくれないので、とりあえず毎回buildにしてみた。

ひととおり動作を調整したけど、fdkaacはttLibCにはいってないので、未実装状態
これは、別でソースコードをいれておいて、対応可能にしておきたいところ。
ttLibGoをもう一度作り直す。
buildでcmakeを実施して、色々やります。cmake ../ttLibC
みたいな感じ。
libttLibC.soかlibttLibC.soがあるならbuildしなくて良い

で、このdlopenでこのプログラムを読み込んで色々テストすることにする。
先に必要な関数を全てloadしてから利用する形にしておくか・・・

ポインタは強制解放もできるけど、atexitでも解放するような感じにしておこうと思う。
一番はじめに必要となる動作を全て読み込んで・・・として、使えない動作は、使えない・・・と表示される的な・・・感じで

この方法だと、全てのdecoder、encoderが実装可能なことがわかったけど、とりあえず先にbuildシステムを全部切り替えるのを目標にしとこうと思う。

fdkaacの実装はttLibCに乗っかってないので、かけないね・・・後で考えるか・・・

とりあえず、reader decoder encoderまで実装できた。
あとは、resamplerとwriterとframe周りの処理だが・・・

frameの応答がnullの場合のgoのオブジェクト復元データがちょっとアレなんだが・・・
