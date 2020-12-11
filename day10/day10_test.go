package day10

import (
	"strings"
	"testing"
)

func TestComputePart1(t *testing.T) {
	rawInput := `16
10
15
5
1
11
7
19
6
12
4
`

	answer, err := Compute(strings.NewReader(rawInput), false)
	if err != nil {
		t.Error(err)
	}
	if answer != 35 {
		t.Errorf("Expected 35, got %v", answer)
	}
}

func TestComputePart1pt2(t *testing.T) {
	rawInput := `28
33
18
42
31
14
46
20
48
47
24
23
49
45
19
38
39
11
1
32
25
35
8
17
7
9
4
2
34
10
3
`

	answer, err := Compute(strings.NewReader(rawInput), false)
	if err != nil {
		t.Error(err)
	}
	if answer != 220 {
		t.Errorf("Expected 220, got %v", answer)
	}
}

func TestComputePart2(t *testing.T) {
	rawInput := `16
10
15
5
1
11
7
19
6
12
4
`

	answer, err := Compute(strings.NewReader(rawInput), true)
	if err != nil {
		t.Error(err)
	}
	if answer != 8 {
		t.Errorf("Expected 8, got %v", answer)
	}
}

func TestComputePart2pt2(t *testing.T) {
	rawInput := `28
33
18
42
31
14
46
20
48
47
24
23
49
45
19
38
39
11
1
32
25
35
8
17
7
9
4
2
34
10
3
`

	answer, err := Compute(strings.NewReader(rawInput), true)
	if err != nil {
		t.Error(err)
	}
	if answer != 19208 {
		t.Errorf("Expected 19208, got %v", answer)
	}
}
