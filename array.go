package go_

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
