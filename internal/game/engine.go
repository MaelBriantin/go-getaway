package game

import (
	"fmt"
	"strconv"
)

func Start() {
	for {
		fmt.Println("Please enter your character's name:")
		fmt.Print("> ")
		var name string
		fmt.Scanln(&name)
		character := CreateCharacter(name)
		fmt.Printf("Welcome, %s! Your adventure begins now.\n", character.Name)

		result := loop(character)
		if result == "exit" {
			break
		}
	}
}

func loop(character Character) string {
	for {
		character.Hud()
		var cmd string
		fmt.Scanln(&cmd)

		switch cmd {
		case "restart":
			fmt.Println("Restarting the game...")
			return "restart"
		case "heal":
			heal(&character)
		case "stats":
			stats(character)
		case "hurt":
			character = hurt(&character)
			if !character.IsCharacterAlive() {
				return gameOver()
			}
		case "exit":
			fmt.Println("Exiting the game. Goodbye!")
			return "exit"
		case "help":
			help()
		default:
			fmt.Println("Unknown command. Please type 'help' for a list of commands.")
		}
	}
}

func help() {
	fmt.Println("Available commands:")
	fmt.Println("  - stats: Show your character's stats")
	fmt.Println("  - hurt: Take damage from an enemy")
	fmt.Println("  - heal: Gain health from a potion")
	fmt.Println("  - exit: Exit the game")
	fmt.Println("  - help: Show this help message")
}

func heal(character *Character) {
	fmt.Println("Enter health amount to gain (max "+ strconv.Itoa(character.MaxHealth-character.Health) + "):")
	var health int
	fmt.Print("> ")
	fmt.Scanln(&health)
	if health <= 0 {
		fmt.Println("Health amount must be greater than 0.")
		return
	}
	if health > (character.MaxHealth - character.Health) {
		health = character.MaxHealth - character.Health
	}
	character.GainHealth(health)
	fmt.Printf("%s gained %d HP. Current HP: %d\n", character.Name, health, character.Health)
}

func stats(character Character) {
	stats := character.GetCharacterStats()
	order := []string{"Name", "Health", "Attack"}
	for _, key := range order {
		fmt.Printf("%s: %s\n", key, stats[key])
	}
}

func hurt(character *Character) Character {
	fmt.Println("Enter damage amount:")
	var damage int
	fmt.Print("> ")
	fmt.Scanln(&damage)
	character.TakeDamage(damage)
	fmt.Printf("%s took %d damage. Current HP: %d\n", character.Name, damage, character.Health)
	return *character
}

func (c *Character) Hud() {
	healt := c.Health
	maxHealth := c.MaxHealth
	healthBar := ""
	for range healt {
		healthBar += "❤️ "
	}
	for i := healt; i < maxHealth; i++ {
		healthBar += "🩶 "
	}
	fmt.Printf("\n")
	fmt.Printf("%s: %s\n", c.Name, healthBar)
	fmt.Printf("> ")
}

func gameOver() string {
	fmt.Println("💀")
	fmt.Println("You have died. Game Over!")
	return "exit"
}
