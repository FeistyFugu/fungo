package templates

const GroupBy =
`package {{.PackageName}}

func {{.FunctionName}}(items []{{.T1}}, calculation func({{.T1}}) ({{.T2}}, error)) (map[{{.T2}}][]{{.T1}}, error) {
	groups := make(map[{{.T2}}][]{{.T1}})
	for _, item := range items {
		res, err := calculation(item)
		if err != nil {
			return groups, err
		}
		if _, ok := groups[res]; !ok {
			groups[res] = make([]{{.T1}}, 20)
		}
		groups[res] = append(groups[res], item)
	}
	return groups, nil
}`