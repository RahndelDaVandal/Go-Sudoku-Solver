package puzzle

import (
	"fmt"
)

type Puzzle []int
type Puzzles []Puzzle

func (p *Puzzle) Show(){
  buf := ""
  colCount := 0
  rowCount := 0
  for _, n := range *p{
    if rowCount == 3{
      buf += "---------------------\n"
      rowCount = 0
    }

    if colCount == 8{
      buf += fmt.Sprint(n) + "\n"
      colCount = 0
      rowCount ++
    }else if colCount == 2 || colCount == 5{
      buf += fmt.Sprint(n) + " | "
      colCount ++
    } else{ 
      buf += fmt.Sprint(n) + " "
      colCount ++
    }
  }
  buf += "\n"
  fmt.Printf(buf)
}
