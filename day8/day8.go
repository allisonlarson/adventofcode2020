package day8

import (
	"bufio"
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
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		parts := strings.Fields(scanner.Text())
		mod, err := strconv.Atoi(parts[1])
		if err != nil {
			return 0, nil
		}
		ins = append(ins, &Instruction{cmd: parts[0], mod: mod})
	}
	acc := 0
	i := 0
	//max := len(ins)
	currentIn := ins[i]
	var oldIn *Instruction
	for !currentIn.run {
		oldIn = currentIn
		switch currentIn.cmd {
		case "acc":
			acc += currentIn.mod
			i++
			currentIn = ins[i]
			oldIn.run = true
		case "nop":
			i++
			currentIn = ins[i]
			oldIn.run = true
		case "jmp":
			i = i + currentIn.mod
			currentIn = ins[i]
			oldIn.run = true
		}
	}
	return acc, nil
}
