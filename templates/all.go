package templates

const All =
`package {{.PackageName}}

func {{.FunctionName}}(items []{{.T1}}, predicate func({{.T1}}) (bool, error)) (bool, error) {
	for _, item := range items {
		res, err := predicate(item)
		if !res || err != nil {
			return false, err
		}
	}
	return true, nil
}`