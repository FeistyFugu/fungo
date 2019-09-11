package templates

const Filter =
`package {{.PackageName}}

func {{.FunctionName}}(source []{{.T1}}, predicate func({{.T1}}) (bool, error)) ([]{{.T1}}, error) {
	var target []{{.T1}}
	for _, item := range source {
		result, err := predicate(item)
		if err != nil {
			return target, err
		}
		if result {
			target = append(target, item)
		}
	}
	return target, nil
}`