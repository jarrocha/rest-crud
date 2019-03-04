// +build !testing

package view

import "text/template"

var TPL *template.Template

func Initialize() {
	TPL = template.Must(template.ParseGlob("view/*.gohtml"))
}
