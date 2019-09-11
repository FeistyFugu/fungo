package templates

const Reduce =
`package {{.PackageName}}

func {{.FunctionName}}(items []{{.T1}}, calculation func({{.T2}}, {{.T1}}) ({{.T2}}, error)) ({{.T2}}, error) {
	var acc {{.T2}}
	var err error
	for _, item := range items {
		acc, err = calculation(acc, item)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}`