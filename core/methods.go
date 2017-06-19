package main

import(
  "fmt"
)

//struct, similar to classes
type Person struct {
  name string
  age int
}

// method to Person reciever
func (p Person) printMe(ctx string) {
  fmt.Println(ctx, p)
}

// pointer based method
func (p *Person) change(){
  p.age = p.age + 1 // will change the incoming
}

// value based method
func (p Person) tryAndNotChange(){
  p.age = p.age + 1 // should not change the incoming
}

func main(){
  p1 := Person{
    name: "Person1",
    age: 35, // comma important for go compiler
  }

  p2 := Person{
    name: "Person2",
    age: 39,
  }

  p1.printMe("original p1")
  p2.printMe("original p2")

  // value based method calls, not expected to change the source
  p1.tryAndNotChange()
  p1.printMe("p1 after value based change method")

  p2.tryAndNotChange()
  p2.printMe("p2 after value based change method")

  // reference based method calls
  p1.change()
  p1.printMe("p1 after pointer based change method")

  p2.change()
  p2.printMe("p2 after pointer based change method")

}
