import argparse
import json
import os
from typing import Dict

CHALLENGE_PATH = "challenges.json"

parser = argparse.ArgumentParser()
parser.add_argument('day', type=int)
parser.add_argument('name', type=str)


def load_challenges(path: str) -> Dict[int, str]:
    vals = {}
    with open(path, 'r') as f:
        vals = json.load(f)
    return {int(k): v for k, v in vals.items()}


def build_main(challenges: Dict[int, str]) -> str:
    imports = '\n'.join([f'\t"advent/solutions/d{i:02d}{name}"' for i, name in challenges.items()])
    names = '\n'.join([f'\t\t"{name}",' for _, name in challenges.items()])
    calls = build_challenge_calls(challenges)
    return f"""package main

import (
{imports}
	"flag"
	"fmt"
	"time"
)

func main() {{
	var challenge string
	flag.StringVar(&challenge, "challenge", "campcleanup", "name or number of challenge")
	all := flag.Bool("all", false, "display all results")
	flag.Parse()

	completed := []string{{
{names}
	}}
	if *all {{
		previous := time.Now()
		fmt.Println("Start Time: ", time.Now())
		for _, c := range completed {{
			s := RunChallenge(c)
			current := time.Now()
			fmt.Println(s, " Duration/ms: ", float64(current.Sub(previous).Microseconds())/1000)
			previous = current
		}}
	}} else {{
		fmt.Println(RunChallenge(challenge))
	}}
}}

func RunChallenge(challenge string) string {{
	var res string
	switch challenge {{
{calls}
    }}
	return res
}}
"""


def build_challenge_calls(challenges: Dict[int, str]) -> str:
    res = []
    for day, name in challenges.items():
        number = str(day)
        full_number = f'{day:02d}'
        template = f"""\tcase "{name}", "{number}":
        input := "inputs/d{full_number}{name}.txt"
		A, B := d{full_number}{name}.Run(input)
		res = fmt.Sprintf("{name} Results A: %s B: %s", A, B)
"""
        res.append(template)
    return ''.join(res)


def build_package_file(number, name: str) -> str:
    return f"""package d{number}{name}

import (
	"bufio"
	"log"
    "os"
)

func Run(path string) (string, string) {{
	lines := utils.LoadAsStrings(path)
	return "A", "B"
}}
"""

def build_test_file(number, name: str) -> str:
    return f"""package d{number}{name}

import (
	"testing"

	"github.com/stretchr/testify/assert"
)
"""


def store_challenges(path, challenges: Dict[int, str]) -> None:
    with open(path, 'w') as f:
        return f.write(json.dumps(challenges))


def main() -> None:
    args = parser.parse_args()
    challenges = load_challenges(CHALLENGE_PATH)
    if args.day in challenges:
        return
    challenges[int(args.day)] = args.name
    number = f'{args.day:02d}'
    with open('main.go', 'w') as f:
        f.write(build_main(challenges))
    os.mkdir(f'solutions/d{number}{args.name}')
    with open(f'solutions/d{number}{args.name}/{args.name}.go', 'w') as f:
        f.write(build_package_file(number, args.name))
    with open(f'solutions/d{number}{args.name}/{args.name}_test.go', 'w') as f:
        f.write(build_test_file(number, args.name))
    with open(f'inputs/d{number}{args.name}.txt', 'w') as f:
        f.write('')
    store_challenges(CHALLENGE_PATH, challenges)

if __name__ == "__main__":
    main()
