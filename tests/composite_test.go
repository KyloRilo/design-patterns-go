package tests

import (
	"testing"

	comp "github.com/KyloRilo/design-patterns-go/patterns/composite"
)

func TestComposite(t *testing.T) {
	file1 := &comp.File{Name: "File1"}
	file2 := &comp.File{Name: "File2"}
	file3 := &comp.File{Name: "File3"}

	folder1 := &comp.Folder{
		Name: "Folder1",
	}

	folder1.Add(file1)

	folder2 := &comp.Folder{
		Name: "Folder2",
	}
	folder2.Add(file2)
	folder2.Add(file3)
	folder2.Add(folder1)

	folder2.Search("rose")
}
