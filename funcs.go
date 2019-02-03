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
	case []int:
		return len(a.([]int)) == 0
	case []int8:
		return len(a.([]int8)) == 0
	case []int16:
		return len(a.([]int16)) == 0
	case []int32:
		return len(a.([]int32)) == 0
	case []int64:
		return len(a.([]int64)) == 0
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
	case []uint8:
		return len(a.([]uint8)) == 0
	case []uint16:
		return len(a.([]uint16)) == 0
	case []uint32:
		return len(a.([]uint32)) == 0
	case []uint64:
		return len(a.([]uint64)) == 0
	case float32:
		return a.(float32) == 0
	case float64:
		return a.(float64) == 0
	case []float32:
		return len(a.([]float32)) == 0
	case []float64:
		return len(a.([]float64)) == 0
	case string:
		return a == ""
	case []string:
		return len(a.([]string)) == 0
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
	if !IsDefaultValue(arr1) {
		return arr1
	}
	if !IsDefaultValue(arr2) {
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

func OrIntArray(arr1, arr2 []int, arrs ...[]int) []int {
	if !IsDefaultValue(arr1) {
		return arr1
	}
	if !IsDefaultValue(arr2) {
		return arr2
	}
	for _, arr := range arrs {
		if len(arr) != 0 {
			return arr
		}
	}
	return []int{}
}

func OrInt8Array(arr1, arr2 []int8, arrs ...[]int8) []int8 {
	if !IsDefaultValue(arr1) {
		return arr1
	}
	if !IsDefaultValue(arr2) {
		return arr2
	}
	for _, arr := range arrs {
		if len(arr) != 0 {
			return arr
		}
	}
	return []int8{}
}

func OrInt16Array(arr1, arr2 []int16, arrs ...[]int16) []int16 {
	if !IsDefaultValue(arr1) {
		return arr1
	}
	if !IsDefaultValue(arr2) {
		return arr2
	}
	for _, arr := range arrs {
		if len(arr) != 0 {
			return arr
		}
	}
	return []int16{}
}

func OrInt32Array(arr1, arr2 []int32, arrs ...[]int32) []int32 {
	if !IsDefaultValue(arr1) {
		return arr1
	}
	if !IsDefaultValue(arr2) {
		return arr2
	}
	for _, arr := range arrs {
		if len(arr) != 0 {
			return arr
		}
	}
	return []int32{}
}

func OrInt64Array(arr1, arr2 []int64, arrs ...[]int64) []int64 {
	if !IsDefaultValue(arr1) {
		return arr1
	}
	if !IsDefaultValue(arr2) {
		return arr2
	}
	for _, arr := range arrs {
		if len(arr) != 0 {
			return arr
		}
	}
	return []int64{}
}

func OrUIntArray(arr1, arr2 []uint, arrs ...[]uint) []uint {
	if !IsDefaultValue(arr1) {
		return arr1
	}
	if !IsDefaultValue(arr2) {
		return arr2
	}
	for _, arr := range arrs {
		if len(arr) != 0 {
			return arr
		}
	}
	return []uint{}
}

func OrUInt8Array(arr1, arr2 []uint8, arrs ...[]uint8) []uint8 {
	if !IsDefaultValue(arr1) {
		return arr1
	}
	if !IsDefaultValue(arr2) {
		return arr2
	}
	for _, arr := range arrs {
		if len(arr) != 0 {
			return arr
		}
	}
	return []uint8{}
}

func OrUInt16Array(arr1, arr2 []uint16, arrs ...[]uint16) []uint16 {
	if !IsDefaultValue(arr1) {
		return arr1
	}
	if !IsDefaultValue(arr2) {
		return arr2
	}
	for _, arr := range arrs {
		if len(arr) != 0 {
			return arr
		}
	}
	return []uint16{}
}

func OrUInt32Array(arr1, arr2 []uint32, arrs ...[]uint32) []uint32 {
	if !IsDefaultValue(arr1) {
		return arr1
	}
	if !IsDefaultValue(arr2) {
		return arr2
	}
	for _, arr := range arrs {
		if len(arr) != 0 {
			return arr
		}
	}
	return []uint32{}
}

func OrUInt64Array(arr1, arr2 []uint64, arrs ...[]uint64) []uint64 {
	if !IsDefaultValue(arr1) {
		return arr1
	}
	if !IsDefaultValue(arr2) {
		return arr2
	}
	for _, arr := range arrs {
		if len(arr) != 0 {
			return arr
		}
	}
	return []uint64{}
}

func OrFloat32Array(arr1, arr2 []float32, arrs ...[]float32) []float32 {
	if !IsDefaultValue(arr1) {
		return arr1
	}
	if !IsDefaultValue(arr2) {
		return arr2
	}
	for _, arr := range arrs {
		if len(arr) != 0 {
			return arr
		}
	}
	return []float32{}
}

func OrFloat64Array(arr1, arr2 []float64, arrs ...[]float64) []float64 {
	if !IsDefaultValue(arr1) {
		return arr1
	}
	if !IsDefaultValue(arr2) {
		return arr2
	}
	for _, arr := range arrs {
		if len(arr) != 0 {
			return arr
		}
	}
	return []float64{}
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
