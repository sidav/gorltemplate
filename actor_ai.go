package main

type actorAi struct {
	aiType int
	aiState int

	targetX, targetY int
	vectorX, vectorY int
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
	if !CURRENTLEVEL.isTilePassable(a.x+ai.vectorX, a.y+ai.vectorY) {
		ai.vectorX, ai.vectorY = rnd.RandInRange(-1, 1), rnd.RandInRange(-1, 1)
	}
	a.setIntent(INTENT_MOVE_OR_OPEN_DOOR, ai.vectorX, ai.vectorY, 10)
}
