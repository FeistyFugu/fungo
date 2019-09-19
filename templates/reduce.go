package templates

const Reduce =
`package {{.PackageName}}

func {{.FunctionName}}(initAcc {{.T2}}, items []{{.T1}}, calculation func({{.T2}}, {{.T1}}) ({{.T2}}, error)) ({{.T2}}, error) {
	acc := initAcc
	var err error
	for _, item := range items {
		acc, err = calculation(acc, item)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}`