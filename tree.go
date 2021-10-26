// (c) Gon Y. Yi 2021 <https://gonyyi.com/copyright>
// Last Update: 10/26/2021

package textdraw

import (
   "encoding/json"
   "fmt"
   "io"
)

type TreeConfig struct {
   Indentation       int      // how many spaces for each
   PrintMaster       bool     // if print master or not
   PrintMasterBox    bool     // if true, draw a box to the master
   MasterBoxComments []string // comments goes inside the box
   MasterAtBottom    bool
   SortSubTree       bool

   BoxConnectBothDirection bool
   strIndentation          string
}

const (
   treeLineMaster   = "── "
   treeLineSubFirst = "┌─ " // will be used only when reverse
   treeLineSubMore  = "├─ "
   treeLineSubLast  = "└─ "
   treeLinePassing  = "│  "
   treeLineBlank    = "   "
)

func NewTreeConfig() *TreeConfig {
   return &TreeConfig{
      Indentation:             8,
      PrintMaster:             true,
      PrintMasterBox:          true,
      MasterBoxComments:       []string{},
      MasterAtBottom:          false,
      SortSubTree:             false,
      BoxConnectBothDirection: false,
   }
}

func NewTree(name string, sub ...Tree) Tree {
   return Tree{Name: name, Sub: sub}
}

func NewTreeFromJSON(b []byte) (*Tree, error) {
   var t Tree
   if err := json.Unmarshal(b, &t); err != nil {
      return nil, err
   }
   return &t, nil
}

type Tree struct {
   Name string `json:"name"`
   Sub  []Tree `json:"sub,omitempty"`
}

func (Tree) NewSub(name string, sub ...Tree) Tree {
   return NewTree(name, sub...)
}

func (s Tree) AddSub(sub ...Tree) Tree {
   s.Sub = append(s.Sub, sub...)
   return s
}

func (s Tree) subMaxIndex() int {
   return len(s.Sub) - 1
}

func (s Tree) Follow(config *TreeConfig, dst io.Writer) {
   config.strIndentation = string(repeatByte(' ', config.Indentation))

   printMaster := func(position BoxConnPosition) {
      if config.PrintMaster {
         if config.PrintMasterBox {
            boxName := []string{s.Name}
            if len(config.MasterBoxComments) > 0 {
               boxName = append(boxName, config.MasterBoxComments...)
            }
            box := NewBox(boxName...)
            if config.BoxConnectBothDirection {
               box = box.AddConnector(BOX_CONN_TOP, config.Indentation)
               box = box.AddConnector(BOX_CONN_BOTTOM, config.Indentation)
            } else {
               box = box.AddConnector(position, config.Indentation)
            }
            fmt.Fprintf(dst, "%s\n", box.String())
         } else {
            fmt.Fprintf(dst, "%s\n", treeLineMaster+s.Name)
         }
      }
   }

   // If master to be at the bottom
   if config.MasterAtBottom {
      s.followSub("", config, dst, true)
      printMaster(BOX_CONN_TOP)
      return
   } else {
      // Normal order
      printMaster(BOX_CONN_BOTTOM)
      s.followSub("", config, dst, true)
      return
   }
}

func (s Tree) follow(prefix string, index, maxIndex int, config *TreeConfig, dst io.Writer, firstDepth bool) {
   if config.MasterAtBottom && firstDepth {
      // If reverse and first depth, then the line should go all the way down.
      fmt.Fprintf(dst, "%s\n", s.String(prefix, index, maxIndex, true))
      prefix += treeLinePassing

   } else {
      fmt.Fprintf(dst, "%s\n", s.String(prefix, index, maxIndex, false))
      if len(s.Sub) > 0 && index != maxIndex {
         prefix += treeLinePassing
      } else {
         prefix += treeLineBlank
      }
   }

   s.followSub(prefix, config, dst, false)
}

func (s Tree) followSub(prefix string, config *TreeConfig, dst io.Writer, firstDepth bool) {
   // prefix += config.strIndentation

   if config.SortSubTree {
      sortTrees(s.Sub)
   }

   for subIdx, v := range s.Sub {
      v.follow(prefix+config.strIndentation, subIdx, s.subMaxIndex(), config, dst, firstDepth)
   }
}

// String will print current line only
func (s Tree) String(prefix string, index, maxIndex int, reverse bool) string {
   if reverse {
      if index == maxIndex {
         // return prefix + treeLineSubLast + s.Name
         return prefix + treeLineSubMore + s.Name
      }
      // First record
      if index == 0 && maxIndex > 0 {
         return prefix + treeLineSubFirst + s.Name
      }

      return prefix + treeLineSubMore + s.Name
   }

   if index == maxIndex {
      return prefix + treeLineSubLast + s.Name
   }
   return prefix + treeLineSubMore + s.Name
}

