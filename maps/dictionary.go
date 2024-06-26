package main

import "errors"

func Search(a map[string]string, b string) string {
	return ""
} 

type Dictionary map[string]string

var ErrNotFound = errors.New("could not find the word you were looking for")

func (d Dictionary) Search(word string) (string, error) {
	definiton, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return definiton, nil
}
