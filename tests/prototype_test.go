package tests

import (
	"fmt"
	"testing"

	proto "github.com/KyloRilo/design-patterns-go/patterns/prototype"
)

func TestPrototype(t *testing.T) {
	file1 := &proto.File{Name: "File1"}
	file2 := &proto.File{Name: "File2"}
	file3 := &proto.File{Name: "File3"}

	folder1 := &proto.Folder{
		Children: []proto.Inode{file1},
		Name:     "Folder1",
	}

	folder2 := &proto.Folder{
		Children: []proto.Inode{folder1, file2, file3},
		Name:     "Folder2",
	}
	fmt.Println("\nPrinting hierarchy for Folder2")
	folder2.Print("  ")

	cloneFolder := folder2.Clone()
	fmt.Println("\nPrinting hierarchy for clone Folder")
	cloneFolder.Print("  ")
}
