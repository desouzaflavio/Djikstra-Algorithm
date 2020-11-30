package main

import "fmt"

type Node struct {
	data   int
	weight int
	next   *Node
	prev   *Node
}

func (nd *Node) initList() {
	nd.next = new(Node)
	nd.prev = new(Node)
}

func (nd *Node) setData(d int) {
	nd.data = d
}

func (nd *Node) setWeight(value int) {
	nd.weight = value
}

func (nd *Node) getAddr() *Node {
	return nd
}

func (nd *Node) getNextAddr() *Node {
	return nd.next
}

func (nd *Node) getPrevAddr() *Node {
	return nd.prev
}

func main() {

	var n Node

	n.initList()
	n.setWeight(5)
	n.setData(10)

	n1 := n.next
	n1.prev = n.getAddr()

	fmt.Println("Peso do Nó:", n.weight)
	fmt.Println("Peso do Nó:", n.data)
	fmt.Printf("Endereço do N[0]: %p\n", &n)
	fmt.Printf("Endereço do N[0]: %p\n", n.getAddr())
	fmt.Printf("Endereço previo do N[1]: %p\n", n1.prev)
	fmt.Printf("Endereço do próximo nó: %p\n", n.getNextAddr())
	fmt.Printf("Endereço do nó anteior: %p\n", n.getPrevAddr())
}
