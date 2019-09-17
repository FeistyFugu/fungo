// This source file was generated by Fungo. DO NOT MODIFY.
//
// Generated on: 2019-09-17 09:20:22
// Fungo version: 0.0.1

package fan_out

import "errors"

func FanOut(items []string, dispatch func(int, string) (int, string, error), merge func(int, string) (bool, error)) (err error) {
	type dispatchResult struct {
		index  int
		result string
		error  error
	}

	size := len(items)
	if len(items) < 1 {
		return errors.New("slice contains no items")
	}

	out := make(chan dispatchResult, size)
	done := make(chan bool, size)

	for i, item := range items {
		go func(i int, item string) {
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
}