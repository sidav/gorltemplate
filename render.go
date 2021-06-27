package main

import cw "gorltemplate/console_wrapper"

type renderer struct {
	consWid, consHeight                          int
	viewportCenterScreenX, viewportCenterScreenY int
	viewportCenterGlobalX, viewportCenterGlobalY int
}

func (r *renderer) render(cx, cy int) {
	r.viewportCenterGlobalX, r.viewportCenterGlobalY = cx, cy
	r.consWid, r.consHeight = cw.GetConsoleSize()
	r.viewportCenterScreenX, r.viewportCenterScreenY = r.consWid/2, r.consHeight/2

	cw.SetColor(cw.WHITE, cw.BLACK)
	cw.Clear_console()

	r.renderLevel()
	r.renderActors()
	r.renderUI()

	cw.Flush_console()
}

func (r *renderer) renderLevel() {
	for sx := 0; sx < r.consWid; sx++ {
		for sy := 0; sy < r.consHeight; sy++ {
			gx, gy := r.viewportCoordsToGlobal(sx, sy)
			if CURRENTLEVEL.coordsValid(gx, gy) {
				cw.SetColor(CURRENTLEVEL.tiles[gx][gy].data.fgcolor, CURRENTLEVEL.tiles[gx][gy].data.bgcolor)
				cw.PutChar(CURRENTLEVEL.tiles[gx][gy].data.char, sx, sy)
			}
		}
	}
}

func (r *renderer) renderActors() {
	for _, a := range CURRENTLEVEL.actors {
		ax, ay := r.globalCoordsToViewport(a.x, a.y)
		cw.SetFgColor(a.data.color)
		cw.PutChar(a.data.char, ax, ay)
	}
}

func (r *renderer) renderUI() {
	cw.SetColor(cw.BLACK, cw.GREEN)
	cw.PutString("UI rendered", 0, r.consHeight-1)
}

func (r *renderer) globalCoordsToViewport(gx, gy int) (int, int) {
	return gx - r.viewportCenterGlobalX + r.viewportCenterScreenX, gy - r.viewportCenterGlobalY + r.viewportCenterScreenY
}

func (r *renderer) viewportCoordsToGlobal(vx, vy int) (int, int) {
	return vx + r.viewportCenterGlobalX - r.viewportCenterScreenX, vy + r.viewportCenterGlobalY - r.viewportCenterScreenY
}
