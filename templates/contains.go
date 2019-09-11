package templates

const Contains =
`package {{.PackageName}}

func {{.FunctionName}}(list1 []{{.T1}}, list2 []{{.T2}}, predicate func(a {{.T1}}, b {{.T2}}) (bool, error)) (bool, int, error) {
	size1 := len(list1)
	size2 := len(list2)
	if (size1 == 0 || size2 == 0) || (size1 < size2) {
		return false, -1, nil
	}
	for i, _ := range list1 {
		bound := i + size2
		if bound > size1 {
			return false, -1, nil
		}
		slice := list1[i:bound]
		foundCount := 0
		for j, _ := range slice {
			result, err := predicate(slice[j], list2[j])
			if err != nil {
				return false, -1, err
			}
			if !result {
				break
			}
			foundCount++
			if foundCount == size2 {
				return true, i, nil
			}
		}
	}
	return false, -1, nil
}`