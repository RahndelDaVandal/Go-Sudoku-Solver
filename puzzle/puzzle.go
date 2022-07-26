package puzzle

import (
	"fmt"
	"strings"
)

type Puzzle []int
type Puzzles []Puzzle

func (p *Puzzle) Print(){
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

func GenGridString() string{
  var g string
  lines := map[int]string{
    0 : "╔═══╤═══╦═══╗",
    1 : "║ . │ . ║ . ║",
    2 : "╟───┼───╫───╢",
    3 : "╠═══╪═══╬═══╣",
    4 : "╚═══╧═══╩═══╝",
  }
  g += expandLine(lines[0])
  g += midGridSection(lines)
  g += expandLine(lines[3])
  g += midGridSection(lines)
  g += expandLine(lines[3])
  g += midGridSection(lines)
  g += expandLine(lines[4])
  return g
}

func expandLine(l string) string{
  r := []rune(l)
  begin := string(r[0])
  end := string(r[9:13]) + "\n"
  midSep := string(r[5:9])
  midSlice := []string{}
  sliceStr := string(r[1:5]) + string(r[1:5])
  for i:=0; i<3; i++{midSlice = append(midSlice, sliceStr)}
  return begin + strings.Join(midSlice, midSep) + end
}

func midGridSection(m map[int]string) string{
  pattern := []int{1, 2, 1, 2, 1,}
  output := ""
  for _, v := range pattern {
    output += expandLine(m[v])
  }
  return output
}

