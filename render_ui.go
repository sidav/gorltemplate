package main

import (
	"fmt"
	cw "gorltemplate/console_wrapper"
)

func (r *renderer) renderUI() {
	cw.SetColor(cw.BLACK, cw.GREEN)
	r.putStringBottomLeft(fmt.Sprintf("TICK: %d", GAMETICK), 0)
	r.putStringBottomRight("UI rendered", 0)
	cw.SetBgColor(cw.BLACK)
	log.Render(r.consHeight - 3)
}

func (r *renderer) putStringBottomLeft(str string, heightFromBottom int) {
	cw.PutString(str, 0, r.consHeight-heightFromBottom-1)
}

func (r *renderer) putStringBottomRight(str string, heightFromBottom int) {
	length := len(str)
	cw.PutString(str, r.consWid-length, r.consHeight-heightFromBottom-1)
}
