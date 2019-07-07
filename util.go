package ttLibGo

/*
#include <stdint.h>
#include <stdlib.h>
extern void *StdMap_make();
extern void StdMap_putString(void *mp, const char *key, const char *value);
extern void StdMap_putUint32(void *mp, const char *key, uint32_t value);
extern void StdMap_putUint64(void *mp, const char *key, uint64_t value);
extern void StdMap_putStringList(void *ptr, const char *key, const char *value);
extern void StdMap_close(void *mp);
*/
import "C"
import (
	"fmt"
	"reflect"
	"unsafe"
)

var mapUtil = struct {
	fromMap func(data map[string]interface{}) unsafe.Pointer
	close   func(ptr unsafe.Pointer)
}{
	fromMap: func(data map[string]interface{}) unsafe.Pointer {
		stdMap := C.StdMap_make()
		for key, value := range data {
			ckey := C.CString(key)
			defer C.free(unsafe.Pointer(ckey))
			switch reflect.TypeOf(value).Kind() {
			case reflect.String:
				cvalue := C.CString(value.(string))
				defer C.free(unsafe.Pointer(cvalue))
				C.StdMap_putString(stdMap, ckey, cvalue)
			case reflect.Uint32:
				C.StdMap_putUint32(stdMap, ckey, C.uint32_t(value.(uint32)))
			case reflect.Uint64:
				C.StdMap_putUint64(stdMap, ckey, C.uint64_t(value.(uint64)))
			case reflect.Slice:
				for _, v := range value.([]string) {
					cval := C.CString(v)
					defer C.free(unsafe.Pointer(cval))
					C.StdMap_putStringList(stdMap, ckey, cval)
				}
			default:
				fmt.Println("unknown:")
				fmt.Println(value)
				fmt.Println(reflect.TypeOf(value).Kind())
			}
		}
		return stdMap
	},
	close: func(ptr unsafe.Pointer) {
		C.StdMap_close(ptr)
	},
}
