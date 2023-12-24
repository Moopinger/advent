package main

import (
	"strings"
)

type Map struct {
	Pattern string //RLL
	Nodes   map[string]Node
}

type Node struct {
	Name  string
	Left  string
	Right string
}

func NewMap(map_data []byte) Map {
	//var nodes []Node
	var pattern string
	nodes := make(map[string]Node)

	data_lines := strings.Split(string(map_data), "\n")

	pattern = data_lines[0]

	node_data := data_lines[2:]

	for _, node := range node_data {
		myNode := newNode(node)
		nodes[myNode.Name] = myNode
		//nodes = append(nodes, myNode)

	}

	m := Map{
		Pattern: pattern,
		Nodes:   nodes,
	}

	return m

}

func newNode(node_data string) Node {
	var name string
	var left string
	var right string
	//0123456789
	//TJF = (TXF, NGK)
	name = node_data[:3]
	left = node_data[7:10]
	right = node_data[12:15]

	n := Node{
		Name:  name,
		Left:  left,
		Right: right,
	}

	return n
}
