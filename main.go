package main

import (
	"advent/solutions/d01trebuchet"
	"advent/solutions/d02cubeconundrum"
	"advent/solutions/d03gearratios"
	"advent/solutions/d04scratchcards"
	"advent/solutions/d05seedfertiliser"
	"advent/solutions/d06waitforit"
	"advent/solutions/d07camelcards"
	"advent/solutions/d08hauntedwasteland"
	"advent/solutions/d09miragemaintenance"
	"advent/solutions/d10pipemaze"
	"advent/solutions/d11cosmicexpansion"
	"advent/solutions/d12hotsprings"
	"advent/solutions/d13pointsofincidence"
	"flag"
	"fmt"
	"time"
)

func main() {
	var challenge string
	flag.StringVar(&challenge, "challenge", "trebuchet", "name or number of challenge")
	all := flag.Bool("all", false, "display all results")
	flag.Parse()

	completed := []string{
		"trebuchet",
		"cubeconundrum",
		"gearratios",
		"scratchcards",
		"seedfertiliser",
		"waitforit",
		"camelcards",
		"hauntedwasteland",
		"miragemaintenance",
		"pipemaze",
		"cosmicexpansion",
		"hotsprings",
		"pointsofincidence",
	}
	if *all {
		previous := time.Now()
		fmt.Println("Start Time: ", time.Now())
		for _, c := range completed {
			s := RunChallenge(c)
			current := time.Now()
			fmt.Println(s, " Duration/ms: ", float64(current.Sub(previous).Microseconds())/1000)
			previous = current
		}
	} else {
		fmt.Println(RunChallenge(challenge))
	}
}

func RunChallenge(challenge string) string {
	var res string
	switch challenge {
	case "trebuchet", "1":
        input := "inputs/d01trebuchet.txt"
		A, B := d01trebuchet.Run(input)
		res = fmt.Sprintf("trebuchet Results A: %s B: %s", A, B)
	case "cubeconundrum", "2":
        input := "inputs/d02cubeconundrum.txt"
		A, B := d02cubeconundrum.Run(input)
		res = fmt.Sprintf("cubeconundrum Results A: %s B: %s", A, B)
	case "gearratios", "3":
        input := "inputs/d03gearratios.txt"
		A, B := d03gearratios.Run(input)
		res = fmt.Sprintf("gearratios Results A: %s B: %s", A, B)
	case "scratchcards", "4":
        input := "inputs/d04scratchcards.txt"
		A, B := d04scratchcards.Run(input)
		res = fmt.Sprintf("scratchcards Results A: %s B: %s", A, B)
	case "seedfertiliser", "5":
        input := "inputs/d05seedfertiliser.txt"
		A, B := d05seedfertiliser.Run(input)
		res = fmt.Sprintf("seedfertiliser Results A: %s B: %s", A, B)
	case "waitforit", "6":
        input := "inputs/d06waitforit.txt"
		A, B := d06waitforit.Run(input)
		res = fmt.Sprintf("waitforit Results A: %s B: %s", A, B)
	case "camelcards", "7":
        input := "inputs/d07camelcards.txt"
		A, B := d07camelcards.Run(input)
		res = fmt.Sprintf("camelcards Results A: %s B: %s", A, B)
	case "hauntedwasteland", "8":
        input := "inputs/d08hauntedwasteland.txt"
		A, B := d08hauntedwasteland.Run(input)
		res = fmt.Sprintf("hauntedwasteland Results A: %s B: %s", A, B)
	case "miragemaintenance", "9":
        input := "inputs/d09miragemaintenance.txt"
		A, B := d09miragemaintenance.Run(input)
		res = fmt.Sprintf("miragemaintenance Results A: %s B: %s", A, B)
	case "pipemaze", "10":
        input := "inputs/d10pipemaze.txt"
		A, B := d10pipemaze.Run(input)
		res = fmt.Sprintf("pipemaze Results A: %s B: %s", A, B)
	case "cosmicexpansion", "11":
        input := "inputs/d11cosmicexpansion.txt"
		A, B := d11cosmicexpansion.Run(input)
		res = fmt.Sprintf("cosmicexpansion Results A: %s B: %s", A, B)
	case "hotsprings", "12":
        input := "inputs/d12hotsprings.txt"
		A, B := d12hotsprings.Run(input)
		res = fmt.Sprintf("hotsprings Results A: %s B: %s", A, B)
	case "pointsofincidence", "13":
        input := "inputs/d13pointsofincidence.txt"
		A, B := d13pointsofincidence.Run(input)
		res = fmt.Sprintf("pointsofincidence Results A: %s B: %s", A, B)

    }
	return res
}
