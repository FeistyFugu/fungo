package templates

const FindLast =
`package {{.PackageName}}

func {{.FunctionName}}(items []{{.T1}}, predicate func({{.T1}}) (bool, error)) (int, {{.T1}}, error) {
	for i := len(items) - 1; i >= 0; i-- {
		res, err := predicate(items[i])
		if err != nil {
			return i, items[i], err
		}
		if res {
			return i, items[i], nil
		}
	}
	var emptyVal {{.T1}}
	return -1, emptyVal, nil
}`