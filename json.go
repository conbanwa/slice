package slice

import (
	"encoding/json"
	"net/url"
	"unsafe"
)

func BodySlice(body []byte, limit ...int) string {
	return Bytes2String(Slice(body, limit...))
}

func BodyCut(body []byte, max ...int) string {
	return Bytes2String(Cut(body, max...))
}

func Bytes2String(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func String2Bytes(b string) []byte {
	return *(*[]byte)(unsafe.Pointer(&b))
}

func ValuesToJson(values url.Values) []byte {
	b, err := ValuesToJsonSafe(values)
	Panic(err)
	return b
}

func ValuesToJsonSafe(values url.Values) ([]byte, error) {
	param := make(map[string]any)
	for k, v := range values {
		if len(v) == 1 {
			param[k] = v[0]
		} else {
			param[k] = v
		}
	}
	return json.Marshal(param)
}

func Panic(err error) {
	if err != nil {
		panic(err)
	}
}
