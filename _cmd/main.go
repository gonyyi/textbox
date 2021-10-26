package main

import (
   "os"
   "github.com/gonyyi/textdraw"
)

func main() {
   var b textdraw.Box
   b = textdraw.NewBox(".:.:.", "Welcome to GAT Gateway", "gat.local.gonyyi.com", "10.0.20.10")
   b.Output(os.Stdout)
}
