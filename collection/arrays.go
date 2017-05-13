package collection

//去除数组中重复元素
func UniqueString(slice []string) []string {
	if len(slice) == 0 {
		return slice
	}

	values := make([]string, 0, len(slice))

	m := make(map[string]bool)
	for _, v := range slice {
		if m[v] == false {
			m[v] = true
			values = append(values, v)
		}
	}

	return values
}
