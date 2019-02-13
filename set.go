package go_

type HashSet map[interface{}]bool

func (set HashSet) Add(key interface{}) {
	set[key] = true
}

func (set HashSet) Contain(key interface{}) bool {
	return set[key]
}

func (set HashSet) LoadWithDefault(key interface{}, defVal bool) bool {
	return Or(set.Contain(key), defVal).(bool)
}

func (set HashSet) Remove(key interface{}) {
	delete(set, key)
}

func (set HashSet) Keys() []interface{} {
	res := make([]interface{}, 0, len(set))
	for k := range set {
		res = append(res, k)
	}
	return res
}

func (set HashSet) Size() int {
	return len(set)
}

func (set HashSet) ForEach(f func(interface{})) {
	for k := range set {
		f(k)
	}
}

func (set HashSet) Map(f func(interface{}) interface{}) []interface{} {
	res := make([]interface{}, 0, len(set))
	for k := range set {
		res = append(res, f(k))
	}
	return res
}
