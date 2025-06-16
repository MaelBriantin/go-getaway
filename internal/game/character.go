package game

import (
	"math/rand"
	"strconv"
	"strings"
)

type Character struct {
	Name   string
	MaxHealth int
	Health int
	Attack int
}

func CreateCharacter(name string) Character {
	health := 3 + rand.Intn(2)
	return Character{
		Name:   strings.ToUpper(name),
		Health: health,
		MaxHealth: health,
		Attack: 1 + rand.Intn(2),
	}
}

func (c Character) GetCharacterStats() map[string]string {
	return map[string]string{
		"Name":   c.Name,
		"Health": strconv.Itoa(c.Health),
		"Attack": strconv.Itoa(c.Attack),
	}
}

func (c *Character) IsCharacterAlive() bool {
	return c.Health > 0
}

func (c *Character) TakeDamage(damage int) {
	if damage < 0 {
		return
	}
	c.Health -= damage
	if c.Health < 0 {
		c.Health = 0
	}
}

func (c *Character) GainHealth(amount int) {
	if amount < 0 {
		return
	}
	c.Health += amount
	if c.Health > c.MaxHealth {
		c.Health = c.MaxHealth
	}
}
