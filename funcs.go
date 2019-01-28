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
	case float32:
		return a.(float32) == 0
	case float64:
		return a.(float64) == 0
	case string:
		return a == ""
	}
	return false
}

func NotEmpty(s string) bool {
	return !IsDefaultValue(s)
}

func NotNilNorEmpty(s *string) bool {
	return s != nil && !IsDefaultValue(*s)
}

// Return the first string if it's not empty, otherwise return second one. Short cut for golang if else
func OrString(s1, s2 string) string {
	if NotEmpty(s1) {
		return s1
	}
	return s2
}

// Return the first string array if it's not empty, otherwise return second one. Short cut for golang if else
func OrStringArray(arr1, arr2 []string) []string {
	if len(arr1) == 0 {
		return arr2
	}
	return arr1
}
