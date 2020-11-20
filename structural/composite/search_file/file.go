package search_file

import "fmt"

type file struct {
	name string
}

func (f *file) search(s string) {
	fmt.Printf("Searching for keyword %s in file %s\n", s, f.name)
}

func (f *file) getName() string {
	return f.name
}
