package main

func (l *level) tryPerformMeleeAttack(attacker *actor, targetVectorX, targetVectorY int) bool {
	ax, ay := attacker.getCoords()
	tx, ty := ax + targetVectorX, ay + targetVectorY
	target := CURRENTLEVEL.getFirstActorAtCoords(tx, ty)
	if target != nil && target.team != attacker.team && target.team >= 0 {
		performMeleeAttack(attacker, target)
		return true
	}
	return false
}

func performMeleeAttack(attacker, target *actor) {
	// rewrite logic if neccessary
	damage := 1
	if attacker.inv.itemInHands != nil {
		damage = attacker.inv.itemInHands.asWeapon.damage
	}
	defense := 0
	if target.inv.equippedArmor != nil {
		defense = target.inv.equippedArmor.asArmor.defense
	}
	finalDamage := damage - defense
	target.hp -= finalDamage
	log.AppendMessagef("%s is hit by %s, %d dmg (%d att-%d def)", target.data.name, attacker.data.name, finalDamage, damage, defense)
}
