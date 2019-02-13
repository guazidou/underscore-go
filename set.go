package go_

type HashSet map[interface{}]bool

func NewHashSet(size int) HashSet {
	return make(map[interface{}]bool, size)
}

func (set HashSet) Add(key interface{}) {
	set[key] = true
}

func (set HashSet) Contains(key interface{}) bool {
	return set[key]
}

func (set HashSet) LoadWithDefault(key interface{}, defVal bool) bool {
	return Or(set.Contains(key), defVal).(bool)
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

func (set HashSet) Intersect(another HashSet) HashSet {
	res := NewHashSet(len(set))
	set.ForEach(func(k interface{}) {
		if another.Contains(k) {
			res.Add(k)
		}
	})
	return res
}

func (set HashSet) Union(another HashSet) HashSet {
	res := NewHashSet(len(set) + len(another))
	set.ForEach(res.Add)
	another.ForEach(res.Add)
	return res
}

func (set HashSet) Diff(another HashSet) HashSet {
	res := NewHashSet(len(set))
	res.ForEach(func(k interface{}) {
		if !another.Contains(k) {
			res.Add(k)
		}
	})
	return res
}
