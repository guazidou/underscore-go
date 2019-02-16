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
