package main

import cw "gorltemplate/console_wrapper"

func (r *renderer) renderUI() {
	cw.SetColor(cw.BLACK, cw.GREEN)
	cw.PutString("UI rendered", 0, r.consHeight-1)
}
