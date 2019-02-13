package go_

type HashSet map[interface{}]bool

func NewHashSet(size int) HashSet {
	return make(map[interface{}]bool, size)
}

func (set HashSet) Add(ele interface{}) {
	set[ele] = true
}

func (set HashSet) Contains(ele interface{}) bool {
	return set[ele]
}

func (set HashSet) SetIfNotExist(ele interface{}) bool {
	pre := set.Contains(ele)
	set.Add(ele)
	return pre
}

func (set HashSet) Remove(ele interface{}) {
	delete(set, ele)
}

func (set HashSet) Keys() []interface{} {
	res := make([]interface{}, 0, len(set))
	for e := range set {
		res = append(res, e)
	}
	return res
}

func (set HashSet) Size() int {
	return len(set)
}

func (set HashSet) Replace(ele1, ele2 interface{}) {
	set.Remove(ele1)
	set.Add(ele2)
}

func (set HashSet) ForEach(f func(interface{})) {
	for k := range set {
		f(k)
	}
}

func (set HashSet) Map(f func(interface{}) interface{}) []interface{} {
	res := make([]interface{}, 0, len(set))
	for e := range set {
		res = append(res, f(e))
	}
	return res
}

func (set HashSet) Intersect(another HashSet) HashSet {
	res := NewHashSet(len(set))
	set.ForEach(func(ele interface{}) {
		if another.Contains(ele) {
			res.Add(ele)
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
	res.ForEach(func(ele interface{}) {
		if !another.Contains(ele) {
			res.Add(ele)
		}
	})
	return res
}
