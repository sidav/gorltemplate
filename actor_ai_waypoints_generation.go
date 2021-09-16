package main

func (ai *actorAi) initAndGenerateWaypoints(actor *actor) {
	const MAXWAYPOINTS = 10
	const MINWAYPOINTS = 4


	ai.waypoints = make([][2]int, 0)
	ai.waypoints = append(ai.waypoints, [2]int{actor.x, actor.y})
	for i := 1; i < MAXWAYPOINTS; i++ {
		found, x, y := selectCoordsForNewWaypointOrthogonally(ai.waypoints[i-1][0], ai.waypoints[i-1][1], i)
		if found {
			ai.waypoints = append(ai.waypoints, [2]int{x, y})
		} else {
			if i > MINWAYPOINTS {
				return
			}
			ai.waypoints = ai.waypoints[:i-1]
			i--
			if i == 0 {
				// panic("Waypoints can't be created!")
				return
			}
		}
	}
}

func selectCoordsForNewWaypointOrthogonally(prevx, prevy, newWaypointNumber int) (bool, int, int) {
	const MINWAYPOINTRADIUS = 3
	const MAXWAYPOINTRADIUS = 100
	candidates := make([][3]int, 0) // [x, y, weight]

	nextIsHorizontal := newWaypointNumber % 2 == 0
	if nextIsHorizontal {
		for x := prevx; x <= prevx+MAXWAYPOINTRADIUS; x++ {
			if !CURRENTLEVEL.isTilePassable(x, prevy) {
				break
			}
			if x < prevx+MINWAYPOINTRADIUS {
				continue
			}
			candidates = append(candidates, [3]int{x, prevy, (prevx-x)*(prevx-x)})
		}
		for x := prevx; x >= prevx-MAXWAYPOINTRADIUS; x-- {
			if !CURRENTLEVEL.isTilePassable(x, prevy) {
				break
			}
			if x > prevx-MINWAYPOINTRADIUS {
				continue
			}
			candidates = append(candidates, [3]int{x, prevy, (prevx-x)*(prevx-x)})
		}
	} else {
		for y := prevy; y <= prevy+MAXWAYPOINTRADIUS; y++ {
			if !CURRENTLEVEL.isTilePassable(prevx, y) {
				break
			}
			if y < prevy+MINWAYPOINTRADIUS {
				continue
			}
			candidates = append(candidates, [3]int{prevx, y, (prevy-y)*(prevy-y)})
		}
		for y := prevy; y >= prevy-MAXWAYPOINTRADIUS; y-- {
			if !CURRENTLEVEL.isTilePassable(prevx, y) {
				break
			}
			if y > prevy-MINWAYPOINTRADIUS {
				continue
			}
			candidates = append(candidates, [3]int{prevx, y, (prevy-y)*(prevy-y)})
		}
	}
	if len(candidates) > 0 {
		randIndex := rnd.SelectRandomIndexFromWeighted(len(candidates), func(i int) int {return candidates[i][2]})
		return true, candidates[randIndex][0], candidates[randIndex][1]
	}
	return false, 0, 0
}

// path- and distance-based waypoint generation
func areCoordsGoodForWaypointMethod1() {
	const MAXWAYPOINTPATHLENGTH = 50
}
