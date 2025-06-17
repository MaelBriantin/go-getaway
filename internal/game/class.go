package game

import (
	"fmt"
	"github.com/MaelBriantin/go-getaway/internal/utils"
)

type Class struct {
	Name        string
	Description string
	Health      uint
	Attack      uint
	Arcane      uint
}

var Warrior = Class{
	Name:        "Warrior",
	Description: "A strong and resilient fighter, skilled in melee combat.",
	Health:      10,
	Attack:      6,
	Arcane:			 0,
}

var Mage = Class{
	Name:        "Mage",
	Description: "A master of the arcane arts, capable of casting powerful spells.",
	Health:      6,
	Attack:      2,
	Arcane:			 8,
}

var Rogue = Class{
	Name:        "Rogue",
	Description: "A stealthy and agile character, adept at sneaking and dealing critical hits.",
	Health:      8,
	Attack:      4,
	Arcane:			 2,
}

var Classes = []Class{Warrior, Mage, Rogue}

func GetClassByName(name string) (Class, error) {
	capitalizedName := utils.CapitalizeFirst(name)
	switch capitalizedName {
	case "Warrior":
		return Warrior, nil
	case "Mage":
		return Mage, nil
	case "Rogue":
		return Rogue, nil
	default:
		return Class{}, fmt.Errorf("unknown class: %s", name)
	}
}