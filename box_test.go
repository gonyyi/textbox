// (c) Gon Y. Yi 2021 <https://gonyyi.com/copyright>
// Last Update: 10/26/2021

package textbox_test

import (
   "os"
   "github.com/gonyyi/textbox"
   "testing"
)

func TestNewBox(t *testing.T) {
   var b textbox.Box
   b = textbox.NewBox("Gon")
   b.AddConnector(textbox.BOX_CONN_TOP, 3).Output(os.Stdout)
   b.AddConnector(textbox.BOX_CONN_LEFT, 1).Output(os.Stdout)
   b.AddConnector(textbox.BOX_CONN_RIGHT, 1).Output(os.Stdout)
   b.AddConnector(textbox.BOX_CONN_BOTTOM, 3).Output(os.Stdout)

   b = textbox.NewBox("Blah Blah v1.0.0", "(C) Gon Y. Yi 2021")
   b.AddConnector(textbox.BOX_CONN_TOP, 3).Output(os.Stdout)
   b.AddConnector(textbox.BOX_CONN_LEFT, 1).Output(os.Stdout)
   b.AddConnector(textbox.BOX_CONN_RIGHT, 1).Output(os.Stdout)
   b.AddConnector(textbox.BOX_CONN_BOTTOM, 3).Output(os.Stdout)

   b = textbox.NewBox(".:.:.", "Welcome to GAT Gateway", "gat.local.gonyyi.com", "10.0.20.10")
   b.Output(os.Stdout)
}

