package main

type actor struct {
	data   *actorStatic
	intent *actorIntent

	team int // -1 means neutral, 0 means enemy to everyone
	ai   *actorAi

	x, y int
	hp   int

	inv *inventory
}

func (a *actor) init() {
	a.hp = a.data.maxhp
	if a.inv == nil {
		a.inv = &inventory{}
	}
}

func (a *actor) isTimeToAct() bool {
	return a.intent == nil || a.intent.turnToComplete <= GAMETICK
}

func (a *actor) getCoords() (int, int) {
	return a.x, a.y
}

func (a *actor) setIntent(intType intentTypeCode, vx, vy, duration int) {
	if a.intent == nil {
		a.intent = &actorIntent{}
	}
	a.intent.intentType = intType
	a.intent.vx = vx
	a.intent.vy = vy
	a.intent.turnToComplete = GAMETICK + duration
}

//func (a *actor) spendTime(amount int) {
//	a.tickToAct = GAMETICK + amount
//}
