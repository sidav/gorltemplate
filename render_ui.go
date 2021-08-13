package main

import (
	"fmt"
	cw "gorltemplate/console_wrapper"
)

func (r *renderer) renderUI() {
	cw.SetColor(cw.BLACK, cw.GREEN)
	cw.PutString(fmt.Sprintf("TICK: %d", GAMETICK), 0, r.consHeight-2)
	cw.PutString("UI rendered", 0, r.consHeight-1)
}
