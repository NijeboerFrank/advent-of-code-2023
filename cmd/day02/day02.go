package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/alecthomas/participle/v2"
)

const (
	BLUE_MAX  = 14
	RED_MAX   = 12
	GREEN_MAX = 13
)

type Game struct {
	Number int   `"Game" @Int ":"`
	Sets   []Set `@@+`
}

type Set struct {
	Grabs []CubeGrab `@@+";"?`
}

type CubeGrab struct {
	Red   int `@Int "red"","?`
	Blue  int `| @Int "blue"","?`
	Green int `| @Int "green"","?`
}

func process_line(line string, parser *participle.Parser[Game]) Game {
	ast, err := parser.ParseString("", line)
	if err != nil {
		log.Print(err)
		log.Fatal("Couldn't parse string")
	}
	return *ast
}

// Return the value of the game number if the game was possible, 0 otherwise.
func get_game_value(game Game) int {
	for _, set := range game.Sets {
		for _, grab := range set.Grabs {
			if grab.Blue > BLUE_MAX {
				return 0
			} else if grab.Red > RED_MAX {
				return 0
			} else if grab.Green > GREEN_MAX {
				return 0
			}
		}
	}
	return game.Number
}

func process(filename string) int {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	parser, err := participle.Build[Game]()
	if err != nil {
		log.Print(err)
		log.Fatal("Could not build parser")
	}

	scanner := bufio.NewScanner(f)
	games := []Game{}

	for scanner.Scan() {
		games = append(games, process_line(scanner.Text(), parser))
	}

	ret := 0
	for _, game := range games {
		plus := get_game_value(game)
		ret += plus
	}
	return ret
}

func main() {
	answer := process("input2.txt")
	fmt.Printf("The answer is: %d\n", answer)
}
