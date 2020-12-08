package day7

import (
	"aoc/util"
	"io"
	"strconv"
	"strings"
)

type Bag struct {
	color    string
	within   []*Bag
	contains []*Bag
	rules    map[string]int
}

func NewBag(color string) *Bag {
	return &Bag{color: color, within: []*Bag{}, rules: map[string]int{}, contains: []*Bag{}}
}

func Compute(r io.Reader, runPartTwo bool) (int, error) {
	input, err := util.ReadInput(r)
	if err != nil {
		return 0, err
	}
	rules := strings.Split(strings.TrimSpace(input), "\n")
	bags := map[string]*Bag{}
	for _, rule := range rules {
		bagsInRule := strings.Split(rule, " bags contain ")
		color := strings.TrimSpace(bagsInRule[0])
		_, ok := bags[color]
		if !ok {
			bags[color] = NewBag(color)
		}
		subjectBag := bags[color]
		if strings.Contains(rule, "contain no other bags.") {
			continue
		}
		containedBags := strings.Split(strings.NewReplacer("bags", "", "bag", "", ".", "").Replace(bagsInRule[1]), " , ")
		for _, bag := range containedBags {
			bagCount, err := strconv.Atoi(string(bag[0]))
			if err != nil {
				return 0, err
			}
			bagColor := strings.TrimSpace(string(bag[2:]))

			_, ok := bags[bagColor]
			if !ok {
				bags[bagColor] = NewBag(bagColor)
			}
			bag := bags[bagColor]
			bag.within = append(bag.within, subjectBag)
			subjectBag.contains = append(subjectBag.contains, bag)
			subjectBag.rules[bagColor] = bagCount
		}
	}

	if runPartTwo {
		return CountContainedBags(bags["shiny gold"], 0), nil
	}

	return len(CountWithinBags(bags["shiny gold"], map[string]int{})), nil
}
func CountContainedBags(bag *Bag, i int) int {
	if len(bag.contains) == 0 {
		return i
	}

	for _, containedBag := range bag.contains {
		i += bag.rules[containedBag.color]
		for j := 0; j < bag.rules[containedBag.color]; j++ {
			i = CountContainedBags(containedBag, i)
		}
	}
	return i
}

func CountWithinBags(bag *Bag, traveled map[string]int) map[string]int {
	if len(bag.within) == 0 {
		return traveled
	}

	for _, bagWithin := range bag.within {
		// Don't care about the value here, just don't want to double count the same color
		traveled[bagWithin.color] += bagWithin.rules[bag.color]
		CountWithinBags(bagWithin, traveled)
	}
	return traveled
}
