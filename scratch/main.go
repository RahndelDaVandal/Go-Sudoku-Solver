package main

import "fmt"

type puzzle []int
type puzzles struct{}

func main(){
  nums := puzzle{2,0,5,6,4,0,3,0,}

  fmt.Println(nums, " : ", len(nums))
  nums.rmZeros()
  fmt.Println(nums, " : ", len(nums))
}

func (n *puzzle) rmZeros() {
  var o []int
  for _, v := range *n {
    if v != 0{
      o = append(o, v)
    }
  }
  *n = o
}
