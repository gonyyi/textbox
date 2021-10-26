package main

import (
   "os"
   "github.com/gonyyi/textbox"
)

func main() {
   var b textbox.Box
   b = textbox.NewBox(".:.:.", "Welcome to GAT Gateway", "gat.local.gonyyi.com", "10.0.20.10")
   b.Output(os.Stdout)
}
