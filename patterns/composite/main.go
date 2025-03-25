package composite

import "fmt"

type Component interface {
	Search(string)
}

type Folder struct {
	components []Component
	Name       string
}

func (f *Folder) Search(kw string) {
	fmt.Printf("Serching recursively for keyword %s in folder %s\n", kw, f.Name)
	for _, composite := range f.components {
		composite.Search(kw)
	}
}

func (f *Folder) Add(c Component) {
	f.components = append(f.components, c)
}

type File struct {
	Name string
}

func (f *File) Search(kw string) {
	fmt.Printf("Searching for keyword %s in file %s\n", kw, f.Name)
}

func (f *File) GetName() string {
	return f.Name
}
