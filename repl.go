package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(">")
	
		scanner.Scan()
		text := scanner.Text()

		cleaned := cleanInput(text)
		args := []string{}
		if len(cleaned) > 1 {
			args = cleaned[1:]
		}

		if len(cleaned) == 0 {
			continue
		}

		commandName := cleaned[0]
		availableCommands := getCommands()

		command, ok := availableCommands[commandName];

		if !ok {
			fmt.Println("invalid command")
			continue
		}

		err := command.callback(cfg, args...)
		if err != nil {
			fmt.Println(err)
		}
	}
}

type cliCommand struct {
	name string
	description string
	callback func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name: "help",
			description: "shows available commands",
			callback: callbackHelp,
		},
		"map": {
			name: "map",
			description: "shows locations",
			callback: callbackMap,
		},
		"mapb": {
			name: "mapb",
			description: "shows previous locations",
			callback: callbackMapB,
		},
		"explore": {
			name: "explore",
			description: "shows pokemons in specific location",
			callback: callbackExplore,
		},
		"catch": {
			name: "catch",
			description: "catches pokemon",
			callback: callbackCatchPokemon,
		},
		"inspect": {
			name: "inspect",
			description: "shows pokemon details",
			callback: callbackInspectPokemon,
		},
		"pokedex": {
			name: "pokedex",
			description: "shows all caught pokemons",
			callback: callbackPokedex,
		},
		"exit": {
			name: "exit",
			description: "exits program",
			callback: callbackExit,
		},
	}
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	return strings.Fields(lowered)
}