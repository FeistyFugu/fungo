package templates

const FindFirst =
`package {{.PackageName}}

func {{.FunctionName}}(items []{{.T1}}, predicate func({{.T1}}) (bool, error)) (int, {{.T1}}, error) {
	for index, item := range items {
		res, err := predicate(item)
		if err != nil {
			return index, item, err
		}
		if res {
			return index, item, nil
		}
	}
	var emptyVal {{.T1}}
	return -1, emptyVal, nil
}`