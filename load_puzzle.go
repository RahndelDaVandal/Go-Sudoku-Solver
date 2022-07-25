package main

import (
	"fmt"
	"io/ioutil"
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
