package main

import (
	"fmt"
)

type StringIterator func() (str string, hasNext bool, err error)

/*
NewStringArrayIterator is exported factory function
that returns function implementing StringIterator type
*/
func NewStringArrayIterator(arr []string) StringIterator {
	i := 0
	return func() (_ string, hasNext bool, err error) {
		if i == len(arr) {
			return "", false, nil
		}
		str := arr[i]
		i++
		return str, true, nil
	}
}

func main() {
	var s string
	var hasNext bool
	var err error

	siter := NewStringArrayIterator([]string{"A", "B"})

	for {
		s, hasNext, err = siter()
		if err != nil || !hasNext {
			break
		}
		fmt.Printf("%s\n", s)
	}
}
