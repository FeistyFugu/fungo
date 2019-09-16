package templates

const Match =
`package {{.PackageName}}

func {{.FunctionName}}(list1 []{{.T1}}, list2 []{{.T2}}, compare func(a {{.T1}}, b {{.T2}}) (bool, error)) (bool, error) {
	if len(list1) != len(list2) {
		return false, nil
	}
	for i := range list1 {
		r, err := compare(list1[i], list2[i])
		if !r || err != nil {
			return r, err
		}
	}
	return true, nil
}`