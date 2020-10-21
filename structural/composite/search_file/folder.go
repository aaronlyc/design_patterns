package search_file

import "fmt"

type folder struct {
	components []component
	name       string
}

func (f *folder) search(s string) {
	fmt.Printf("searching recursively for keyword %s in folder %s\n", s, f.name)
	for _, component := range f.components {
		component.search(s)
	}
}

func (f *folder) add(c component) {
	f.components = append(f.components, c)
}
