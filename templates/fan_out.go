package templates

const FanOut =
`package {{.PackageName}}

import "errors"

func {{.FunctionName}}(items []{{.T1}}, dispatch func(int, {{.T1}}) (int, {{.T2}}, error), merge func(int, {{.T2}}) (bool, error)) (err error) {
	type dispatchResult struct {
		index  int
		result {{.T2}}
		error  error
	}

	size := len(items)
	if len(items) < 1 {
		return errors.New("slice contains no items")
	}

	out := make(chan dispatchResult, size)
	done := make(chan bool, size)

	for i, item := range items {
		go func(i int, item {{.T1}}) {
			select {
			case <-done:
				return
			default:
				index, result, e := dispatch(i, item)
				out <- dispatchResult{index, result, e}
			}
		}(i, item)
	}
	count := 0
	for res := range out {
		count ++
		if res.error != nil {
			err = res.error
			break
		}
		stop, e := merge(res.index, res.result)
		if e != nil || stop || count >= size {
			err = e
			break
		}
	}
	done <- true
	close(done)
	return err
}`