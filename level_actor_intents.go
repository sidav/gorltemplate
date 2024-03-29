package main

func (l *level) tryExecuteActorsIntent(a *actor) bool {
	if a.intent == nil {
		return false
	}
	if !a.intent.isInstant() && a.intent.turnToComplete > GAMETICK {
		return true
	}
	switch a.intent.intentType {
	case INTENT_MOVE_ONLY:
		if l.tryMoveActorByVector(a, a.intent.vx, a.intent.vy) {
			a.intent.intentType = INTENT_WAIT
			return true
		}
	case INTENT_MOVE_OR_OPEN_DOOR:
		if l.tryActivateSwitchAsActor(a, a.intent.vx, a.intent.vy) {
			a.intent.intentType = INTENT_WAIT
			return true
		}
		if l.tryOpenDoorForActor(a, a.intent.vx, a.intent.vy) {
			a.intent.intentType = INTENT_WAIT
			return true
		}
		if l.tryMoveActorByVector(a, a.intent.vx, a.intent.vy) {
			a.intent.intentType = INTENT_WAIT
			return true
		}
		if l.tryPerformMeleeAttack(a, a.intent.vx, a.intent.vy) {
			a.intent.intentType = INTENT_WAIT
			return true
		}
	case INTENT_MELEE_ATTACK:
		if l.tryPerformMeleeAttack(a, a.intent.vx, a.intent.vy) {
			a.intent.intentType = INTENT_WAIT
			return true
		}
	}
	return false
}
