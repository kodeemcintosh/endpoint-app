package main

import (
	"fmt"
)

const (
	Create CommandType = 0
	Delete CommandType = 1
	Move   CommandType = 2
	List   CommandType = 3
)

func (cType CommandType) String() string {
	types := [...]string{
		"CREATE",
		"DELETE",
		"MOVE",
		"LIST",
	}

	return types[cType]
}

func (c Command) Execute() {
	switch c.Type {
	case Create.String():
		p := Path(c.Args[0])

		fmt.Printf("%s %s\n", Create.String(), string(p))
		HandleCreate(p)

	case Delete.String():
		p := Path(c.Args[0])

		fmt.Printf("%s %s\n", Delete.String(), string(p))
		HandleDelete(p)

	case Move.String():
		p := Path(c.Args[0])
		d := Path(c.Args[1])

		fmt.Printf("%s %s %s\n", Move.String(), string(p), string(d))
		HandleMove(p, d)

	case List.String():
		fmt.Println(List.String())
		HandleList()

	}
}
