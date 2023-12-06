package main

import (
	"advent/solutions/d01trebuchet"
	"advent/solutions/d02cubeconundrum"
	"advent/solutions/d03gearratios"
	"advent/solutions/d04scratchcards"
	"advent/solutions/d05seedfertiliser"
	"advent/solutions/d06waitforit"
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

    }
	return res
}
