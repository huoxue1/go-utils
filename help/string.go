package help

import (
	"encoding/json"
	"fmt"
	"strconv"
	"unsafe"
)

func StringToByteSlice(s string) []byte {

	tmp1 := (*[2]uintptr)(unsafe.Pointer(&s))

	tmp2 := [3]uintptr{tmp1[0], tmp1[1], tmp1[1]}

	return *(*[]byte)(unsafe.Pointer(&tmp2))

}

func ByteSliceToString(bytes []byte) string {

	return *(*string)(unsafe.Pointer(&bytes))

}

func ToString(data any) string {
	if content, ok := data.(fmt.Stringer); ok {
		return content.String()
	}
	switch data.(type) {
	case int:
		return strconv.Itoa(data.(int))
	case int32:
		return strconv.FormatInt(int64(data.(int32)), 10)
	case int16:
		return strconv.FormatInt(int64(data.(int16)), 10)
	case int64:
		return strconv.FormatInt(data.(int64), 10)
	case float64:
		return strconv.FormatFloat(data.(float64), 'E', -1, 64)
	case string:
		return data.(string)
	case []byte:
		return ByteSliceToString(data.([]byte))
	default:
		data, err := json.Marshal(data)
		if err != nil {
			return ""
		}
		return ByteSliceToString(data)
	}
}
