// This source file was generated by Fungo. DO NOT MODIFY.
//
// Generated on: 2019-09-11 18:10:22
// Fungo version: 0.0.1

package filter

func Filter(source []string, predicate func(string) (bool, error)) ([]string, error) {
	var target []string
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