package main

import(
  "fmt"
)

func main(){
  //inline creation
  slice := []int {2,4,6,8,10,12}

  // creation by var declaration
  var subSlices [][]int //slice of subSlices

  //using append to fill subSlices
  subSlices = append(subSlices, slice[1:3], slice[:3], slice[3:])

  // loop through all subslices and print them
  for i, v := range subSlices {
    fmt.Println(i, " = ", v)
  }

  // internals
  // define a slice
  original := []int{2,4,6,8,10}
  fmt.Println("Cap", cap(original), ", length ", len(original))

  //sublicing does not affect the underlying original structure
  sub1 := original[2:3] // capacity is 3, even though len is 1
  fmt.Println("sub1", sub1, ", cap ", cap(sub1), ", len ", len(sub1))

  // change the subslice, changes original
  sub1[0] = 55
  fmt.Println("sub ", sub1, " original ", original)

  // copy will create a branch new slice
  sub2 := make([]int, 2)
  n := copy(sub2, original[2:4])
  fmt.Println(" copied ", n, " items", " sub2 ", sub2, " cap ", cap(sub2))

  // changing copied subslice won't affect original
  sub2[0] = 100
  fmt.Println(" sub2 ", sub2, " original ", original)

  //append
  sliceCombined := []int{1}
  sliceCombined = append(sliceCombined, original...)
  fmt.Println("sliceCombined", sliceCombined)

  //modify
  //remove #1
  sliceCombined = append(sliceCombined[:1], sliceCombined[2:]...)
  fmt.Println("sliceCombined", sliceCombined)

}
