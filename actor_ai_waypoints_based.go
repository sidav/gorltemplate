package main

func (a *actor) getDirectionVectorToTarget(tx, ty int) (int, int) {
	vx := 0
	vy := 0 
	if tx < a.x {
		vx = -1 
	} else if tx > a.x {
		vx = 1 
	}
	if ty < a.y {
		vy = -1
	} else if ty > a.y {
		vy = 1
	}
	return vx, vy
}

func (ai *actorAi) setNextWaypoint() {
	if ai.currentWaypointIndex == len(ai.waypoints)-1 && ai.traverseDirection == 1 {
		ai.traverseDirection = -1
	}
	if ai.currentWaypointIndex == 0 && ai.traverseDirection == -1 {
		ai.traverseDirection = 1
	}
	ai.currentWaypointIndex += ai.traverseDirection
}

func (a *actor) navigateThroughWaypoints() {
	ai := a.ai
	nextx, nexty := ai.waypoints[ai.currentWaypointIndex][0], ai.waypoints[ai.currentWaypointIndex][1]
	if a.x == nextx && a.y == nexty {
		ai.setNextWaypoint()
		nextx, nexty = ai.waypoints[ai.currentWaypointIndex][0], ai.waypoints[ai.currentWaypointIndex][1]
	}
	vx, vy := a.getDirectionVectorToTarget(nextx, nexty)
	a.setIntent(INTENT_MOVE_OR_OPEN_DOOR, vx, vy, 10)
}
