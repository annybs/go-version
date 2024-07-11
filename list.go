package version

// List is a slice of versions that implements sort.Interface.
type List []*Version

func (list List) Len() int {
	return len(list)
}

func (list List) Less(i, j int) bool {
	return list[i].Less(list[j])
}

func (list List) Swap(i, j int) {
	a := list[i]
	b := list[j]
	list[i] = b
	list[j] = a
}
