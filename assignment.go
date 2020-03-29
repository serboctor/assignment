package main

import "errors"

// CountTokens counts tokens representing a complete word in a list of strings.
func CountTokens(list []string) (int, error) {

	token := false
	counter := 0
	for _, v := range list {
		if len(v) != 1 {
			return -1, errors.New("list elements must be single elements")
		}

		if v == " " && token == true {
			token = false
		} else if v != " " && token == false {
			counter++
			token = true
		}
	}

	return counter, nil
}
