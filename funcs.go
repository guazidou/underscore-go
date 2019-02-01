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
func OrString(s1, s2 string, ss ...string) string {
	if NotEmpty(s1) {
		return s1
	}
	if NotEmpty(s2) {
		return s2
	}
	for _, s := range ss {
		if NotEmpty(s) {
			return s
		}
	}
	return ""
}

// Return the first string array if it's not empty, otherwise return second one. Short cut for golang if else
func OrStringArray(arr1, arr2 []string, arrs ...[]string) []string {
	if len(arr1) != 0 {
		return arr1
	}
	if len(arr2) != 0 {
		return arr2
	}
	for _, arr := range arrs {
		if len(arr) != 0 {
			return arr
		}
	}
	return []string{}
}

func OrBool(a, b bool, bs ...bool) bool {
	if a || b {
		return a || b
	}
	for _, v := range bs {
		if v {
			return v
		}
	}
	return false
}

func OrInt64(a, b int64, is ...int64) int64 {
	if !IsDefaultValue(a) {
		return a
	}
	if !IsDefaultValue(b) {
		return b
	}
	for _, v := range is {
		if !IsDefaultValue(v) {
			return v
		}
	}
	return 0
}

func OrInt32(a, b int32, is ...int32) int32 {
	if !IsDefaultValue(a) {
		return a
	}
	if !IsDefaultValue(b) {
		return b
	}
	for _, v := range is {
		if !IsDefaultValue(v) {
			return v
		}
	}
	return 0
}

func OrInt16(a, b int16, is ...int16) int16 {
	if !IsDefaultValue(a) {
		return a
	}
	if !IsDefaultValue(b) {
		return b
	}
	for _, v := range is {
		if !IsDefaultValue(v) {
			return v
		}
	}
	return 0
}

func OrInt8(a, b int8, is ...int8) int8 {
	if !IsDefaultValue(a) {
		return a
	}
	if !IsDefaultValue(b) {
		return b
	}
	for _, v := range is {
		if !IsDefaultValue(v) {
			return v
		}
	}
	return 0
}

func OrInt(a, b int, is ...int) int {
	if !IsDefaultValue(a) {
		return a
	}
	if !IsDefaultValue(b) {
		return b
	}
	for _, v := range is {
		if !IsDefaultValue(v) {
			return v
		}
	}
	return 0
}

func OrFloat64(a, b float64, is ...float64) float64 {
	if !IsDefaultValue(a) {
		return a
	}
	if !IsDefaultValue(b) {
		return b
	}
	for _, v := range is {
		if !IsDefaultValue(v) {
			return v
		}
	}
	return 0
}

func OrFloat32(a, b float32, is ...float32) float32 {
	if !IsDefaultValue(a) {
		return a
	}
	if !IsDefaultValue(b) {
		return b
	}
	for _, v := range is {
		if !IsDefaultValue(v) {
			return v
		}
	}
	return 0
}

func Or(a, b interface{}, vals ...interface{}) interface{} {
	if !IsDefaultValue(a) {
		return a
	}
	if !IsDefaultValue(b) {
		return b
	}
	for _, val := range vals {
		if !IsDefaultValue(val) {
			return val
		}
	}
	return nil
}
