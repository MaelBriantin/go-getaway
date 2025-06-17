package game

import (
	"strconv"
	"strings"
)

type Character struct {
	Class     Class
	Name      string
	MaxHealth uint
	Health    uint
	Attack    uint
	Arcane    uint
}

func CreateCharacter(name string, class Class) Character {
	return Character{
		Name:      strings.ToUpper(name),
		Class:     class,
		Health:    class.Health,
		MaxHealth: class.Health,
		Attack:    class.Attack,
		Arcane:    class.Arcane,
	}
}

func (c Character) GetCharacterStats() map[string]string {
	return map[string]string{
		"Name":   c.Name,
		"Class":  c.Class.Name,
		"Health": strconv.Itoa(int(c.Health)),
		"Attack": strconv.Itoa(int(c.Attack)),
	}
}

func (c *Character) IsCharacterAlive() bool {
	return c.Health > 0
}

func (c *Character) TakeDamage(damage int) {
	if damage < 0 {
		return
	}
	c.Health -= uint(damage)
}

func (c *Character) GainHealth(amount int) {
	if amount < 0 {
		return
	}
	c.Health += uint(amount)
	if c.Health > c.MaxHealth {
		c.Health = c.MaxHealth
	}
}
