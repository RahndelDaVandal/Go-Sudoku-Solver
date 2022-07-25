package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)


func loadPuzzles(fileName string) []string {
  f, err := ioutil.ReadFile(fileName)
  if err != nil{
    fmt.Println("ERROR", err)
  }
  puzzles := strings.Split(string(f), "\n")
  puzzles = rmBlankPuzzles(puzzles)
  return puzzles
}

func rmBlankPuzzles(s []string) []string{
  var o []string
  for _, v := range s{
    if v != "" && len(v) == 81 {o = append(o, v)}
  }
  return o
}

func main(){
  fileName := os.Args[1]
  puzzles := loadPuzzles(fileName)

  fmt.Println(len(puzzles[0]))

  // puzzle := []rune(puzzles[0])

  // fmt.Println(len(puzzle))

  intPuzzle := make([]int, 0)

  for _, v := range puzzles[0]{
    n, _ := strconv.Atoi(string(v))
    intPuzzle = append(intPuzzle, n) 
  }
  fmt.Println(len(intPuzzle))
}

