package main

func (r *renderer) globalCoordsToViewport(gx, gy int) (int, int) {
	return gx - r.viewportCenterGlobalX + r.viewportCenterScreenX, gy - r.viewportCenterGlobalY + r.viewportCenterScreenY
}

func (r *renderer) viewportCoordsToGlobal(vx, vy int) (int, int) {
	return vx + r.viewportCenterGlobalX - r.viewportCenterScreenX, vy + r.viewportCenterGlobalY - r.viewportCenterScreenY
}
