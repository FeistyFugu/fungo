package templates

const Max =
`package {{.PackageName}}

import "errors"

func Max(items []{{.T1}}, comparison func({{.T1}}, {{.T1}}) (int, error)) ({{.T1}}, error) {
	var max {{.T1}}
	if len(items) == 0 {
		return max, errors.New("slice contains no items")
	}
	max = items[0]
	for i := 1; i < len(items); i++ {
		res, err := comparison(items[i], max)
		if err != nil {
			return max, err
		}
		if res > 0 {
			max = items[i]
		}
	}
	return max, nil
}`
