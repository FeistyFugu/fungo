// This source file was generated by Fungo. DO NOT MODIFY.
//
// Generated on: 2019-09-16 12:02:02
// Fungo version: 0.0.1

package filter

func intEqual(list1 []int, list2 []int, compare func(a int, b int) (bool, error)) (bool, error) {
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
}