package main

func Search(a map[string]string, b string) string {
	return ""
}

type Dictionary map[string]string

func (d Dictionary) Search(word string) string {
	return d[word]
}
