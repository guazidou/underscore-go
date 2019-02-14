package go_

type Set interface {
	Add(element interface{})
	Contains(element interface{}) bool
	Remove(element interface{})
	SetIfNotExist(element interface{}) bool
	Values() []interface{}
	Size() int
	Replace(element1, element2 interface{})
	ForEach(f func(element interface{}))
	Map(f func(element interface{}) interface{}) Set
	Union(another Set) Set
	Diff(another Set) Set
}

type HashSet map[interface{}]bool

func NewHashSet(size int) HashSet {
	return make(map[interface{}]bool, size)
}

func NewHashSetFromInt64List(data []int64) HashSet {
	set := NewHashSet(len(data))
	for _, v := range data {
		set.Add(v)
	}
	return set
}

func NewHashSetFromStringList(data []string) HashSet {
	set := NewHashSet(len(data))
	for _, v := range data {
		set.Add(v)
	}
	return set
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

func (set HashSet) Values() []interface{} {
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

func (set HashSet) Map(f func(interface{}) interface{}) Set {
	res := NewHashSet(set.Size())
	for e := range set {
		res.Add(f(e))
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

func (set HashSet) Union(another Set) Set {
	res := NewHashSet(set.Size() + another.Size())
	set.ForEach(res.Add)
	another.ForEach(res.Add)
	return res
}

func (set HashSet) Diff(another Set) Set {
	res := NewHashSet(len(set))
	res.ForEach(func(ele interface{}) {
		if !another.Contains(ele) {
			res.Add(ele)
		}
	})
	return res
}
