package go_

func IsDefaultValue(a interface{}) bool {
	switch a.(type) {
	case int:
		return a.(int) == 0
	case int8:
		return a.(int8) == 0
	case int16:
		return a.(int16) == 0
	case int32:
		return a.(int32) == 0
	case int64:
		return a.(int64) == 0
	case uint:
		return a.(uint) == 0
	case uint8:
		return a.(uint8) == 0
	case uint16:
		return a.(uint16) == 0
	case uint32:
		return a.(uint32) == 0
	case uint64:
		return a.(uint64) == 0
	case string:
		return a == ""
	}
	return false
}

func NotEmpty(s string) bool {
	return !IsDefaultValue(s)
}
