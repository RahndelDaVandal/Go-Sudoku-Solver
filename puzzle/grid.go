package puzzle

import (
	"strings"
)

var Grid = GenGridString()

var lines = map[int][]rune{
  0 : []rune("╔═══╤═══╦═══╗"),
  1 : []rune("║ . │ . ║ . ║"),
  2 : []rune("╟───┼───╫───╢"),
  3 : []rune("╠═══╪═══╬═══╣"),
  4 : []rune("╚═══╧═══╩═══╝"),
}

func GenGridString() string{
  mid := func() string{
    output := ""
    for _, v := range []int{1, 2, 1, 2, 1,} {
      output += expandLine(lines[v])
    }
    return output
  }

  g := expandLine(lines[0])
  for i := 1; i < 6; i++ {
    if i % 2 == 0 {
      g += expandLine(lines[3])
    } else {
      g += mid()
    }
  }
  g += expandLine(lines[4])

  return g
}

func expandLine(l []rune) string{
  midSep := string(l[5:9])
  midStr := string(l[1:5]) + string(l[1:5])
  midSlice := []string{midStr, midStr, midStr}
  midSection := strings.Join(midSlice, midSep)

  if l[0] != lines[4][0]{
    return string(l[0]) + midSection + string(l[9:13]) + "\n"
  }
  return string(l[0]) + midSection + string(l[9:13])
}
