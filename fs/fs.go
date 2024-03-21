package lpfs

import (
	"fmt"
	"strings"
)

type Fs struct {
	root *Dir
}

type Handler interface {
	Read()
	Write()
}

type Dir struct {
	name   string
	childs map[string]*Node
}

type File struct {
	name    string
	handler Handler
}

type Node interface {
	isDir() bool
}

func (dir *Dir) isDir() bool {
	return true
}

func (file *File) isDir() bool {
	return false
}

func Fs() *Fs {
	fs := &Fs{
		root: Dir("/"),
	}
}

func Dir(name string) *Dir {
	return &Dir{
		name:   name,
		childs: make(map[string]*Node),
	}
}

func (node *Node) walk(path string) (*Node, error) {
	elems := strings.Split(strings.Trim(path, "/"), "/")
	p := node
	for i, v := range elems {
		if v == "" {
			// Skip unwanted subsequent "/" (e.g. "/foo//bar")
			continue
		}
		if p.isDir() {
			p = p.childs[v]
		} else {
			return nil, fmt.Errorf(path)
		}
	}
	return p, nil
}

func (dir *Dir) newChild(node *Node) error {
	if _, ok := dir.childs[node.name]; ok {
		return fmt.Errorf(node.name)
	}
	dir.childs[node.name] = node
	return nil
}
