package main

import(
  "fmt"
)

func main(){
  var map1 map[string]int
  map1 = make(map[string]int) // intialize
  map1["key1"] =  0

  // inline map
  outputs := map[string]map[string]int {
    "var declared map" : map1,
  }

  // print all map key values
  for k,v := range outputs {
    fmt.Println(k, " has ", v)
  }

  //check if a map has key
  v, ok := map1["invalid key"]
  if ok {
    fmt.Println("Value", v)
  } else {
    fmt.Println("Key not found")
  }

  // add new key/val
  map1["delme"] = 100
  fmt.Println("Map after add ", map1)

  //delete map key
  delete(map1, "delme")
  fmt.Println("Map after delete ", map1)

}
