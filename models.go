package main

import ()

type Path string

type Command struct {
	Type string
	Args []string
}

type CommandType int

type Node struct {
	Name     string
	Depth    int
	Children NodeList
}

type NodeList []*Node
