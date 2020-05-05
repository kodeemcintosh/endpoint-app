package main

import ()

func HandleCreate(p Path) {
	nl := GenerateNodeList(p)

	root.CreateNode(nl)
}

func HandleDelete(p Path) {
	nl := GenerateNodeList(p)

	root.DeleteNode(nl, p)
}

func HandleMove(p Path, destination Path) {
	nl := GenerateNodeList(p)
	dNl := GenerateNodeList(destination)

	node := root.ReturnNode(nl)
	node.Depth = len(dNl) + 1

	node.UpdateChildren()

	dNl = append(dNl, &node)

	root.CreateNode(dNl)
}

func HandleList() {
	root.Walk()
}
