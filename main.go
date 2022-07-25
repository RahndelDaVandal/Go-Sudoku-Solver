package main

import (
	"fmt"
	"os"
)

func main(){
  fileName := os.Args[1]
  puzzles := loadPuzzles(fileName)
  for _, p := range puzzles{
    fmt.Println(p)
  }
}

