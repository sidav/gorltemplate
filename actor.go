package main

type actor struct {
	data *actorStatic
	x, y int
	hp int
}

type actorStatic struct {
	maxhp int
	char rune
	color int
}
