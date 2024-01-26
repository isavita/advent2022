package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type File struct {
	size int
}

type Directory struct {
	files       map[string]*File
	directories map[string]*Directory
}

func newDirectory() *Directory {
	return &Directory{
		files:       make(map[string]*File),
		directories: make(map[string]*Directory),
	}
}

func (d *Directory) totalSize() int {
	size := 0
	for _, f := range d.files {
		size += f.size
	}
	for _, dir := range d.directories {
		size += dir.totalSize()
	}
	return size
}

func main() {
	file, err := os.Open("day7/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	root := newDirectory()
	currentDir := root
	directoryStack := []*Directory{root}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "$ cd") {
			path := strings.TrimSpace(line[4:])
			if path == "/" {
				currentDir = root
				directoryStack = []*Directory{root}
			} else if path == ".." {
				directoryStack = directoryStack[:len(directoryStack)-1]
				currentDir = directoryStack[len(directoryStack)-1]
			} else {
				if _, exists := currentDir.directories[path]; !exists {
					currentDir.directories[path] = newDirectory()
				}
				currentDir = currentDir.directories[path]
				directoryStack = append(directoryStack, currentDir)
			}
		} else if strings.HasPrefix(line, "dir") {
			dirName := strings.TrimSpace(line[4:])
			currentDir.directories[dirName] = newDirectory()
		} else {
			parts := strings.Fields(line)
			if len(parts) == 2 {
				size, _ := strconv.Atoi(parts[0])
				fileName := parts[1]
				currentDir.files[fileName] = &File{size: size}
			}
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	sumSizes := 0
	var calculateSizes func(*Directory)
	calculateSizes = func(d *Directory) {
		dirSize := d.totalSize()
		if dirSize <= 100000 {
			sumSizes += dirSize
		}
		for _, dir := range d.directories {
			calculateSizes(dir)
		}
	}
	calculateSizes(root)

	println(sumSizes)
}
