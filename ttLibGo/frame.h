/*
他のライブラリでも同じ処理を使いたいので
この動作だけframe.c frame.hに抜き出しておく
文字列からframeTypeを取り出すだけなら、別にgoで記述しておいてもよさそうだけど・・・
 */
#include <stdlib.h>
#include <string.h>
#include <ttLibC/frame/frame.h>

/**
 * Frame_getFrameType
 *  codecTypeの名前からttLibC_Frame_Typeのデータを参照します
 * @param type codecTypeの名称
 */
ttLibC_Frame_Type Frame_getFrameType(const char *type);
 