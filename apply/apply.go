// This source file was generated by Fungo. DO NOT MODIFY.
//
// Generated on: 2019-09-17 09:20:22
// Fungo version: 0.0.1

package apply

func Apply(source []string, transform func(string) (string, error)) ([]string, error) {
	size := len(source)
	target := make([]string, size)
	for i := range source {
		result, err := transform(source[i])
		if err != nil {
			return target, err
		}
		target[i] = result
	}
	return target, nil
}