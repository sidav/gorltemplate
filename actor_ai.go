package main

const (
	AI_TYPE_SIMPLE int = iota
	AI_TYPE_WAYPOINT_BASED

	AI_STATE_WANDERING = iota
	AI_STATE_ATTACKING
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
	ai := a.ai

	if a.canAiSeePlayer() {
		ai.aiState = AI_STATE_ATTACKING
		ai.targetX, ai.targetY = PLAYERCONTROLLER.player.getCoords()
		log.AppendMessagef("%s notices you!", a.data.name)
	} else {
		// ai.aiState = AI_STATE_WANDERING
	}
}

// act according to aiState. TODO: actually consider aiState.
func (a *actor) aiBehave() {
	ai := a.ai

	if ai.aiState == AI_STATE_ATTACKING {
		a.aiAttack()
		return
	}

	if ai.aiType == AI_TYPE_WAYPOINT_BASED {
		a.navigateThroughWaypoints()
	} else {
		if !CURRENTLEVEL.isTilePassableFor(a.x+ai.vectorX, a.y+ai.vectorY, a) {
			ai.vectorX, ai.vectorY = rnd.RandInRange(-1, 1), rnd.RandInRange(-1, 1)
		}
		a.setIntent(INTENT_MOVE_OR_OPEN_DOOR, ai.vectorX, ai.vectorY, 10)
	}
}

func (a *actor) aiAttack() {
	vx, vy := CURRENTLEVEL.getVectorToNextCellForPathTo(a, a.ai.targetX, a.ai.targetY)
	log.AppendMessagef("%d, %d", vx, vy)
	a.setIntent(INTENT_MOVE_OR_OPEN_DOOR, vx, vy, 10)
}

func (a *actor) canAiSeePlayer() bool {
	return playerFovMap[a.x][a.y]
}
