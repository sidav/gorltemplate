package main

func (l *level) executeActorsIntent(a *actor) {
	if a.intent == nil {
		return
	}
	if a.intent.intentType != INTENT_MOVE_OR_OPEN_DOOR && a.intent.turnToComplete > GAMETICK {
		return
	}
	switch a.intent.intentType {
	case INTENT_MOVE_OR_OPEN_DOOR:
		if l.tryActivateSwitchAsActor(a, a.intent.vx, a.intent.vy) {
			a.intent.intentType = INTENT_WAIT
			return
		}
		if l.tryOpenDoorForActor(a, a.intent.vx, a.intent.vy) {
			a.intent.intentType = INTENT_WAIT
			return
		}
		if l.tryMoveActorByVector(a, a.intent.vx, a.intent.vy) {
			a.intent.intentType = INTENT_WAIT
			return
		}
	}
}
