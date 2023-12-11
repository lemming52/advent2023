# Advent of Code 2023 Solutions

My set of solutions for [Advent of Code 2023](https://adventofcode.com/2023), written in Golang because _I want more practice_.

> Something is wrong with global snow production[.](https://www.youtube.com/watch?v=H45Aki9udK4)

## Running

My inputs from the challenges are all stored in the `inputs` directory, and at the time of writing these files are effectively hardcoded into the running.

To run a particular day (i.e. _trebuchet_, ...) use either the name or the day
```sh
go run main.go -challenge trebuchet
go run main.go -challenge 1
```

To run all days
```
go run main.go -all
```

### Challenge Days

Day | Challenge |Day | Challenge
----|-----------|----|----------
1 | `trebuchet` | 14 | ` `
2 | `cubeconundrum` | 15 | ` `
3 | `gearratios` | 16 | ` `
4 | `scratchcards` | 17 | ` `
5 | `seedfertiliser` | 18 | ` `
6 | `waitforit` | 19 | ` `
7 | `camelcards` | 20 | ` `
8 | `hauntedwasteland` | 21 | ` `
9 | `miragemaintenance` | 22 | ` `
10 | `pipemaze` | 23 | ` `
11 | `cosmicexpansion` | 24 | ` `
12 | ` ` | 25 | ` `
13 | ` `

### Adding new template

To template out a new day, from the root directory run
```sh
python3 scripts/template.py <day_number> <challenge_name>
```
