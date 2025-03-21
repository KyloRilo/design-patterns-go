package main

import (
	"fmt"
	"testing"
)

type Inode interface {
	print(string)
	clone() Inode
}

type File struct {
	name string
}

func (f *File) print(indent string) {
	fmt.Println(indent + f.name)
}

func (f *File) clone() Inode {
	return &File{name: f.name + "_clone"}
}

type Folder struct {
	children []Inode
	name     string
}

func (f *Folder) print(indent string) {
	fmt.Println(indent + f.name)
	for _, i := range f.children {
		i.print(indent + indent)
	}
}

func (f *Folder) clone() Inode {
	var tempChildren []Inode
	cloneFolder := &Folder{name: f.name + "_clone"}
	for _, i := range f.children {
		copy := i.clone()
		tempChildren = append(tempChildren, copy)
	}

	cloneFolder.children = tempChildren
	return cloneFolder
}

func TestPrototype(t *testing.T) {
	file1 := &File{name: "File1"}
	file2 := &File{name: "File2"}
	file3 := &File{name: "File3"}

	folder1 := &Folder{
		children: []Inode{file1},
		name:     "Folder1",
	}

	folder2 := &Folder{
		children: []Inode{folder1, file2, file3},
		name:     "Folder2",
	}
	fmt.Println("\nPrinting hierarchy for Folder2")
	folder2.print("  ")

	cloneFolder := folder2.clone()
	fmt.Println("\nPrinting hierarchy for clone Folder")
	cloneFolder.print("  ")
}
