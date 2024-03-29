// This source file was generated by Fungo. DO NOT MODIFY.
//
// Generated on: 2019-09-17 15:29:12
// Fungo version: 0.0.1

package reduce

func Reduce(initAcc map[string][]Person, items []Person, calculation func(map[string][]Person, Person) (map[string][]Person, error)) (map[string][]Person, error) {
	acc := initAcc
	var err error
	for _, item := range items {
		acc, err = calculation(acc, item)
		if err != nil {
			return acc, err
		}
	}
	return acc, nil
}