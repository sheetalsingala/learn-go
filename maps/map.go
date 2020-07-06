package main

import "errors"

type Dictionary map[string]string 

func (d Dictionary) Search(word string) (string, error){
	definition, pk := d[word]
	if !pk {
		return "", errors.New("Could not find word")
	}
	return definition, nil
}
