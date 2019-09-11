package templates

const Apply =
`package {{.PackageName}}

func {{.FunctionName}}(source []{{.T1}}, transform func({{.T1}}) ({{.T2}}, error)) ([]{{.T2}}, error) {
	size := len(source)
	target := make([]{{.T2}}, size)
	for i := range source {
		result, err := transform(source[i])
		if err != nil {
			return target, err
		}
		target[i] = result
	}
	return target, nil
}`