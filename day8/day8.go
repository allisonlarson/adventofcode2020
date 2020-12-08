package day8

import (
	"bufio"
	"errors"
	"io"
	"strconv"
	"strings"
)

type Instruction struct {
	cmd string
	mod int
	run bool
}

func Compute(r io.Reader, runPartTwo bool) (int, error) {
	ins := []*Instruction{}
	trades := []*Instruction{}
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		parts := strings.Fields(scanner.Text())
		mod, err := strconv.Atoi(parts[1])
		if err != nil {
			return 0, nil
		}
		in := &Instruction{cmd: parts[0], mod: mod}
		ins = append(ins, in)
		if in.cmd == "nop" || in.cmd == "jmp" {
			trades = append(trades, in)
		}
	}
	acc, err := runCmds(ins)
	if !runPartTwo {
		return acc, nil
	}

	// Brute force the answer
	for _, trade := range trades {
		for _, ins := range ins {
			ins.run = false
		}

		switch trade.cmd {
		case "nop":
			trade.cmd = "jmp"
		case "jmp":
			trade.cmd = "nop"
		}

		acc, err = runCmds(ins)
		if err == nil {
			return acc, err
		}

		switch trade.cmd {
		case "nop":
			trade.cmd = "jmp"
		case "jmp":
			trade.cmd = "nop"
		}
	}

	return 0, errors.New("Couldn't find a solution")
}

func runCmds(ins []*Instruction) (int, error) {
	acc := 0
	i := 0
	currentIn := ins[i]
	for !currentIn.run {
		switch currentIn.cmd {
		case "acc":
			acc += currentIn.mod
			i++
		case "nop":
			i++
		case "jmp":
			i = i + currentIn.mod
		}

		if i >= len(ins) {
			return acc, nil
		}
		currentIn.run = true
		currentIn = ins[i]
	}
	return acc, errors.New("Exited due to infinite loop")
}
