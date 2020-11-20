package clone_folder

import (
	"fmt"
	"testing"
)

func TestInode(t *testing.T) {
	file1 := &file{name: "file1"}
	file2 := &file{name: "file2"}
	file3 := &file{name: "file3"}

	folder1 := &folder{
		children: []inode{file1},
		name: "Folder1",
	}

	folder2 := &folder{
		children: []inode{folder1, file2, file3},
		name: "Folder2",
	}

	fmt.Println("\nPrinting hierarchy for Folder2")
	folder2.print(" ")

	cloneFolder := folder2.clone()
	fmt.Println("\nPrinting hierarchy for clone Folder")
	cloneFolder.print(" ")
}