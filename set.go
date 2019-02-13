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
