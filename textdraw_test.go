// (c) Gon Y. Yi 2021 <https://gonyyi.com/copyright>
// Last Update: 10/26/2021

package textbox

import (
   "testing"
)

func TestOrderedMapKey(t *testing.T) {
   m1 := make(map[string]map[string][]string)
   m2 := make(map[string]int)
   m3 := make(map[string]string)

   m1["m1-1"] = nil
   m1["m1-3"] = nil
   m1["m1-2"] = nil

   m2["m2-3"] = 3
   m2["m2-1"] = 1
   m2["m2-2"] = 2

   m3["m3-1"] = "a"
   m3["m3-3"] = "c"
   m3["m3-2"] = "b"

   f := func(name string, s []string) {
      println(name)
      for i, v := range s {
         println("\t", i, v)
      }
   }

   f("m1", orderedMapKey(m1))
   f("m2", orderedMapKey(m2))
   f("m3", orderedMapKey(m3))
}

