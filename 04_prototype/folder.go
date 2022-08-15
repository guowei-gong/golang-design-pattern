package prototype

import "fmt"

type Folder struct {
	name     string
	children []Inode
}

func (f *Folder) print(indentation string) {
	fmt.Println(indentation + f.name)
}

func (f *Folder) clone() Inode {
	var tempChildren []Inode
	cloneFolder := &Folder{
		name: f.name + "_clone",
	}
	for _, i := range f.children {
		copy := i.clone()
		tempChildren = append(tempChildren, copy)
	}
	cloneFolder.children = tempChildren
	return cloneFolder
}
