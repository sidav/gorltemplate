package main

import cw "gorltemplate/console_wrapper"

type renderer struct {
	consWid, consHeight                          int
	viewportCenterScreenX, viewportCenterScreenY int
	viewportCenterGlobalX, viewportCenterGlobalY int
}

func (r *renderer) render(cx, cy int, fovMap [][]bool) {
	r.viewportCenterGlobalX, r.viewportCenterGlobalY = cx, cy
	r.consWid, r.consHeight = cw.GetConsoleSize()
	r.viewportCenterScreenX, r.viewportCenterScreenY = r.consWid/2, r.consHeight/2

	cw.SetColor(cw.WHITE, cw.BLACK)
	cw.Clear_console()

	r.renderLevel(fovMap)
	r.renderActors(fovMap)
	r.renderUI()

	cw.Flush_console()
}

func (r *renderer) renderLevel(fovMap [][]bool) {
	for sx := 0; sx < r.consWid; sx++ {
		for sy := 0; sy < r.consHeight; sy++ {
			gx, gy := r.viewportCoordsToGlobal(sx, sy)
			if CURRENTLEVEL.coordsValid(gx, gy) {
				if fovMap[gx][gy] {
					cw.SetColor(CURRENTLEVEL.tiles[gx][gy].data.fgcolor, CURRENTLEVEL.tiles[gx][gy].data.bgcolor)
					cw.PutChar(CURRENTLEVEL.tiles[gx][gy].data.char, sx, sy)

				} else if CURRENTLEVEL.tiles[gx][gy].wasSeenPreviously {
					fg, bg := CURRENTLEVEL.tiles[gx][gy].data.getUnseenColors()
					cw.SetColor(fg, bg)
					cw.PutChar(CURRENTLEVEL.tiles[gx][gy].data.char, sx, sy)
				}
			}
		}
	}
	cw.SetColor(cw.WHITE, cw.BLACK)
}

func (r *renderer) renderActors(fovMap [][]bool) {
	for _, a := range CURRENTLEVEL.actors {
		if fovMap[a.x][a.y] {
			ax, ay := r.globalCoordsToViewport(a.x, a.y)
			cw.SetFgColor(a.data.color)
			cw.PutChar(a.data.char, ax, ay)
		}
	}
	cw.SetColor(cw.WHITE, cw.BLACK)
}
