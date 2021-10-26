// (c) Gon Y. Yi 2021 <https://gonyyi.com/copyright>
// Last Update: 10/26/2021

package textdraw

import "io"

type Box [][]rune

func (b Box) Output(w io.Writer) {
   if w != nil {
      w.Write([]byte(b.string(true)))
   }
}
func (b Box) String() string {
   return b.string(false)
}
func (b Box) string(endingNewline bool) string {
   var buf []rune
   for _, line := range b {
      buf = append(buf, line...)
      buf = append(buf, '\n')
   }
   if endingNewline == false {
      if lBuf := len(buf); lBuf > 0 {
         return string(buf[:lBuf-1])
      }
   }
   return string(buf)
}

type FormatFn func([]rune) []rune

func (b Box) Format(f FormatFn) Box {
   for idx, line := range b {
      b[idx] = f(line)
   }
   return b
}

func (b Box) AddPrefix(prefix string) Box {
   var pre = []rune(prefix)
   return b.Format(func(line []rune) []rune {
      return append(pre, line...)
   })
}

func (b Box) AddConnector(posBits BoxConnPosition, index int) Box {
   lastIdx := len(b) - 1
   for idx, line := range b {
      if idx == 0 && posBits == BOX_CONN_TOP { // FIRST LINE
         line[index] = HORIZONTAL_UP
      } else if idx == lastIdx && posBits == BOX_CONN_BOTTOM { // LAST LINE
         line[index] = HORIZONTAL_DOWN
      } else if idx != 0 && idx != lastIdx && idx == index { // MIDDLE LINES (LEFT and RIGHT)
         if posBits == BOX_CONN_LEFT {
            line[0] = VERTICAL_LEFT
         }
         if posBits == BOX_CONN_RIGHT {
            line[len(line)-1] = VERTICAL_RIGHT
         }
      } else {
         continue // if nothing was chosen, pass
      }
      b[idx] = line // update the line
   }
   return b
}

func NewBox(s ...string) Box {
   var out [][]rune
   var maxWidth int
   // Get maxWidth
   for _, line := range s {
      if lLine := len(line); maxWidth < lLine {
         maxWidth = lLine
      }
   }

   // Draw TOP
   out = append(out, append([]rune{TOP_LEFT}, append(repeatRune(HORIZONTAL, maxWidth+2), TOP_RIGHT)...))
   // Draw Body
   for _, v := range s {
      out = append(out, append(append([]rune{VERTICAL, ' '},
         append([]rune(v), repeatRune(' ', maxWidth-len(v))...)...), ' ', VERTICAL))
   }
   // Draw Bottom
   out = append(out, append([]rune{BOTTOM_LEFT}, append(repeatRune(HORIZONTAL, maxWidth+2), BOTTOM_RIGHT)...))
   return out
}
