package prototype

import (
	"fmt"
)

type Inode interface {
	Print(string)
	Clone() Inode
}

type File struct {
	Name string
}

func (f *File) Print(indent string) {
	fmt.Println(indent + f.Name)
}

func (f *File) Clone() Inode {
	return &File{Name: f.Name + "_clone"}
}

type Folder struct {
	Children []Inode
	Name     string
}

func (f *Folder) Print(indent string) {
	fmt.Println(indent + f.Name)
	for _, i := range f.Children {
		i.Print(indent + indent)
	}
}

func (f *Folder) Clone() Inode {
	var tempChildren []Inode
	cloneFolder := &Folder{Name: f.Name + "_clone"}
	for _, i := range f.Children {
		copy := i.Clone()
		tempChildren = append(tempChildren, copy)
	}

	cloneFolder.Children = tempChildren
	return cloneFolder
}
