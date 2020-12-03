package day3

import (
	"bufio"
	"io"
)

type route struct {
	xSlope   int
	ySlope   int
	x        int
	treesHit int
}

func newRoute(xSlope, ySlope int) *route {
	return &route{
		xSlope: xSlope,
		ySlope: ySlope,
		x:      xSlope,
	}
}

func Compute(r io.Reader, runPartTwo bool) (int, error) {
	routes := []*route{}
	routes = append(routes, newRoute(3, 1))
	if runPartTwo {
		routes = append(routes, newRoute(1, 1))
		routes = append(routes, newRoute(5, 1))
		routes = append(routes, newRoute(7, 1))
		routes = append(routes, newRoute(1, 2))
	}

	treeMap := []string{}
	scanner := bufio.NewScanner(r)
	i := 0
	for scanner.Scan() {
		treeMap = append(treeMap, scanner.Text())
	}
	i++

	for _, route := range routes {
		for i := 0; i < len(treeMap); i += route.ySlope {
			if i >= route.ySlope && i%route.ySlope == 0 {
				if treeMap[i][route.x%len(treeMap[i])] == '#' {
					route.treesHit += 1
				}
				route.x += route.xSlope
			}
		}

	}
	if runPartTwo {
		total := 1
		for _, route := range routes {
			total = total * route.treesHit
		}
		return total, scanner.Err()
	}
	return routes[0].treesHit, scanner.Err()
}
