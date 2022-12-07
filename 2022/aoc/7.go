package aoc

import (
	"fmt"
	"strconv"
	"strings"
)

type Input7 *File

type Runner7 struct{}

func (r Runner7) FmtInput(input string) Input7 {
	commands := strings.Split(input[2:], "\n$ ")
	var root, curr *File
	for _, cmd := range commands {
		lines := strings.Split(cmd, "\n")
		args := strings.Split(lines[0], " ")
		switch args[0] {
		case "cd":
			switch args[1] {
			case "/":
				if root == nil {
					root = &File{id: "/", size: 0}
				}
				curr = root
			case "..":
				if curr.parent != nil {
					curr = curr.parent
				}
			default:
				file := findInFile(curr, args[1])
				if file != nil {
					curr = file
				}
			}
		case "ls":
			for _, line := range lines[1:] {
				lineArgs := strings.Split(line, " ")
				var id string
				var size int
				switch lineArgs[0] {
				case "dir":
					id = lineArgs[1]
				default:
					id = lineArgs[1]
					size, _ = strconv.Atoi(lineArgs[0])
				}
				curr.children = append(curr.children, &File{id: id, size: size, parent: curr})
			}
		default:
			panic(fmt.Errorf("command not found: %v", args[0]))
		}
	}
	return root
}

func (r Runner7) Run1(input Input7, _ bool) int {
	var root *File = input
	root.CalcSize()

	return sumLowerThan100K(root)
}

func (r Runner7) Run2(input Input7, _ bool) int {
	var root *File = input
	root.CalcSize()
	diskSpace := 70000000
	spaceToBeFreed := 30000000 - (diskSpace - root.size)
	potentialFolders := []int{}

	findHigherThan(root, spaceToBeFreed, &potentialFolders)
	min := potentialFolders[0]
	for _, folder := range potentialFolders {
		if folder < min {
			min = folder
		}
	}

	return min
}

type File struct {
	size     int
	id       string
	children []*File
	parent   *File
}

func (f File) IsFile() bool {
	return len(f.children) <= 0
}

func (f *File) CalcSize() int {
	if f.IsFile() {
		return f.size
	}
	size := 0
	for _, child := range f.children {
		size += child.CalcSize()
	}
	f.size = size
	return size
}

func (f File) String() string {
	var sb strings.Builder

	f.printWithTabs(&sb, 0)

	return sb.String()
}

func (f File) printWithTabs(sb *strings.Builder, count int) {
	for i := 0; i < count; i++ {
		sb.WriteRune(' ')
	}
	sb.WriteString(fmt.Sprintf("- %s ", f.id))
	if f.IsFile() {
		sb.WriteString(fmt.Sprintf("(file, size=%d)", f.size))
	} else {
		sb.WriteString("(dir)")
	}
	sb.WriteRune('\n')
	for _, child := range f.children {
		child.printWithTabs(sb, count+2)
	}
}

func findInFile(file *File, id string) *File {
	for _, child := range file.children {
		if id == child.id {
			return child
		}
	}

	for _, child := range file.children {
		file := findInFile(child, id)
		if file != nil {
			return file
		}
	}

	return nil
}

func sumLowerThan100K(f *File) int {
	if f.IsFile() {
		return 0
	}
	sum := 0
	if f.size < 100000 {
		sum += f.size
	}
	for _, child := range f.children {
		sum += sumLowerThan100K(child)
	}
	return sum
}

func findHigherThan(f *File, min int, result *[]int) {
	if f.IsFile() {
		return
	}
	if f.size >= min {
		*result = append(*result, f.size)
	}
	for _, child := range f.children {
		findHigherThan(child, min, result)
	}
}
