package main

import(
  "fmt"
)

type IntWrapper int

func (i IntWrapper) printVal(){
  fmt.Println("Value is", i)
}

func (iw *IntWrapper) setValue(i int){
  *iw = IntWrapper(i)
}

func main(){
  iw := new (IntWrapper)
  iw.setValue(100)
  iw.printVal()

  i := IntWrapper(10)
  i.printVal()
}
