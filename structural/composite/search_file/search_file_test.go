package search_file

import (
	. "github.com/onsi/ginkgo"
)

var _ = Describe("SearchFile", func() {
	file1 := &file{name: "file1"}
	file2 := &file{name: "file2"}
	file3 := &file{name: "file3"}
	file4 := &file{name: "file4"}
	folder1 := &folder{
		name: "Folder1",
	}
	folder1.add(file1)

	It("测试folder1:", func() {
		folder1.search("hello")
	})

	It("测试folder2:", func() {
		folder2 := &folder{
			name: "Folder2",
		}
		folder2.add(folder1)
		folder2.add(file2)
		folder2.add(file3)
		folder2.add(file4)
		folder2.search("World")
	})
})
