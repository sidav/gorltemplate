package main

type actor struct {
	data *actorStatic
	intent *actorIntent

	team int
	ai   *actorAi

	x, y      int
	hp        int
}

func (a *actor) isTimeToAct() bool {
	return a.intent == nil || a.intent.turnToComplete <= GAMETICK
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

func (a *actor) decide() {
	if a.ai == nil {
		return
	}
}
