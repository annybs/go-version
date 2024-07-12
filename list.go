package version

// List is a slice of versions that implements sort.Interface.
type List []*Version

// Match tests versions against a constraint and returns a new List of matching versions only.
func (list List) Match(c *Constraint) List {
	filtered := List{}

	for _, v := range list {
		if v.Match(c) {
			filtered = append(filtered, v)
		}
	}

	return filtered
}

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
