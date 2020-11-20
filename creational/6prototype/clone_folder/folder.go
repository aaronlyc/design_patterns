/*
folder克隆文件夹的具体实现
*/
package clone_folder

import "fmt"

type folder struct {
	children []inode    // 文件夹中的内容
	name string
}

func (f *folder) print(indentation string) {
	fmt.Println(indentation + f.name)
	for _, child := range f.children {
		child.print(indentation + indentation)
	}
}

func (f *folder) clone() inode {
	cloneFolder := &folder{name: f.name + "_clone"}
	var tempChildren []inode
	for _, child := range f.children {
		childCopy := child.clone()
		tempChildren = append(tempChildren, childCopy)
	}
	cloneFolder.children = tempChildren
	return cloneFolder
}

