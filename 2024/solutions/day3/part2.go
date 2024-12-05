package day3

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/gomsim/Advent-of-code/shared/exit"
	"github.com/gomsim/Advent-of-code/shared/input"
	"github.com/gomsim/Advent-of-code/shared/register"
)

func init() {
	register.Part2("2024", "day3", Part2)
}

var opPattern = regexp.MustCompile(`(mul|do|don't)\((\d+(,\d+)*)?\)`)

const (
	mul  = "mul"
	do   = "do"
	dont = "don't"
)

type operation struct {
	op   string
	args []int
}

func Part2(file string) {
	in := input.String(file)

	exec := parseProgram(in)
	output := run(exec)

	fmt.Println(output)
}

func parseProgram(input string) (exec []operation) {
	for _, match := range opPattern.FindAllString(input, -1) {
		tokens := strings.Split(match, "(")
		switch tokens[0] {
		case mul:
			exec = append(exec, operation{
				op:   mul,
				args: parseArgs(match),
			})
		case do:
			exec = append(exec, operation{
				op:   do,
				args: nil,
			})
		case dont:
			exec = append(exec, operation{
				op:   dont,
				args: nil,
			})
		}
	}
	return exec
}

func parseArgs(input string) (args []int) {
	washPattern := regexp.MustCompile(`[a-z']+\(|\)`)
	washed := washPattern.ReplaceAllString(input, "")
	for _, arg := range strings.Split(washed, ",") {
		asInt, err := strconv.Atoi(arg)
		exit.If(err)
		args = append(args, asInt)
	}

	return args
}

func run(exec []operation) (acc int) {
	enabled := true
	for _, op := range exec {
		switch op.op {
		case mul:
			if enabled {
				acc += op.args[0] * op.args[1]
			}
		case do:
			enabled = true
		case dont:
			enabled = false
		}
	}
	return acc
}
