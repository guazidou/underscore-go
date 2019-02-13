package go_

type HashSet map[interface{}]interface{}

func (set HashSet) Put(key interface{}, value interface{}) {
	set[key] = value
}

func (set HashSet) Load(key interface{}) interface{} {
	return set[key]
}

func (set HashSet) LoadWithDefault(key interface{}, defVal interface{}) interface{} {
	return Or(set.Load(key), defVal)
}

func (set HashSet) Delete(key interface{}) interface{} {
	val := set[key]
	delete(set, key)
	return val
}
