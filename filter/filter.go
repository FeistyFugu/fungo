// This source file was generated by Fungo. DO NOT MODIFY.
//
// Generated on: 2019-09-17 10:06:29
// Fungo version: 0.0.1

package filter

func Filter(source []int, predicate func(int) (bool, error)) ([]int, error) {
	var target []int
	for _, item := range source {
		result, err := predicate(item)
		if err != nil {
			return target, err
		}
		if result {
			target = append(target, item)
		}
	}
	return target, nil
}