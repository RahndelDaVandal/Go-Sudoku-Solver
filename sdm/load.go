package sdm

import (
	"bufio"
	"strconv"
	"log"
	"os"
  "sudoku_solver/puzzle"
)

type strPuzzles []string

func Load(fileName string) puzzle.Puzzles{
  var sP strPuzzles
  sP = readInSdmFile(fileName)
  sP.rmBlankPuzzles()
  return convertStringToInt(sP)
}

func readInSdmFile(fileName string) []string{
  file, err := os.Open(fileName)
  if err != nil{
    log.Fatalf("Failed to open %v", fileName)
  }

  scanner := bufio.NewScanner(file)
  scanner.Split(bufio.ScanLines)
  
  var sP strPuzzles
  for scanner.Scan(){
    sP = append(sP, scanner.Text())
  }
  file.Close()
  return sP
}

func (sP *strPuzzles) rmBlankPuzzles() {
  var o []string
  for _, v := range *sP{
    if v != "" && len(v) == 81 {o = append(o, v)}
  }
  *sP = o
}

func convertStringToInt(sP strPuzzles) puzzle.Puzzles{
  puzzles := make(puzzle.Puzzles, 0)

  for _, v := range sP{
    puzzles = append(puzzles, stringToInt(v))
  }
  return puzzles
}

func stringToInt(s string) puzzle.Puzzle{
  var p puzzle.Puzzle
  for _, v := range s{
    n, _ := strconv.Atoi(string(v))
    p = append(p, n)
  }
  return p
}
