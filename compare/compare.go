// This source file was generated by Fungo. DO NOT MODIFY.
//
// Generated on: 2019-09-17 15:29:12
// Fungo version: 0.0.1

package compare

func Compare(list1 []string, list2 []int, comparison func(a string, b int) (bool, error)) (bool, error) {
	if len(list1) != len(list2) {
		return false, nil
	}
	for i := range list1 {
		r, err := comparison(list1[i], list2[i])
		if !r || err != nil {
			return r, err
		}
	}
	return true, nil
}