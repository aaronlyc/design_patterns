/*
定义原型接口
*/
package clone_folder

type inode interface {
	print(string)
	clone() inode
}