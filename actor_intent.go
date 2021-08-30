package main

// Intent represents an "undergoing" action, which takes more than one tick.
// The whole mechanism is needed to implement a "wait now, act later" mechanic for multi-turn actions, instead of
// my usual "act now, wait later" approach.

type intentTypeCode uint8

const (
	INTENT_NOTHING intentTypeCode = iota
	INTENT_WAIT
	INTENT_MOVE_OR_OPEN_DOOR
	INTENT_MOVE_ONLY
	INTENT_ATTACK
)

type actorIntent struct {
	intentType intentTypeCode
	vx, vy int
	turnToComplete int
}

func (aInt *actorIntent) isInstant() bool {
	return aInt.intentType == INTENT_MOVE_ONLY || aInt.intentType == INTENT_MOVE_OR_OPEN_DOOR
}
