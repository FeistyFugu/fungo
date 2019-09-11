package templates

const Min =
`package min

import "errors"

func Min(items []{{.T1}}, comparison func({{.T1}}, {{.T1}}) (int, error)) ({{.T1}}, error) {
	var min {{.T1}}
	if len(items) == 0 {
		return min, errors.New("slice contains no items")
	}
	min = items[0]
	for i := 1; i < len(items); i++ {
		res, err := comparison(items[i], min)
		if err != nil {
			return min, err
		}
		if res < 0 {
			min = items[i]
		}
	}
	return min, nil
}`
