package reduce

//go:generate fungo -template=Reduce -fileName=reduce -t1=Person -t2=map[string][]Person

import (
	"fmt"
	"testing"
)

type Person struct {
	FirstName, LastName string
}

func TestReduce(t *testing.T) {
	persons := []Person{{"John", "Johnson"}, {"Julia", "Johnson"}, {"Steve", "Stevenson"}, {"Stacy", "Stevenson"}}
	result := make(map[string][]Person)
	result, _ = Reduce(result, persons, func(families map[string][]Person, p Person) (map[string][]Person, error) {
		if _, ok := families[p.LastName]; !ok {
			families[p.LastName] = make([]Person, 0)
		}
		families[p.LastName] = append(families[p.LastName], p)
		return families, nil
	})
	fmt.Println(result)
}
