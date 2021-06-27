package main

type actor struct {
	data *actorStatic

	team int
	ai   *actorAi

	tickToAct int
	x, y      int
	hp        int
}

func (a *actor) isTimeToAct() bool {
	return GAMETICK <= a.tickToAct
}

func (a *actor) spendTime(amount int) {
	a.tickToAct = GAMETICK + amount
}

func (a *actor) act() {
	if a.ai == nil {
		return
	}
}
