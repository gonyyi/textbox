// (c) Gon Y. Yi 2021 <https://gonyyi.com/copyright>
// Last Update: 10/26/2021

package textbox_test

import (
   "os"
   "github.com/gonyyi/textbox"
   "testing"
)

const test = `{"name":"m1", "sub":[
   {"name":"m1-3", "sub":[
      {"name":"m1-3-3"},
      {"name":"m1-3-1"},
      {"name":"m1-3-2"}
   ]},
   {"name":"m1-1", "sub":[
      {"name":"m1-1-1"},
      {"name":"m1-1-2", "sub":[{"name":"m1-1-2-1"}, {"name":"m1-1-2-1", "sub":[{"name":"m1-1-2-1-1"}, {"name":"m1-1-2-1-2"}]}]}
   ]},
   {"name":"m1-2", "sub":[
      {"name":"m1-2-3"},
      {"name":"m1-2-1"},
      {"name":"m1-2-2"}
   ]}
]}`

func TestNewTreeFromJSON(t *testing.T) {
   c := textbox.NewTreeConfig()
   c.Indentation = 2
   tr, err := textbox.NewTreeFromJSON([]byte(test))
   if err != nil {println(err.Error()); return}
   tr.Follow(c, os.Stdout)

}
func TestNewTree(t *testing.T) {
   c:=textbox.NewTreeConfig()
   c.Indentation = 2

   tr := textbox.NewTree("Gon")
   tr = tr.AddSub(tr.NewSub("Name", tr.NewSub("LastName", tr.NewSub("Yi"))))
   tr = tr.AddSub(tr.NewSub("SSN", tr.NewSub("Number", tr.NewSub("123-45-6789")), tr.NewSub("Issued", tr.NewSub("01/01/2001"))), tr.NewSub("Issued", tr.NewSub("01/01/2001")))
   tr.Follow(c, os.Stdout)
}

