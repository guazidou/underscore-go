package go_

import (
	"bytes"
	"strconv"
)

func JoinErrors(errs []error, sep string) string {
	n := len(sep) * (len(errs) - 1)
	for i := 0; i < len(errs); i++ {
		n += len(errs[i].Error())
	}

	b := make([]byte, n)
	bp := copy(b, errs[0].Error())
	for _, s := range errs[1:] {
		bp += copy(b[bp:], sep)
		bp += copy(b[bp:], s.Error())
	}
	return string(b)
}

func JointInt64(data []int64, sep string) string {
	l := 0
	sl := len(sep)
	dl := len(data)
	for _, v := range data {
		for v > 0 {
			v /= 10
			l++
		}
		l += sl
	}
	l -= sl
	buffer := bytes.NewBuffer(make([]byte, 0, l))
	for i, v := range data {
		buffer.WriteString(strconv.FormatInt(v, 10))
		if i != dl-1 {
			buffer.WriteString(sep)
		}
	}
	return buffer.String()
}

// This method returns an map composed from key-value pairs. It will ignore the pair that are not formed in key, value format.
func FromPairs(data ...[]string) map[string]string {
	res := make(map[string]string, len(data))
	for _, pair := range data {
		if len(pair) != 2 {
			continue
		}
		res[pair[0]] = pair[1]
	}
	return res
}

func ZipObject(key []string, value []string) map[string]string {
	if len(key) < len(value) {
		return map[string]string{} // return empty map when given key is shorter than value
	}
	res := make(map[string]string, len(key))
	for i := range key {
		res[key[i]] = value[i]
	}
	return res
}

func UnzipObject(data map[string]string) ([]string, []string) {
	key := make([]string, 0, len(data))
	value := make([]string, 0, len(data))

	for k, v := range data {
		key = append(key, k)
		value = append(value, v)
	}
	return key, value
}

func CountBy(data []interface{}, transFunc func(interface{}) interface{}) map[interface{}]int64 {
	res := make(map[interface{}]int64)
	for _, v := range data {
		t := transFunc(v)
		res[t] = res[t] + 1
	}
	return res
}

func DiffInt64(data []int64, exclude ...[]int64) []int64 {
	set := NewHashSetFromInt64List(data)
	for _, arr := range exclude {
		for _, v := range arr {
			if set.Contains(v) {
				set.Remove(v)
			}
		}
	}
	res := make([]int64, 0, set.Size())
	set.ForEach(func(element interface{}) {
		res = append(res, element.(int64))
	})
	return res
}

func DiffString(data []string, exclude ...[]string) []string {
	set := NewHashSetFromStringList(data)
	for _, arr := range exclude {
		for _, v := range arr {
			if set.Contains(v) {
				set.Remove(v)
			}
		}
	}
	res := make([]string, 0, set.Size())
	set.ForEach(func(element interface{}) {
		res = append(res, element.(string))
	})
	return res
}
