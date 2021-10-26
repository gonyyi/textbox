// (c) Gon Y. Yi 2021 <https://gonyyi.com/copyright>
// Last Update: 10/26/2021

package textbox

import (
   "reflect"
   "sort"
)

// ┌ ─ ┬ ─ ┐ TOP_LEFT(┌), HORIZONTAL, HORIZONTAL_DOWN(┬), HORIZONTAL, TOP_RIGHT(┐)
// │       │ VERTICAL
// ├       ┤ VERTICAL_RIGHT(├), VERTICAL_LEFT(┤)
// └ ─ ┴ ─ ┘ BOTTOM_LEFT(└), HORIZONTAL, HORIZONTAL_UP(┴), HORIZONTAL, BOTTOM_RIGHT(┘)
const (
   HORIZONTAL      = '─'
   HORIZONTAL_DOWN = '┬'
   HORIZONTAL_UP   = '┴'
   TOP_LEFT        = '┌'
   TOP_RIGHT       = '┐'
   BOTTOM_LEFT     = '└'
   BOTTOM_RIGHT    = '┘'
   VERTICAL        = '│'
   VERTICAL_LEFT   = '┤'
   VERTICAL_RIGHT  = '├'
)

type BoxConnPosition uint8

const (
   BOX_CONN_TOP BoxConnPosition = iota
   BOX_CONN_BOTTOM
   BOX_CONN_LEFT
   BOX_CONN_RIGHT
)

func repeatRune(b rune, n int) []rune {
   var out []rune
   for i := 0; i < n; i++ {
      out = append(out, b)
   }
   return out
}

func repeatByte(b byte, n int) []byte {
   var out []byte
   for i := 0; i < n; i++ {
      out = append(out, b)
   }
   return out
}

// sortTrees will sort by name and then index.
func sortTrees(trees []Tree) {
   sort.Slice(trees, func(i, j int) bool {
      if trees[i].Name < trees[j].Name {
         return true
      }
      if trees[i].Name > trees[j].Name {
         return false
      }
      return i < j
   })
}

func orderedMapKey(someMap interface{}) []string {
   var out []string
   switch reflect.TypeOf(someMap).Kind() {
   case reflect.Map:
      s := reflect.ValueOf(someMap)
      for _, v := range s.MapKeys() {
         out = append(out, v.String())
      }
   }
   sort.Strings(out)
   return out
}

