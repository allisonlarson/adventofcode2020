package day7

import (
	"aoc/util"
	"fmt"
	"io"
	"strings"
)

type Bag struct {
	color  string
	within []*Bag
}

func Compute(r io.Reader, runPartTwo bool) (int, error) {
	input, err := util.ReadInput(r)
	if err != nil {
		return 0, err
	}
	rules := strings.Split(strings.TrimSpace(input), "\n")
	//lastRule := len(rules) - 1
	//for i := 0; i < len(rules)/2; i++ {
	//	rules[i], rules[lastRule-i] = rules[lastRule-i], rules[i]
	//}
	bags := map[string]*Bag{}
	for _, rule := range rules {
		bagsInRule := strings.Split(rule, " bags contain ")
		color := strings.TrimSpace(bagsInRule[0])
		subjectBag, ok := bags[color]
		if !ok {
			subjectBag = &Bag{color: color, within: []*Bag{}}
			bags[color] = subjectBag
		}
		if strings.Contains(rule, "contain no other bags.") {
			continue
		}
		containedBags := strings.Split(strings.NewReplacer("bags", "", "bag", "", ".", "").Replace(bagsInRule[1]), " , ")
		for _, bag := range containedBags {
			fmt.Println(bag)
			//bagCount, err := strconv.Atoi(string(bag[0]))
			//if err != nil {
			//	return 0, err
			//}
			bagColor := strings.TrimSpace(string(bag[2:]))

			bag, ok := bags[bagColor]
			if ok {
				bag.within = append(bag.within, subjectBag)
				fmt.Printf("appending to bag: %v %v\n", bag, subjectBag)
				//bag.within[bagCount] = subjectBag
				//subjectBag.bags[bagCount] = bag
			} else {
				bag = &Bag{color: bagColor, within: []*Bag{subjectBag}}
				bags[bagColor] = bag

				fmt.Printf("new bag: %v %v\n", bag, subjectBag)
				//bag.within[bagCount] = subjectBag
			}
		}
	}
	//fmt.Println(bags["shiny gold"])
	//fmt.Println(bags["bright white"])
	fmt.Println()
	traveledBags := map[string]int{}

	RecursiveCount(bags["shiny gold"], traveledBags)
	return len(traveledBags), nil
}

func RecursiveCount(bag *Bag, traveled map[string]int) map[string]int {
	if len(bag.within) == 0 {
		return traveled
	}

	for _, bagWithin := range bag.within {
		traveled[bagWithin.color] = 0
		RecursiveCount(bagWithin, traveled)
	}
	return traveled
}
