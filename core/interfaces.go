package main
import(
  "fmt"
)

type Node interface {
  SetValue(int)
  GetValue() int
}

type MyNode struct {
    value int
}

func (n MyNode) SetValue(val int){
  n.value = val
}

func (n MyNode) GetValue() int {
  return n.value
}
func main(){
  var n Node
  _, ok := n.(*MyNode)
  fmt.Println("is of type MyNode:", ok)//false

  n = new(MyNode)
  _, ok = n.(*MyNode) // is of MyNode pointer type]
  fmt.Println("is of type MyNode:", ok) // true

  n = MyNode{value:10}
  _, ok = n.(MyNode) // is of MyNode  type
  fmt.Println("is of type MyNode:", ok)//true

}
