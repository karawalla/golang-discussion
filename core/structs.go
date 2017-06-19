package main

import (
  "fmt"
)

type Node struct {
  value int
  next *Node
}

func (n *Node) SetValue(val int){
  n.value = val
}

func (n Node) GetValue() int{
  return n.value
}

func main(){
  n := new(Node)
  n.SetValue(10)

  fmt.Println("Node ", n.GetValue())
}
