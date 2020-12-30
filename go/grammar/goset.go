package grammar

// golang没有泛型 所以这里使用interface{}
type MapSet map[interface{}] struct {}

func (set *MapSet)Add(e interface{}) bool  {
	if _, found := (*set)[e]; found {
		return false
	}
	(*set)[e] = struct{}{}
	return true
}