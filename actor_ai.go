package main

const (
	AI_TYPE_SIMPLE int = iota
	AI_TYPE_WAYPOINT_BASED
)

type actorAi struct {
	aiType int
	aiState int

	targetX, targetY int
	vectorX, vectorY int

	waypoints [][2]int
	currentWaypointIndex int
	traverseDirection int
}

func (a *actor) initializeAI(aiType int) {
	a.ai = &actorAi{
		aiType:    aiType,
	}
	if aiType == AI_TYPE_WAYPOINT_BASED {
		a.ai.initAndGenerateWaypoints(a)
		a.ai.traverseDirection = 1
	}
}

func (a *actor) aiAct() {
	if a.ai == nil || !a.isTimeToAct() {
		return
	}
	a.aiObserve()
	a.aiBehave()
}

// maybe change current ai state?
func (a *actor) aiObserve() {

}

// act according to aiState. TODO: actually consider aiState.
func (a *actor) aiBehave() {
	ai := a.ai
	if ai.aiType == AI_TYPE_WAYPOINT_BASED {
		a.navigateThroughWaypoints()
	} else {
		if !CURRENTLEVEL.isTilePassable(a.x+ai.vectorX, a.y+ai.vectorY) {
			ai.vectorX, ai.vectorY = rnd.RandInRange(-1, 1), rnd.RandInRange(-1, 1)
		}
		a.setIntent(INTENT_MOVE_OR_OPEN_DOOR, ai.vectorX, ai.vectorY, 10)
	}
}
