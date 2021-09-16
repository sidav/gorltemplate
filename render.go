package main

import (
	cw "gorltemplate/console_wrapper"
	"strconv"
)

type renderer struct {
	ignoreFOV bool
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
				staticData := CURRENTLEVEL.tiles[gx][gy].getStaticData()
				if fovMap[gx][gy] || r.ignoreFOV {
					cw.SetColor(staticData.fgcolor, staticData.bgcolor)
					cw.PutChar(staticData.char, sx, sy)

				} else if CURRENTLEVEL.tiles[gx][gy].wasSeenPreviously {
					fg, bg := staticData.getUnseenColors()
					cw.SetColor(fg, bg)
					cw.PutChar(staticData.char, sx, sy)
				}
			}
		}
	}
	cw.SetColor(cw.WHITE, cw.BLACK)
}

func (r *renderer) renderActors(fovMap [][]bool) {
	for _, a := range CURRENTLEVEL.actors {
		if a.ai != nil && r.ignoreFOV {
			for i := 0; i < len(a.ai.waypoints); i++ {
				aix, aiy := r.globalCoordsToViewport(a.ai.waypoints[i][0], a.ai.waypoints[i][1])
				cw.SetColor(cw.MAGENTA, cw.BLACK)
				cw.PutString(strconv.Itoa(i), aix, aiy)
			}
		}
		if fovMap[a.x][a.y] || r.ignoreFOV {
			ax, ay := r.globalCoordsToViewport(a.x, a.y)
			cw.SetFgColor(a.data.color)
			cw.PutChar(a.data.char, ax, ay)
		}
	}
	cw.SetColor(cw.WHITE, cw.BLACK)
}
